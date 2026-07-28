output "private_vpc_connection" {
  description = "VPC peering connection name — Cloud SQL depends on this existing"
  value       = google_service_networking_connection.private_vpc_connection.network
}