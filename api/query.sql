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

-- name: CreateSeason :exec
-- Insert a new season into the database
INSERT INTO seasons (
	id,
	competition_id,
	start_date,
	end_date,
	rounds,
	created_at,
	updated_at,
	deleted_at
)
VALUES (
	@id,
	@competition_id,
	@start_date,
	@end_date,
	@rounds,
	@created_at,
	@updated_at,
	@deleted_at
);

-- name: GetSeason :one
-- Fetch a season by id, excluding soft-deleted seasons
SELECT
	id,
	competition_id,
	start_date,
	end_date,
	rounds,
	created_at,
	updated_at,
	deleted_at
FROM
	seasons
WHERE
	id = @id
AND
	deleted_at IS NULL;

-- name: GetSeasons :many
-- Fetch all seasons for a competition, excluding soft-deleted seasons
SELECT
	id,
	competition_id,
	start_date,
	end_date,
	rounds,
	created_at,
	updated_at,
	deleted_at
FROM
	seasons
WHERE
	competition_id = @competition_id
AND
	deleted_at IS NULL;

-- name: UpdateSeason :exec
-- Update an existing season by id
UPDATE seasons
SET
	competition_id = @competition_id,
	start_date = @start_date,
	end_date = @end_date,
	rounds = @rounds,
	updated_at = @updated_at
WHERE
	id = @id
AND	
	deleted_at IS NULL;

-- name: DeleteSeason :exec
-- Soft delete a season
UPDATE seasons
SET
	deleted_at = @deleted_at
WHERE
	id = @id
AND
	deleted_at IS NULL;

-- name: DeleteSeasonsByCompetitionID :exec
-- Soft delete all seasons for a competition
UPDATE seasons
SET
    deleted_at = @deleted_at
WHERE
    competition_id = @competition_id
AND
    deleted_at IS NULL;

-- name: CreateTeam :exec
-- Insert a new team into the database
INSERT INTO teams (
	id,
	name,
	abbreviation,
	location,
	created_at,
	updated_at,
	deleted_at
)
VALUES (
	@id,
	@name,
	@abbreviation,
	@location,
	@created_at,
	@updated_at,
	@deleted_at
);

-- name: GetTeam :one
-- Fetch a team by id, excluding soft-deleted teams
SELECT
	id,
	name,
	abbreviation,
	location,
	created_at,
	updated_at,
	deleted_at
FROM
	teams
WHERE
	id = @id
AND
	deleted_at IS NULL;	

-- name: GetTeams :many
-- Fetch all teams for a competition, excluding soft-deleted teams
SELECT
	id,
	name,
	abbreviation,
	location,
	created_at,
	updated_at,
	deleted_at
FROM
	teams
WHERE
	deleted_at IS NULL;	

-- name: UpdateTeam :exec
-- Update an existing team by id
UPDATE teams
SET
	name = @name,
	abbreviation = @abbreviation,
	location = @location,
	updated_at = @updated_at
WHERE
	id = @id
AND
	deleted_at IS NULL;

-- name: DeleteTeam :exec
-- Soft delete a team
UPDATE teams
SET
	deleted_at = @deleted_at
WHERE
	id = @id
AND
	deleted_at IS NULL;

-- name: CreateSeasonTeams :exec
-- Insert a new season_teams relationship
INSERT INTO season_teams (
  id,
  team_id,
  season_id,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  @id,
  @team_id,
  @season_id,
  @created_at,
  @updated_at,
  @deleted_at
);

-- name: GetSeasonTeams :many
-- Fetch all season_teams
SELECT
  id,
  team_id,
  season_id,
  created_at,
  updated_at,
  deleted_at
FROM
  season_teams
WHERE
  season_id = @season_id
AND
  deleted_at IS NULL;

-- name: DeleteSeasonTeam :exec
-- Soft delete a team_season record
UPDATE season_teams
SET
  deleted_at = @deleted_at
WHERE
  id = @id
AND
  deleted_at IS NULL;

-- name: DeleteSeasonTeamsBySeasonID :exec
-- Soft delete all team_season records for a given season
UPDATE season_teams
SET
  deleted_at = @deleted_at
WHERE
  season_id = @season_id
AND
  deleted_at IS NULL;

-- name: CreateGame :exec
-- Insert a new game into the database
INSERT INTO games (
    id,
    season_id,
    round,
    date,
    home_team_id,
    away_team_id,
    home_score,
    away_score,
    status,
    created_at,
    updated_at,
    deleted_at
)
VALUES (
    @id,
    @season_id,
    @round,
    @date,
    @home_team_id,
    @away_team_id,
    @home_score,
    @away_score,
    @status,
    @created_at,
    @updated_at,
    @deleted_at
);

-- name: GetGame :one
-- Fetch a game by id, excluding soft-deleted games
SELECT
    id,
    season_id,
    round,
    date,
    home_team_id,
    away_team_id,
    home_score,
    away_score,
    status,
    created_at,
    updated_at,
    deleted_at
FROM
    games
WHERE
    id = @id
AND
    deleted_at IS NULL;

-- name: GetGames :many
-- Fetch all games for a season, excluding soft-deleted games
SELECT
    id,
    season_id,
    round,
    date,
    home_team_id,
    away_team_id,
    home_score,
    away_score,
    status,
    created_at,
    updated_at,
    deleted_at
FROM
    games
WHERE
    season_id = @season_id
AND
    deleted_at IS NULL;

-- name: UpdateGame :exec
-- Update an existing game by id
UPDATE games
SET
    round = @round,
    date = @date,
    home_team_id = @home_team_id,
    away_team_id = @away_team_id,
    home_score = @home_score,
    away_score = @away_score,
    status = @status,
    updated_at = @updated_at
WHERE
    id = @id
AND
    deleted_at IS NULL;

-- name: DeleteGame :exec
-- Soft delete a game
UPDATE games
SET
    deleted_at = @deleted_at
WHERE
    id = @id
AND
    deleted_at IS NULL;

-- name: DeleteGamesByCompetitionID :exec
-- Soft delete all games belonging to a competition via seasons
UPDATE games
SET
    deleted_at = @deleted_at
WHERE
    season_id IN (
        SELECT id
        FROM seasons
        WHERE competition_id = @competition_id
          AND deleted_at IS NULL
    )
AND
    deleted_at IS NULL;

-- name: DeleteGamesBySeasonID :exec
-- Soft delete all games for a given season
UPDATE games
SET
  deleted_at = @deleted_at
WHERE
  season_id = @season_id
AND
  deleted_at IS NULL;
