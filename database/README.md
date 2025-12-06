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

## Todo:

- Split competition name and season sponsor.
- Rounds and Finals seperations.
- Shield challenges (side comps)
- Bonus points
