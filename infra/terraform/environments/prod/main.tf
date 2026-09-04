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
  repository_id = "gainline-prod"
  environment   = "prod"
}

module "networking" {
  source               = "../../modules/networking"
  project_id           = var.project_id
  region               = var.region
  environment          = "prod"
  use_default_network  = false
  subnet_cidr          = "10.20.0.0/20"
}

module "gke" {
  source         = "../../modules/gke"
  project_id     = var.project_id
  region         = var.region
  zone           = var.zone
  cluster_name   = "gainline-prod"
  network        = module.networking.network_name
  subnetwork     = module.networking.subnetwork_name
  machine_type   = "e2-small"
  node_count     = 1
  min_node_count = 1
  max_node_count = 4
  disk_size_gb   = 20
  environment    = "prod"
}

module "redis" {
  source              = "../../modules/redis"
  project_id          = var.project_id
  region              = var.region
  instance_name       = "gainline-prod"
  memory_size_gb      = 1
  environment         = "prod"
  authorized_network  = module.networking.network_self_link
}

module "db_secret" {
  source      = "../../modules/db-secret"
  project_id  = var.project_id
  environment = "prod"
  k8s_namespace = "gainline-prod"
}

module "sql" {
  source                 = "../../modules/sql"
  project_id             = var.project_id
  region                 = var.region
  instance_name          = "gainline-prod"
  tier                   = "db-f1-micro"
  database_password      = module.db_secret.password
  environment            = "prod"
  private_vpc_connection = module.networking.private_vpc_connection
  network_self_link      = module.networking.network_self_link
}

module "workload_identity" {
  source      = "../../modules/workload-identity"
  project_id  = var.project_id
  github_org  = "bradley-adams"
  github_repo = "gainline"
  environment = "prod"
}