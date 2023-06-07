# Auto-generated by 'happy infra'. Do not edit
# Make improvements in happy, so that everyone can benefit.
module "stack" {
  source           = "git@github.com:chanzuckerberg/happy//terraform/modules/happy-stack-eks?ref=main"
  image_tag        = var.image_tag
  stack_name       = var.stack_name
  k8s_namespace    = var.k8s_namespace
  image_tags       = jsondecode(var.image_tags)
  stack_prefix     = "/${var.stack_name}"
  app_name         = "hapi"
  deployment_stage = "prod"
  services = {
    hapi = {
      desired_count         = 2
      name                  = "hapi"
      platform_architecture = "arm64"
      port                  = 3001
      service_type          = "EXTERNAL"
    }
  }
  routing_method = "CONTEXT"
  additional_env_vars_from_secrets = {
    items = ["hapi-prod-ssm-secrets"]
  }
}
