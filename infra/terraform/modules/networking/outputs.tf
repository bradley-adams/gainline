output "private_vpc_connection" {
  description = "VPC peering connection name — Cloud SQL depends on this existing"
  value       = google_service_networking_connection.private_vpc_connection.network
}

output "network_self_link" {
  description = "Self link of the VPC network in use (shared default, or this environment's dedicated network)"
  value       = local.network_self_link
}

output "network_name" {
  description = "Name of the VPC network in use"
  value       = local.network_name
}

output "subnetwork_name" {
  description = "Name of the subnetwork in use"
  value       = local.subnetwork_name
}