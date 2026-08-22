variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
}

variable "instance_name" {
  description = "Memorystore Redis instance name"
  type        = string
}

variable "memory_size_gb" {
  description = "Redis memory size in GB"
  type        = number
  default     = 1
}

variable "environment" {
  description = "Environment name (dev or prod)"
  type        = string
}

variable "authorized_network" {
  description = "Self link of the VPC network Redis should be reachable from. Defaults to the shared 'default' network for backward compatibility."
  type        = string
  default     = null
}