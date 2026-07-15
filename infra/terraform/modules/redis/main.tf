resource "google_redis_instance" "redis" {
  project        = var.project_id
  name           = var.instance_name
  tier           = "BASIC"
  memory_size_gb = var.memory_size_gb
  region         = var.region

  labels = {
    environment = var.environment
  }
}