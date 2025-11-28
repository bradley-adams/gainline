-- Insert NPC competition
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    '44dd315c-1abc-43aa-9843-642f920190d1', 'National Provincial Championship', now(), now()
);

-- Insert 2025 season
INSERT INTO seasons (id, competition_id, start_date, end_date, rounds, sponsor, created_at, updated_at)
VALUES (
    '9300778f-cce0-4efe-af6c-e399d8170315',
    '44dd315c-1abc-43aa-9843-642f920190d1',
    TIMESTAMPTZ '2025-07-31 21:10+12',
    TIMESTAMPTZ '2025-10-25 23:10+12',
    12,
    'Bunnings',
    now(),
    now()
);

-- Insert 2025 teams
INSERT INTO teams (id, name, abbreviation, location, created_at, updated_at)
VALUES
('013952a5-87e1-4d26-a312-09b2aff54241', 'Auckland', 'AUK', 'Auckland', now(), now()),
('b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'Bay Of Plenty', 'BOP', 'Tauranga', now(), now()),
('f192a9ce-dce2-4389-8491-1a193ac7699e', 'Canterbury', 'CAN', 'Christchurch', now(), now()),
('6b5c3642-c026-4e89-81f7-024c40638f9a', 'Counties Manukau', 'CMK', 'Pukekohe', now(), now()),
('dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'Hawke''s Bay', 'HKB', 'Napier', now(), now()),
('636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'Manawatū', 'MAN', 'Palmerston North', now(), now()),
('e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'North Harbour', 'NHB', 'Albany', now(), now()),
('7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'Northland', 'NOR', 'Whangārei', now(), now()),
('a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'Otago', 'OTA', 'Dunedin', now(), now()),
('15c76909-f78a-4d89-bc19-7c80265e1e08', 'Southland', 'STL', 'Invercargill', now(), now()),
('bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 'Taranaki', 'TAR', 'New Plymouth', now(), now()),
('19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'Tasman', 'TAS', 'Nelson', now(), now()),
('7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'Waikato', 'WAI', 'Hamilton', now(), now()),
('ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'Wellington', 'WEL', 'Wellington', now(), now());

-- Connect teams to 2025 season
INSERT INTO season_teams (id, season_id, team_id, created_at, updated_at, deleted_at)
VALUES
('8fc1cc1b-7de1-464a-b3c5-7db1806f3661', '9300778f-cce0-4efe-af6c-e399d8170315', '013952a5-87e1-4d26-a312-09b2aff54241', now(), now(), NULL), -- Auckland
('6e45c89b-ef37-4593-b6c1-1fe4a054deb7', '9300778f-cce0-4efe-af6c-e399d8170315', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', now(), now(), NULL), -- Bay Of Plenty
('e587e373-986d-4be6-a894-70bf62367455', '9300778f-cce0-4efe-af6c-e399d8170315', 'f192a9ce-dce2-4389-8491-1a193ac7699e', now(), now(), NULL), -- Canterbury
('ae444444-4444-4444-4444-444444444444', '9300778f-cce0-4efe-af6c-e399d8170315', '6b5c3642-c026-4e89-81f7-024c40638f9a', now(), now(), NULL), -- Counties Manukau
('9a6d021d-1daa-4320-a6ea-4b62dd8ac5c5', '9300778f-cce0-4efe-af6c-e399d8170315', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', now(), now(), NULL), -- Hawke's Bay
('84d4696e-67ae-4836-a3a9-336c7bb4c4fe', '9300778f-cce0-4efe-af6c-e399d8170315', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', now(), now(), NULL), -- Manawatū
('ad73b809-eb3b-48af-9d50-e26951f52702', '9300778f-cce0-4efe-af6c-e399d8170315', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', now(), now(), NULL), -- North Harbour
('575c1e41-04f4-47cc-9179-19cd0e0227a9', '9300778f-cce0-4efe-af6c-e399d8170315', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', now(), now(), NULL), -- Northland
('cbcf4c8a-4172-4a38-8ea0-19dd2282df1f', '9300778f-cce0-4efe-af6c-e399d8170315', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', now(), now(), NULL), -- Otago
('a3d1dd91-2bfc-4819-9e52-7e5cb7ea2fdd', '9300778f-cce0-4efe-af6c-e399d8170315', '15c76909-f78a-4d89-bc19-7c80265e1e08', now(), now(), NULL), -- Southland
('3071577c-0130-4b75-b778-0f843981aff0', '9300778f-cce0-4efe-af6c-e399d8170315', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', now(), now(), NULL), -- Taranaki
('fdc63492-8ff1-4610-92f1-82f629524404', '9300778f-cce0-4efe-af6c-e399d8170315', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', now(), now(), NULL), -- Tasman
('f414e700-8b43-4870-812d-783a5b9ddb2d', '9300778f-cce0-4efe-af6c-e399d8170315', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', now(), now(), NULL), -- Waikato
('869a0f34-f15d-45ef-b13f-eb551050a849', '9300778f-cce0-4efe-af6c-e399d8170315', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', now(), now(), NULL); -- Wellington

-- Insert 2025 games
INSERT INTO games (
    id, season_id, round, date,
    home_team_id, away_team_id,
    home_score, away_score,
    status, created_at, updated_at
)
VALUES
-- Round 1
('940ef044-fbaf-4e87-8026-6d0e33eab20f', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-07-31 21:10+12',
 '013952a5-87e1-4d26-a312-09b2aff54241', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 35, 36, 'finished', now(), now()),

('ce82f220-dcae-4b0a-a9ea-1e1b220a098d', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-01 21:10+12',
 '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 38, 25, 'finished', now(), now()),

('1e966827-7f69-4274-8ee8-323d6276a47d', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 16:05+12',
 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 15, 33, 'finished', now(), now()),

('3d34c43c-ba21-4cea-8964-a4c1df89faf9', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 18:35+12',
 '15c76909-f78a-4d89-bc19-7c80265e1e08', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 15, 24, 'finished', now(), now()),

('b26dc5a0-f061-4d23-bf19-cdc03b1d0787', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-02 21:10+12',
 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 23, 3, 'finished', now(), now()),

('ca38f91b-17d2-4c01-9b8f-69b1a0923371', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-03 16:05+12',
 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 37, 7, 'finished', now(), now()),

('d656fe21-abd3-4f1b-8e5b-7e560c519877', '9300778f-cce0-4efe-af6c-e399d8170315', 1, '2025-08-03 18:35+12',
 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', '6b5c3642-c026-4e89-81f7-024c40638f9a', 54, 14, 'finished', now(), now()),

-- Round 2
('9ba6c29c-63c7-468e-9c85-ff3a86bd2b8f', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-08 21:10+12',
 '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '15c76909-f78a-4d89-bc19-7c80265e1e08', 22, 17, 'finished', now(), now()),
 
('d7b67afd-8e02-4880-9cda-f5db2ead6260', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 16:05+12',
 '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 24, 35, 'finished', now(), now()),

('73d3e3a0-175b-44c9-bc19-3080d2da1b8a', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 16:05+12',
 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 24, 46, 'finished', now(), now()),

('f796cef5-619e-4ec8-bfd1-2be4b748386b', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 18:35+12',
 '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 49, 17, 'finished', now(), now()),

('8dd237ae-623b-4fcf-8b55-aad2f2a76b1a', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-09 21:10+12',
 'f192a9ce-dce2-4389-8491-1a193ac7699e', '013952a5-87e1-4d26-a312-09b2aff54241', 19, 15, 'finished', now(), now()),

('0b8c5035-389e-4fb3-9e39-e52260784285', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-10 16:05+12',
 '6b5c3642-c026-4e89-81f7-024c40638f9a', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 22, 39, 'finished', now(), now()),

('cd4c92df-e18a-451d-af6b-2f120919d504', '9300778f-cce0-4efe-af6c-e399d8170315', 2, '2025-08-10 18:35+12',
 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 21, 27, 'finished', now(), now()),

-- Round 3
('15b6b6e2-2783-4d5d-b6f9-b1e8fac5298e', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-15 21:10+12',
 '15c76909-f78a-4d89-bc19-7c80265e1e08', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 29, 22, 'finished', now(), now()),

('35ce17ae-a222-48e4-bb00-2b63a4712322', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-16 16:05+12',
 '7e5abf68-8358-4c20-b6a4-f64ef264c13c', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 14, 28, 'finished', now(), now()),
 
('8fc60e11-2133-4224-91ed-aa1264199154', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-16 16:05+12',
 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 7, 21, 'finished', now(), now()),

('50c9c723-ec51-47b9-a4e7-db8130377d15', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-16 18:35+12',
 '013952a5-87e1-4d26-a312-09b2aff54241', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 8, 50, 'finished', now(), now()),

('8a9c75c8-e74d-4590-a09f-8b70951a03bd', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-16 21:10+12',
 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 36, 22, 'finished', now(), now()),

('0419caa0-5cbd-49b1-aad0-4c8c505e11e5', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-17 16:05+12',
 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 41, 46, 'finished', now(), now()),

('ad43c079-9502-4033-9eb2-c7e030f0cac0', '9300778f-cce0-4efe-af6c-e399d8170315', 3, '2025-08-17 18:35+12',
 '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', '6b5c3642-c026-4e89-81f7-024c40638f9a', 27, 26, 'finished', now(), now()),

-- Round 4
('f7119686-8a76-4d67-ac91-53e121bbefdd', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-21 21:10+12',
 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', '15c76909-f78a-4d89-bc19-7c80265e1e08', 25, 30, 'finished', now(), now()),

('d5b1c2f1-fa32-4ff7-8f6d-c53f2159e7e6', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-22 21:10+12',
 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 'f192a9ce-dce2-4389-8491-1a193ac7699e', 7, 26, 'finished', now(), now()),

('9a2f7e0a-03d4-4cfa-bdca-5c7230fae60e', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-23 16:05+12',
 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 22, 23, 'finished', now(), now()),

('47570ea7-e50f-4fd8-aca6-946ad07e58fb', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-23 18:35+12',
 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 19, 43, 'finished', now(), now()),

('fee66798-9326-4cc5-a0aa-b3df992b3b8a', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-23 21:10+12',
 '013952a5-87e1-4d26-a312-09b2aff54241', '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 24, 43, 'finished', now(), now()),

('337db57b-55bd-4cfa-ab5a-55b38b6ddebf', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-24 15:35+12',
 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 33, 19, 'finished', now(), now()),

('54cf1c34-c79d-48e2-92e6-ad195eea49e1', '9300778f-cce0-4efe-af6c-e399d8170315', 4, '2025-08-24 18:35+12',
 '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 38, 28, 'finished', now(), now()),

-- Round 5
('87751cc2-1833-41f3-884d-f038335b39e5', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 7, 19, 'finished', now(), now()),

('5c38b91e-bf62-4db7-b8bf-e9f44cb9f148', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 'f192a9ce-dce2-4389-8491-1a193ac7699e', '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 53, 14, 'finished', now(), now()),

('284af7c8-2d67-4f1a-bed6-ce8f01cefb51', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 '6b5c3642-c026-4e89-81f7-024c40638f9a', 'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 22, 43, 'finished', now(), now()),

('e88c6ca3-3d47-4318-85eb-637838b42ebf', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 27, 22, 'finished', now(), now()),

('1093f483-8147-42f8-ac4c-f26468320f69', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 31, 27, 'finished', now(), now()),

('120edadc-472a-4798-a999-c1c5f7a9ff69', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', '013952a5-87e1-4d26-a312-09b2aff54241', 36, 17, 'finished', now(), now()),

('e5c0a0f9-6dad-46b6-9f51-eaf6e5598e07', '9300778f-cce0-4efe-af6c-e399d8170315', 5, '2025-08-08 21:10+12',
 '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', '15c76909-f78a-4d89-bc19-7c80265e1e08', 10, 25, 'finished', now(), now());