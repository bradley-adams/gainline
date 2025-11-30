# Gainline DB Migrations

## create a migration

```
migrate create -ext sql -dir database/migrations -seq a_new_migration
```

## Todo:

- Seed the full 2025 NPC with finals as additional rounds
- Split competition name and season sponsor.
- Rounds and Finals seperations.
- Shield challenges (side comps)
- Bonus points

### Finals series:

Adding this column to the season table so it can be selected when adding a game.

````finals_rounds JSONB
-- [
--   { "name": "Quarter-final", "order": 1 },
--   { "name": "Semi-final", "order": 2 },
--   { "name": "Final", "order": 3 }
-- ]```
````
