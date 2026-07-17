variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
  default     = "australia-southeast1"
}

variable "zone" {
  description = "GCP zone"
  type        = string
  default     = "australia-southeast1-a"
}

variable "database_password" {
  description = "Cloud SQL database password"
  type        = string
  sensitive   = true
}