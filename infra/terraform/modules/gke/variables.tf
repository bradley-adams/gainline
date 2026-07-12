variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
}

variable "zone" {
  description = "GCP zone"
  type        = string
}

variable "cluster_name" {
  description = "GKE cluster name"
  type        = string
}

variable "network" {
  description = "VPC network name"
  type        = string
  default     = "default"
}

variable "subnetwork" {
  description = "VPC subnetwork name"
  type        = string
  default     = "default"
}

variable "node_pool_name" {
  description = "Name of the node pool"
  type        = string
  default     = "default-pool"
}

variable "machine_type" {
  description = "Machine type for nodes"
  type        = string
  default     = "e2-small"
}

variable "node_count" {
  description = "Initial number of nodes"
  type        = number
  default     = 1
}

variable "min_node_count" {
  description = "Minimum number of nodes for autoscaling"
  type        = number
  default     = 1
}

variable "max_node_count" {
  description = "Maximum number of nodes for autoscaling"
  type        = number
  default     = 2
}

variable "disk_size_gb" {
  description = "Disk size per node in GB"
  type        = number
  default     = 20
}

variable "environment" {
  description = "Environment name (dev or prod)"
  type        = string
}