package db_handler

import (
	"context"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
)

func CreateCompetition(
	ctx context.Context,
	q Queries,
	params db.CreateCompetitionParams,
) error {
	return q.CreateCompetition(ctx, params)
}

func GetCompetition(
	ctx context.Context,
	q Queries,
	id uuid.UUID,
) (db.Competition, error) {
	return q.GetCompetition(ctx, id)
}

func GetCompetitions(
	ctx context.Context,
	q Queries,
) ([]db.Competition, error) {
	return q.GetCompetitions(ctx)
}

func UpdateCompetition(
	ctx context.Context,
	q Queries,
	params db.UpdateCompetitionParams,
) error {
	return q.UpdateCompetition(ctx, params)
}

func DeleteCompetition(
	ctx context.Context,
	q Queries,
	params db.DeleteCompetitionParams,
) error {
	return q.DeleteCompetition(ctx, params)
}

func CreateSeason(
	ctx context.Context,
	q Queries,
	params db.CreateSeasonParams,
) error {
	return q.CreateSeason(ctx, params)
}

func GetSeason(
	ctx context.Context,
	q Queries,
	id uuid.UUID,
) (db.Season, error) {
	return q.GetSeason(ctx, id)
}

func GetSeasons(
	ctx context.Context,
	q Queries,
	competitionID uuid.UUID,
) ([]db.Season, error) {
	return q.GetSeasons(ctx, competitionID)
}

func UpdateSeason(
	ctx context.Context,
	q Queries,
	params db.UpdateSeasonParams,
) error {
	return q.UpdateSeason(ctx, params)
}

func DeleteSeason(
	ctx context.Context,
	q Queries,
	params db.DeleteSeasonParams,
) error {
	return q.DeleteSeason(ctx, params)
}

func CreateTeam(
	ctx context.Context,
	q Queries,
	params db.CreateTeamParams,
) error {
	return q.CreateTeam(ctx, params)
}

func GetTeam(
	ctx context.Context,
	q Queries,
	id uuid.UUID,
) (db.Team, error) {
	return q.GetTeam(ctx, id)
}

func GetTeams(
	ctx context.Context,
	q Queries,
) ([]db.Team, error) {
	return q.GetTeams(ctx)
}

func UpdateTeam(
	ctx context.Context,
	q Queries,
	params db.UpdateTeamParams,
) error {
	return q.UpdateTeam(ctx, params)
}

func DeleteTeam(
	ctx context.Context,
	q Queries,
	params db.DeleteTeamParams,
) error {
	return q.DeleteTeam(ctx, params)
}
