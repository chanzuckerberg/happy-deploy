locals {
  create_ingress    = (var.routing.service_type == "EXTERNAL" || var.routing.service_type == "INTERNAL")
  listen_ports_base = [{ "HTTP" : 80 }]
  listen_ports_tls  = [merge(local.listen_ports_base[0], { "HTTPS" : 443 })]
  ingress_base_annotations = {
    "alb.ingress.kubernetes.io/backend-protocol"     = "HTTP"
    "alb.ingress.kubernetes.io/healthcheck-path"     = var.health_check_path
    "alb.ingress.kubernetes.io/healthcheck-protocol" = "HTTP"
    # All ingresses are "internet-facing" so we need them all to listen on TLS
    "alb.ingress.kubernetes.io/listen-ports" = jsonencode(local.listen_ports_tls)
    # All ingresses are "internet-facing". If a service_type was marked "INTERNAL", it will be protected using OIDC.
    "alb.ingress.kubernetes.io/scheme"                  = "internet-facing"
    "alb.ingress.kubernetes.io/subnets"                 = join(",", var.cloud_env.public_subnets)
    "alb.ingress.kubernetes.io/success-codes"           = var.routing.success_codes
    "alb.ingress.kubernetes.io/tags"                    = var.tags_string
    "alb.ingress.kubernetes.io/target-group-attributes" = "deregistration_delay.timeout_seconds=60"
    "alb.ingress.kubernetes.io/target-type"             = "instance"
    "kubernetes.io/ingress.class"                       = "alb"
    "alb.ingress.kubernetes.io/group.name"              = var.routing.group_name
    "alb.ingress.kubernetes.io/group.order"             = var.routing.priority
  }
  ingress_tls_annotations = {
    "alb.ingress.kubernetes.io/actions.redirect" = <<EOT
        {"Type": "redirect", "RedirectConfig": {"Protocol": "HTTPS", "Port": "443", "StatusCode": "HTTP_301"}}
      EOT
    "alb.ingress.kubernetes.io/certificate-arn"  = var.certificate_arn
    "alb.ingress.kubernetes.io/ssl-policy"       = "ELBSecurityPolicy-TLS-1-2-2017-01"
  }
  ingress_auth_annotations = {
    "alb.ingress.kubernetes.io/auth-type"                       = "oidc"
    "alb.ingress.kubernetes.io/auth-on-unauthenticated-request" = "authenticate"
    "alb.ingress.kubernetes.io/auth-idp-oidc"                   = jsonencode(var.routing.oidc_config)
  }
  ingress_annotations = (
    var.routing.service_type == "EXTERNAL" ?
    merge(local.ingress_tls_annotations, local.ingress_base_annotations) :
    merge(local.ingress_tls_annotations, local.ingress_auth_annotations, local.ingress_base_annotations)
  )
}

resource "kubernetes_ingress_v1" "ingress" {
  count = local.create_ingress ? 1 : 0
  metadata {
    name      = var.ingress_name
    namespace = var.k8s_namespace
    labels = {
      app = var.ingress_name
    }

    annotations = local.ingress_annotations
  }

  spec {
    rule {
      http {
        path {
          backend {
            service {
              name = "redirect"
              port {
                name = "use-annotation"
              }
            }
          }

          path = "/*"
        }
      }
    }


    rule {

      host = var.routing.host_match
      http {
        path {
          path = var.routing.path
          backend {
            service {
              name = var.routing.service_name
              port {
                number = var.routing.service_port
              }
            }
          }
        }
      }
    }
  }
}
