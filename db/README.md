# Gainline DB Migrations

Database migrations and tooling for managing the Gainline Postgres database.

## Release Tagging

Releases are tagged **from `main` only** using an interactive Make target.

The process will:

- Show the current tag
- Prompt for the next version
- Prevent duplicate tags
- Ensure you are on `main`
- Ensure the working tree is clean
- Create and push an annotated tag

### Tag a release

```bash
make release-tag
```

Example flow:

```
📌 Current tag: v0.1.0

Enter new tag (vX.Y.Z): v0.2.0
Confirm release tag 'v0.2.0'? (y/N): y
Tagging release v0.2.0…

```

## Creating a Migration

Create a new sequential SQL migration using the Makefile:

```bash
make migrate-create name=add_new_table
```

This will generate a pair of migration files (up/down) under:

```
migrations/
```

## Running Migrations

All migration commands are wrapped by the Makefile and run against the database
defined by `DB_URL` in the Makefile (local Postgres by default).

### Apply migrations (up)

Run all pending migrations:

```bash
make migrate-up
```

### Roll back migrations (down)

Roll back the most recent migration:

```bash
make migrate-down
```

### Roll back all migrations (DANGER)

Roll back all applied migrations. This will reset the database schema.

```bash
make migrate-down-all
```

### Clear dirty migration (DANGER)

If a migration fails and leaves the database in a dirty state, force the migration
version manually:

```bash
make migrate-force version=1
```

## Todo:

- Creating DB connection package.
- Adding season sponsor column. (Remove from competition name).
- Handle shield challenge games (Side competitions/trophies).
- Points tables and bonus points.
