variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "environment" {
  description = "Environment name (dev or prod)"
  type        = string
}

variable "k8s_namespace" {
  description = "Kubernetes namespace the api pod runs in"
  type        = string
  default     = "default"
}

variable "k8s_service_account" {
  description = "Kubernetes service account name the api pod runs as. Must match the ServiceAccount created by the api Helm chart."
  type        = string
  default     = "api"
}