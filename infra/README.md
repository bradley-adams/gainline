# Gainline Infra

Terraform + Helm for the Gainline platform (GKE, Cloud SQL, Memorystore, Artifact Registry).

## Environments

- **Dev**: https://dev.34.87.247.234.nip.io
- **Prod**: https://prod.34.40.161.175.nip.io

Each has its own VPC, GKE cluster, Cloud SQL instance, and Redis instance — nothing shared.

If a cluster's load balancer ever gets recreated (usually from a cluster rebuild), the
ingress IP changes, which breaks three things until fixed manually:

- **Auth0** — add the new URL to Callback/Logout/Web Origins/CORS for that app.
- **`ui` workflow** — `api_url` is hardcoded per environment and baked into the UI build.
- **This README** — the links above.

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
root module and provisions a registry, a GKE cluster, Redis, Cloud SQL, networking, a
Secret Manager secret for the DB password, and the workload identity bindings GitHub
Actions and the API pod use.

Cloud SQL instances need `edition = "ENTERPRISE"` set explicitly (already in the `sql`
module). GCP defaults new instances to `ENTERPRISE_PLUS`, which doesn't support the
`db-f1-micro` tier we're on. Without it, instance creation fails outright.

### Get cluster creds

```bash
gcloud container clusters get-credentials $(terraform output -raw cluster_name) \
  --region <region> --project <project_id>
```

### Cluster add-ons (needed on every new/rebuilt cluster)

Not managed by Terraform. Installed directly via Helm/kubectl. A fresh or rebuilt
cluster has none of this, and nothing works until all three are in place.

**cert-manager**

```bash
helm repo add jetstack https://charts.jetstack.io
helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager --create-namespace \
  --set installCRDs=true
```

**ingress-nginx**

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

Then apply the shared ClusterIssuer:

```bash
kubectl apply -f helm/cert-manager/cluster-issuer.yaml
```

**Secrets Store CSI driver** (open-source, not GKE's managed add-on — that one can't
sync to a K8s Secret, only mount as a file, and the app needs `DB_PASSWORD` as an env var):

```bash
helm repo add secrets-store-csi-driver https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts
helm install csi-secrets-store secrets-store-csi-driver/secrets-store-csi-driver \
  --namespace kube-system \
  --set syncSecret.enabled=true

kubectl apply -f https://raw.githubusercontent.com/GoogleCloudPlatform/secrets-store-csi-driver-provider-gcp/main/deploy/provider-gcp-plugin.yaml
```

If `kubectl get pods -n kube-system | grep secrets-store` shows a pod stuck `Pending`
with `Insufficient memory`, the autoscaler won't fix it on its own — resize manually:

```bash
gcloud container clusters resize <cluster-name> \
  --node-pool default-pool --num-nodes 2 \
  --region <region> --project <project_id>
```

### Deploy a service

Normal deploys go through GitHub Actions. For manual/debug deploys:

```bash
helm upgrade --install api ./helm/api -f ./helm/api/values-dev.yaml \
  --set image.tag=<git-sha> --set migrate.image.tag=<git-sha> \
  --namespace gainline-dev --create-namespace
```

Same for `gamestate` and `ui`, swap `values-dev.yaml` -> `values-prod.yaml` for prod.
Always pass `image.tag` explicitly — it has no default, so a missing tag fails loudly
instead of silently deploying a broken `:latest` pod.

If a rebuild leaves `dbHost`/`redisHost` pointing at a stale IP (connection timeouts,
not auth errors), check them against `terraform output sql_private_ip` / `redis_host`.

## Rebuilding a cluster from scratch

If a cluster's ever destroyed and recreated (network change, accidental deletion,
whatever), the full sequence to get back to a working state is:

1. `terraform apply` the environment.
2. Get cluster creds.
3. Install cert-manager, ingress-nginx, apply the ClusterIssuer, install the Secrets
   Store CSI driver — see "Cluster add-ons" above.
4. Redeploy `api`, `gamestate`, `ui` via Helm (or trigger the GitHub Actions workflows).
5. Check the new ingress IP (`kubectl get ingress -n gainline-<env>`) and update Auth0 +
   the `ui` workflow's hardcoded `api_url` if it changed — see "Environments" above.

## Todo

- No CI running terraform plan on PRs yet.
