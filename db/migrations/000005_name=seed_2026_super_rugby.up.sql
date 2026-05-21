-- Seed Super Rugby Pacific 2026 competition, season, teams, stages, and all round/finals games

-- Insert Super Rugby Pacific competition
INSERT INTO competitions (
    id,
    name,
    created_at,
    updated_at
)
VALUES (
    'b3f77b8d-25e5-4817-aed4-d023160cd7ed',
    'Super Rugby Pacific',
    now(),
    now()
);

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
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'b3f77b8d-25e5-4817-aed4-d023160cd7ed',
    TIMESTAMPTZ '2026-02-13 19:05+13',
    TIMESTAMPTZ '2026-05-30 14:35+12',
    now(),
    now()
);

-- Insert 2026 teams
INSERT INTO teams (
    id,
    name,
    abbreviation,
    location,
    created_at,
    updated_at
)
VALUES
(
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    'Blues',
    'BLU',
    'Auckland',
    now(),
    now()
),
(
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    'Chiefs',
    'CHI',
    'Hamilton',
    now(),
    now()
),
(
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    'Hurricanes',
    'HUR',
    'Wellington',
    now(),
    now()
),
(
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    'Crusaders',
    'CRU',
    'Christchurch',
    now(),
    now()
),
(
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    'Highlanders',
    'HIG',
    'Dunedin',
    now(),
    now()
),
(
    '17412805-230b-49b6-838b-dcc485034022',
    'ACT Brumbies',
    'BRU',
    'Canberra',
    now(),
    now()
),
(
    'ace4c47c-6b89-4375-85ff-347e50514622',
    'Queensland Reds',
    'RED',
    'Brisbane',
    now(),
    now()
),
(
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    'NSW Waratahs',
    'WAR',
    'Sydney',
    now(),
    now()
),
(
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    'Western Force',
    'FOR',
    'Perth',
    now(),
    now()
),
(
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    'Fijian Drua',
    'DRU',
    'Fiji',
    now(),
    now()
),
(
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    'Moana Pasifika',
    'MOA',
    'Pacific Islands',
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
    '5a928315-9585-4555-948b-1e02e6d8799d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    now(),
    now(),
    NULL -- Blues
),
(
    '901e3bd7-2ea9-44f4-847f-e429e56ba6ab',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    now(),
    now(),
    NULL -- Chiefs
),
(
    '5812e230-7a75-4fc7-9845-e5df6dfa4bfa',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    now(),
    now(),
    NULL -- Hurricanes
),
(
    '48c4f651-9768-4f80-bda9-cdbdd69fe780',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    now(),
    now(),
    NULL -- Crusaders
),
(
    '0496abf5-2769-41a0-9258-f68e3819abe2',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    now(),
    now(),
    NULL -- Highlanders
),
(
    'dbeb3fb7-763d-44dc-a527-f74ab9836541',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '17412805-230b-49b6-838b-dcc485034022',
    now(),
    now(),
    NULL -- ACT Brumbies
),
(
    '01e2b583-36f7-4883-9c24-2b5137ae8ae8',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    now(),
    now(),
    NULL -- Queensland Reds
),
(
    '3a4f729f-12f5-4121-a3c7-33542dbea724',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    now(),
    now(),
    NULL -- NSW Waratahs
),
(
    '7a8332d3-3049-491e-adf8-f6b3fc15dfca',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    now(),
    now(),
    NULL -- Western Force
),
(
    'bf70e474-9ab0-42c9-af9a-8d9ffcb58e3d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    now(),
    now(),
    NULL -- Fijian Drua
),
(
    '76df12d9-b012-4ce7-8ef4-34b10658b37f',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    now(),
    now(),
    NULL -- Moana Pasifika
);

-- Create 16 regular rounds and 3 finals stages for 2026 season
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
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 1',
    'regular',
    1,
    now(),
    now()
),
(
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 2',
    'regular',
    2,
    now(),
    now()
),
(
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 3',
    'regular',
    3,
    now(),
    now()
),
(
    'cb1906c4-d3e2-462f-9fc6-d13edc0e3d82',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 4',
    'regular',
    4,
    now(),
    now()
),
(
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 5',
    'regular',
    5,
    now(),
    now()
),
(
    'c7dffdee-1fc4-4231-a5f2-7b8775f5aa53',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 6',
    'regular',
    6,
    now(),
    now()
),
(
    'db77a2e3-2ca7-4241-979e-f6ccc35cae5a',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 7',
    'regular',
    7,
    now(),
    now()
),
(
    '19dbef45-4662-4858-ae9a-aa0a5acfa00d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 8',
    'regular',
    8,
    now(),
    now()
),
(
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 9',
    'regular',
    9,
    now(),
    now()
),
(
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 10',
    'regular',
    10,
    now(),
    now()
),
(
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 11',
    'regular',
    11,
    now(),
    now()
),
(
    '1aa5321e-cc15-4e5a-bc8b-86559d5a2432',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 12',
    'regular',
    12,
    now(),
    now()
),
(
    'ce4c3fb3-be49-4b22-a737-cd26361121d5',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 13',
    'regular',
    13,
    now(),
    now()
),
(
    '407e8805-79e2-4295-a051-20907918411e',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 14',
    'regular',
    14,
    now(),
    now()
),
(
    '5e6fea46-4397-4f37-ba84-71c20a23a193',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 15',
    'regular',
    15,
    now(),
    now()
),
(
    'e551e625-a067-4dc9-a59e-f155fd3ade1f',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Round 16',
    'regular',
    16,
    now(),
    now()
),
(
    '85d6bf83-e40a-4937-be7f-109bc1cd7279',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Qualifying Finals',
    'finals',
    17,
    now(),
    now()
),
(
    '29824133-1005-4e2d-a036-a23c72178ff4',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Semi Finals',
    'finals',
    18,
    now(),
    now()
),
(
    '42ddbb76-8b57-40c2-b5ce-2583247edad2',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'Grand Final',
    'finals',
    19,
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

-- Round 1 (13–14 Feb 2026)
(
    '9398dbb6-ac63-4f5c-b321-0eac630355f1',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    TIMESTAMPTZ '2026-02-13 19:05+13',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    17,
    16,
    'finished',
    now(),
    now() -- HIG v CRU
),
(
    'a64652da-1f76-40cf-afd3-d0ad1d576b24',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    TIMESTAMPTZ '2026-02-13 19:35+11',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    25,
    24,
    'finished',
    now(),
    now() -- WAR v RED
),
(
    'd4284a58-7c42-4e7a-a05b-5a257f9fed93',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    TIMESTAMPTZ '2026-02-14 15:35+11',
    '17412805-230b-49b6-838b-dcc485034022',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    40,
    15,
    'finished',
    now(),
    now() -- BRU v FOR
),
(
    'bf7cc7a7-af70-4353-aa4b-4465af448bc9',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    TIMESTAMPTZ '2026-02-14 19:05+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    35,
    27,
    'finished',
    now(),
    now() -- CHI v HUR
),
(
    '73cf0621-d591-409c-9d03-f021874f536d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '77cbe7eb-772a-485b-9b77-4f8624c31183',
    TIMESTAMPTZ '2026-02-14 19:35+13',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    38,
    17,
    'finished',
    now(),
    now() -- BLU v DRU
),

-- Round 2 (20–22 Feb 2026)
(
    '97eefe5f-aeec-4b30-8d6f-775777a8ed3f',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    TIMESTAMPTZ '2026-02-20 19:05+13',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '17412805-230b-49b6-838b-dcc485034022',
    36,
    24,
    'finished',
    now(),
    now() -- HUR v BRU
),
(
    'eecf9933-483b-4ac7-b9a3-8f30f1048caf',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    TIMESTAMPTZ '2026-02-20 19:35+11',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    33,
    19,
    'finished',
    now(),
    now() -- DRU v WAR
),
(
    '5ffc0d93-99f6-4f81-8915-261d6f1a78ba',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    TIMESTAMPTZ '2026-02-21 19:05+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    44,
    10,
    'finished',
    now(),
    now() -- CHI v HIG
),
(
    '7515369c-d44a-41e1-b02a-f77871602072',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    TIMESTAMPTZ '2026-02-21 16:35+8',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    21,
    27,
    'finished',
    now(),
    now() -- FOR v BLU
),
(
    'ee57277f-bcac-4964-8bfe-16f9e76c1971',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0bdda8e-e212-4d91-af76-2674232d6225',
    TIMESTAMPTZ '2026-02-22 15:35+11',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    13,
    24,
    'finished',
    now(),
    now() -- RED v CRU
),

-- Round 3 (27–28 Feb 2026)
(
    '2b7bf2fe-500a-4066-873b-aa2f5d149d38',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    TIMESTAMPTZ '2026-02-27 19:05+13',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    35,
    17,
    'finished',
    now(),
    now() -- HIG v MOA
),
(
    '4adca519-3b4b-472c-80c3-a16e718255b9',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    TIMESTAMPTZ '2026-02-27 18:35+13',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    24,
    22,
    'finished',
    now(),
    now() -- MOA v FOR
),
(
    '093abc2e-7128-479a-bf17-1e36962d6be3',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    TIMESTAMPTZ '2026-02-28 15:35+11',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '17412805-230b-49b6-838b-dcc485034022',
    37,
    27,
    'finished',
    now(),
    now() -- RED v BRU
),
(
    'dd58d736-7694-4ba9-834f-a774bab8ab2d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    TIMESTAMPTZ '2026-02-28 19:05+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    22,
    28,
    'finished',
    now(),
    now() -- CHI v CRU
),
(
    'cd958976-d5bd-4bd6-b921-b6ba3dd40b4d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'a0c0b59e-645f-4f87-b9b5-b913f402a9d3',
    TIMESTAMPTZ '2026-02-28 19:35+13',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    29,
    17,
    'finished',
    now(),
    now() -- BLU v HUR
),

-- Round 4 (6–7 Mar 2026)
(
    'f68f631d-0a57-43cb-88de-1b9645da24d4',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'cb1906c4-d3e2-462f-9fc6-d13edc0e3d82',
    TIMESTAMPTZ '2026-03-06 19:05+11',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    19,
    59,
    'finished',
    now(),
    now() -- WAR v HUR
),
(
    '00836c62-4766-4f6c-ab8a-eb4ed9d10157',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'cb1906c4-d3e2-462f-9fc6-d13edc0e3d82',
    TIMESTAMPTZ '2026-03-06 19:35+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    57,
    24,
    'finished',
    now(),
    now() -- CHI v MOA
),
(
    'ba70ab16-80cc-4d1d-be10-5bd8bdc9c075',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'cb1906c4-d3e2-462f-9fc6-d13edc0e3d82',
    TIMESTAMPTZ '2026-03-07 16:35+11',
    '17412805-230b-49b6-838b-dcc485034022',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    31,
    34,
    'finished',
    now(),
    now() -- BRU v RED
),
(
    'f0864660-edc0-4847-a52b-1b6d0b2c0468',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'cb1906c4-d3e2-462f-9fc6-d13edc0e3d82',
    TIMESTAMPTZ '2026-03-07 19:05+13',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    29,
    13,
    'finished',
    now(),
    now() -- BLU v CRU
),

-- Round 5 (13–15 Mar 2026)
(
    'e4e439d0-407b-406d-8118-8e13830e05c2',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    TIMESTAMPTZ '2026-03-13 19:05+13',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    31,
    23,
    'finished',
    now(),
    now() -- HUR v FOR
),
(
    '4f089f2d-fed6-4aec-85a6-73fa994b0d35',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    TIMESTAMPTZ '2026-03-14 19:05+13',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    29,
    18,
    'finished',
    now(),
    now() -- CRU v HIG
),
(
    '2e043911-b181-414e-8122-e7b0cd8006ec',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    TIMESTAMPTZ '2026-03-14 19:05+10',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '17412805-230b-49b6-838b-dcc485034022',
    42,
    27,
    'finished',
    now(),
    now() -- DRU v BRU
),
(
    '9a6104e8-e752-48b0-a3dc-6b68736e8637',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    TIMESTAMPTZ '2026-03-14 19:05+11',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    26,
    17,
    'finished',
    now(),
    now() -- RED v WAR
),
(
    '10d6d1d7-3ab9-4905-9504-598b2a84cb2f',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'd5cc735b-c5df-42b1-932d-88d32a80f963',
    TIMESTAMPTZ '2026-03-15 19:05+13',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    43,
    7,
    'finished',
    now(),
    now() -- BLU v MOA
),

-- Round 6 (20–21 Mar 2026)
(
    '87a736d7-5b2e-46a5-bcfa-7b6e2c87d223',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'c7dffdee-1fc4-4231-a5f2-7b8775f5aa53',
    TIMESTAMPTZ '2026-03-20 19:05+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    42,
    24,
    'finished',
    now(),
    now() -- CHI v BLU
),
(
    '055f2c51-2fbf-4cba-a05b-a23ee05ea8c1',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'c7dffdee-1fc4-4231-a5f2-7b8775f5aa53',
    TIMESTAMPTZ '2026-03-21 19:05+13',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    20,
    22,
    'finished',
    now(),
    now() -- HUR v CRU
),
(
    '537f6720-e66e-4c68-aacf-cbc2b0b32fa6',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'c7dffdee-1fc4-4231-a5f2-7b8775f5aa53',
    TIMESTAMPTZ '2026-03-21 19:05+11',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    27,
    36,
    'finished',
    now(),
    now() -- WAR v DRU
),

-- Round 7 (27–28 Mar 2026)
(
    '4f68f682-6302-4c8b-be2a-87b01ccd88cf',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'db77a2e3-2ca7-4241-979e-f6ccc35cae5a',
    TIMESTAMPTZ '2026-03-27 19:05+13',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    17,
    38,
    'finished',
    now(),
    now() -- HIG v CHI
),
(
    '850fbc50-462f-4d89-a57b-90465e0cf2f8',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'db77a2e3-2ca7-4241-979e-f6ccc35cae5a',
    TIMESTAMPTZ '2026-03-28 16:35+8',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    14,
    24,
    'finished',
    now(),
    now() -- FOR v CHI
),
(
    'dd61e303-586a-428f-9be4-adf46f4f3ed6',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'db77a2e3-2ca7-4241-979e-f6ccc35cae5a',
    TIMESTAMPTZ '2026-03-28 19:05+13',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    40,
    15,
    'finished',
    now(),
    now() -- BLU v DRU
),
(
    '02e710f8-b823-4ee0-9b5c-bb293390a20d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'db77a2e3-2ca7-4241-979e-f6ccc35cae5a',
    TIMESTAMPTZ '2026-03-28 19:05+13',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    45,
    28,
    'finished',
    now(),
    now() -- HUR v RED
),

-- Round 8 (3–4 Apr 2026)
(
    '2117d834-a904-46d0-9bab-4e052b458a9d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '19dbef45-4662-4858-ae9a-aa0a5acfa00d',
    TIMESTAMPTZ '2026-04-03 19:05+13',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    69,
    26,
    'finished',
    now(),
    now() -- CRU v DRU
),
(
    '29add84a-8039-4f19-9170-e24c4112ef8a',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '19dbef45-4662-4858-ae9a-aa0a5acfa00d',
    TIMESTAMPTZ '2026-04-04 19:05+13',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    42,
    14,
    'finished',
    now(),
    now() -- CHI v WAR
),
(
    'ac6e460b-a11c-403a-ba25-35c45ec4657b',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '19dbef45-4662-4858-ae9a-aa0a5acfa00d',
    TIMESTAMPTZ '2026-04-04 19:05+11',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    19,
    42,
    'finished',
    now(),
    now() -- RED v FOR
),

-- Round 9 (10–11 Apr 2026)
(
    '90c20198-3f3d-4012-aea5-3c475414245d',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    TIMESTAMPTZ '2026-04-10 19:05+12',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '17412805-230b-49b6-838b-dcc485034022',
    10,
    14,
    'finished',
    now(),
    now() -- HIG v BRU
),
(
    '89d26aee-709b-4ef3-a375-d4c06b8f00a5',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    TIMESTAMPTZ '2026-04-11 15:05+12',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    31,
    26,
    'finished',
    now(),
    now() -- RED v CRU
),
(
    'ca7e843f-8dbd-4c7e-832d-c945ff7eb616',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    TIMESTAMPTZ '2026-04-11 19:05+12',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    42,
    19,
    'finished',
    now(),
    now() -- HUR v BLU
),
(
    '85f2edd8-1c6a-4b0b-9dfb-cc906e843831',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    TIMESTAMPTZ '2026-04-11 19:05+12',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    17,
    62,
    'finished',
    now(),
    now() -- MOA v CHI
),
(
    'ff906112-422a-4200-b3db-402b870c360f',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'f4956b41-335d-49af-9e18-8b0c4545159d',
    TIMESTAMPTZ '2026-04-11 19:05+10',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    24,
    22,
    'finished',
    now(),
    now() -- DRU v FOR
),

-- Round 10 (17–18 Apr 2026)
(
    '3ef9d905-4e9c-44ab-9e39-2e36d0277dbd',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    TIMESTAMPTZ '2026-04-17 19:05+12',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    47,
    40,
    'finished',
    now(),
    now() -- BLU v HIG
),
(
    'e387c002-d0c1-427d-a26a-506f78737eb4',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    TIMESTAMPTZ '2026-04-17 19:35+11',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    29,
    14,
    'finished',
    now(),
    now() -- WAR v MOA
),
(
    '29a3354e-4712-4d84-94fa-1163c1fa1737',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    TIMESTAMPTZ '2026-04-18 19:05+12',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CHI v HUR
),
(
    '549702c9-7914-416a-bcd6-017ea964dbf4',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    TIMESTAMPTZ '2026-04-18 19:35+10',
    '17412805-230b-49b6-838b-dcc485034022',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BRU v DRU
),
(
    '3c6286dc-b148-4d2d-9c8e-48d31eee275e',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '2f953c79-614f-435e-b6da-8aa844eb4b37',
    TIMESTAMPTZ '2026-04-18 19:55+8',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- FOR v CRU
),

-- Round 11 (24–26 Apr 2026) — ANZAC / Super Round at Te Kaha, Christchurch
(
    '6889eaa8-b1cf-4c9e-8db7-bb28df123573',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    TIMESTAMPTZ '2026-04-24 19:35+12',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CRU v WAR
),
(
    '9cb15728-9b8b-416f-a593-74f0789ec6a0',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    TIMESTAMPTZ '2026-04-25 17:05+12',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '17412805-230b-49b6-838b-dcc485034022',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HUR v BRU
),
(
    '2220b1ef-05e4-4499-80c4-d15dbd2854f1',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    TIMESTAMPTZ '2026-04-25 19:35+12',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BLU v RED
),
(
    '5d06164e-6dde-4471-bcb4-614828bc5bff',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    TIMESTAMPTZ '2026-04-26 14:00+12',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HIG v MOA
),
(
    'a42f0efa-76ec-439d-9285-4a140066b3ba',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '35691741-2c7b-44d5-84e3-6c5ca20dc992',
    TIMESTAMPTZ '2026-04-26 16:30+12',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CHI v DRU
),

-- Round 12 (1–2 May 2026)
(
    'bacf68fa-5816-414e-ae99-ea2686fca930',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '1aa5321e-cc15-4e5a-bc8b-86559d5a2432',
    TIMESTAMPTZ '2026-05-01 19:05+12',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HUR v CHI
),
(
    '21c78e8d-3987-46bf-9f9e-5589f1141915',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '1aa5321e-cc15-4e5a-bc8b-86559d5a2432',
    TIMESTAMPTZ '2026-05-02 16:35+10',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- DRU v HIG
),
(
    '0e40f4f5-b8ca-42bb-8d73-308c51a8acab',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '1aa5321e-cc15-4e5a-bc8b-86559d5a2432',
    TIMESTAMPTZ '2026-05-02 19:05+10',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- FOR v RED
),

-- Round 13 (8–9 May 2026)
(
    '91ea0862-ab54-4847-b7f0-a691530d12ee',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'ce4c3fb3-be49-4b22-a737-cd26361121d5',
    TIMESTAMPTZ '2026-05-08 19:05+12',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '17412805-230b-49b6-838b-dcc485034022',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BLU v BRU
),
(
    'dfc00ef3-ee70-47e7-99a6-914a960aaaf1',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'ce4c3fb3-be49-4b22-a737-cd26361121d5',
    TIMESTAMPTZ '2026-05-09 16:35+10',
    'ace4c47c-6b89-4375-85ff-347e50514622',
    '2047449a-bc30-43e8-87e4-307d3c4d88a0',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- RED v MOA
),
(
    '5576bbd1-267c-49bb-ad6a-cdcffc69ef43',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'ce4c3fb3-be49-4b22-a737-cd26361121d5',
    TIMESTAMPTZ '2026-05-09 19:05+12',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CRU v CHI
),

-- Round 14 (15–16 May 2026)
(
    '05b69118-eb68-4c71-bb87-34a0d1a756d1',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '407e8805-79e2-4295-a051-20907918411e',
    TIMESTAMPTZ '2026-05-15 19:05+12',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HIG v HUR
),
(
    '7b598720-d398-4429-8fd8-8eb7289ae5ef',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '407e8805-79e2-4295-a051-20907918411e',
    TIMESTAMPTZ '2026-05-16 19:05+10',
    '83cc840d-4e4f-4ce7-b98b-b28ce2048ac7',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- DRU v HUR
),

-- Round 15 (22–23 May 2026)
(
    '73e420bf-e473-47fe-9dad-4823e744291a',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '5e6fea46-4397-4f37-ba84-71c20a23a193',
    TIMESTAMPTZ '2026-05-22 19:05+12',
    '7c7d5ef8-6bce-44da-a057-2b7226e72cb9',
    '160c56fa-fb81-4d4d-ba16-0e6578e3eeee',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CHI v FOR
),
(
    '323883cb-f7e7-4cf0-88d5-6095c58abe08',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '5e6fea46-4397-4f37-ba84-71c20a23a193',
    TIMESTAMPTZ '2026-05-23 16:35+11',
    '17412805-230b-49b6-838b-dcc485034022',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BRU v CRU
),
(
    '24f9fd4b-ccd8-46f3-93d2-07876296e528',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    '5e6fea46-4397-4f37-ba84-71c20a23a193',
    TIMESTAMPTZ '2026-05-23 19:05+12',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    '78d1c75a-348f-4b0f-ab08-4b176975f4ab',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- BLU v WAR
),

-- Round 16 (29–30 May 2026)
(
    '6f31cd12-caed-4845-8742-0590b29edbfd',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'e551e625-a067-4dc9-a59e-f155fd3ade1f',
    TIMESTAMPTZ '2026-05-29 19:05+12',
    '4c2328ca-ab04-48bd-8385-ef9abe0d336f',
    '80bb746c-08bd-4c9f-8485-3806ec094c17',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- HUR v HIG
),
(
    '1eeb42f6-869e-4b8f-84ea-2d655d68f663',
    'ac2f106c-5837-4781-9f6a-8b5d5a6d217f',
    'e551e625-a067-4dc9-a59e-f155fd3ade1f',
    TIMESTAMPTZ '2026-05-30 14:35+12',
    '578410e9-2a85-4b2a-a85c-5a92a50b006d',
    'd60a381b-ed5c-4a93-a466-eb74bf7a8cf4',
    NULL,
    NULL,
    'scheduled',
    now(),
    now() -- CRU v BLU
);