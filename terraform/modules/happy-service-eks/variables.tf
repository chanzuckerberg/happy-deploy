variable "cpu" {
  type        = string
  description = "CPU shares (1cpu=1000m) per pod"
  default     = "100m"
}

variable "cpu_requests" {
  type        = string
  description = "CPU shares (1cpu=1000m) requested per pod"
  default     = "10m"
}

variable "memory" {
  type        = string
  description = "Memory in megabits per pod"
  default     = "100Mi"
}

variable "memory_requests" {
  type        = string
  description = "Memory requests per pod"
  default     = "10Mi"
}

variable "cmd" {
  type        = list(string)
  description = "Command to run"
  default     = []
}

variable "args" {
  type        = list(string)
  description = "Args to pass to the command"
  default     = []
}

variable "image_tag" {
  type        = string
  description = "The image tag to deploy"
}

variable "image_pull_policy" {
  type        = string
  description = "The image pull policy to use"
  default     = "IfNotPresent"
}

variable "desired_count" {
  type        = number
  description = "How many instances of this task should we run across our cluster?"
  default     = 2
}

variable "max_count" {
  type        = number
  description = "The maximum number of instances of this task that should be running across our cluster"
  default     = 2
}

variable "scaling_cpu_threshold_percentage" {
  type        = number
  description = "The CPU threshold percentage at which we should scale up"
  default     = 80
}

variable "stack_name" {
  type        = string
  description = "Happy Path stack name"
}

variable "cloud_env" {
  type = object({
    public_subnets : list(string),
    private_subnets : list(string),
    database_subnets : list(string),
    database_subnet_group : string,
    vpc_id : string,
    vpc_cidr_block : string,
  })
  description = "Typically data.terraform_remote_state.cloud-env.outputs"
}

variable "deployment_stage" {
  type        = string
  description = "The name of the deployment stage of the Application"
  default     = "dev"
}

variable "health_check_path" {
  type        = string
  description = "path to use for health checks"
  default     = "/"
}

variable "wait_for_steady_state" {
  type        = bool
  description = "Whether Terraform should block until the service is in a steady state before exiting"
  default     = true
}

variable "k8s_namespace" {
  type        = string
  description = "K8S namespace for this service"
}

variable "certificate_arn" {
  type        = string
  description = "ACM certificate ARN to attach to the load balancer listener"
}

variable "container_name" {
  type        = string
  description = "The name of the container"
}

variable "service_endpoints" {
  type        = map(string)
  default     = {}
  description = "Service endpoints to be injected for service discovery"
}

variable "period_seconds" {
  type        = number
  default     = 3
  description = "The period in seconds used for the liveness and readiness probes."
}

variable "initial_delay_seconds" {
  type        = number
  default     = 30
  description = "The initial delay in seconds for the liveness and readiness probes."
}

variable "platform_architecture" {
  type        = string
  description = "The platform to deploy to (valid values: `amd64`, `arm64`). Defaults to `amd64`."
  default     = "amd64"

  validation {
    condition     = var.platform_architecture == "amd64" || var.platform_architecture == "arm64"
    error_message = "Must be one of `amd64` or `arm64`."
  }
}

variable "aws_iam_policy_json" {
  type        = string
  default     = ""
  description = "The AWS IAM policy to give to the pod."
}

variable "eks_cluster" {
  type = object({
    cluster_id : string,
    cluster_arn : string,
    cluster_endpoint : string,
    cluster_ca : string,
    cluster_oidc_issuer_url : string,
    cluster_version : string,
    worker_iam_role_name : string,
    worker_security_group : string,
    oidc_provider_arn : string,
  })
  description = "eks-cluster module output"
}

variable "additional_env_vars" {
  type        = map(string)
  description = "Additional environment variables to add to the task definition"
  default     = {}
}

variable "additional_env_vars_from_config_maps" {
  type = object({
    items : optional(list(string), []),
    prefix : optional(string, ""),
  })
  default = {
    items  = []
    prefix = ""
  }
  description = "Additional environment variables to add to the container from the following config maps"
}

variable "additional_env_vars_from_secrets" {
  type = object({
    items : optional(list(string), []),
    prefix : optional(string, ""),
  })
  default = {
    items  = []
    prefix = ""
  }
  description = "Additional environment variables to add to the container from the following secrets"
}

variable "additional_volumes_from_secrets" {
  type = object({
    items : optional(list(string), []),
    base_dir : optional(string, "/var"),
  })
  default = {
    items    = []
    base_dir = "/var"
  }
  description = "Additional volumes to add to the container from the following secrets"
}

variable "additional_volumes_from_config_maps" {
  type = object({
    items : optional(list(string), []),
  })
  default = {
    items = []
  }
  description = "Additional volumes to add to the container from the following config maps"
}

variable "routing" {
  type = object({
    method : optional(string, "DOMAIN")
    host_match : string
    group_name : string
    alb : optional(object({
      name : string,
      listener_port : number,
    }), null)
    priority : number
    path : optional(string, "/*")
    service_name : string
    port : number
    service_port : number
    service_scheme : optional(string, "HTTP")
    scheme : optional(string, "HTTP")
    success_codes : optional(string, "200-499")
    service_type : string
    service_mesh : bool
    allow_mesh_services : optional(list(object({
      service : string,
      stack : string
    })), null)
    oidc_config : optional(object({
      issuer : string
      authorizationEndpoint : string
      tokenEndpoint : string
      userInfoEndpoint : string
      secretName : string
      }), {
      issuer                = ""
      authorizationEndpoint = ""
      tokenEndpoint         = ""
      userInfoEndpoint      = ""
      secretName            = ""
    })
    bypasses : optional(map(object({
      paths   = optional(set(string), [])
      methods = optional(set(string), [])
    })))
  })
  description = "Routing configuration for the ingress"

  validation {
    condition     = var.routing.service_mesh == true || var.routing.allow_mesh_services == null
    error_message = "The allow_mesh_services option is only supported if service_mesh is enabled on the stack"
  }
}

variable "sidecars" {
  type = map(object({
    image : string
    tag : string
    port : optional(number, 80)
    scheme : optional(string, "HTTP")
    memory : optional(string, "100Mi")
    cpu : optional(string, "100m")
    image_pull_policy : optional(string, "IfNotPresent")
    health_check_path : optional(string, "/")
    initial_delay_seconds : optional(number, 30),
    period_seconds : optional(number, 3),
  }))
  default     = {}
  description = "Map of sidecar containers to be deployed alongside the service"

  validation {
    condition = alltrue([for k, v in var.sidecars : (
      v.scheme == "HTTP" ||
      v.scheme == "HTTPS"
    )])
    error_message = "The scheme argument needs to be 'HTTP' or 'HTTPS'."
  }
}

variable "tags" {
  description = "Standard tags to attach to all happy services"
  type = object({
    env : string,
    owner : string,
    project : string,
    service : string,
    managedBy : string,
  })
  default = {
    env       = "ADDTAGS"
    managedBy = "ADDTAGS"
    owner     = "ADDTAGS"
    project   = "ADDTAGS"
    service   = "ADDTAGS"
  }
}

variable "regional_wafv2_arn" {
  type        = string
  description = "A WAF to protect the EKS Ingress if needed"
  default     = null
}

variable "additional_pod_labels" {
  type        = map(string)
  description = "Additional labels to add to the pods."
  default     = {}
}
