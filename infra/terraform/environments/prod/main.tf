terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

module "registry" {
  source        = "../../modules/registry"
  project_id    = var.project_id
  region        = var.region
  repository_id = "gainline-prod"
  environment   = "prod"
}

module "gke" {
  source         = "../../modules/gke"
  project_id     = var.project_id
  region         = var.region
  zone           = var.zone
  cluster_name   = "gainline-prod"
  machine_type   = "e2-small"
  node_count     = 1
  min_node_count = 1
  max_node_count = 3
  disk_size_gb   = 20
  environment    = "prod"
}

module "sql" {
  source            = "../../modules/sql"
  project_id        = var.project_id
  region            = var.region
  instance_name     = "gainline-prod"
  tier              = "db-f1-micro"
  database_password = var.database_password
  environment       = "prod"
}

module "redis" {
  source         = "../../modules/redis"
  project_id     = var.project_id
  region         = var.region
  instance_name  = "gainline-prod"
  memory_size_gb = 1
  environment    = "prod"
}

module "workload_identity" {
  source      = "../../modules/workload-identity"
  project_id  = var.project_id
  github_org  = "bradley-adams"
  github_repo = "gainline"
  environment = "prod"
}