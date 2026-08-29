# Generates the Cloud SQL password so it never has to be hand-typed into tfvars.
resource "random_password" "db" {
  length  = 32
  special = false # Cloud SQL / connection strings are picky about some special chars
}

resource "google_secret_manager_secret" "db_password" {
  project   = var.project_id
  secret_id = "${var.environment}-db-password"

  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "db_password" {
  secret      = google_secret_manager_secret.db_password.id
  secret_data = random_password.db.result
}

# GCP service account the API pod runs as (via Workload Identity), so it can
# read the secret above through the Secret Manager CSI driver.
resource "google_service_account" "api_workload" {
  project      = var.project_id
  account_id   = "api-${var.environment}"
  display_name = "API workload ${var.environment}"
}

resource "google_secret_manager_secret_iam_member" "api_secret_accessor" {
  project   = var.project_id
  secret_id = google_secret_manager_secret.db_password.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.api_workload.email}"
}

# Lets the Kubernetes service account (var.k8s_namespace/var.k8s_service_account)
# impersonate this GCP service account via GKE Workload Identity.
resource "google_service_account_iam_member" "workload_identity_binding" {
  service_account_id = google_service_account.api_workload.name
  role                = "roles/iam.workloadIdentityUser"
  member              = "serviceAccount:${var.project_id}.svc.id.goog[${var.k8s_namespace}/${var.k8s_service_account}]"
}