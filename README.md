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

## Todo

- Implement search across core entities (teams, games, seasons, competitions).
- Authentication via Auth0 with basic roles (admin/user) and enforcement
- Deploy.
