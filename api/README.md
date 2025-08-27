# gainline api

go run .

http://localhost:8080

curl http://localhost:8080/health

http://localhost:8080/swagger/index.html#/

```
docker run --rm -v $(pwd)/migrations:/migrations \
  --network gainline_default migrate/migrate \
  -path=/migrations -database "postgres://gainline:gainline@gainline-db:5432/gainline?sslmode=disable" up
```

```
sqlc generate
```

```
swag init -g http/handlers/http.go
```

```
swag fmt
```

```
mockgen -destination=/home/bradley/Personal/gainline/api/db/db_handler/mock/db.go -package=mock_db github.com/bradley-adams/gainline/db/db_handler DB,Queries
```

## Todo:

- Middleware to enforce ownership (e.g., only owners/admins can edit/delete).
- Request validation layer (schema validation for all endpoints).
- Expand DB handler testing (unit + integration with mock DB).
- Error response standardisation (consistent shape for errors).
