-- Seed deletable test data

INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd',
    'Deletable Test Competition',
    now(),
    now()
);

INSERT INTO seasons (
    id,
    competition_id,
    start_date,
    end_date,
    created_at,
    updated_at
)
VALUES (
    'fe04fe69-834f-42be-9821-04e53e8de26d',
    'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd',
    TIMESTAMPTZ '2025-07-31 21:10+12',
    TIMESTAMPTZ '2025-10-25 23:10+12',
    now(),
    now()
);

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
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0001',
    'Deletable Test Team A',
    'DTA',
    'Test City A',
    now(),
    now()
),
(
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0002',
    'Deletable Test Team B',
    'DTB',
    'Test City B',
    now(),
    now()
);

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
    'e38a003f-35b3-4f4a-95ea-1bd047d3c158',
    'fe04fe69-834f-42be-9821-04e53e8de26d',
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0001',
    now(),
    now(),
    NULL
),
(
    'e7a22e34-27e2-437c-a0c8-a21bacf57b75',
    'fe04fe69-834f-42be-9821-04e53e8de26d',
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0002',
    now(),
    now(),
    NULL
);

INSERT INTO stages (
    id,
    season_id,
    name,
    stage_type,
    order_index,
    created_at,
    updated_at
)
VALUES (
    'b138fab0-39fc-4eb5-9c10-44a918ed3952',
    'fe04fe69-834f-42be-9821-04e53e8de26d',
    'Round 1',
    'regular',
    1,
    now(),
    now()
);

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
VALUES (
    '30f8181f-0a44-4ad7-a163-3ef2d29e504e',
    'fe04fe69-834f-42be-9821-04e53e8de26d',
    'b138fab0-39fc-4eb5-9c10-44a918ed3952',
    TIMESTAMPTZ '2025-07-31 21:10+12',
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0001',
    '2c6f1e7b-1d3e-4e0a-9c4b-3e5e0b9f0002',
    10,
    12,
    'finished',
    now(),
    now()
);
