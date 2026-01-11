-- Seed competition used for DELETE endpoint testing (safe to delete)
INSERT INTO competitions (id, name, created_at, updated_at)
VALUES (
    'a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd',
    'Deletable Test Competition',
    now(),
    now()
);