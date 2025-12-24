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
	Create(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonAggregate, error)
	GetAll(ctx context.Context, competitionID uuid.UUID) ([]SeasonAggregate, error)
	Get(ctx context.Context, competitionID, seasonID uuid.UUID) (SeasonAggregate, error)
	Update(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonAggregate, error)
	Delete(ctx context.Context, seasonID uuid.UUID) error
}

// seasonService is the concrete implementation backed by db_handler.DB.
type seasonService struct {
	db db_handler.DB
}

func NewSeasonService(db db_handler.DB) SeasonService {
	return &seasonService{db: db}
}

func (s *seasonService) Create(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonAggregate, error) {
	var season SeasonAggregate
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		season, err = createSeason(ctx, queries, req, competitionID)
		return err
	})
	if err != nil {
		return SeasonAggregate{}, err
	}
	return season, nil
}

func (s *seasonService) GetAll(ctx context.Context, competitionID uuid.UUID) ([]SeasonAggregate, error) {
	var seasons []SeasonAggregate
	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		seasons, err = getSeasons(ctx, queries, competitionID)
		return err
	})
	if err != nil {
		return nil, err
	}
	return seasons, nil
}

func (s *seasonService) Get(ctx context.Context, competitionID, seasonID uuid.UUID) (SeasonAggregate, error) {
	var season SeasonAggregate
	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		season, err = getSeason(ctx, queries, seasonID)
		return err
	})
	if err != nil {
		return SeasonAggregate{}, err
	}
	return season, nil
}

func (s *seasonService) Update(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonAggregate, error) {
	var season SeasonAggregate
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		season, txErr = updateSeason(ctx, queries, req, competitionID, seasonID)
		return txErr
	})
	if err != nil {
		return SeasonAggregate{}, err
	}
	return season, nil
}

func (s *seasonService) Delete(ctx context.Context, seasonID uuid.UUID) error {
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		return deleteSeason(ctx, queries, seasonID)
	})
	if err != nil {
		return err
	}
	return nil
}

type SeasonAggregate struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	StartDate     time.Time
	EndDate       time.Time
	Stages        []db.Stage
	Teams         []db.Team
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     zero.Time
}

func ToSeasonResponse(s SeasonAggregate) api.SeasonResponse {
	var teams []api.TeamResponse
	for _, team := range s.Teams {
		teams = append(teams, api.ToTeamResponse(team))
	}

	var stages []api.StageResponse
	for _, stage := range s.Stages {
		stages = append(stages, api.ToStageResponse(stage))
	}

	return api.SeasonResponse{
		ID:            s.ID,
		CompetitionID: s.CompetitionID,
		StartDate:     s.StartDate,
		EndDate:       s.EndDate,
		Stages:        stages,
		Teams:         teams,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		DeletedAt:     zero.TimeFrom(s.DeletedAt.Time),
	}
}

func createSeason(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID uuid.UUID) (SeasonAggregate, error) {
	now := time.Now()
	seasonID := uuid.New()

	if err := insertSeason(ctx, queries, seasonID, competitionID, req, now); err != nil {
		return SeasonAggregate{}, err
	}

	teamIDs := dedupeUUIDs(req.Teams)
	if err := ensureTeamsExist(ctx, queries, teamIDs); err != nil {
		return SeasonAggregate{}, errors.Wrap(err, "teams do not all exist")
	}

	if err := ensureSeasonHasTeams(ctx, queries, seasonID, teamIDs, now, nil); err != nil {
		return SeasonAggregate{}, err
	}

	if err := createStages(ctx, queries, seasonID, req.Stages, now); err != nil {
		return SeasonAggregate{}, err
	}

	return getSeason(ctx, queries, seasonID)
}

func insertSeason(ctx context.Context, queries db_handler.Queries, seasonID, competitionID uuid.UUID, req *api.SeasonRequest, now time.Time) error {
	params := db.CreateSeasonParams{
		ID:            seasonID,
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
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

func createStages(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	stages []api.StageRequest,
	now time.Time,
) error {
	for _, stage := range stages {
		if err := createStage(ctx, queries, seasonID, stage, now); err != nil {
			return err
		}
	}
	return nil
}

func getSeason(ctx context.Context, queries db_handler.Queries, seasonID uuid.UUID) (SeasonAggregate, error) {
	season, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return SeasonAggregate{}, errors.Wrap(err, "unable to get season")
	}

	return buildSeasonAggregate(ctx, queries, season)
}

func buildSeasonAggregate(
	ctx context.Context,
	queries db_handler.Queries,
	season db.Season,
) (SeasonAggregate, error) {

	teams, err := getSeasonTeams(ctx, queries, season.ID)
	if err != nil {
		return SeasonAggregate{}, err
	}

	stages, err := queries.GetStagesBySeasonID(ctx, season.ID)
	if err != nil {
		return SeasonAggregate{}, errors.Wrap(err, "unable to get season stages")
	}

	return SeasonAggregate{
		ID:            season.ID,
		CompetitionID: season.CompetitionID,
		StartDate:     season.StartDate,
		EndDate:       season.EndDate,
		Teams:         teams,
		Stages:        stages,
		CreatedAt:     season.CreatedAt,
		UpdatedAt:     season.UpdatedAt,
		DeletedAt:     zero.TimeFrom(season.DeletedAt.Time),
	}, nil
}

func getSeasonTeams(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
) ([]db.Team, error) {

	rows, err := queries.GetSeasonTeams(ctx, seasonID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get season teams")
	}

	teams := make([]db.Team, 0, len(rows))
	for _, row := range rows {
		team, err := queries.GetTeam(ctx, row.TeamID)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get team")
		}
		teams = append(teams, team)
	}

	return teams, nil
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

func getSeasons(ctx context.Context, queries db_handler.Queries, competitionID uuid.UUID) ([]SeasonAggregate, error) {
	seasons, err := queries.GetSeasons(ctx, competitionID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get seasons")
	}

	seasonAggregates := make([]SeasonAggregate, 0, len(seasons))
	for _, season := range seasons {
		seasonAggregate, err := buildSeasonAggregate(ctx, queries, season)
		if err != nil {
			return nil, err
		}
		seasonAggregates = append(seasonAggregates, seasonAggregate)
	}

	return seasonAggregates, nil
}

func updateSeason(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (SeasonAggregate, error) {
	now := time.Now()

	if err := updateSeasonFields(ctx, queries, req, competitionID, seasonID, now); err != nil {
		return SeasonAggregate{}, err
	}

	if err := syncSeasonTeams(ctx, queries, seasonID, req.Teams, now); err != nil {
		return SeasonAggregate{}, errors.Wrap(err, "unable to sync season teams")
	}

	if err := syncSeasonStages(ctx, queries, seasonID, req.Stages, now); err != nil {
		return SeasonAggregate{}, errors.Wrap(err, "unable to sync season stages")
	}

	return getSeason(ctx, queries, seasonID)
}

func updateSeasonFields(ctx context.Context, queries db_handler.Queries, req *api.SeasonRequest, competitionID, seasonID uuid.UUID, now time.Time) error {
	params := db.UpdateSeasonParams{
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		UpdatedAt:     now,
		ID:            seasonID,
	}
	if err := queries.UpdateSeason(ctx, params); err != nil {
		return errors.Wrap(err, "unable to update season")
	}
	return nil
}

func syncSeasonStages(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	stages []api.StageRequest,
	now time.Time,
) error {
	existingStages, err := queries.GetStagesBySeasonID(ctx, seasonID)
	if err != nil {
		return errors.Wrap(err, "unable to get existing stages")
	}

	existingMap := make(map[uuid.UUID]db.Stage, len(existingStages))
	for _, stage := range existingStages {
		existingMap[stage.ID] = stage
	}

	requestedSet := make(map[uuid.UUID]struct{})

	for _, stage := range stages {
		if stage.ID == nil {
			if err := createStage(ctx, queries, seasonID, stage, now); err != nil {
				return err
			}
			continue
		}

		stageID := *stage.ID
		requestedSet[stageID] = struct{}{}

		if _, exists := existingMap[stageID]; !exists {
			return errors.Errorf("unable to update stage %s for season %s", stageID, seasonID)
		}

		if err := updateStage(ctx, queries, stageID, stage, now); err != nil {
			return err
		}
	}

	return removeExtraSeasonStages(ctx, queries, existingStages, requestedSet, now)
}

func createStage(
	ctx context.Context,
	queries db_handler.Queries,
	seasonID uuid.UUID,
	stage api.StageRequest,
	now time.Time,
) error {
	params := db.CreateStageParams{
		ID:         uuid.New(),
		SeasonID:   seasonID,
		Name:       stage.Name,
		StageType:  db.StageType(stage.StageType),
		OrderIndex: stage.OrderIndex,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  sql.NullTime{Valid: false},
	}

	if err := queries.CreateStage(ctx, params); err != nil {
		return errors.Wrapf(err, "unable to create stage %q", stage.Name)
	}

	return nil
}

func updateStage(
	ctx context.Context,
	queries db_handler.Queries,
	stageID uuid.UUID,
	stage api.StageRequest,
	now time.Time,
) error {
	params := db.UpdateStageParams{
		ID:         stageID,
		Name:       stage.Name,
		StageType:  db.StageType(stage.StageType),
		OrderIndex: stage.OrderIndex,
		UpdatedAt:  now,
	}

	if err := queries.UpdateStage(ctx, params); err != nil {
		return errors.Wrapf(err, "unable to update stage %s", stageID)
	}

	return nil
}

func removeExtraSeasonStages(
	ctx context.Context,
	queries db_handler.Queries,
	existingStages []db.Stage,
	requestedSet map[uuid.UUID]struct{},
	now time.Time,
) error {
	for _, stage := range existingStages {
		if _, keep := requestedSet[stage.ID]; keep {
			continue
		}

		params := db.DeleteStageParams{
			ID:        stage.ID,
			DeletedAt: sql.NullTime{Time: now, Valid: true},
		}

		if err := queries.DeleteStage(ctx, params); err != nil {
			return errors.Wrapf(err, "unable to delete stage %s", stage.ID)
		}
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

	if err := queries.DeleteStagesBySeasonID(ctx, db.DeleteStagesBySeasonIDParams{
		SeasonID:  seasonID,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}); err != nil {
		return errors.Wrap(err, "unable to delete stages for season")
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
