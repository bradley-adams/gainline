terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.0"
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
  repository_id = "gainline-dev"
  environment   = "dev"
}

module "gke" {
  source         = "../../modules/gke"
  project_id     = var.project_id
  region         = var.region
  zone           = var.zone
  cluster_name   = "gainline-dev"
  machine_type   = "e2-small"
  node_count     = 1
  min_node_count = 1
  max_node_count = 2
  disk_size_gb   = 20
  environment    = "dev"
}

module "redis" {
  source         = "../../modules/redis"
  project_id     = var.project_id
  region         = var.region
  instance_name  = "gainline-dev"
  memory_size_gb = 1
  environment    = "dev"
}

module "workload_identity" {
  source      = "../../modules/workload-identity"
  project_id  = var.project_id
  github_org  = "bradley-adams"
  github_repo = "gainline"
  environment = "dev"
}

module "networking" {
  source     = "../../modules/networking"
  project_id = var.project_id
}

module "db_secret" {
  source      = "../../modules/db-secret"
  project_id  = var.project_id
  environment = "dev"
  k8s_namespace = "gainline-dev"
}

module "sql" {
  source                 = "../../modules/sql"
  project_id             = var.project_id
  region                 = var.region
  instance_name          = "gainline-dev"
  tier                   = "db-f1-micro"
  database_password      = module.db_secret.password
  environment            = "dev"
  private_vpc_connection = module.networking.private_vpc_connection
}