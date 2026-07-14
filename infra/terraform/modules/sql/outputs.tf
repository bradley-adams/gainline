output "instance_name" {
  description = "Cloud SQL instance name"
  value       = google_sql_database_instance.postgres.name
}

output "connection_name" {
  description = "Cloud SQL connection name used by Cloud SQL proxy"
  value       = google_sql_database_instance.postgres.connection_name
}

output "public_ip" {
  description = "Cloud SQL public IP address"
  value       = google_sql_database_instance.postgres.public_ip_address
}

output "database_name" {
  description = "Database name"
  value       = google_sql_database.database.name
}