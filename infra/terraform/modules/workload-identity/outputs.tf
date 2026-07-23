output "workload_identity_provider" {
  description = "Workload Identity Provider resource name — used in GitHub Actions workflow"
  value       = google_iam_workload_identity_pool_provider.github.name
}

output "service_account_email" {
  description = "Service account email — used in GitHub Actions workflow"
  value       = google_service_account.github_actions.email
}