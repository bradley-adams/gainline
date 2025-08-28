package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
	"github.com/pkg/errors"
)

type SeasonWithTeams struct {
	ID            uuid.UUID `json:"id"`
	CompetitionID uuid.UUID `json:"competition_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Rounds        int32     `json:"rounds"`
	Teams         []db.Team `json:"teams"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     zero.Time `json:"deleted_at"`
}

func ToSeasonResponse(s SeasonWithTeams) api.SeasonResponse {
	var teams []api.TeamResponse
	for _, team := range s.Teams {
		teams = append(teams, api.ToTeamResponse(team))
	}

	return api.SeasonResponse{
		ID:            s.ID,
		CompetitionID: s.CompetitionID,
		StartDate:     s.StartDate,
		EndDate:       s.EndDate,
		Rounds:        s.Rounds,
		Teams:         teams,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		DeletedAt:     zero.TimeFrom(s.DeletedAt.Time),
	}
}

func CreateSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
) (SeasonWithTeams, error) {
	var seasonWithTeams SeasonWithTeams

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		seasonWithTeams, err = createSeason(ctx, queries, req, competitionID)
		return err
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed creating season")
	}

	return seasonWithTeams, nil
}

func createSeason(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
) (SeasonWithTeams, error) {
	now := time.Now()
	seasonID := uuid.New()

	err := insertSeason(ctx, queries, seasonID, competitionID, req, now)
	if err != nil {
		return SeasonWithTeams{}, err
	}

	teamIDs := dedupeUUIDs(req.Teams)
	err = ensureTeamsExist(ctx, queries, teamIDs)
	if err != nil {
		return SeasonWithTeams{}, err
	}

	err = ensureSeasonHasTeams(ctx, queries, seasonID, teamIDs, now, nil)
	if err != nil {
		return SeasonWithTeams{}, err
	}

	return getSeasonWithTeams(ctx, queries, seasonID)
}

func insertSeason(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID, competitionID uuid.UUID,
	req *api.SeasonRequest,
	now time.Time,
) error {
	createSeasonParams := db.CreateSeasonParams{
		ID:            seasonID,
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Rounds:        req.Rounds,
		CreatedAt:     now,
		UpdatedAt:     now,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	err := queries.CreateSeason(ctx, createSeasonParams)
	if err != nil {
		return errors.Wrap(err, "unable to create season")
	}
	return nil
}

func ensureTeamsExist(ctx context.Context, queries db_handler.Queries, teamIDs []uuid.UUID) error {
	for _, id := range teamIDs {
		_, err := queries.GetTeam(ctx, id)
		if err != nil {
			return errors.Wrapf(err, "unable to get team %s", id.String())
		}
	}
	return nil
}

func ensureSeasonHasTeams(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	teamIDs []uuid.UUID,
	now time.Time,
	existingMap map[uuid.UUID]db.GetSeasonTeamsRow,
) error {
	for _, teamID := range teamIDs {
		if existingMap != nil {
			_, already := existingMap[teamID]
			if already {
				continue
			}
		}

		createSeasonTeamsParams := db.CreateSeasonTeamsParams{
			ID:        uuid.New(),
			SeasonID:  seasonID,
			TeamID:    teamID,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
		}

		err := queries.CreateSeasonTeams(ctx, createSeasonTeamsParams)
		if err != nil {
			return errors.Wrapf(err, "unable to add team %s to season %s", teamID.String(), seasonID.String())
		}
	}
	return nil
}

func getSeasonWithTeams(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
) (SeasonWithTeams, error) {
	season, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "unable to get season")
	}

	return buildSeasonWithTeams(ctx, queries, season)
}

func buildSeasonWithTeams(
	ctx context.Context,
	queries db_handler.Queries,
	season db.Season,
) (SeasonWithTeams, error) {
	seasonTeams, err := queries.GetSeasonTeams(ctx, season.ID)
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "unable to get season teams")
	}

	var teams []db.Team
	for _, st := range seasonTeams {
		team, err := queries.GetTeam(ctx, st.TeamID)
		if err != nil {
			return SeasonWithTeams{}, errors.Wrap(err, "unable to get team")
		}
		teams = append(teams, team)
	}

	return SeasonWithTeams{
		ID:            season.ID,
		CompetitionID: season.CompetitionID,
		StartDate:     season.StartDate,
		EndDate:       season.EndDate,
		Rounds:        season.Rounds,
		Teams:         teams,
		CreatedAt:     season.CreatedAt,
		UpdatedAt:     season.UpdatedAt,
		DeletedAt:     zero.TimeFrom(season.DeletedAt.Time),
	}, nil
}

func dedupeUUIDs(in []uuid.UUID) []uuid.UUID {
	m := make(map[uuid.UUID]struct{}, len(in))
	for _, id := range in {
		m[id] = struct{}{}
	}
	out := make([]uuid.UUID, 0, len(m))
	for id := range m {
		out = append(out, id)
	}
	return out
}

func GetSeasons(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
) ([]SeasonWithTeams, error) {
	var seasons []SeasonWithTeams

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		seasons, err = getSeasons(ctx, queries, competitionID)
		return err
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting seasons")
	}

	return seasons, nil
}

func getSeasons(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
) ([]SeasonWithTeams, error) {
	var seasonsWithTeams []SeasonWithTeams

	seasons, err := queries.GetSeasons(ctx, competitionID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get seasons")
	}

	for _, s := range seasons {
		swt, err := buildSeasonWithTeams(ctx, queries, s)
		if err != nil {
			return nil, err
		}
		seasonsWithTeams = append(seasonsWithTeams, swt)
	}

	return seasonsWithTeams, nil
}

func GetSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (SeasonWithTeams, error) {
	var season SeasonWithTeams

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		season, err = getSeasonWithTeams(ctx, queries, seasonID)
		return err
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed getting season")
	}

	return season, nil
}

func UpdateSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (SeasonWithTeams, error) {
	var seasonWithTeams SeasonWithTeams

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		seasonWithTeams, txErr = updateSeason(ctx, queries, req, competitionID, seasonID)
		return txErr
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed updating season")
	}

	return seasonWithTeams, nil
}

func updateSeason(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (SeasonWithTeams, error) {
	now := time.Now()

	err := updateSeasonFields(ctx, queries, req, competitionID, seasonID, now)
	if err != nil {
		return SeasonWithTeams{}, err
	}

	// dedupe + sync teams (adds missing and removes extras)
	err = syncSeasonTeams(ctx, queries, seasonID, req.Teams, now)
	if err != nil {
		return SeasonWithTeams{}, err
	}

	return getSeasonWithTeams(ctx, queries, seasonID)
}

func updateSeasonFields(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.SeasonRequest,
	competitionID, seasonID uuid.UUID,
	now time.Time,
) error {
	updateSeasonParams := db.UpdateSeasonParams{
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Rounds:        req.Rounds,
		UpdatedAt:     now,
		ID:            seasonID,
	}

	err := queries.UpdateSeason(ctx, updateSeasonParams)
	if err != nil {
		return errors.Wrap(err, "unable to update season")
	}
	return nil
}

func syncSeasonTeams(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	rawTeamIDs []uuid.UUID,
	now time.Time,
) error {
	requestedTeamIDs := dedupeUUIDs(rawTeamIDs)
	err := ensureTeamsExist(ctx, queries, requestedTeamIDs)
	if err != nil {
		return err
	}

	existingLinks, requestedSet, existingMap, err := buildSeasonTeamMaps(ctx, queries, seasonID, requestedTeamIDs)
	if err != nil {
		return err
	}

	err = ensureSeasonHasTeams(ctx, queries, seasonID, requestedTeamIDs, now, existingMap)
	if err != nil {
		return err
	}

	return removeExtraSeasonTeams(ctx, queries, seasonID, existingLinks, requestedSet)
}

func buildSeasonTeamMaps(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	requestedTeamIDs []uuid.UUID,
) ([]db.GetSeasonTeamsRow, map[uuid.UUID]struct{}, map[uuid.UUID]db.GetSeasonTeamsRow, error) {
	existingLinks, err := queries.GetSeasonTeams(ctx, seasonID)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "unable to get season's existing teams")
	}

	requestedSet := make(map[uuid.UUID]struct{}, len(requestedTeamIDs))
	for _, id := range requestedTeamIDs {
		requestedSet[id] = struct{}{}
	}

	existingMap := make(map[uuid.UUID]db.GetSeasonTeamsRow, len(existingLinks))
	for _, st := range existingLinks {
		existingMap[st.TeamID] = st
	}

	return existingLinks, requestedSet, existingMap, nil
}

func removeExtraSeasonTeams(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	existingLinks []db.GetSeasonTeamsRow,
	requestedSet map[uuid.UUID]struct{},
) error {
	for _, st := range existingLinks {
		_, keep := requestedSet[st.TeamID]
		if keep {
			continue
		}

		deleteParams := db.DeleteSeasonTeamParams{
			ID:        st.ID,
			DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
		}

		err := queries.DeleteSeasonTeam(ctx, deleteParams)
		if err != nil {
			return errors.Wrapf(err, "unable to remove team %s from season %s", st.TeamID.String(), seasonID.String())
		}
	}
	return nil
}

func DeleteSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) error {
	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		return deleteSeason(ctx, queries, competitionID, seasonID)
	})
	if err != nil {
		return errors.Wrap(err, "failed deleting season")
	}

	return nil
}

func deleteSeason(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) error {
	now := time.Now()

	err := softDeleteSeasonDependencies(ctx, queries, seasonID, now)
	if err != nil {
		return err
	}

	err = softDeleteSeason(ctx, queries, seasonID, now)
	if err != nil {
		return err
	}

	return nil
}

func softDeleteSeasonDependencies(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	now time.Time,
) error {
	// Delete games
	deleteGamesBySeasonIDParams := db.DeleteGamesBySeasonIDParams{
		DeletedAt: sql.NullTime{Time: now, Valid: true},
		SeasonID:  seasonID,
	}
	err := queries.DeleteGamesBySeasonID(ctx, deleteGamesBySeasonIDParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete games for season")
	}

	// Delete season teams
	deleteSeasonTeamsBySeasonIDParams := db.DeleteSeasonTeamsBySeasonIDParams{
		DeletedAt: sql.NullTime{Time: now, Valid: true},
		SeasonID:  seasonID,
	}
	err = queries.DeleteSeasonTeamsBySeasonID(ctx, deleteSeasonTeamsBySeasonIDParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete season teams for season")
	}

	return nil
}

func softDeleteSeason(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	now time.Time,
) error {
	deleteSeasonParams := db.DeleteSeasonParams{
		ID:        seasonID,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}
	err := queries.DeleteSeason(ctx, deleteSeasonParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete season")
	}

	return nil
}
