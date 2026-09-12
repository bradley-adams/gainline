# gainline

A rugby season scheduler and live game tracker.

## Using Make Commands

All common Docker and database operations are wrapped in the Makefile.

```bash
make
```

## Application Lifecycle

### Build containers

```bash
make build
```

### Start all services

```bash
make up
```

### Rebuild and start services

```bash
make rebuild
```

### Stop all services

```bash
make down
```

### Restart everything

```bash
make restart
```

### View logs

```bash
make logs
```

## Database Commands

### Start database only

```bash
make db-up
```

### Stop database only

```bash
make db-stop
```

### Run migrations

```bash
make migrate
```

### Reset database (⚠ Destructive)

Removes the database volume, recreates the database,
and runs migrations.

```bash
make db-reset
```

## Redis Commands

### Start Redis only

```bash
make redis-up
```

### Stop Redis only

```bash
make redis-stop
```

### Open a redis-cli session

```bash
make redis-cli
```

## Cleaning Everything

Remove all containers, volumes, and orphans:

```bash
make clean
```

## Services

| Service            | Port  | Description                       |
| ------------------ | ----- | --------------------------------- |
| gainline-ui        | 4200  | Angular frontend                  |
| gainline-api       | 8080  | REST API                          |
| gainline-gamestate | 50051 | gRPC live game state service      |
| gainline-db        | 5432  | PostgreSQL                        |
| gainline-redis     | 6379  | Redis pub/sub for live game state |

## Environments & Deployment

Runs on GKE, one dev and one prod environment, each with its own VPC, Cloud SQL, and Memorystore Redis instance. Infra is Terraform, apps deploy via Helm.

- Dev: https://dev.34.87.247.234.nip.io
- Prod: https://prod.34.40.161.175.nip.io

Deploys happen through GitHub Actions (.github/workflows/), one workflow per service (api, gamestate, ui), triggered on push to main (deploys to dev) or manually via workflow_dispatch (choose dev or prod).

Auth is handled by Auth0 in both environments, with separate Auth0 applications per environment.

For infra setup, cluster rebuild steps, and everything Terraform/Helm related, see the infra repo's README.

## Todo

- Implement search across core entities (teams, games, seasons, competitions).
