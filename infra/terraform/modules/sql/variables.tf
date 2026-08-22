variable "project_id" {
  description = "GCP project ID"
  type        = string
}

variable "region" {
  description = "GCP region"
  type        = string
}

variable "instance_name" {
  description = "Cloud SQL instance name"
  type        = string
}

variable "database_version" {
  description = "Postgres version"
  type        = string
  default     = "POSTGRES_16"
}

variable "tier" {
  description = "Cloud SQL machine tier"
  type        = string
  default     = "db-f1-micro"
}

variable "database_name" {
  description = "Database name"
  type        = string
  default     = "gainline"
}

variable "database_user" {
  description = "Database user"
  type        = string
  default     = "gainline"
}

variable "database_password" {
  description = "Database password"
  type        = string
  sensitive   = true
}

variable "environment" {
  description = "Environment name (dev or prod)"
  type        = string
}

variable "private_vpc_connection" {
  description = "VPC peering connection — ensures networking is set up before Cloud SQL"
  type        = string
}

variable "network_self_link" {
  description = "Self link of the VPC network to attach Cloud SQL's private IP to. Defaults to the shared 'default' network for backward compatibility."
  type        = string
  default     = null
}