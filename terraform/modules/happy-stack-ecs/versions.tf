terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.45"
    }
    datadog = {
      source  = "datadog/datadog"
      version = ">= 3.20.0"
    }
    happy = {
      source  = "chanzuckerberg/happy"
      version = ">= 0.52.0"
    }
  }
  required_version = ">= 1.3"
}
