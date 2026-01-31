# Gainline api

## Getting Started

### Run the App Locally:

```bash
go run .
```

### Curl health:

```bash
curl http://localhost:8080/health
```

### Open Swagger UI:

```bash
http://localhost:8080/swagger/index.html#/
```

### Get latest version of DB pkg

```bash
go get github.com/bradley-adams/gainline/db@latest
```

### Using a local version of the `db` package

To develop against a local copy of the `db` module:

In your `go.mod`, uncomment or add:

```bash
go replace github.com/bradley-adams/gainline/db => ../db
```

Then run:

```bash
go mod tidy
```

Remember to remove the replace directive before committing or releasing.

### Database Migrations:

```bash
docker run --rm -v $(pwd)/migrations:/migrations \
  --network gainline_default migrate/migrate \
  -path=/migrations -database "postgres://gainline:gainline@gainline-db:5432/gainline?sslmode=disable" up
```

### SQL Code Generation

```bash
sqlc generate
```

### Swagger Documentation

```bash
swag init -g http/handlers/http.go
```

```bash
swag fmt
```

### DB Mock Generation

```bash
mockgen -destination=/home/bradley/Personal/gainline/api/db/db_handler/mock/db.go -package=mock_db github.com/bradley-adams/gainline/db/db_handler DB,Queries
```

I think I have to update the mock command to this. Keeping the one above for now incase I am mucking up.

```bash
mockgen \
  -destination=./db/db_handler/mock/db.go \
  -package=mock_db \
  github.com/bradley-adams/gainline/db/db_handler DB,Queries
```

## Todo:

- Create/Update competition swagger default violates unique constraint.
- Expand DB handler testing (integration with mock DB).
- Error response standardisation (consistent shape for errors).
- Aggregates should be assembled using aggregate-shaped queries (Season)
