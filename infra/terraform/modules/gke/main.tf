resource "google_container_cluster" "cluster" {
  name     = var.cluster_name
  location = var.zone

  # Remove default node pool and manage it separately
  # so we can configure autoscaling properly
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = var.network
  subnetwork = var.subnetwork

  deletion_protection = false
}

resource "google_container_node_pool" "nodes" {
  name       = var.node_pool_name
  location   = var.zone
  cluster    = google_container_cluster.cluster.name
  node_count = var.node_count

  autoscaling {
    min_node_count = var.min_node_count
    max_node_count = var.max_node_count
  }

  node_config {
    machine_type = var.machine_type
    disk_size_gb = var.disk_size_gb
    disk_type    = "pd-standard"

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]

    labels = {
      environment = var.environment
    }
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }
}