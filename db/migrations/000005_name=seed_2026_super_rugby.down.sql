-- Remove seeded Super Rugby Pacific 2026 competition, season, teams, stages, and all round/finals games

-- Remove 2026 games
DELETE FROM games
WHERE season_id = 'ac2f106c-5837-4781-9f6a-8b5d5a6d217f';

-- Remove 2026 stages (regular + finals)
DELETE FROM stages
WHERE season_id = 'ac2f106c-5837-4781-9f6a-8b5d5a6d217f';

-- Remove 2026 season-team links
DELETE FROM season_teams
WHERE season_id = 'ac2f106c-5837-4781-9f6a-8b5d5a6d217f';

-- Remove 2026 season
DELETE FROM seasons
WHERE id = 'ac2f106c-5837-4781-9f6a-8b5d5a6d217f';

-- Remove 2026 teams
DELETE FROM teams
WHERE id IN (
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '17412805-230b-49b6-838b-dcc485034022',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0'
);

-- Remove Super Rugby Pacific competition
DELETE FROM competitions
WHERE id = 'b3f77b8d-25e5-4817-aed4-d023160cd7ed';