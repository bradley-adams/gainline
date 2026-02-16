# gainline

A rugby season scheduler.

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

Stops services and removes the database volume.

```bash
make db-reset
```

## Cleaning Everything

Remove all containers, volumes, and orphans:

```bash
make clean
```

## Todo

- Run tests in PRs
- Implement search across core entities (teams, games, seasons, competitions).
- Add pagination to API responses and frontend tables.
- Write a full migration seeder (seasons, teams, games, scores).
- Authentication via Auth0 (basic login flow for admins).
- Implement user setup (user + admin roles, role enforcement).
- User setup. Just user and admin roles to start.
- Decide on and implement deployment maybe.
- Metrics and logs somewhere.
