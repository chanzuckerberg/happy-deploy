# Auto-generated by 'happy infra'. Do not edit
# Make improvements in happy, so that everyone can benefit.
module "stack" {
  source           = "git@github.com:chanzuckerberg/happy//terraform/modules/happy-stack-eks?ref=alokshin/cerbos-sidecar"
  image_tag        = var.image_tag
  stack_name       = var.stack_name
  k8s_namespace    = var.k8s_namespace
  image_tags       = jsondecode(var.image_tags)
  stack_prefix     = "/${var.stack_name}"
  app_name         = "cerbos-sidecar"
  deployment_stage = "rdev"
  services = {
    frontend = {
      cpu                              = "100m"
      desired_count                    = 1
      health_check_path                = "/_cerbos/health"
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
      sidecars = {
        echo = {
          cpu                   = "100m"
          health_check_path     = "/"
          health_check_scheme   = "HTTP"
          image                 = "ealen/echo-server"
          image_pull_policy     = "IfNotPresent"
          initial_delay_seconds = 30
          memory                = "128Mi"
          period_seconds        = 3
          port                  = 80
          tag                   = "0.7.0"
          volume_mounts = {
            config = {
              mount_path = "/config"
              read_only  = true
            }
            policies = {
              mount_path = "/policies"
              read_only  = null
            }
            sock = {
              mount_path = "/sock"
              read_only  = null
            }
          }
        }
      }
      success_codes = "200-399"
      synthetics    = false
      volumes = {
        certs = {
          ref  = "$-cerbos-sidecar-demo"
          type = "SECRET"
        }
        config = {
          ref  = "$-cerbos-sidecar-demo"
          type = "CONFIGMAP"
        }
        policies = {
          ref  = ""
          type = "EMPTY_DIR"
        }
        sock = {
          ref  = ""
          type = "EMPTY_DIR"
        }
      }
    }
  }
  create_dashboard = false
  routing_method   = "CONTEXT"
}
