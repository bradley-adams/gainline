output "registry_url" {
  description = "Artifact Registry URL"
  value       = module.registry.repository_url
}

output "cluster_name" {
  description = "GKE cluster name"
  value       = module.gke.cluster_name
}

output "cluster_endpoint" {
  description = "GKE cluster endpoint"
  value       = module.gke.cluster_endpoint
  sensitive   = true
}

output "sql_connection_name" {
  description = "Cloud SQL connection name"
  value       = module.sql.connection_name
}

output "sql_public_ip" {
  description = "Cloud SQL public IP"
  value       = module.sql.public_ip
}

output "redis_host" {
  description = "Redis host"
  value       = module.redis.host
}

output "redis_port" {
  description = "Redis port"
  value       = module.redis.port
}

output "workload_identity_provider" {
  description = "Workload Identity Provider for GitHub Actions"
  value       = module.workload_identity.workload_identity_provider
}

output "service_account_email" {
  description = "Service account email for GitHub Actions"
  value       = module.workload_identity.service_account_email
}