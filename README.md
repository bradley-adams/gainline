# gainline

A rugby season scheduler.

## Build and Run:

```
docker compose build
docker compose up -d
```

### Up just the DB & migrate:

```
docker compose up -d gainline-db
docker compose run --rm gainline-migrate
```

### Stop and Remove old data

```
docker stop gainline-db
docker rm gainline-db
docker volume rm gainline-data
```

## Todo:

- Run tests in PRs
- Implement search across core entities (teams, games, seasons, competitions).
- Add pagination to API responses and frontend tables.
- Write a full migration seeder (seasons, teams, games, scores).
- Authentication via Auth0 (basic login flow for admins).
- Implement user setup (user + admin roles, role enforcement).
- User setup. Just user and admin roles to start.
- Decide on and implement deployment maybe?
- Metrics and logs maybe?
