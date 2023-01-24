data "aws_region" "current" {}
data "kubernetes_secret" "integration_secret" {
  metadata {
    name      = "integration-secret"
    namespace = var.k8s_namespace
  }
}

locals {
  # kubernetes_secret resource is always marked sensitive, which makes things a little difficult
  # when decoding pieces of the integration secret later. Mark the whole thing as nonsensitive and only
  # output particular fields as sensitive in this modules outputs (for instance, the RDS password)
  secret       = jsondecode(nonsensitive(data.kubernetes_secret.integration_secret.data.integration_secret))
  external_dns = local.secret["external_zone_name"]
  internal_dns = local.secret["internal_zone_name"]

  stack_host_match = var.stack_ingress.service_type == "INTERNAL" ? try(join(".", [var.stack_name, "internal", local.external_dns]), "") : try(join(".", [var.stack_name, local.external_dns]), "")

  service_definitions = { for k, v in var.services : k => merge(v, {
    host_match   = v.service_type == "INTERNAL" ? try(join(".", ["${var.stack_name}-${k}", "internal", local.external_dns]), "") : try(join(".", ["${var.stack_name}-${k}", local.external_dns]), "")
    service_name = "${var.stack_name}-${k}"
  }) }

  task_definitions = { for k, v in var.tasks : k => merge(v, {
    task_name = "${var.stack_name}-${k}"
  }) }

  backends = [for k, v in var.stack_ingress.backends : merge(v, {
    service_name = "${var.stack_name}-${v.service_name}"
  })]

  external_endpoints = concat([for k, v in local.service_definitions :
    v.service_type == "EXTERNAL" && v.create_ingress ?
    {
      "EXTERNAL_${upper(k)}_ENDPOINT" = try(join("", ["https://", v.service_name, ".", local.external_dns]), "")
    }
    : {
      "INTERNAL_${upper(k)}_ENDPOINT" = try(join("", ["https://", v.service_name, ".", local.internal_dns]), "")
    }
  ])

  stack_external_endpoint = var.stack_ingress.create_ingress ? (
    {
      "EXTERNAL__STACK__ENDPOINT" = try(join("", ["https://", try(join(".", [var.stack_name, local.external_dns]), "")]), "")
    }
  ) : {}


  private_endpoints = concat([for k, v in local.service_definitions :
    {
      "PRIVATE_${upper(k)}_ENDPOINT" = "http://${v.service_name}.${var.k8s_namespace}.svc.cluster.local:${v.port}"
    }
  ])

  flat_external_endpoints = zipmap(
    flatten(
      [for item in local.external_endpoints : keys(item)]
    ),
    flatten(
      [for item in local.external_endpoints : values(item)]
    )
  )

  flat_private_endpoints = zipmap(
    flatten(
      [for item in local.private_endpoints : keys(item)]
    ),
    flatten(
      [for item in local.private_endpoints : values(item)]
    )
  )

  service_endpoints = merge(local.flat_external_endpoints, local.flat_private_endpoints, local.stack_external_endpoint)

  db_env_vars = merge(flatten(
    [for dbname, dbcongif in local.secret["dbs"] : [
      for varname, value in dbcongif : { upper(replace("${dbname}_${varname}", "/[^a-zA-Z0-9_]/", "_")) : value }
    ]]
  )...)

  stack_tags = {
    happy_env          = var.deployment_stage,
    happy_stack_name   = var.stack_name,
    happy_region       = data.aws_region.current.name,
    happy_service_type = var.stack_ingress.service_type,
    happy_last_applied = timestamp(),
  }

  stack_tags_string = join(",", [for key, val in local.stack_tags : "${key}=${val}"])
}

module "services" {
  for_each              = local.service_definitions
  source                = "../happy-service-eks"
  image                 = join(":", [local.secret["ecrs"][each.key]["url"], lookup(var.image_tags, each.key, var.image_tag)])
  container_name        = each.value.name
  stack_name            = var.stack_name
  desired_count         = each.value.desired_count
  service_type          = each.value.service_type
  memory                = each.value.memory
  cpu                   = each.value.cpu
  health_check_path     = each.value.health_check_path
  k8s_namespace         = var.k8s_namespace
  cloud_env             = local.secret["cloud_env"]
  certificate_arn       = local.secret["certificate_arn"]
  oauth_certificate_arn = local.secret["oauth_certificate_arn"]
  deployment_stage      = var.deployment_stage
  service_endpoints     = local.service_endpoints
  aws_iam_policy_json   = each.value.aws_iam_policy_json
  eks_cluster           = local.secret["eks_cluster"]
  additional_env_vars   = local.db_env_vars
  routing = {
    method       = var.routing_method
    host_match   = var.routing_method == "DOMAIN" ? local.stack_host_match : each.value.host_match
    group_name   = var.routing_method == "DOMAIN" ? "stack-${var.stack_name}" : "service-${each.value.service_name}"
    priority     = each.value.priority
    path         = each.value.path
    service_name = each.value.service_name
    service_port = each.value.port
  }
}

module "tasks" {
  for_each          = local.task_definitions
  source            = "../happy-task-eks"
  task_name         = each.value.task_name
  image             = each.value.image
  cpu               = each.value.cpu
  memory            = each.value.memory
  cmd               = each.value.cmd
  remote_dev_prefix = var.stack_prefix
  deployment_stage  = var.deployment_stage
  k8s_namespace     = var.k8s_namespace
  stack_name        = var.stack_name
}
