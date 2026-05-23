-- Remove seeded National Provincial Championship 2026 season, stages, and all round/finals games

-- Remove 2026 games
DELETE FROM games
WHERE season_id = '142b57f9-c52e-4b1b-96f0-e1de6427afaf';

-- Remove 2026 stages (regular + finals)
DELETE FROM stages
WHERE season_id = '142b57f9-c52e-4b1b-96f0-e1de6427afaf';

-- Remove 2026 season-team links
DELETE FROM season_teams
WHERE season_id = '142b57f9-c52e-4b1b-96f0-e1de6427afaf';

-- Remove 2026 season
DELETE FROM seasons
WHERE id = '142b57f9-c52e-4b1b-96f0-e1de6427afaf';