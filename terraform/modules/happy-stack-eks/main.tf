data "kubernetes_secret" "integration_secret" {
  metadata {
    name      = "integration-secret"
    namespace = var.k8s_namespace
  }
}

resource "validation_error" "mix_of_internal_and_external_services" {
  condition = length(local.external_services) > 0 && length(local.internal_services) > 0 && var.routing_method == "DOMAIN"
  summary   = "Invalid mix of INTERNAL and EXTERNAL services"
  details   = "With DOMAIN routing, a mix of EXTERNAL and INTERNAL services is not permitted; only EXTERNAL and PRIVATE can be mixed"
}

locals {
  # kubernetes_secret resource is always marked sensitive, which makes things a little difficult
  # when decoding pieces of the integration secret later. Mark the whole thing as nonsensitive and only
  # output particular fields as sensitive in this modules outputs (for instance, the RDS password)
  secret       = jsondecode(nonsensitive(data.kubernetes_secret.integration_secret.data.integration_secret))
  external_dns = local.secret["external_zone_name"]

  service_defs = { for k, v in var.services : k => merge(v, {
    external_stack_host_match   = try(join(".", [var.stack_name, local.external_dns]), "")
    external_service_host_match = try(join(".", ["${var.stack_name}-${k}", local.external_dns]), "")

    stack_host_match   = try(join(".", [var.stack_name, local.external_dns]), "")
    service_host_match = try(join(".", ["${var.stack_name}-${k}", local.external_dns]), "")
  }) }

  service_definitions = { for k, v in local.service_defs : k => merge(v, {
    external_host_match = var.routing_method == "CONTEXT" ? v.external_stack_host_match : v.external_service_host_match
    host_match          = var.routing_method == "CONTEXT" ? v.stack_host_match : v.service_host_match
    group_name          = var.routing_method == "CONTEXT" ? "stack-${var.stack_name}" : "service-${var.stack_name}-${k}"
    service_name        = "${var.stack_name}-${k}"
  }) }

  external_services = [for v in var.services : v if v.service_type == "EXTERNAL"]
  internal_services = [for v in var.services : v if v.service_type == "INTERNAL"]

  task_definitions = { for k, v in var.tasks : k => merge(v, {
    task_name = "${var.stack_name}-${k}"
  }) }

  external_endpoints = concat([for k, v in local.service_definitions :
    v.service_type == "EXTERNAL" ?
    {
      "EXTERNAL_${upper(replace(k, "-", "_"))}_ENDPOINT" = try(join("", ["https://", v.external_host_match]), "")
      "${upper(replace(k, "-", "_"))}_ENDPOINT"          = try(join("", ["https://", v.host_match]), "")
    }
    : {
      "EXTERNAL_${upper(replace(k, "-", "_"))}_ENDPOINT" = try(join("", ["https://", v.external_host_match]), "")
      "INTERNAL_${upper(replace(k, "-", "_"))}_ENDPOINT" = try(join("", ["https://", v.host_match]), "")
      "${upper(replace(k, "-", "_"))}_ENDPOINT"          = try(join("", ["https://", v.host_match]), "")
    }
  ])

  private_endpoints = concat([for k, v in local.service_definitions :
    {
      "PRIVATE_${upper(replace(k, "-", "_"))}_ENDPOINT" = "http://${v.service_name}.${var.k8s_namespace}.svc.cluster.local:${v.port}"
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

  service_endpoints = merge(local.flat_external_endpoints, local.flat_private_endpoints)

  db_env_vars = merge(flatten(
    [for dbname, dbcongif in local.secret["dbs"] : [
      for varname, value in dbcongif : { upper(replace("${dbname}_${varname}", "/[^a-zA-Z0-9_]/", "_")) : value }
    ]]
  )...)

  oidc_config_secret_name = "${var.stack_name}-oidc-config"
  issuer_domain           = try(local.secret["oidc_config"]["idp_url"], "todofindissuer.com")
  issuer_url              = "https://${local.issuer_domain}"
  oidc_config = {
    issuer                = local.issuer_url
    authorizationEndpoint = "${local.issuer_url}/oauth2/v1/authorize"
    tokenEndpoint         = "${local.issuer_url}/oauth2/v1/token"
    userInfoEndpoint      = "${local.issuer_url}/oauth2/v1/userinfo"
    secretName            = local.oidc_config_secret_name
  }
}

resource "kubernetes_secret" "oidc_config" {
  metadata {
    name      = local.oidc_config_secret_name
    namespace = var.k8s_namespace
  }

  data = {
    clientID     = try(local.secret["oidc_config"]["client_id"], "")
    clientSecret = try(local.secret["oidc_config"]["client_secret"], "")
  }
}

module "services" {
  for_each              = local.service_definitions
  source                = "../happy-service-eks"
  image                 = join(":", [local.secret["ecrs"][each.key]["url"], lookup(var.image_tags, each.key, var.image_tag)])
  container_name        = each.value.name
  stack_name            = var.stack_name
  desired_count         = each.value.desired_count
  memory                = each.value.memory
  cpu                   = each.value.cpu
  health_check_path     = each.value.health_check_path
  k8s_namespace         = var.k8s_namespace
  cloud_env             = local.secret["cloud_env"]
  certificate_arn       = local.secret["certificate_arn"]
  deployment_stage      = var.deployment_stage
  service_endpoints     = local.service_endpoints
  aws_iam_policy_json   = each.value.aws_iam_policy_json
  eks_cluster           = local.secret["eks_cluster"]
  initial_delay_seconds = each.value.initial_delay_seconds
  period_seconds        = each.value.period_seconds
  routing = {
    method        = var.routing_method
    host_match    = each.value.host_match
    group_name    = each.value.group_name
    priority      = each.value.priority
    path          = each.value.path
    service_name  = each.value.service_name
    service_port  = each.value.port
    success_codes = each.value.success_codes
    service_type  = each.value.service_type
    oidc_config   = local.oidc_config
  }

  additional_env_vars                  = merge(local.db_env_vars, var.additional_env_vars, local.stack_configs)
  additional_env_vars_from_config_maps = var.additional_env_vars_from_config_maps
  additional_env_vars_from_secrets     = var.additional_env_vars_from_secrets

  tags = local.secret["tags"]
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
