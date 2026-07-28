variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "network" {
  description = "VPC network name"
  type        = string
  default     = "default"
}