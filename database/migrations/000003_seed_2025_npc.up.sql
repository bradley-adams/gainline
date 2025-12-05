-- Insert NPC competition
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    '44dd315c-1abc-43aa-9843-642f920190d1', 'National Provincial Championship', now(), now()
);

-- Insert 2025 season
INSERT INTO seasons (id, competition_id, start_date, end_date, rounds, created_at, updated_at)
VALUES (
    '9300778f-cce0-4efe-af6c-e399d8170315',
    '44dd315c-1abc-43aa-9843-642f920190d1',
    TIMESTAMPTZ '2025-07-31 21:10+12',
    TIMESTAMPTZ '2025-10-25 23:10+12',
    10,
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
('9bc1d9b6-73e5-479d-9152-7bbeb73b0411', '9300778f-cce0-4efe-af6c-e399d8170315', '6b5c3642-c026-4e89-81f7-024c40638f9a', now(), now(), NULL), -- Counties Manukau
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

-- Create 10 round stages for 2025 season
INSERT INTO stages (id, season_id, name, stage_type, order_index)
VALUES
('eab15533-dea6-4a3d-8a95-d38e4fba2d5a', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 1', 'regular', 1),
('847bcffb-30f1-42c6-be61-2807c3032566', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 2', 'regular', 2),
('559272f0-94a2-4909-b5ac-b09a26b8f8b8', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 3', 'regular', 3),
('18b9655c-1eee-41f9-999b-1254abad43d6', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 4', 'regular', 4),
('5bc23a00-153d-42f6-be4c-558210ea541b', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 5', 'regular', 5),
('85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 6', 'regular', 6),
('e53435fc-717e-4fce-9b85-c99c606ae3ce', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 7', 'regular', 7),
('f6a305fb-c036-404a-9a48-0bf9f4c2ac39', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 8', 'regular', 8),
('cd95a74b-d643-4913-b391-77b933edbd8f', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 9', 'regular', 9),
('eff747cf-af9a-44d4-a348-efa5e1c099a3', '9300778f-cce0-4efe-af6c-e399d8170315', 'Round 10', 'regular', 10),
('6ad67327-958b-491d-bcd2-f9cc50f0330a', '9300778f-cce0-4efe-af6c-e399d8170315', 'Quarterfinal', 'finals', 11),
('a1ce69e8-838e-49f3-b222-0820750dc292', '9300778f-cce0-4efe-af6c-e399d8170315', 'Semifinal', 'finals', 12),
('c8b2bbf9-dc43-4834-a8a3-722090d857c2', '9300778f-cce0-4efe-af6c-e399d8170315', 'Final', 'finals', 13);


-- Insert 2025 games
INSERT INTO games (
    id, season_id, stage_id, date,
    home_team_id, away_team_id,
    home_score, away_score,
    status, created_at, updated_at
)
VALUES
-- Round 1
(
    '940ef044-fbaf-4e87-8026-6d0e33eab20f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-07-31 21:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    35, 
    36, 
    'finished', 
    now(), 
    now()
),
(
    'ce82f220-dcae-4b0a-a9ea-1e1b220a098d', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-01 21:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    38, 
    25, 
    'finished', 
    now(), 
    now()
),
(
    '1e966827-7f69-4274-8ee8-323d6276a47d', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-02 16:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    15, 
    33, 
    'finished', 
    now(), 
    now()
),
(
    '3d34c43c-ba21-4cea-8964-a4c1df89faf9', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-02 18:35+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    15, 
    24, 
    'finished', 
    now(), 
    now()
),
(
    'b26dc5a0-f061-4d23-bf19-cdc03b1d0787', 
    '9300778f-cce0-4efe-af6c-e399d8170315',
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-02 21:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    23, 
    3, 
    'finished', 
    now(), 
    now()
),
(
    'ca38f91b-17d2-4c01-9b8f-69b1a0923371', 
    '9300778f-cce0-4efe-af6c-e399d8170315',
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-03 16:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    37, 
    7, 
    'finished', 
    now(), 
    now()
),
(
    'd656fe21-abd3-4f1b-8e5b-7e560c519877', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eab15533-dea6-4a3d-8a95-d38e4fba2d5a', 
    '2025-08-03 18:35+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    54, 
    14, 
    'finished', 
    now(), 
    now()
),

-- Round 2
(
    '9ba6c29c-63c7-468e-9c85-ff3a86bd2b8f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-08 21:10+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    22, 
    17, 
    'finished', 
    now(), 
    now()
),
(
    'd7b67afd-8e02-4880-9cda-f5db2ead6260', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-09 16:05+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    24, 
    35, 
    'finished', 
    now(), 
    now()
),
(
    '73d3e3a0-175b-44c9-bc19-3080d2da1b8a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-09 16:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    24, 
    46, 
    'finished', 
    now(), 
    now()
),
(
    'f796cef5-619e-4ec8-bfd1-2be4b748386b', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-09 18:35+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    49, 
    17, 
    'finished', 
    now(), 
    now()
),
(
    '8dd237ae-623b-4fcf-8b55-aad2f2a76b1a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-09 21:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    19, 
    15, 
    'finished', 
    now(), 
    now()
),
(
    '0b8c5035-389e-4fb3-9e39-e52260784285', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-10 16:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    22, 
    39, 
    'finished', 
    now(), 
    now()
),
(
    'cd4c92df-e18a-451d-af6b-2f120919d504', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '847bcffb-30f1-42c6-be61-2807c3032566', 
    '2025-08-10 18:35+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    21, 
    27, 
    'finished', 
    now(), 
    now()
),

-- Round 3
(
    '15b6b6e2-2783-4d5d-b6f9-b1e8fac5298e', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-15 21:10+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    29, 
    22, 
    'finished', 
    now(), 
    now()
),
(
    '35ce17ae-a222-48e4-bb00-2b63a4712322', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-16 16:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    14, 
    28, 
    'finished', 
    now(), 
    now()
),
(
    '8fc60e11-2133-4224-91ed-aa1264199154', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-16 16:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    7, 
    21, 
    'finished', 
    now(), 
    now()
),
(
    '50c9c723-ec51-47b9-a4e7-db8130377d15', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-16 18:35+12',
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    8, 
    50, 
    'finished', 
    now(), 
    now()
),
(
    '8a9c75c8-e74d-4590-a09f-8b70951a03bd', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-16 21:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    36, 
    22, 
    'finished', 
    now(), 
    now()
),
(
    '0419caa0-5cbd-49b1-aad0-4c8c505e11e5', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-17 16:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    41, 
    46, 
    'finished', 
    now(), 
    now()
),
(
    'ad43c079-9502-4033-9eb2-c7e030f0cac0', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '559272f0-94a2-4909-b5ac-b09a26b8f8b8', 
    '2025-08-17 18:35+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    27, 
    26, 
    'finished', 
    now(), 
    now()
),

-- Round 4
(
    'f7119686-8a76-4d67-ac91-53e121bbefdd', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-21 21:10+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    25, 
    30, 
    'finished', 
    now(), 
    now()
),
(
    'd5b1c2f1-fa32-4ff7-8f6d-c53f2159e7e6', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-22 21:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    7, 
    26, 
    'finished', 
    now(), 
    now()
),
(
    '9a2f7e0a-03d4-4cfa-bdca-5c7230fae60e', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-23 16:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    22, 
    23, 
    'finished', 
    now(), 
    now()
),
(
    '47570ea7-e50f-4fd8-aca6-946ad07e58fb', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-23 18:35+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    19, 
    43, 
    'finished', 
    now(), 
    now()
),
(
    'fee66798-9326-4cc5-a0aa-b3df992b3b8a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-23 21:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    24, 
    43, 
    'finished', 
    now(), 
    now()
),
(
    '337db57b-55bd-4cfa-ab5a-55b38b6ddebf', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-24 15:35+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    33, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    '54cf1c34-c79d-48e2-92e6-ad195eea49e1', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '18b9655c-1eee-41f9-999b-1254abad43d6', 
    '2025-08-24 18:35+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    38, 
    28, 
    'finished', 
    now(), 
    now()),

-- Round 5
(
    '87751cc2-1833-41f3-884d-f038335b39e5', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-29 21:10+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    7, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    '5c38b91e-bf62-4db7-b8bf-e9f44cb9f148', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-30 16:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    53, 
    14, 
    'finished', 
    now(), 
    now()
),
(
    '284af7c8-2d67-4f1a-bed6-ce8f01cefb51', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-30 16:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    22, 
    43, 
    'finished', 
    now(), 
    now()
),
(
    'e88c6ca3-3d47-4318-85eb-637838b42ebf', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-30 18:35+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    27, 
    22, 
    'finished', 
    now(), 
    now()
),
(
    '1093f483-8147-42f8-ac4c-f26468320f69', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-30 21:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    31, 
    27, 
    'finished', 
    now(), 
    now()
),
(
    '120edadc-472a-4798-a999-c1c5f7a9ff69', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-31 16:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    36, 
    17, 
    'finished', 
    now(), 
    now()
),
(
    'e5c0a0f9-6dad-46b6-9f51-eaf6e5598e07', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '5bc23a00-153d-42f6-be4c-558210ea541b', 
    '2025-08-31 18:35+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    10, 
    25, 
    'finished', 
    now(), 
    now()
),

-- Round 6
(
    '140d54f6-1533-4265-a48a-c0d087ed16fc', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-04 21:10+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    22, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    '79878f7e-769e-441e-8710-b379df20f730', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-05 21:10+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    29, 
    10, 
    'finished', 
    now(), 
    now()
),
(
    'cc8b240a-9508-4217-a167-43c575ab23ef', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-06 16:05:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    36, 
    26, 
    'finished', 
    now(), 
    now()
),
(
    '262c7e9b-74cb-4df8-a390-f4e3af8a38cb', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-06 16:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    45, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    '58a14de3-f7a8-4b37-ae26-7af352a9e50f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-06 16:05+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    14, 
    54, 
    'finished', 
    now(), 
    now()
),
(
    '1185ee20-d153-48c9-94f8-732e5fd5cd35', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-07 16:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    17, 
    43, 
    'finished', 
    now(), 
    now()
),
(
    'a2ee305f-0cb7-40f2-8473-171b54666cd6', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '85767fc4-5bc0-4e3f-87cb-1e05dc1981f6', 
    '2025-09-08 16:35+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    21, 
    29, 
    'finished', 
    now(), 
    now()
),

-- Round 7
(
    '62ff6277-bc6c-4e2d-b311-a4f57b5bcd13', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-11 21:10+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    43, 
    26, 
    'finished', 
    now(), 
    now()
),
(
    'a862d5ab-1f2a-4277-8f91-589347f5fe1a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-12 21:10+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    29, 
    28, 
    'finished', 
    now(), 
    now()
),
(
    '75fde3c9-dd5c-43ea-8d78-3858dc76f704', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-13 16:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    21, 
    24, 
    'finished', 
    now(), 
    now()
),
(
    '406a2c92-df7d-4f63-ac64-52c1ee4d012b', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-13 18:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    28, 
    26, 
    'finished', 
    now(), 
    now()
),
(
    '7e2a9a37-1ebf-4780-9dd7-38e55b387b5f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-14 16:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    52, 
    29, 
    'finished', 
    now(), 
    now()
),
(
    'ff6be7f9-36b7-4f85-9915-70d301d668dd', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-14 16:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    31, 
    25, 
    'finished', 
    now(), 
    now()
),
(
    '68101d89-673e-49ff-94ec-b33619dbeea7', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'e53435fc-717e-4fce-9b85-c99c606ae3ce', 
    '2025-09-14 18:35+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    10, 
    64, 
    'finished', 
    now(), 
    now()
),

-- Round 8
(
    'ad35b833-7a47-4ef2-b43c-8a6cdf4387de', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-19 21:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    38, 
    24, 
    'finished', 
    now(), 
    now()
),
(
    'eaccd3aa-dadb-47a7-ba69-59bb21b8d64c', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-20 16:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    75, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    'eab628cd-fc15-4663-87e7-4ca6b28ff781', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-20 16:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    49, 
    28, 
    'finished', 
    now(), 
    now()
),
(
    '1997364a-66b9-4acc-90ee-2acfb16a2358', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-20 18:40+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    36, 
    38, 
    'finished', 
    now(), 
    now()
),
(
    'b5eda992-689e-4ae5-90bf-188764fb81a8', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-20 21:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    24, 
    29, 
    'finished', 
    now(), 
    now()
),
(
    'cd39fb99-ce04-4457-9c5e-e34021d096d8', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-21 16:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    21, 
    22, 
    'finished', 
    now(), 
    now()
),
(
    '2ee1d0aa-c9ba-4a5e-89f0-88cb94055cd2', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-21 18:20+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    19, 
    55, 
    'finished', 
    now(), 
    now()
),

-- Round 9
(
    'c5732e49-8de5-4ca0-be3e-a7fa9511d294', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'cd95a74b-d643-4913-b391-77b933edbd8f', 
    '2025-09-25 21:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    45, 
    28, 
    'finished', 
    now(), 
    now()
),
(
    'adea27e0-2648-4256-9cb7-6bb56e2de5bc', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-26 21:10+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    38, 
    55, 
    'finished', 
    now(), 
    now()
),
(
    '477d8f8c-c903-4ec6-845a-3b8c2e3c163f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-27 15:10+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    41, 
    26, 
    'finished', 
    now(), 
    now()
),
(
    'dafe49bc-c348-48ae-be19-68cb2294b904', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'cd95a74b-d643-4913-b391-77b933edbd8f', 
    '2025-09-27 16:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    39, 
    20, 
    'finished', 
    now(), 
    now()
),
(
    '46894429-5989-4a79-b19f-7f7060f87a84', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-09-28 15:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    48, 
    24, 
    'finished', 
    now(), 
    now()
),
(
    '3e53b537-0743-4ea2-a584-c178fce29a0a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'cd95a74b-d643-4913-b391-77b933edbd8f', 
    '2025-09-28 15:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    41, 
    5, 
    'finished', 
    now(), 
    now()
),
(
    '9c849f9d-ce78-4bfd-8aeb-55b2b1914500', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'cd95a74b-d643-4913-b391-77b933edbd8f', 
    '2025-09-28 17:35+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    19, 
    19, 
    'finished', 
    now(), 
    now()
),

-- Round 10
(
    '3b4f97e0-ede4-40a2-a59b-254496d693a7', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-03 21:10+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7', 
    '15c76909-f78a-4d89-bc19-7c80265e1e08', 
    15, 
    14, 
    'finished', 
    now(), 
    now()
),
(
    '762b84d8-17fe-43cf-81af-ac2bd56faa8e', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-04 15:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    26, 
    33, 
    'finished', 
    now(), 
    now()
),
(
    '98f6e1bf-be79-4952-bf8a-4924c3872bcb', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-04 15:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241', 
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    17, 
    51, 
    'finished', 
    now(), 
    now()
),
(
    '0d60bdcc-effc-458c-b0bf-1a87a089684a', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-04 17:35+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c', 
    25, 
    19, 
    'finished', 
    now(), 
    now()
),
(
    'c03d6ccb-ef22-490a-90d0-77e710bb65d0', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-04 21:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    10, 
    38, 
    'finished', 
    now(), 
    now()
),
(
    'bce365a1-b6f2-454a-842a-2c5c51b79a15', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'eff747cf-af9a-44d4-a348-efa5e1c099a3', 
    '2025-10-05 15:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    41, 
    49, 
    'finished', 
    now(), 
    now()
),
(
    '1dfcadb0-de4d-4be0-a4e7-443aa766a526', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'f6a305fb-c036-404a-9a48-0bf9f4c2ac39', 
    '2025-10-05 17:35+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    34, 
    14, 
    'finished', 
    now(), 
    now()
),

-- Quarterfinals
(
    '2aac44fc-b7dd-48ff-8211-569543ac078f', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '6ad67327-958b-491d-bcd2-f9cc50f0330a', 
    '2025-10-10 21:10+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97', 
    44,
    41,
    'finished', 
    now(), 
    now()
),
(
    '0757454f-3586-440e-bfc4-1250b647cbcf', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '6ad67327-958b-491d-bcd2-f9cc50f0330a', 
    '2025-10-11 17:10+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f', 
    27,
    7,
    'finished', 
    now(), 
    now()
),
(
    '44fbc8b7-f653-43f2-ac05-2c6baabe5ef1', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '6ad67327-958b-491d-bcd2-f9cc50f0330a', 
    '2025-10-11 21:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f', 
    26,
    12,
    'finished', 
    now(), 
    now()
),
(
    '213dfa4c-96ce-4bcb-8a17-e07da9e4a3b5', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    '6ad67327-958b-491d-bcd2-f9cc50f0330a', 
    '2025-10-12 15:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    '6b5c3642-c026-4e89-81f7-024c40638f9a', 
    23,
    15,
    'finished', 
    now(), 
    now()
),

-- Semifinals
(
    '3222cfcc-428c-4feb-989b-5a79d0228433', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'a1ce69e8-838e-49f3-b222-0820750dc292', 
    '2025-10-17 20:15+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a', 
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532', 
    41,
    17,
    'finished', 
    now(), 
    now()
),
(
    'bd86578c-442d-4487-ae67-4e9cf31632f1', 
    '9300778f-cce0-4efe-af6c-e399d8170315', 
    'a1ce69e8-838e-49f3-b222-0820750dc292', 
    '2025-10-18 20:15+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e', 
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1', 
    43,
    19,
    'finished', 
    now(), 
    now()
);
