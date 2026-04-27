-- National Provincial Championship 2024

-- Competition (reuses same competition ID as 2025 NPC)
-- NOTE: If you have already run the 2025 NPC seed, the competition already exists.
-- Only insert if it doesn't already exist, or use a separate competition record.
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES ('44dd315c-1abc-43aa-9843-642f920190d1', 'National Provincial Championship', now(), now())
ON CONFLICT (id) DO NOTHING;

INSERT INTO seasons (id, competition_id, start_date, end_date, created_at, updated_at)
VALUES (
    'aa000000-0000-4000-8000-000000000001',
    '44dd315c-1abc-43aa-9843-642f920190d1',
    TIMESTAMPTZ '2024-08-09 17:35+12',
    TIMESTAMPTZ '2024-10-26 15:05+13',
    now(), now()
);

INSERT INTO teams (id, name, abbreviation, location, created_at, updated_at)
VALUES
('013952a5-87e1-4d26-a312-09b2aff54241', 'Auckland',         'AUK', 'Auckland',        now(), now()),
('b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'Bay Of Plenty',    'BOP', 'Tauranga',         now(), now()),
('f192a9ce-dce2-4389-8491-1a193ac7699e', 'Canterbury',       'CAN', 'Christchurch',     now(), now()),
('6b5c3642-c026-4e89-81f7-024c40638f9a', 'Counties Manukau', 'CMK', 'Pukekohe',         now(), now()),
('dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'Hawke''s Bay',     'HKB', 'Napier',           now(), now()),
('636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'Manawatū',         'MAN', 'Palmerston North', now(), now()),
('e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'North Harbour',    'NHB', 'Albany',           now(), now()),
('7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'Northland',        'NOR', 'Whangārei',        now(), now()),
('a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'Otago',            'OTA', 'Dunedin',          now(), now()),
('15c76909-f78a-4d89-bc19-7c80265e1e08', 'Southland',        'STL', 'Invercargill',     now(), now()),
('bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 'Taranaki',         'TAR', 'New Plymouth',     now(), now()),
('19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'Tasman',           'TAS', 'Nelson',           now(), now()),
('7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'Waikato',          'WAI', 'Hamilton',         now(), now()),
('ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'Wellington',       'WEL', 'Wellington',       now(), now())
ON CONFLICT (id) DO NOTHING;

INSERT INTO season_teams (id, season_id, team_id, created_at, updated_at, deleted_at)
VALUES
('bb000000-0000-4000-8000-000000000001', 'aa000000-0000-4000-8000-000000000001', '013952a5-87e1-4d26-a312-09b2aff54241', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000002', 'aa000000-0000-4000-8000-000000000001', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000003', 'aa000000-0000-4000-8000-000000000001', 'f192a9ce-dce2-4389-8491-1a193ac7699e', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000004', 'aa000000-0000-4000-8000-000000000001', '6b5c3642-c026-4e89-81f7-024c40638f9a', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000005', 'aa000000-0000-4000-8000-000000000001', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000006', 'aa000000-0000-4000-8000-000000000001', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000007', 'aa000000-0000-4000-8000-000000000001', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000008', 'aa000000-0000-4000-8000-000000000001', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000009', 'aa000000-0000-4000-8000-000000000001', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000010', 'aa000000-0000-4000-8000-000000000001', '15c76909-f78a-4d89-bc19-7c80265e1e08', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000011', 'aa000000-0000-4000-8000-000000000001', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000012', 'aa000000-0000-4000-8000-000000000001', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000013', 'aa000000-0000-4000-8000-000000000001', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', now(), now(), NULL),
('bb000000-0000-4000-8000-000000000014', 'aa000000-0000-4000-8000-000000000001', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', now(), now(), NULL);

INSERT INTO stages (id, season_id, name, stage_type, order_index, created_at, updated_at)
VALUES
('npc24s001', 'aa000000-0000-4000-8000-000000000001', 'Round 1',       'regular', 1,  now(), now()),
('npc24s002', 'aa000000-0000-4000-8000-000000000001', 'Round 2',       'regular', 2,  now(), now()),
('npc24s003', 'aa000000-0000-4000-8000-000000000001', 'Round 3',       'regular', 3,  now(), now()),
('npc24s004', 'aa000000-0000-4000-8000-000000000001', 'Round 4',       'regular', 4,  now(), now()),
('npc24s005', 'aa000000-0000-4000-8000-000000000001', 'Round 5',       'regular', 5,  now(), now()),
('npc24s006', 'aa000000-0000-4000-8000-000000000001', 'Round 6',       'regular', 6,  now(), now()),
('npc24s007', 'aa000000-0000-4000-8000-000000000001', 'Round 7',       'regular', 7,  now(), now()),
('npc24s008', 'aa000000-0000-4000-8000-000000000001', 'Round 8',       'regular', 8,  now(), now()),
('npc24s009', 'aa000000-0000-4000-8000-000000000001', 'Round 9',       'regular', 9,  now(), now()),
('npc24s010', 'aa000000-0000-4000-8000-000000000001', 'Quarterfinals', 'finals',  10, now(), now()),
('npc24s011', 'aa000000-0000-4000-8000-000000000001', 'Semifinals',    'finals',  11, now(), now()),
('npc24s012', 'aa000000-0000-4000-8000-000000000001', 'Final',         'finals',  12, now(), now());

INSERT INTO games (id, season_id, stage_id, date, home_team_id, away_team_id, home_score, away_score, status, created_at, updated_at)
VALUES
-- Round 1
('npc24g001', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-09 17:35+12', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '6b5c3642-c026-4e89-81f7-024c40638f9a', 31, 15, 'finished', now(), now()),
('npc24g002', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-09 19:35+12', '013952a5-87e1-4d26-a312-09b2aff54241', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 21, 29, 'finished', now(), now()),
('npc24g003', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-10 14:05+12', 'f192a9ce-dce2-4389-8491-1a193ac7699e', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 34, 21, 'finished', now(), now()),
('npc24g004', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-10 14:05+12', '15c76909-f78a-4d89-bc19-7c80265e1e08', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 22, 13, 'finished', now(), now()),
('npc24g005', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-10 16:35+12', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 21, 36, 'finished', now(), now()),
('npc24g006', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-11 14:05+12', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 32, 41, 'finished', now(), now()),
('npc24g007', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-11 16:35+12', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 21, 54, 'finished', now(), now()),
('npc24g008', 'aa000000-0000-4000-8000-000000000001', 'npc24s001', TIMESTAMPTZ '2024-08-14 19:05+12', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '6b5c3642-c026-4e89-81f7-024c40638f9a', 44, 31, 'finished', now(), now()),
-- Round 2
('npc24g009', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-16 19:05+12', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', '013952a5-87e1-4d26-a312-09b2aff54241', 27, 25, 'finished', now(), now()),
('npc24g010', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-17 14:05+12', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 35, 18, 'finished', now(), now()),
('npc24g011', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-17 14:05+12', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 22, 7,  'finished', now(), now()),
('npc24g012', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-17 16:35+12', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '15c76909-f78a-4d89-bc19-7c80265e1e08', 31, 17, 'finished', now(), now()),
('npc24g013', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-18 14:05+12', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 26, 19, 'finished', now(), now()),
('npc24g014', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-18 14:05+12', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 24, 20, 'finished', now(), now()),
('npc24g015', 'aa000000-0000-4000-8000-000000000001', 'npc24s002', TIMESTAMPTZ '2024-08-18 16:35+12', '6b5c3642-c026-4e89-81f7-024c40638f9a', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 20, 26, 'finished', now(), now()),
-- Round 3
('npc24g016', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-23 19:05+12', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 55, 30, 'finished', now(), now()),
('npc24g017', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-24 14:05+12', '6b5c3642-c026-4e89-81f7-024c40638f9a', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 3,  48, 'finished', now(), now()),
('npc24g018', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-24 16:35+12', '013952a5-87e1-4d26-a312-09b2aff54241', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 21, 27, 'finished', now(), now()),
('npc24g019', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-24 19:05+12', '15c76909-f78a-4d89-bc19-7c80265e1e08', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 24, 39, 'finished', now(), now()),
('npc24g020', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-25 14:05+12', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 31, 26, 'finished', now(), now()),
('npc24g021', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-25 14:05+12', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 39, 31, 'finished', now(), now()),
('npc24g022', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-25 16:35+12', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 43, 29, 'finished', now(), now()),
('npc24g023', 'aa000000-0000-4000-8000-000000000001', 'npc24s003', TIMESTAMPTZ '2024-08-28 19:05+12', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 21, 27, 'finished', now(), now()),
-- Round 4
('npc24g024', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-08-30 19:05+12', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '15c76909-f78a-4d89-bc19-7c80265e1e08', 26, 31, 'finished', now(), now()),
('npc24g025', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-08-31 14:05+12', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', '6b5c3642-c026-4e89-81f7-024c40638f9a', 33, 36, 'finished', now(), now()),
('npc24g026', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-08-31 14:05+12', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 22, 18, 'finished', now(), now()),
('npc24g027', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-08-31 16:35+12', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', '013952a5-87e1-4d26-a312-09b2aff54241', 39, 21, 'finished', now(), now()),
('npc24g028', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-08-31 19:05+12', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 34, 15, 'finished', now(), now()),
('npc24g029', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-09-01 14:05+12', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 21, 46, 'finished', now(), now()),
('npc24g030', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-09-01 16:35+12', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 26, 38, 'finished', now(), now()),
('npc24g031', 'aa000000-0000-4000-8000-000000000001', 'npc24s004', TIMESTAMPTZ '2024-09-04 19:05+12', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 34, 19, 'finished', now(), now()),
-- Round 5
('npc24g032', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-06 19:05+12', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 68, 14, 'finished', now(), now()),
('npc24g033', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-07 14:05+12', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', '15c76909-f78a-4d89-bc19-7c80265e1e08', 36, 12, 'finished', now(), now()),
('npc24g034', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-07 14:05+12', '013952a5-87e1-4d26-a312-09b2aff54241', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 36, 32, 'finished', now(), now()),
('npc24g035', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-07 16:35+12', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 16, 34, 'finished', now(), now()),
('npc24g036', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-07 19:05+12', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 24, 25, 'finished', now(), now()),
('npc24g037', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-08 14:05+12', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 25, 19, 'finished', now(), now()),
('npc24g038', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-08 16:35+12', '6b5c3642-c026-4e89-81f7-024c40638f9a', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 25, 14, 'finished', now(), now()),
('npc24g039', 'aa000000-0000-4000-8000-000000000001', 'npc24s005', TIMESTAMPTZ '2024-09-11 19:05+12', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 28, 32, 'finished', now(), now()),
-- Round 6
('npc24g040', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-13 19:05+12', '15c76909-f78a-4d89-bc19-7c80265e1e08', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 29, 41, 'finished', now(), now()),
('npc24g041', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-14 14:05+12', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 33, 20, 'finished', now(), now()),
('npc24g042', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-14 14:05+12', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 58, 19, 'finished', now(), now()),
('npc24g043', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-14 16:35+12', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 50, 5,  'finished', now(), now()),
('npc24g044', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-15 14:05+12', '6b5c3642-c026-4e89-81f7-024c40638f9a', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 45, 17, 'finished', now(), now()),
('npc24g045', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-15 14:05+12', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 28, 15, 'finished', now(), now()),
('npc24g046', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-15 16:35+12', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '013952a5-87e1-4d26-a312-09b2aff54241', 17, 24, 'finished', now(), now()),
('npc24g047', 'aa000000-0000-4000-8000-000000000001', 'npc24s006', TIMESTAMPTZ '2024-09-18 19:05+12', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', '15c76909-f78a-4d89-bc19-7c80265e1e08', 26, 21, 'finished', now(), now()),
-- Round 7
('npc24g048', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-20 19:05+12', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 19, 63, 'finished', now(), now()),
('npc24g049', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-21 13:05+12', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 47, 24, 'finished', now(), now()),
('npc24g050', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-21 14:05+12', 'f192a9ce-dce2-4389-8491-1a193ac7699e', '6b5c3642-c026-4e89-81f7-024c40638f9a', 36, 28, 'finished', now(), now()),
('npc24g051', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-21 14:05+12', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 30, 25, 'finished', now(), now()),
('npc24g052', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-22 13:05+12', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 25, 27, 'finished', now(), now()),
('npc24g053', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-22 14:05+12', '013952a5-87e1-4d26-a312-09b2aff54241', '15c76909-f78a-4d89-bc19-7c80265e1e08', 27, 19, 'finished', now(), now()),
('npc24g054', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-22 14:05+12', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 10, 28, 'finished', now(), now()),
('npc24g055', 'aa000000-0000-4000-8000-000000000001', 'npc24s007', TIMESTAMPTZ '2024-09-25 19:05+12', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 31, 25, 'finished', now(), now()),
-- Round 8
('npc24g056', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-27 19:05+12', '6b5c3642-c026-4e89-81f7-024c40638f9a', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 51, 12, 'finished', now(), now()),
('npc24g057', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-28 14:05+12', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 53, 13, 'finished', now(), now()),
('npc24g058', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-28 14:05+12', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '013952a5-87e1-4d26-a312-09b2aff54241', 36, 35, 'finished', now(), now()),
('npc24g059', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-28 16:35+12', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 47, 31, 'finished', now(), now()),
('npc24g060', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-29 14:05+12', '15c76909-f78a-4d89-bc19-7c80265e1e08', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 14, 38, 'finished', now(), now()),
('npc24g061', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-29 14:05+12', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 65, 19, 'finished', now(), now()),
('npc24g062', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-09-29 16:35+12', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 33, 31, 'finished', now(), now()),
('npc24g063', 'aa000000-0000-4000-8000-000000000001', 'npc24s008', TIMESTAMPTZ '2024-10-02 19:05+13', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', '013952a5-87e1-4d26-a312-09b2aff54241', 31, 17, 'finished', now(), now()),
-- Round 9
('npc24g064', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-04 19:05+13', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 28, 31, 'finished', now(), now()),
('npc24g065', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-05 14:05+13', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', '6b5c3642-c026-4e89-81f7-024c40638f9a', 26, 45, 'finished', now(), now()),
('npc24g066', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-05 16:35+13', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 46, 28, 'finished', now(), now()),
('npc24g067', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-05 16:35+13', '15c76909-f78a-4d89-bc19-7c80265e1e08', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 59, 35, 'finished', now(), now()),
('npc24g068', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-05 19:05+13', 'f192a9ce-dce2-4389-8491-1a193ac7699e', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 36, 19, 'finished', now(), now()),
('npc24g069', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-06 14:05+13', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 29, 42, 'finished', now(), now()),
('npc24g070', 'aa000000-0000-4000-8000-000000000001', 'npc24s009', TIMESTAMPTZ '2024-10-06 16:35+13', '013952a5-87e1-4d26-a312-09b2aff54241', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 24, 26, 'finished', now(), now()),
-- Quarterfinals
('npc24g071', 'aa000000-0000-4000-8000-000000000001', 'npc24s010', TIMESTAMPTZ '2024-10-11 19:05+13', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', '6b5c3642-c026-4e89-81f7-024c40638f9a', 29, 14, 'finished', now(), now()),
('npc24g072', 'aa000000-0000-4000-8000-000000000001', 'npc24s010', TIMESTAMPTZ '2024-10-12 14:05+13', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 19, 17, 'finished', now(), now()),
('npc24g073', 'aa000000-0000-4000-8000-000000000001', 'npc24s010', TIMESTAMPTZ '2024-10-12 19:05+13', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 14, 15, 'finished', now(), now()),
('npc24g074', 'aa000000-0000-4000-8000-000000000001', 'npc24s010', TIMESTAMPTZ '2024-10-13 14:05+13', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 14, 62, 'finished', now(), now()),
-- Semifinals
('npc24g075', 'aa000000-0000-4000-8000-000000000001', 'npc24s011', TIMESTAMPTZ '2024-10-19 16:10+13', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 32, 20, 'finished', now(), now()),
('npc24g076', 'aa000000-0000-4000-8000-000000000001', 'npc24s011', TIMESTAMPTZ '2024-10-19 19:10+13', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 29, 24, 'finished', now(), now()),
-- Final
('npc24g077', 'aa000000-0000-4000-8000-000000000001', 'npc24s012', TIMESTAMPTZ '2024-10-26 15:05+13', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 23, 20, 'finished', now(), now());