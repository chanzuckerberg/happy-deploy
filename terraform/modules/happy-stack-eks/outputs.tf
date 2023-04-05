output "service_endpoints" {
  value       = local.service_endpoints
  description = "The URL endpoints for services"
  sensitive   = false
}

output "task_arns" {
  value       = { for name, task in module.tasks : name => task.task_definition_arn }
  description = "ARNs for all the tasks"
}

output "dashboard" {
  value = {
    id  = var.create_dashboard ? datadog_dashboard_json.stack_dashboard[0].id : ""
    url = var.create_dashboard ? datadog_dashboard_json.stack_dashboard[0].url : ""
  }
}

output "service_ecrs" {
  value = { for k, v in module.services : k => v.ecr.url }
}

output "target_group_arns" {
  value = { for k, v in module.services : k => [for x in v.target_groups : x.arn] }
}