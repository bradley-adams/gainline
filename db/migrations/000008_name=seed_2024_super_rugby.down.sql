-- Remove seeded Super Rugby Pacific 2024 season, teams, stages, and all round/finals games

-- Remove 2024 games
DELETE FROM games
WHERE season_id = '6e07a7c9-5d3d-4806-a9b5-35978c3750ca';

-- Remove 2024 stages (regular + finals)
DELETE FROM stages
WHERE season_id = '6e07a7c9-5d3d-4806-a9b5-35978c3750ca';

-- Remove 2024 season-team links
DELETE FROM season_teams
WHERE season_id = '6e07a7c9-5d3d-4806-a9b5-35978c3750ca';

-- Remove 2024 season
DELETE FROM seasons
WHERE id = '6e07a7c9-5d3d-4806-a9b5-35978c3750ca';

-- Remove Melbourne Rebels (wound up after 2024, not owned by any other migration)
DELETE FROM teams
WHERE id = 'de073a60-5b56-4f72-9946-61cdbb5fa937';