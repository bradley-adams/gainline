-- Remove seeded test data

-- Remove test game
DELETE FROM games
WHERE id = '30f8181f-0a44-4ad7-a163-3ef2d29e504e';

-- Remove test stage
DELETE FROM stages
WHERE id = 'b138fab0-39fc-4eb5-9c10-44a918ed3952';

-- Remove test season-team links
DELETE FROM season_teams
WHERE season_id = 'fe04fe69-834f-42be-9821-04e53e8de26d';

-- Remove test season
DELETE FROM seasons
WHERE id = 'fe04fe69-834f-42be-9821-04e53e8de26d';

-- Remove test teams
DELETE FROM teams
WHERE id IN (
    'a8a3868f-7064-4769-a11c-1f8df795f6cc',
    'a1dd4e4c-1d72-48ea-850e-a9715f2940e3'
);

-- Remove test competition
DELETE FROM competitions
WHERE id = 'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd';