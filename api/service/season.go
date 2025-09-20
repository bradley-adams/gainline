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

// SeasonService defines the contract for season-related operations.
type SeasonService interface {
	Create(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonWithTeams, error)
	GetAll(ctx context.Context, competitionID uuid.UUID) ([]SeasonWithTeams, error)
	Get(ctx context.Context, competitionID, seasonID uuid.UUID) (SeasonWithTeams, error)
	Update(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonWithTeams, error)
	Delete(ctx context.Context, seasonID uuid.UUID) error
}

// seasonService is the concrete implementation backed by db_handler.DB.
type seasonService struct {
	db db_handler.DB
}

func NewSeasonService(db db_handler.DB) SeasonService {
	return &seasonService{db: db}
}

func (s *seasonService) Create(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonWithTeams, error) {
	var season SeasonWithTeams
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		season, err = createSeason(ctx, queries, req, competitionID)
		return err
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed creating season")
	}
	return season, nil
}

func (s *seasonService) GetAll(ctx context.Context, competitionID uuid.UUID) ([]SeasonWithTeams, error) {
	var seasons []SeasonWithTeams
	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		seasons, err = getSeasons(ctx, queries, competitionID)
		return err
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting seasons")
	}
	return seasons, nil
}

func (s *seasonService) Get(ctx context.Context, competitionID, seasonID uuid.UUID) (SeasonWithTeams, error) {
	var season SeasonWithTeams
	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		season, err = getSeasonWithTeams(ctx, queries, seasonID)
		return err
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed getting season")
	}
	return season, nil
}

func (s *seasonService) Update(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonWithTeams, error) {
	var season SeasonWithTeams
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		season, txErr = updateSeason(ctx, queries, req, competitionID, seasonID)
		return txErr
	})
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "failed updating season")
	}
	return season, nil
}

func (s *seasonService) Delete(ctx context.Context, seasonID uuid.UUID) error {
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		return deleteSeason(ctx, queries, seasonID)
	})
	if err != nil {
		return errors.Wrap(err, "failed deleting season")
	}
	return nil
}

type SeasonWithTeams struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	StartDate     time.Time
	EndDate       time.Time
	Rounds        int32
	Teams         []db.Team
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     zero.Time
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

// --- Internal functions ---

func createSeason(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonWithTeams, error) {
	now := time.Now()
	seasonID := uuid.New()

	if err := insertSeason(ctx, queries, seasonID, competitionID, req, now); err != nil {
		return SeasonWithTeams{}, err
	}

	teamIDs := dedupeUUIDs(req.Teams)
	if err := ensureTeamsExist(ctx, queries, teamIDs); err != nil {
		return SeasonWithTeams{}, err
	}

	if err := ensureSeasonHasTeams(ctx, queries, seasonID, teamIDs, now, nil); err != nil {
		return SeasonWithTeams{}, err
	}

	return getSeasonWithTeams(ctx, queries, seasonID)
}

func insertSeason(ctx context.Context, queries db_handler.Queries, seasonID, competitionID uuid.UUID, req *api.SeasonRequest, now time.Time) error {
	params := db.CreateSeasonParams{
		ID:            seasonID,
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Rounds:        req.Rounds,
		CreatedAt:     now,
		UpdatedAt:     now,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	if err := queries.CreateSeason(ctx, params); err != nil {
		return errors.Wrap(err, "unable to create season")
	}
	return nil
}

func ensureTeamsExist(ctx context.Context, queries db_handler.Queries, teamIDs []uuid.UUID) error {
	for _, id := range teamIDs {
		if _, err := queries.GetTeam(ctx, id); err != nil {
			return errors.Wrapf(err, "unable to get team %s", id.String())
		}
	}
	return nil
}

func ensureSeasonHasTeams(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, teamIDs []uuid.UUID, now time.Time, existingMap map[uuid.UUID]db.GetSeasonTeamsRow) error {
	for _, teamID := range teamIDs {
		if existingMap != nil {
			if _, exists := existingMap[teamID]; exists {
				continue
			}
		}

		params := db.CreateSeasonTeamsParams{
			ID:        uuid.New(),
			SeasonID:  seasonID,
			TeamID:    teamID,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
		}
		if err := queries.CreateSeasonTeams(ctx, params); err != nil {
			return errors.Wrapf(err, "unable to add team %s to season %s", teamID, seasonID)
		}
	}
	return nil
}

func getSeasonWithTeams(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID) (SeasonWithTeams, error) {
	season, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return SeasonWithTeams{}, errors.Wrap(err, "unable to get season")
	}
	return buildSeasonWithTeams(ctx, queries, season)
}

func buildSeasonWithTeams(ctx context.Context, queries db_handler.Queries, season db.Season) (SeasonWithTeams, error) {
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

func dedupeUUIDs(ids []uuid.UUID) []uuid.UUID {
	set := make(map[uuid.UUID]struct{}, len(ids))
	for _, id := range ids {
		set[id] = struct{}{}
	}
	out := make([]uuid.UUID, 0, len(set))
	for id := range set {
		out = append(out, id)
	}
	return out
}

func getSeasons(ctx context.Context, queries db_handler.Queries, competitionID uuid.UUID) ([]SeasonWithTeams, error) {
	rawSeasons, err := queries.GetSeasons(ctx, competitionID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get seasons")
	}

	var seasons []SeasonWithTeams
	for _, s := range rawSeasons {
		swt, err := buildSeasonWithTeams(ctx, queries, s)
		if err != nil {
			return nil, err
		}
		seasons = append(seasons, swt)
	}
	return seasons, nil
}

func updateSeason(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonWithTeams, error) {
	now := time.Now()
	if err := updateSeasonFields(ctx, queries, req, competitionID, seasonID, now); err != nil {
		return SeasonWithTeams{}, err
	}
	if err := syncSeasonTeams(ctx, queries, seasonID, req.Teams, now); err != nil {
		return SeasonWithTeams{}, err
	}
	return getSeasonWithTeams(ctx, queries, seasonID)
}

func updateSeasonFields(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID, seasonID uuid.UUID, now time.Time) error {
	params := db.UpdateSeasonParams{
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Rounds:        req.Rounds,
		UpdatedAt:     now,
		ID:            seasonID,
	}
	if err := queries.UpdateSeason(ctx, params); err != nil {
		return errors.Wrap(err, "unable to update season")
	}
	return nil
}

func syncSeasonTeams(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, teamIDs []uuid.UUID, now time.Time) error {
	teamIDs = dedupeUUIDs(teamIDs)
	if err := ensureTeamsExist(ctx, queries, teamIDs); err != nil {
		return err
	}

	existingLinks, requestedSet, existingMap, err := buildSeasonTeamMaps(ctx, queries, seasonID, teamIDs)
	if err != nil {
		return err
	}

	if err := ensureSeasonHasTeams(ctx, queries, seasonID, teamIDs, now, existingMap); err != nil {
		return err
	}

	return removeExtraSeasonTeams(ctx, queries, seasonID, existingLinks, requestedSet)
}

func buildSeasonTeamMaps(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, requestedTeamIDs []uuid.UUID) ([]db.GetSeasonTeamsRow, map[uuid.UUID]struct{}, map[uuid.UUID]db.GetSeasonTeamsRow, error) {
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

func removeExtraSeasonTeams(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, existingLinks []db.GetSeasonTeamsRow, requestedSet map[uuid.UUID]struct{}) error {
	for _, st := range existingLinks {
		if _, keep := requestedSet[st.TeamID]; keep {
			continue
		}
		params := db.DeleteSeasonTeamParams{
			ID:        st.ID,
			DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
		}
		if err := queries.DeleteSeasonTeam(ctx, params); err != nil {
			return errors.Wrapf(err, "unable to remove team %s from season %s", st.TeamID, seasonID)
		}
	}
	return nil
}

func deleteSeason(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID) error {
	now := time.Now()
	if err := softDeleteSeasonDependencies(ctx, queries, seasonID, now); err != nil {
		return err
	}
	if err := softDeleteSeason(ctx, queries, seasonID, now); err != nil {
		return err
	}
	return nil
}

func softDeleteSeasonDependencies(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, now time.Time) error {
	if err := queries.DeleteGamesBySeasonID(ctx, db.DeleteGamesBySeasonIDParams{
		SeasonID:  seasonID,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}); err != nil {
		return errors.Wrap(err, "unable to delete games for season")
	}

	if err := queries.DeleteSeasonTeamsBySeasonID(ctx, db.DeleteSeasonTeamsBySeasonIDParams{
		SeasonID:  seasonID,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}); err != nil {
		return errors.Wrap(err, "unable to delete season teams for season")
	}

	return nil
}

func softDeleteSeason(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID, now time.Time) error {
	params := db.DeleteSeasonParams{
		ID:        seasonID,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}
	if err := queries.DeleteSeason(ctx, params); err != nil {
		return errors.Wrap(err, "unable to delete season")
	}
	return nil
}
