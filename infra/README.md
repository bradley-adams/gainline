# Gainline Infra

Terraform + Helm for the Gainline platform (GKE, Cloud SQL, Memorystore, Artifact Registry).

## Getting Started

### Spin up an environment

Need `terraform` and `gcloud` authenticated against the target GCP project.

State lives in GCS (`gs://gainline-tfstate`, one prefix per environment) instead of
local state files. If that bucket doesn't exist yet:

```bash
gsutil mb -p gainline-503521 -l australia-southeast1 gs://gainline-tfstate
gsutil versioning set on gs://gainline-tfstate
```

```bash
cd terraform/environments/dev
terraform init
terraform plan -var-file=terraform.tfvars
terraform apply -var-file=terraform.tfvars
```

Same thing for `prod`, just `cd` into that folder instead. Each environment is its own
root module and provisions a registry, a GKE cluster, Redis, Cloud SQL, networking, and
the workload identity binding GitHub Actions uses to deploy.

### Get cluster creds

```bash
gcloud container clusters get-credentials $(terraform output -raw cluster_name) \
  --region <region> --project <project_id>
```

### Deploy a service

```bash
helm upgrade --install api ./helm/api -f ./helm/api/values-dev.yaml
```

Same for `gamestate` and `ui`, and swap `values-dev.yaml` -> `values-prod.yaml` for prod.
The `api` chart runs DB migrations itself as a pre-install/pre-upgrade hook, so no need
to run them separately.

### cert-manager

Needs applying once, after cert-manager itself is on the cluster:

```bash
kubectl apply -f helm/cert-manager/cluster-issuer.yaml
```

It's just the shared ClusterIssuer the UI ingress uses to get its Let's Encrypt cert.

## Todo:

- dev networking still rides on the shared default network, prod got its own VPC, bring dev in line.
- No CI running terraform plan on PRs yet.
