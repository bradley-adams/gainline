-- Seed competition used for DELETE endpoint testing (safe to delete)
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd',
    'Deletable Test Competition',
    now(),
    now()
);

-- Seed season used for DELETE endpoint testing (safe to delete)
INSERT INTO seasons (
    id,
    competition_id,
    start_date,
    end_date,
    created_at,
    updated_at
)
VALUES (
    '9300778f-cce0-4efe-af6c-e399d8170315',
    'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd',
    TIMESTAMPTZ '2025-07-31 21:10+12',
    TIMESTAMPTZ '2025-10-25 23:10+12',
    now(),
    now()
);

-- Seed teams used for DELETE endpoint testing (safe to delete)
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