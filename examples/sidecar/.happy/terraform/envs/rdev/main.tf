# This file is autogenerated by 'happy infra generate'. Do not edit manually.
module "stack" {
  source           = "git@github.com:chanzuckerberg/happy//terraform/modules/happy-stack-eks?ref=main"
  image_tag        = var.image_tag
  stack_name       = var.stack_name
  k8s_namespace    = var.k8s_namespace
  image_tags       = jsondecode(var.image_tags)
  stack_prefix     = "/${var.stack_name}"
  app_name         = "sidecar"
  deployment_stage = "rdev"
  services = {
    frontend = {
      cpu                              = "100m"
      desired_count                    = 1
      health_check_path                = "/"
      initial_delay_seconds            = 30
      max_count                        = 1
      memory                           = "128Mi"
      name                             = "frontend"
      path                             = "/*"
      period_seconds                   = 3
      platform_architecture            = "amd64"
      port                             = 3000
      priority                         = 0
      scaling_cpu_threshold_percentage = 80
      service_type                     = "INTERNAL"
      success_codes                    = "200-499"
      synthetics                       = false

      sidecars = {
        echo = {
          image  = "ealen/echo-server"
          tag    = "0.7.0"
          port   = 80
          cpu    = "100m"
          memory = "128Mi"
        }
      }
    }
  }
  create_dashboard = false
  routing_method   = "CONTEXT"
}
