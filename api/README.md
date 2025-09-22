# gainline api

## Getting Started

### Run the App Locally:
```bash
go run .
```

### Curl health:
```
curl http://localhost:8080/health
```

### Open Swagger UI:
```
http://localhost:8080/swagger/index.html#/
```

### Database Migrations:
```
docker run --rm -v $(pwd)/migrations:/migrations \
  --network gainline_default migrate/migrate \
  -path=/migrations -database "postgres://gainline:gainline@gainline-db:5432/gainline?sslmode=disable" up
```

### SQL Code Generation
```
sqlc generate
```

### Swagger Documentation
```
swag init -g http/handlers/http.go
```
```
swag fmt
```

### DB Mock Generation
```
mockgen -destination=/home/bradley/Personal/gainline/api/db/db_handler/mock/db.go -package=mock_db github.com/bradley-adams/gainline/db/db_handler DB,Queries
```

## Todo:
- Game handler tests.
- Check the errors returned by all services.
- Swagger spec default dont delete actual competition seeded data. Maybe seed a second one to delete
- Create/Update competition swagger default violates unique constraint.
- Expand DB handler testing (integration with mock DB).
- Error response standardisation (consistent shape for errors).
