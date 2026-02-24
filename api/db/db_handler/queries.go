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

func GetCompetitionsPaginated(
	ctx context.Context,
	q Queries,
	params db.GetCompetitionsPaginatedParams,
) ([]db.Competition, error) {
	return q.GetCompetitionsPaginated(ctx, params)
}

func CountCompetitions(
	ctx context.Context,
	q Queries,
) (int64, error) {
	return q.CountCompetitions(ctx)
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

func DeleteSeasonsByCompetitionID(
	ctx context.Context,
	q Queries,
	params db.DeleteSeasonsByCompetitionIDParams,
) error {
	return q.DeleteSeasonsByCompetitionID(ctx, params)
}

func CreateStage(
	ctx context.Context,
	q Queries,
	params db.CreateStageParams,
) error {
	return q.CreateStage(ctx, params)
}

func GetStagesBySeasonID(
	ctx context.Context,
	q Queries,
	id uuid.UUID,
) ([]db.Stage, error) {
	return q.GetStagesBySeasonID(ctx, id)
}

func UpdateStage(
	ctx context.Context,
	q Queries,
	params db.UpdateStageParams,
) error {
	return q.UpdateStage(ctx, params)
}

func DeleteStage(
	ctx context.Context,
	q Queries,
	params db.DeleteStageParams,
) error {
	return q.DeleteStage(ctx, params)
}

func DeleteStagesBySeasonID(
	ctx context.Context,
	q Queries,
	params db.DeleteStagesBySeasonIDParams,
) error {
	return q.DeleteStagesBySeasonID(ctx, params)
}

func DeleteStagesByCompetitionID(
	ctx context.Context,
	q Queries,
	params db.DeleteStagesByCompetitionIDParams,
) error {
	return q.DeleteStagesByCompetitionID(ctx, params)
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

func CreateSeasonTeams(
	ctx context.Context,
	q Queries,
	params db.CreateSeasonTeamsParams,
) error {
	return q.CreateSeasonTeams(ctx, params)
}

func GetSeasonTeams(
	ctx context.Context,
	q Queries,
	seasonID uuid.UUID,
) ([]db.GetSeasonTeamsRow, error) {
	return q.GetSeasonTeams(ctx, seasonID)
}

func DeleteSeasonTeam(
	ctx context.Context,
	q Queries,
	params db.DeleteSeasonTeamParams,
) error {
	return q.DeleteSeasonTeam(ctx, params)
}

func DeleteSeasonTeamsBySeasonID(
	ctx context.Context,
	q Queries,
	params db.DeleteSeasonTeamsBySeasonIDParams,
) error {
	return q.DeleteSeasonTeamsBySeasonID(ctx, params)
}

func DeleteSeasonTeamsByCompetitionID(
	ctx context.Context,
	q Queries,
	params db.DeleteSeasonTeamsByCompetitionIDParams,
) error {
	return q.DeleteSeasonTeamsByCompetitionID(ctx, params)
}

func CreateGame(
	ctx context.Context,
	q Queries,
	params db.CreateGameParams,
) error {
	return q.CreateGame(ctx, params)
}

func GetGame(
	ctx context.Context,
	q Queries,
	id uuid.UUID,
) (db.Game, error) {
	return q.GetGame(ctx, id)
}

func GetGames(
	ctx context.Context,
	q Queries,
	seasonID uuid.UUID,
) ([]db.Game, error) {
	return q.GetGames(ctx, seasonID)
}

func UpdateGame(
	ctx context.Context,
	q Queries,
	params db.UpdateGameParams,
) error {
	return q.UpdateGame(ctx, params)
}

func DeleteGame(
	ctx context.Context,
	q Queries,
	params db.DeleteGameParams,
) error {
	return q.DeleteGame(ctx, params)
}

func DeleteGamesByCompetitionID(
	ctx context.Context,
	q Queries,
	params db.DeleteGamesByCompetitionIDParams,
) error {
	return q.DeleteGamesByCompetitionID(ctx, params)
}

func DeleteGamesBySeasonID(
	ctx context.Context,
	q Queries,
	params db.DeleteGamesBySeasonIDParams,
) error {
	return q.DeleteGamesBySeasonID(ctx, params)
}
