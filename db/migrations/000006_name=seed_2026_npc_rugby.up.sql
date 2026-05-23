-- Seed National Provincial Championship 2026 season, stages, and all round/finals games

-- Insert 2026 season
INSERT INTO seasons (
    id,
    competition_id,
    start_date,
    end_date,
    created_at,
    updated_at
)
VALUES (
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '44dd315c-1abc-43aa-9843-642f920190d1', 
    TIMESTAMPTZ '2026-07-30 19:05+12',
    TIMESTAMPTZ '2026-10-25 17:00+13',
    now(),
    now()
);

-- Connect teams to 2026 season
INSERT INTO season_teams (
    id,
    season_id,
    team_id,
    created_at,
    updated_at,
    deleted_at
)
VALUES
(
    'f756baa5-df39-4b50-8037-a7818ff78340',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    now(),
    now(),
    NULL -- Auckland
),
(
    'e80f4323-63d9-4b84-bf0b-f15581d61afd',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    now(),
    now(),
    NULL -- Bay Of Plenty
),
(
    '2f6065da-badc-4807-9463-011ed5aa3d77',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    now(),
    now(),
    NULL -- Canterbury
),
(
    'ee91c884-1af1-43d3-81bd-1e7190292d7a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    now(),
    now(),
    NULL -- Counties Manukau
),
(
    'fc322daf-1ba7-4002-8229-b91eb2023a59',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    now(),
    now(),
    NULL -- Hawke's Bay
),
(
    '11d86c71-5f71-47fb-8188-d6261130bcbc',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    now(),
    now(),
    NULL -- Manawatū
),
(
    '43730085-7371-4391-9df1-52d502930a02',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    now(),
    now(),
    NULL -- North Harbour
),
(
    'c45b696a-7bbc-4bf3-a690-96dfb65467df',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    now(),
    now(),
    NULL -- Northland
),
(
    '3cdf694f-d772-4370-830c-3652b6fb7249',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    now(),
    now(),
    NULL -- Otago
),
(
    '69dc84a7-f5aa-4549-82e3-37b8f43a9f67',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    now(),
    now(),
    NULL -- Southland
),
(
    'aa5c110c-9123-4ca2-8d6e-a587b969cb8f',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    now(),
    now(),
    NULL -- Taranaki
),
(
    '453c4cc7-2beb-4652-b94c-35d31b875601',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    now(),
    now(),
    NULL -- Tasman
),
(
    '982be3d5-77ff-4614-b1b3-6db12fbab84f',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    now(),
    now(),
    NULL -- Waikato
),
(
    '83de1c91-bcbb-4336-8d9a-26a8d054a1b8',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    now(),
    now(),
    NULL -- Wellington
);

-- Create 10 regular rounds and 3 finals stages for 2026 season
INSERT INTO stages (
    id,
    season_id,
    name,
    stage_type,
    order_index,
    created_at,
    updated_at
)
VALUES
(
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 1',
    'regular',
    1,
    now(),
    now()
),
(
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 2',
    'regular',
    2,
    now(),
    now()
),
(
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 3',
    'regular',
    3,
    now(),
    now()
),
(
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 4',
    'regular',
    4,
    now(),
    now()
),
(
    '43c143a3-e801-463b-9441-9abe03183d7a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 5',
    'regular',
    5,
    now(),
    now()
),
(
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 6',
    'regular',
    6,
    now(),
    now()
),
(
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 7',
    'regular',
    7,
    now(),
    now()
),
(
    'bf59da77-f08a-41ba-893b-9839568796bc',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 8',
    'regular',
    8,
    now(),
    now()
),
(
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 9',
    'regular',
    9,
    now(),
    now()
),
(
    '638f5847-72fc-44f4-8dda-d274f3450939',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Round 10',
    'regular',
    10,
    now(),
    now()
),
(
    '8786240f-3a5a-4caf-a2c6-940ff9e79a28',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Quarterfinals',
    'finals',
    11,
    now(),
    now()
),
(
    '2282c4dd-ab69-4261-bd52-2868d3bafcad',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Semifinals',
    'finals',
    12,
    now(),
    now()
),
(
    '1c336224-1d5d-431e-b294-0c0a836d312d',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'Final',
    'finals',
    13,
    now(),
    now()
);

-- Insert 2026 games
INSERT INTO games (
    id,
    season_id,
    stage_id,
    date,
    home_team_id,
    away_team_id,
    home_score,
    away_score,
    status,
    created_at,
    updated_at
)
VALUES

-- Round 1 (30 Jul – 3 Aug 2026)
(
    '60771fc2-85a7-45e3-997a-689cd728e733',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-07-30 19:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BOP v WAI
),
(
    '63ab216a-d18c-40ab-93bf-20194ebe1a43',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-01 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v CAN
),
(
    '23235dab-7f45-47d4-b3df-19b7703db2ff',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-01 19:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v HKB
),
(
    'b96afae6-b709-42d4-ac7e-3f562e243382',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-02 14:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v TAS
),
(
    '9024f4bb-79b7-4726-a351-6535a0577af6',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-02 19:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v NOR
),
(
    '0f156bae-41b0-4ddc-973b-451b20235ee7',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-02 19:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- MAN v CMK
),
(
    'eba5eb72-f4f9-48ea-8db2-52d303fe2449',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'c0289073-1c1a-4157-a93d-5b27220467f1',
    TIMESTAMPTZ '2026-08-03 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NHB v STL
),

-- Round 2 (7–10 Aug 2026)
(
    'b20dbdce-0723-49d8-8cdd-740f7e4d6dc4',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-07 19:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v AUK
),
(
    '6682171a-32b2-47e7-8f05-a85c42a4c9cb',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-08 14:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v WEL
),
(
    '8561b5e6-f371-4fde-b3e0-6eb8bf57ea07',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-08 19:10+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WAI v TAR
),
(
    '44b38a82-396a-4d13-8ad1-360614c44af7',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-09 14:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v CAN
),
(
    'ec890805-c119-4edd-a73f-2e0bf7b54e83',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-09 19:10+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NOR v MAN
),
(
    '3deff7b5-09dc-46ac-ab72-ac539181d118',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-09 19:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v NHB
),
(
    'bf301fc6-e233-4984-9b63-f108b6671ec8',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ae5f7bff-7798-4a3d-8c0c-2546f0bad08a',
    TIMESTAMPTZ '2026-08-10 14:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CMK v STL
),

-- Round 3 (13–17 Aug 2026)
(
    'aaba658e-a3e7-436a-a1ae-b5478afe32d0',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-13 19:10+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v WAI
),
(
    '39fffa8e-902f-4488-9db6-2d22ea94db34',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-15 14:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v BOP
),
(
    'afe70513-411a-4ad9-94eb-ea2ca5f732d7',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-15 19:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v CMK
),
(
    '799d9207-0791-4ca4-9243-0adfdd0b0212',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-16 14:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v NHB
),
(
    '5c33babb-7f8c-4181-af14-2c96c03ea158',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-16 19:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v MAN
),
(
    '68dec969-37a1-43f9-9a4f-75257431b6da',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-16 19:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v STL
),
(
    'cc1a6d5c-0a88-465e-9172-5063ba4e8c13',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '686b7894-3684-4c47-9be7-9fd7db34859e',
    TIMESTAMPTZ '2026-08-17 14:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v NOR
),

-- Round 4 (20–24 Aug 2026)
(
    '37b77c0f-62ee-4d1b-a034-d7fbf2c0c244',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-20 19:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v NOR
),
(
    '7572ee51-f062-447c-9241-62672a4c0c9a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-22 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v BOP
),
(
    'b82c56fc-57c2-4f43-8c4c-90f729318a4f',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-22 19:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v OTA
),
(
    'e6bd1234-0ea7-4fbe-bc65-c75bdd078a64',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-23 14:05+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WAI v CMK
),
(
    '8845fb60-3c8b-4768-8c53-999104bb2b29',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-23 19:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v MAN
),
(
    '0b51f509-7b7a-4fdd-84ef-4c2f68cc2628',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-23 19:10+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NHB v AUK
),
(
    'ed86ceaa-9af3-4ed7-a8ad-1c3e22f7da5e',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'ea4a42cd-5fb7-432b-93be-4610b04f8aa2',
    TIMESTAMPTZ '2026-08-24 14:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v STL
),

-- Round 5 (27–31 Aug 2026)
(
    '2894e2ea-5a6e-43df-bdf7-ffdc3db89c8c',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-28 19:10+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NOR v OTA
),
(
    'c9c28e4b-9a4a-4d2b-a3cb-42439f19810c',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-29 14:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v CAN
),
(
    'e4ee5d65-67fc-4749-9d3a-94bfa36c694f',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-29 19:10+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BOP v CMK
),
(
    '24345335-b790-4199-823c-0f795e87f3dd',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-30 14:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v WEL
),
(
    'ab0ab6c7-3198-490d-8e0f-a3798703aae5',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-30 19:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- MAN v TAS
),
(
    '85ef413b-bce4-4df4-8592-b442eb24f4d1',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-30 19:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v STL
),
(
    '0f2b998f-97c3-440c-b493-0bf3b1d77e67',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '43c143a3-e801-463b-9441-9abe03183d7a',
    TIMESTAMPTZ '2026-08-31 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NHB v WAI
),

-- Round 6 (3–7 Sep 2026)
(
    'a0279a32-5d9d-43e1-a1ac-39556e6b69a4',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-04 19:10+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v HKB
),
(
    'cc2aff21-f38f-46d7-bbc9-2f7d34568d86',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-05 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v NOR
),
(
    '2d3934f2-93b3-4b32-be42-3d1b4c90f400',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-05 19:10+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WAI v CAN
),
(
    '1b91ae25-25b9-4b96-b53c-5716c03f7732',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-06 14:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CMK v TAR
),
(
    'cce243fe-8bf9-4ec7-8ddf-540b3531f945',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-06 19:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v MAN
),
(
    '8dc10da7-34b6-4281-922b-57874f62a854',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-06 19:10+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BOP v AUK
),
(
    'bbf55f79-58e5-48d3-a4ec-1007ceeca7e7',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '7c53af64-38e1-4b70-b311-57aaf150f54a',
    TIMESTAMPTZ '2026-09-07 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NHB v STL
),

-- Round 7 (10–14 Sep 2026)
(
    '91e0e1aa-af18-4f9f-b3c4-9825bff13f68',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-10 19:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v TAR
),
(
    '5c24b9fa-cedd-4c30-be77-f4365cf3277b',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-11 14:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v TAS
),
(
    '76c2552d-d705-4489-96ac-bc66fee0149f',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-11 19:10+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v WAI
),
(
    '1ff5b381-92ae-4740-925c-cb7a315aa636',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-12 14:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BOP v MAN
),
(
    '5838a9c6-774f-4aa3-8c83-848322ad6b43',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-12 19:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v WEL
),
(
    'd2318ff2-ebc7-45a8-8d95-409194f33738',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-13 14:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CMK v NHB
),
(
    '3001be04-11a3-4da1-8df1-3a0b9cc3575e',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '0b677f5c-ecd2-49b8-b67a-4c2e4ba5a152',
    TIMESTAMPTZ '2026-09-13 14:05+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- STL v NOR
),

-- Round 8 (17–21 Sep 2026)
(
    'dc24bfa8-17b3-4d4d-832b-c737dbccb9d1',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-17 19:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v HKB
),
(
    'bca40dbf-a818-4ab4-a6f2-b286df0cb1d8',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-18 19:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v AUK
),
(
    'd7f23081-b4ab-4b57-9369-1fc10e6f2f55',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-19 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v OTA
),
(
    '6387517a-1bcc-478c-8f88-276f9f031d15',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-19 17:05+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WAI v MAN
),
(
    'bb57d63b-a66d-41cb-b460-13a286845aa1',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-19 19:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v CMK
),
(
    'cf00ae4f-4001-4b1e-a979-143186b0bfc3',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-20 14:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NOR v BOP
),
(
    '67a615e9-2d28-41f0-9134-2ee8077b65bf',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    'bf59da77-f08a-41ba-893b-9839568796bc',
    TIMESTAMPTZ '2026-09-20 19:10+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- NHB v STL
),

-- Round 9 (24–28 Sep 2026)
(
    'bf852bcb-f48b-4e35-98bc-efa4fcd91a9c',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-24 19:10+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAS v BOP
),
(
    'b9511ec5-ead7-4e55-896e-b8f3776aca96',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-25 19:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v WAI
),
(
    '967acc3e-1093-45b1-99f9-93bc03be8ed4',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-26 14:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HKB v CAN
),
(
    'f4eb096c-2e1a-43a4-898a-2783f8348686',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-26 19:10+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- MAN v WEL
),
(
    '8ba6cb85-05e6-4e23-ba0c-a70e6c5679a3',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-27 14:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- OTA v CMK
),
(
    '247d35bf-d52b-48e0-9fa9-5ea53fd5d688',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-27 19:10+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- TAR v NHB
),
(
    'ca3bf06e-3a5f-4cc3-bf56-fd68b1155ec5',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '8c6f904b-84b0-442c-a442-1827b266bd65',
    TIMESTAMPTZ '2026-09-28 14:05+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- STL v NOR
),

-- Round 10 (1–4 Oct 2026)
(
    'b4fb471d-d64c-43ca-83e3-ab11af9c691a',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-01 19:10+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WEL v HKB
),
(
    '50f8be67-53f5-4a3c-b2dc-4db5897b939e',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-02 19:10+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- AUK v OTA
),
(
    '130b9341-f2d8-42b2-b9e2-17ad41cb1b36',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-03 14:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BOP v TAR
),
(
    '0ea75f95-90bf-4b89-8138-eb097f2e01ac',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-03 19:10+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- WAI v TAS
),
(
    '6de86d27-177d-4e00-b33c-fcf035801082',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-03 19:10+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CAN v NHB
),
(
    '5725ca41-0c58-4cc5-960b-0667d2f570f1',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-04 14:05+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- MAN v NOR
),
(
    '8be68081-8d16-442f-8d8a-3b1b34854262',
    '142b57f9-c52e-4b1b-96f0-e1de6427afaf',
    '638f5847-72fc-44f4-8dda-d274f3450939',
    TIMESTAMPTZ '2026-10-04 19:10+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CMK v STL
);