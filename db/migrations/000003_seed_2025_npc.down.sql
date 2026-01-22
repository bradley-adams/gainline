-- Remove 2025 games
DELETE FROM games
WHERE season_id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Remove 2025 season-team links
DELETE FROM season_teams
WHERE season_id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Remove 2025 stages (regular + finals)
DELETE FROM stages
WHERE season_id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Remove 2025 teams
DELETE FROM teams
WHERE id IN (
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96'
);

-- Remove 2025 season
DELETE FROM seasons
WHERE id = '9300778f-cce0-4efe-af6c-e399d8170315';

-- Remove NPC competition
DELETE FROM competitions
WHERE id = '44dd315c-1abc-43aa-9843-642f920190d1';
