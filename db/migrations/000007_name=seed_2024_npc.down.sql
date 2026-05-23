-- Remove seeded National Provincial Championship 2024 season, stages, and all round/finals games

-- Remove 2024 games
DELETE FROM games
WHERE season_id = '0827c393-161a-41b4-badf-5678e5c8f153';

-- Remove 2024 stages (regular + finals)
DELETE FROM stages
WHERE season_id = '0827c393-161a-41b4-badf-5678e5c8f153';

-- Remove 2024 season-team links
DELETE FROM season_teams
WHERE season_id = '0827c393-161a-41b4-badf-5678e5c8f153';

-- Remove 2024 season
DELETE FROM seasons
WHERE id = '0827c393-161a-41b4-badf-5678e5c8f153';