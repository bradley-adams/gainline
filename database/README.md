# Gainline DB Migrations

## create a migration

```
migrate create -ext sql -dir database/migrations -seq a_new_migration
```

## run migration

Up:

```
migrate -path database/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" up
```

Down:

```
migrate -path database/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" down
```

Clear Dirty migration:

```
migrate -path database/migrations -database "postgres://gainline:gainline@localhost:5432/gainline?sslmode=disable" force 1
```

## Todo:

- Adding season sponsor column. (Remove from competition name).
- Handle shield challenge games (Side competitions/trophies).
- Points tables and bonus points.
