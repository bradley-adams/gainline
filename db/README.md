# Gainline DB Migrations

Database migrations and Open for a db connection.

## Tagging a Release

Tags must be created from `main` after changes have been merged.

```
git checkout main
git pull origin main
git tag v0.1.0
git push origin v0.1.0
```

## Creating a migration

```
migrate create -ext sql -dir db/migrations -seq a_new_migration
```

## Runnin migrations

Up:

```
migrate -path db/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" up
```

Down:

```
migrate -path db/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" down
```

Clear Dirty migration:

```
migrate -path db/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" force 1
```

## Todo:

- Creating DB connection package.
- Adding season sponsor column. (Remove from competition name).
- Handle shield challenge games (Side competitions/trophies).
- Points tables and bonus points.
