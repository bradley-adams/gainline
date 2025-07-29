-- name: CreateCompetition :exec
-- Insert a new competition into the database
INSERT INTO competitions (
    id,
    name,
    created_at,
    updated_at,
    deleted_at
)
VALUES (
    @id,
    @name,
    @created_at,
    @updated_at,
    @deleted_at
);

-- name: GetCompetition :one
-- Fetch a competition by id, excluding soft-deleted competitions
SELECT 
	id,
	name,
	created_at,
	updated_at,
	deleted_at 
FROM 
	competitions
WHERE 
	id = @id
AND
	deleted_at IS NULL;

-- name: GetCompetitions :many
-- Fetch all competitions, excluding soft-deleted competitions
SELECT
	id,
	name,
	created_at,
	updated_at,
	deleted_at 
FROM
	competitions
WHERE
	deleted_at IS NULL;

-- name: UpdateCompetition :exec
-- Update an existing competition by id
UPDATE competitions
SET
	name = @name
WHERE
	id = @id
AND
	deleted_at IS NULL;

-- name: DeleteCompetition :exec
-- Soft delete a competition
UPDATE competitions
SET
	deleted_at = @deleted_at
WHERE
	id = @id
AND
	deleted_at IS NULL;	