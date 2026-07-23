variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "github_org" {
  description = "GitHub organisation or username"
  type        = string
}

variable "github_repo" {
  description = "GitHub repository name"
  type        = string
}

variable "environment" {
  description = "Environment name (dev or prod)"
  type        = string
}