# Gainline DB Migrations

Database migrations and tooling for managing the Gainline Postgres database.

## Prerequisites

Install `just`:

```bash
sudo apt install just
```

Verify:

```bash
just --version
```

## View All Commands

```bash
just
```

## Release Tagging

Releases are tagged **from `main` only** using an interactive Just recipe.

The process will:

- Show the current tag
- Prompt for the next version
- Prevent duplicate tags
- Ensure you are on `main`
- Ensure the working tree is clean
- Create and push an annotated tag

### Tag a release

```bash
just release-tag
```

Example flow:

```
📌 Current tag: v0.1.0

Enter new tag (vX.Y.Z): v0.2.0
Confirm release tag 'v0.2.0'? (y/N): y
Tagging release v0.2.0…
```

## Creating a Migration

Create a new sequential SQL migration:

```bash
just migrate-create name=add_new_table
```

This will generate a pair of migration files (up/down) under:

```
migrations/
```

## Running Migrations

All migration commands are wrapped by the `justfile` and run against the database
defined by `DB_URL` in the `justfile` (local Postgres by default).

### Apply migrations (up)

Run all pending migrations:

```bash
just migrate-up
```

### Roll back migrations (down)

Roll back the most recent migration:

```bash
just migrate-down
```

### Roll back all migrations (DANGER)

Roll back all applied migrations. This will reset the database schema.

```bash
just migrate-down-all
```

### Clear dirty migration (DANGER)

If a migration fails and leaves the database in a dirty state, force the migration
version manually:

```bash
just migrate-force 1
```

## Todo

- Adding season sponsor column. (Remove from competition name).
- Handle shield challenge games (Side competitions/trophies).
- Points tables and bonus points.
