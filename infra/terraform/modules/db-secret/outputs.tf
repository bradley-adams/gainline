output "password" {
  description = "Generated database password (feed into the sql module)"
  value       = random_password.db.result
  sensitive   = true
}

output "secret_id" {
  description = "Secret Manager secret ID, for referencing from the api Helm chart's SecretProviderClass"
  value       = google_secret_manager_secret.db_password.secret_id
}

output "api_service_account_email" {
  description = "GCP service account email the api Helm chart's ServiceAccount must be annotated with (iam.gke.io/gcp-service-account)"
  value       = google_service_account.api_workload.email
}