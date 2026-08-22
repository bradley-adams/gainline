variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "network" {
  description = "VPC network name to use when use_default_network is true"
  type        = string
  default     = "default"
}

variable "use_default_network" {
  description = "If true, peer against the shared 'default' network (legacy dev behavior). If false, create a dedicated VPC network + subnet for this environment."
  type        = bool
  default     = true
}

variable "environment" {
  description = "Environment name (dev or prod) — used to name the dedicated VPC when use_default_network is false"
  type        = string
  default     = "dev"
}

variable "region" {
  description = "GCP region for the dedicated subnet — required when use_default_network is false"
  type        = string
  default     = null
}

variable "subnet_cidr" {
  description = "CIDR range for the dedicated subnet — only used when use_default_network is false"
  type        = string
  default     = "10.10.0.0/20"
}