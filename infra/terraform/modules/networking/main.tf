# Dedicated VPC network + subnet for this environment (only when not using the shared default network)
resource "google_compute_network" "vpc" {
  count                   = var.use_default_network ? 0 : 1
  project                 = var.project_id
  name                    = "gainline-${var.environment}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  count         = var.use_default_network ? 0 : 1
  project       = var.project_id
  name          = "gainline-${var.environment}"
  region        = var.region
  network       = google_compute_network.vpc[0].id
  ip_cidr_range = var.subnet_cidr
}

locals {
  network_self_link = var.use_default_network ? "projects/${var.project_id}/global/networks/${var.network}" : google_compute_network.vpc[0].id
  network_name       = var.use_default_network ? var.network : google_compute_network.vpc[0].name
  subnetwork_name    = var.use_default_network ? var.network : google_compute_subnetwork.subnet[0].name
}

# Reserve a private IP range for Google managed services (Cloud SQL etc.)
resource "google_compute_global_address" "private_ip_range" {
  project       = var.project_id
  name          = var.use_default_network ? "google-managed-services" : "google-managed-services-${var.environment}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = local.network_self_link
}

# Create the VPC peering connection to Google's managed services network
resource "google_service_networking_connection" "private_vpc_connection" {
  network                 = local.network_self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_range.name]
}