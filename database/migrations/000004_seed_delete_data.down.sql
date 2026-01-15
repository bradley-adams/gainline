-- Cleanup seeded DELETE test season-team links
DELETE FROM season_teams
WHERE season_id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Cleanup seeded DELETE test season
DELETE FROM seasons
WHERE id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Cleanup seeded DELETE test teams
DELETE FROM teams
WHERE id IN (
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0001',
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0002'
);

-- Cleanup seeded DELETE test competition
DELETE FROM competitions
WHERE id = 'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd';

