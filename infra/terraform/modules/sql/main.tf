resource "google_sql_database_instance" "postgres" {
  project          = var.project_id
  name             = var.instance_name
  database_version = var.database_version
  region           = var.region

  depends_on = [var.private_vpc_connection]

  settings {
    tier    = var.tier
    edition = "ENTERPRISE"

    backup_configuration {
      enabled = false
    }

    ip_configuration {
      ipv4_enabled    = false
      private_network = coalesce(var.network_self_link, "projects/${var.project_id}/global/networks/default")
    }
  }

  deletion_protection = false
}

resource "google_sql_database" "database" {
  project  = var.project_id
  name     = var.database_name
  instance = google_sql_database_instance.postgres.name
}

resource "google_sql_user" "user" {
  project  = var.project_id
  name     = var.database_user
  instance = google_sql_database_instance.postgres.name
  password = var.database_password
}