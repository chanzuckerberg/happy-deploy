locals {
  # this is making an assumption that the health check path is accessible to the internet
  # if this is a non-prod OIDC protected stack, make sure to allow the health check endpoint
  # through the OIDC proxy.
  synthetics = { for k, v in local.service_definitions : v.service_name =>
    v.service_type == "EXTERNAL" ? "https://${v.service_name}.${local.external_dns}${v.health_check_path}" : "https://${v.service_name}.${local.internal_dns}${v.health_check_path}"
    if v.synthetics
  }
}

data "datadog_synthetics_locations" "locations" {}

resource "datadog_synthetics_test" "test_api" {
  for_each = local.synthetics
  type     = "api"
  subtype  = "http"
  request_definition {
    method = "GET"
    url    = each.value
  }
  assertion {
    type     = "statusCode"
    operator = "is"
    target   = "200"
  }
  locations = keys(data.datadog_synthetics_locations.locations.locations)
  options_list {
    tick_every = 900

    retry {
      count    = 2
      interval = 300
    }

    monitor_options {
      renotify_interval = 120
    }
  }
  name    = "A website synthetic for the happy stack ${var.deployment_stage} ${var.stack_name} ${each.key} located at ${each.value}"
  message = "Notify @opsgenie-${var.stack_name}-${var.deployment_stage}-${each.key}"
  status  = "live"
}
