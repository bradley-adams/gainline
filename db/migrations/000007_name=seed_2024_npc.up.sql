-- Seed National Provincial Championship 2024 season, stages, and all round/finals games
-- Competition and teams are owned by 000003_seed_2025_npc and are referenced here by ID

-- Insert 2024 season
INSERT INTO seasons (
    id,
    competition_id,
    start_date,
    end_date,
    created_at,
    updated_at
)
VALUES (
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '44dd315c-1abc-43aa-9843-642f920190d1',
    TIMESTAMPTZ '2024-08-01 19:05+12',
    TIMESTAMPTZ '2024-10-19 18:00+13',
    now(),
    now()
);

-- Connect teams to 2024 season
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
    '10100d4b-4ddf-4afe-b99f-49d28e9825f5',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    now(),
    now(),
    NULL -- Auckland
),
(
    '7311dda8-a78f-4150-8cb8-d51c58b7c034',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    now(),
    now(),
    NULL -- Bay Of Plenty
),
(
    '67757d31-819c-4aeb-96ed-5b6807f5f4f1',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    now(),
    now(),
    NULL -- Canterbury
),
(
    '845196d0-fee9-478f-8517-08cfb2098183',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    now(),
    now(),
    NULL -- Counties Manukau
),
(
    'd74a5067-0dac-4463-a204-005ead3d3009',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    now(),
    now(),
    NULL -- Hawke's Bay
),
(
    'fc51185c-bbe5-4870-b8e4-91297b1b1df7',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    now(),
    now(),
    NULL -- Manawatū
),
(
    'cc6946ea-cd57-4f10-b20b-358695bad4b8',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    now(),
    now(),
    NULL -- North Harbour
),
(
    '01ecb69f-3da9-4030-af65-e27867c391bb',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    now(),
    now(),
    NULL -- Northland
),
(
    'a00d312e-67d4-44ef-9937-fddd8bd99fc1',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    now(),
    now(),
    NULL -- Otago
),
(
    'd594aff4-ad42-43be-9824-f8c1c93b5ca4',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    now(),
    now(),
    NULL -- Southland
),
(
    '20f129cd-7aa0-46a5-a824-553880ec8a26',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    now(),
    now(),
    NULL -- Taranaki
),
(
    '6c619844-3f96-4844-be95-fa938559ee22',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    now(),
    now(),
    NULL -- Tasman
),
(
    '716a44c2-96e6-4a38-8cf3-3087fac6902e',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    now(),
    now(),
    NULL -- Waikato
),
(
    'd571c78d-b7b3-4345-b0f2-c9f4c6f7c884',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    now(),
    now(),
    NULL -- Wellington
)
;

-- Create 9 regular rounds and 3 finals stages for 2024 season
-- NOTE: No game rows are inserted for Quarterfinals, Semifinals, or Final as matchups
-- are unknown until the regular season ends.
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
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 1',
    'regular',
    1,
    now(),
    now()
),
(
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 2',
    'regular',
    2,
    now(),
    now()
),
(
    '00d9e2b7-672a-4203-8114-af2198761c79',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 3',
    'regular',
    3,
    now(),
    now()
),
(
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 4',
    'regular',
    4,
    now(),
    now()
),
(
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 5',
    'regular',
    5,
    now(),
    now()
),
(
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 6',
    'regular',
    6,
    now(),
    now()
),
(
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 7',
    'regular',
    7,
    now(),
    now()
),
(
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 8',
    'regular',
    8,
    now(),
    now()
),
(
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Round 9',
    'regular',
    9,
    now(),
    now()
),
(
    '9c4c14d2-5e75-49bb-8805-57a911692441',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Quarterfinals',
    'finals',
    10,
    now(),
    now()
),
(
    'a5f16462-f937-439c-a31b-b5a45256108b',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Semifinals',
    'finals',
    11,
    now(),
    now()
),
(
    '5e85c90c-8ad1-4946-9bac-accb90be99c6',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'Final',
    'finals',
    12,
    now(),
    now()
)
;

-- Insert 2024 games
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

-- Round 1 (1–4 Aug 2024)
(
    '14bdb11a-9739-48ae-9eef-5416e0b28aa3',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-01 19:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    31,
    15,
    'finished',
    now(),
    now() -- TAR v CMK
),
(
    '81f53f81-805f-49a5-8545-72b39cada407',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-02 19:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    21,
    29,
    'finished',
    now(),
    now() -- AUK v WEL
),
(
    '7be0eee1-da7f-4c18-b474-32b1b225fa56',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-03 14:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    34,
    21,
    'finished',
    now(),
    now() -- CAN v NOR
),
(
    '8c8d4147-a93d-414a-a3c9-06b19f53db61',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-03 16:35+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    22,
    13,
    'finished',
    now(),
    now() -- STL v OTA
),
(
    'b689688e-4196-45a1-92ea-e6e274535c40',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-03 19:05+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    21,
    36,
    'finished',
    now(),
    now() -- WAI v BOP
),
(
    '37be74ed-4dbf-4889-9b5c-142602d0f03f',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-04 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    32,
    41,
    'finished',
    now(),
    now() -- NHB v HKB
),
(
    '1dd4deef-454a-44c2-8f6b-3ec89fb8f674',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'bafe22c7-ec42-4274-b7f2-3efa3fd24e43',
    TIMESTAMPTZ '2024-08-04 19:05+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    21,
    54,
    'finished',
    now(),
    now() -- MAN v TAS
),
-- Round 2 (8–11 Aug 2024)
(
    '7a588e34-7f81-4643-9ded-556923748f78',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-08 19:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    27,
    25,
    'finished',
    now(),
    now() -- OTA v AUK
),
(
    '3f6c20c4-a013-4c90-9b33-9f4e2f28a33b',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-09 19:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    35,
    18,
    'finished',
    now(),
    now() -- NOR v MAN
),
(
    '9c87c174-2675-48eb-9de5-60821114a3bd',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-10 14:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    22,
    7,
    'finished',
    now(),
    now() -- TAS v CAN
),
(
    '4c2da2f8-556d-4e6a-9914-48f4f2c314a2',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-10 16:35+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    31,
    17,
    'finished',
    now(),
    now() -- HKB v STL
),
(
    '9669ac65-1c72-4070-a636-ed0fa71a09e4',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-10 19:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    26,
    19,
    'finished',
    now(),
    now() -- WEL v TAR
),
(
    'ed81b230-dfe1-4b85-8b37-fbcdc36d76f7',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-11 14:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    24,
    20,
    'finished',
    now(),
    now() -- BOP v NHB
),
(
    'f71a3f83-e4aa-44b5-9251-64f9727c67c4',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '1d133ee4-f0b1-4276-85b5-0b4df9334f58',
    TIMESTAMPTZ '2024-08-11 19:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    20,
    26,
    'finished',
    now(),
    now() -- CMK v WAI
),
-- Round 3 (15–18 Aug 2024)
(
    'e23524cd-562a-4e01-a864-d30995cecc82',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-15 19:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    55,
    30,
    'finished',
    now(),
    now() -- HKB v NOR
),
(
    'c463e66d-b9ed-407e-8120-b7d7278b7185',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-16 19:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    3,
    48,
    'finished',
    now(),
    now() -- CMK v TAS
),
(
    'f4c86ba7-d712-496d-af14-be2c9dc4574e',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-17 14:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    21,
    27,
    'finished',
    now(),
    now() -- AUK v CAN
),
(
    '897cfae9-7ef2-4a26-8977-5509b815d6e6',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-17 16:35+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    24,
    39,
    'finished',
    now(),
    now() -- STL v TAR
),
(
    '17870745-4e47-4081-81e1-ea96f10e9841',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-17 19:05+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    31,
    26,
    'finished',
    now(),
    now() -- OTA v BOP
),
(
    'bbd32c58-b435-4fee-8681-71619173109a',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-18 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    39,
    31,
    'finished',
    now(),
    now() -- WEL v MAN
),
(
    '8fc9b11c-e254-40e6-b595-af8b807d9ce0',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '00d9e2b7-672a-4203-8114-af2198761c79',
    TIMESTAMPTZ '2024-08-18 19:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    43,
    29,
    'finished',
    now(),
    now() -- NHB v WAI
),
-- Round 4 (22–25 Aug 2024)
(
    '0833da34-df39-4100-b038-822361fbdef7',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-22 19:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    26,
    31,
    'finished',
    now(),
    now() -- NOR v STL
),
(
    'f19037ae-9694-4722-aff1-fc303e90469d',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-23 19:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    33,
    36,
    'finished',
    now(),
    now() -- NHB v CMK
),
(
    '814b83bc-5f3c-40e8-9eca-3235ce7d3713',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-24 14:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    22,
    18,
    'finished',
    now(),
    now() -- TAR v OTA
),
(
    '20d8884e-a3d8-42eb-b974-7c1d431e36d5',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-24 16:35+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    39,
    21,
    'finished',
    now(),
    now() -- WAI v AUK
),
(
    'bc6ab552-ea33-459e-95e3-d7e5963a0df9',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-24 19:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    34,
    15,
    'finished',
    now(),
    now() -- TAS v BOP
),
(
    'bb763cf7-d0ae-435b-b36c-e5ab0a05f1ad',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-25 14:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    21,
    46,
    'finished',
    now(),
    now() -- CAN v WEL
),
(
    '14d4a7b2-c9fb-4795-90de-bd1a4bf0fe36',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a7e3ba5a-edfc-4dac-93d3-11f21d76b9c9',
    TIMESTAMPTZ '2024-08-25 19:05+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    26,
    38,
    'finished',
    now(),
    now() -- MAN v HKB
),
-- Round 5 (29 Aug – 1 Sep 2024)
(
    '45907ba7-506d-4d7c-b420-db26cbb762bf',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-08-29 19:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    68,
    14,
    'finished',
    now(),
    now() -- BOP v MAN
),
(
    '72a02bc5-7c72-4f1d-84f4-be99f90f022c',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-08-30 19:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    36,
    12,
    'finished',
    now(),
    now() -- WEL v STL
),
(
    '29585d62-53bc-4dc6-af1a-7fa83e275f2b',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-08-31 14:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    36,
    32,
    'finished',
    now(),
    now() -- AUK v NHB
),
(
    '46dc72c2-6b08-4925-9833-656880f2a683',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-08-31 16:35+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    16,
    34,
    'finished',
    now(),
    now() -- OTA v CAN
),
(
    '7e65ecc8-fe75-42f9-a5c9-ff32e4477be7',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-08-31 19:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    24,
    25,
    'finished',
    now(),
    now() -- HKB v TAS
),
(
    '5b09c484-10dc-4f65-a1e6-92108d70f845',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-09-01 14:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    25,
    19,
    'finished',
    now(),
    now() -- TAR v WAI
),
(
    '438c5b69-034b-421d-b35e-cc7bdac726dc',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '26824cf7-bacc-406f-89bd-8a20ba05c9ff',
    TIMESTAMPTZ '2024-09-01 19:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    25,
    14,
    'finished',
    now(),
    now() -- CMK v NOR
),
-- Round 6 (5–8 Sep 2024)
(
    '3fa2bff7-d0a6-4b5e-923e-ca6cec6c2c2e',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-05 19:05+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    29,
    41,
    'finished',
    now(),
    now() -- STL v CAN
),
(
    'fc88075c-93ee-4a9c-86dd-1cb2714eb6e8',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-06 19:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    33,
    20,
    'finished',
    now(),
    now() -- BOP v TAR
),
(
    '6d8f462f-d96d-4682-a192-87f0befd90a4',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-07 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    58,
    19,
    'finished',
    now(),
    now() -- NHB v MAN
),
(
    '76f411ee-3bb4-4d85-8837-2bbd6ae5c61b',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-07 16:35+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    50,
    5,
    'finished',
    now(),
    now() -- WAI v HKB
),
(
    '16b6ee06-665e-47b3-be2b-fd8760078536',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-07 19:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    45,
    17,
    'finished',
    now(),
    now() -- CMK v OTA
),
(
    '2c5d3590-9537-405e-9bc1-08a8fed7ef94',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-08 14:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    28,
    15,
    'finished',
    now(),
    now() -- TAS v WEL
),
(
    '8532c3e9-b9a1-4cf4-92fd-3989786ec5ab',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'd8b7ff3c-1e63-4523-bd5d-2d78f95def82',
    TIMESTAMPTZ '2024-09-08 19:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    17,
    24,
    'finished',
    now(),
    now() -- NOR v AUK
),
-- Round 7 (12–15 Sep 2024)
(
    '5b6116b0-85fd-4014-af05-4089d5128d13',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-12 19:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    19,
    63,
    'finished',
    now(),
    now() -- HKB v TAR
),
(
    '1bb66625-92fb-4441-b93e-48019b753fc4',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-13 19:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    47,
    24,
    'finished',
    now(),
    now() -- NOR v NHB
),
(
    '34bd3c44-428f-4650-b77c-7039537f5fce',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-14 14:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    36,
    28,
    'finished',
    now(),
    now() -- CAN v CMK
),
(
    'b3319f80-507d-4726-b5ff-4f773500ee8f',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-14 16:35+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    30,
    25,
    'finished',
    now(),
    now() -- WEL v BOP
),
(
    '54267554-394e-4bfa-b2bf-182974303970',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-14 19:05+12',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    25,
    27,
    'finished',
    now(),
    now() -- WAI v TAS
),
(
    'db25139c-ef5d-4422-a323-c5bb6fdd2e97',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-15 14:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    27,
    19,
    'finished',
    now(),
    now() -- AUK v STL
),
(
    'eebc32b2-1b1e-48c2-b9da-d79d8cf642a6',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9e0ebaa2-b8b6-4b55-9367-04aac9cef5ac',
    TIMESTAMPTZ '2024-09-15 19:05+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    10,
    28,
    'finished',
    now(),
    now() -- MAN v OTA
),
-- Round 8 (19–22 Sep 2024)
(
    '334522cd-ad20-4ae8-bced-3678d26ff767',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-19 19:05+12',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    51,
    12,
    'finished',
    now(),
    now() -- CMK v WEL
),
(
    'f354dca6-dc39-49c9-acb4-372806e4f9bc',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-20 19:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    53,
    13,
    'finished',
    now(),
    now() -- BOP v NOR
),
(
    '5ab5fa97-e7cc-46b2-b195-c16018bfd474',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-21 14:05+12',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    36,
    35,
    'finished',
    now(),
    now() -- HKB v AUK
),
(
    'ad5bba0b-0595-4ca6-af4b-369acadeae8d',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-21 16:35+12',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    47,
    31,
    'finished',
    now(),
    now() -- OTA v TAS
),
(
    '1943a716-b3e9-4a8b-833e-3de7edd2d820',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-21 19:05+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    14,
    38,
    'finished',
    now(),
    now() -- STL v WAI
),
(
    '4321759a-9dc2-46fc-b5ad-9c5a1c815542',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-22 14:05+12',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    65,
    19,
    'finished',
    now(),
    now() -- NHB v CAN
),
(
    'df770c01-c76d-40e9-8e55-e1e2412ee852',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'c863d6d4-7bad-4eee-952e-879ce7cf6a1c',
    TIMESTAMPTZ '2024-09-22 19:05+12',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    33,
    31,
    'finished',
    now(),
    now() -- TAR v MAN
),
-- Round 9 (26–29 Sep 2024)
(
    'b753345f-5d61-496a-a374-5e9f0d1f8b6a',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-26 19:05+12',
    '7e5abf68-8358-4c20-b6a4-f64ef264c13c',
    'a5d930c3-13aa-4a85-b5c9-8f40c2c61c8a',
    28,
    31,
    'finished',
    now(),
    now() -- NOR v OTA
),
(
    '39ae6c0e-c4b0-44cc-86b9-714290beb475',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-27 19:05+12',
    '636f1f87-bc47-4e63-a3de-bf7cb8eb0c22',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    26,
    45,
    'finished',
    now(),
    now() -- MAN v CMK
),
(
    '83632b5e-42ff-472e-8709-174235416ba9',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-28 14:05+12',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    46,
    28,
    'finished',
    now(),
    now() -- WEL v HKB
),
(
    'c9b3f1a7-e2db-4c38-8d08-1ba5a079c3d3',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-28 16:35+12',
    '15c76909-f78a-4d89-bc19-7c80265e1e08',
    'e2d6c2bb-eac6-42d6-8727-4d4cbeb3e3d7',
    59,
    35,
    'finished',
    now(),
    now() -- STL v NHB
),
(
    'f9ab21d0-5fa3-4b98-852a-2eb2bc027db2',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-28 19:05+12',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    36,
    19,
    'finished',
    now(),
    now() -- CAN v WAI
),
(
    '53c0f33a-4eac-45cd-969a-c67a05b284b2',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-29 14:05+12',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    29,
    42,
    'finished',
    now(),
    now() -- TAS v TAR
),
(
    '8b97ee2d-05e6-4225-8164-0f023deff29e',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '016038c6-ec7a-47d9-8d39-7e244e5748e0',
    TIMESTAMPTZ '2024-09-29 19:05+12',
    '013952a5-87e1-4d26-a312-09b2aff54241',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    24,
    26,
    'finished',
    now(),
    now() -- AUK v BOP
),
-- Quarterfinals (5–6 Oct 2024)
(
    '6c372adb-f5e6-4fd5-a57b-a6fb05436a63',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9c4c14d2-5e75-49bb-8805-57a911692441',
    TIMESTAMPTZ '2024-10-05 14:05+13',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '6b5c3642-c026-4e89-81f7-024c40638f9a',
    29,
    14,
    'finished',
    now(),
    now() -- WEL v CMK
),
(
    'ff82e85b-18c7-460f-908a-907b8c9c8fdf',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9c4c14d2-5e75-49bb-8805-57a911692441',
    TIMESTAMPTZ '2024-10-05 19:05+12',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'dedb2044-1d2f-4dc7-84c6-509ec69c82e1',
    19,
    17,
    'finished',
    now(),
    now() -- BOP v HKB
),
(
    '5acede14-83b4-4cd8-95e8-58faac5bbb0d',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9c4c14d2-5e75-49bb-8805-57a911692441',
    TIMESTAMPTZ '2024-10-06 14:05+13',
    'bfe6ec41-e3f0-4f8f-90d2-d7bca66e1a1f',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    14,
    15,
    'finished',
    now(),
    now() -- TAR v WAI
),
(
    'c1f6af57-cd5b-4ea3-ae5d-ca15c277603b',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '9c4c14d2-5e75-49bb-8805-57a911692441',
    TIMESTAMPTZ '2024-10-06 19:05+13',
    '19b3ea1e-0c46-41f3-84ea-490b6b1db30f',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    14,
    62,
    'finished',
    now(),
    now() -- TAS v CAN
),
-- Semifinals (12–13 Oct 2024)
(
    '372b4363-7562-4e7f-9b02-a31cf154113f',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a5f16462-f937-439c-a31b-b5a45256108b',
    TIMESTAMPTZ '2024-10-12 16:05+13',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    'f192a9ce-dce2-4389-8491-1a193ac7699e',
    32,
    20,
    'finished',
    now(),
    now() -- BOP v CAN
),
(
    '8d0d3953-943f-40ce-9784-a94faaac666f',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    'a5f16462-f937-439c-a31b-b5a45256108b',
    TIMESTAMPTZ '2024-10-13 16:05+13',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    '7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97',
    29,
    24,
    'finished',
    now(),
    now() -- WEL v WAI
),
-- Final (19 Oct 2024)
(
    '9967ffa7-6e70-4ac5-9594-3b0c25f9861e',
    '0827c393-161a-41b4-badf-5678e5c8f153',
    '5e85c90c-8ad1-4946-9bac-accb90be99c6',
    TIMESTAMPTZ '2024-10-19 16:05+13',
    'ab4c78b1-5dc6-4a14-8f15-d1f144b81d96',
    'b5c6e9d7-8f11-4ef2-acc6-2e5a97839532',
    23,
    20,
    'finished',
    now(),
    now() -- WEL v BOP
)
;