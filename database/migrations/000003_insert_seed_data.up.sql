-- Insert competition
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    '44dd315c-1abc-43aa-9843-642f920190d1', 'Bunnings NPC', now(), now()
);

-- Insert season
INSERT INTO seasons (id, competition_id, start_date, end_date, rounds, created_at, updated_at)
VALUES (
    '9300778f-cce0-4efe-af6c-e399d8170315',
    '44dd315c-1abc-43aa-9843-642f920190d1',
    TIMESTAMPTZ '2025-07-31 19:10:00+12',
    TIMESTAMPTZ '2025-10-25 21:10:00+12',
    12,
    now(),
    now()
);

-- Insert teams
INSERT INTO teams (id, name, abbreviation, location, created_at, updated_at)
VALUES
('013952a5-87e1-4d26-a312-09b2aff54241', 'Auckland', 'AUK', 'Auckland', now(), now()),
('7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'Waikato', 'WAI', 'Hamilton', now(), now()),
('636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'Manawatū', 'MAN', 'Palmerston North', now(), now()),
('e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'North Harbour', 'NHB', 'Albany', now(), now()),
('ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'Wellington', 'WEL', 'Wellington', now(), now()),
('f192a9ce-dce2-4389-8491-1a193ac7699e', 'Canterbury', 'CAN', 'Christchurch', now(), now()),
('15c76909-f78a-4d89-bc19-7c80265e1e08', 'Southland', 'STL', 'Invercargill', now(), now()),
('a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'Otago', 'OTA', 'Dunedin', now(), now()),
('bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 'Taranaki', 'TAR', 'New Plymouth', now(), now()),
('7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'Northland', 'NOR', 'Whangārei', now(), now()),
('b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'Bay Of Plenty', 'BOP', 'Tauranga', now(), now()),
('19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'Tasman', 'TAS', 'Nelson', now(), now()),
('dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'Hawke''s Bay', 'HKB', 'Napier', now(), now()),
('6b5c3642-c026-4e89-81f7-024c40638f9a', 'Counties Manukau', 'CMK', 'Pukekohe', now(), now());

-- Insert games
INSERT INTO games (
    id, season_id, round, date,
    home_team_id, away_team_id,
    home_score, away_score,
    status, created_at, updated_at
)
VALUES
-- Round 1
('4019a7f3-7741-4d8f-b3e0-1c7f3a0a1a01', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-07-31 19:10+13',
 '013952a5-87e1-4d26-a312-09b2aff54241', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', NULL, NULL, 'scheduled', now(), now()),

('9b00d3e3-299d-4ee2-b3a4-b71f68eb1d28', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-01 19:10+13',
 '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', NULL, NULL, 'scheduled', now(), now()),

('309ccabd-6e4c-4d39-9709-cbcb2ae1a3b7', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 14:05+13',
 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'f192a9ce-dce2-4389-8491-1a193ac7699e', NULL, NULL, 'scheduled', now(), now()),

('30e4793a-4b5e-4fef-986c-f0e1e35754d1', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 16:35+13',
 '15c76909-f78a-4d89-bc19-7c80265e1e08', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', NULL, NULL, 'scheduled', now(), now()),

('ac86ff36-6bfc-4c59-a497-cb6d3cb6cdb9', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 19:10+13',
 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', NULL, NULL, 'scheduled', now(), now()),

('d5295ee1-d1d4-4f71-a73a-6f0a9b981cd9', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-03 14:05+13',
 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', NULL, NULL, 'scheduled', now(), now()),

('7f1a5d72-2417-4d30-b272-b19e215bbbf4', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-03 16:35+13',
 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '6b5c3642-c026-4e89-81f7-024c40638f9a', NULL, NULL, 'scheduled', now(), now()),

-- Round 2
('3e2a0e9f-4eec-464a-8124-cb50d30517e9', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-08 19:10+13',
 '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '15c76909-f78a-4d89-bc19-7c80265e1e08', NULL, NULL, 'scheduled', now(), now()),

('cb40de0d-d396-4e6d-9e1a-87ee8aa16b82', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 14:05+13',
 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', NULL, NULL, 'scheduled', now(), now()),

('e97e292b-684e-43b4-bb1a-b6c75a2f3775', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 14:05+13',
 '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', NULL, NULL, 'scheduled', now(), now()),

('6fc6cfad-7f30-45db-9a6d-5cf63ecf5576', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 16:35+13',
 '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', NULL, NULL, 'scheduled', now(), now()),

('cefe2b54-bca9-42b2-9ae6-e98bc8fffb32', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 19:10+13',
 'f192a9ce-dce2-4389-8491-1a193ac7699e', '013952a5-87e1-4d26-a312-09b2aff54241', NULL, NULL, 'scheduled', now(), now()),

('6c0207ea-f960-4db7-9df5-25a35714d4f6', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-10 14:05+13',
 '6b5c3642-c026-4e89-81f7-024c40638f9a', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', NULL, NULL, 'scheduled', now(), now()),

('7df99e59-896e-4a1f-bc61-f49d6e1d8970', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-10 16:35+13',
 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', NULL, NULL, 'scheduled', now(), now());
