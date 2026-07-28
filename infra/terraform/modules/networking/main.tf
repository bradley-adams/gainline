# Reserve a private IP range for Google managed services (Cloud SQL etc.)
resource "google_compute_global_address" "private_ip_range" {
  project       = var.project_id
  name          = "google-managed-services"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = "projects/${var.project_id}/global/networks/${var.network}"
}

# Create the VPC peering connection to Google's managed services network
resource "google_service_networking_connection" "private_vpc_connection" {
  network                 = "projects/${var.project_id}/global/networks/${var.network}"
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_range.name]
}