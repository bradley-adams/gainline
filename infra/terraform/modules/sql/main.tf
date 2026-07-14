resource "google_sql_database_instance" "postgres" {
  project          = var.project_id
  name             = var.instance_name
  database_version = var.database_version
  region           = var.region

  settings {
    tier = var.tier

    backup_configuration {
      enabled = false
    }

    ip_configuration {
      ipv4_enabled = true
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