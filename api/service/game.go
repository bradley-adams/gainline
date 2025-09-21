package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GameService defines the contract for game-related operations.
type GameService interface {
	Create(ctx context.Context, req *api.GameRequest, season SeasonWithTeams) (db.Game, error)
	GetAll(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error)
	Get(ctx context.Context, gameID uuid.UUID) (db.Game, error)
	Update(ctx context.Context, req *api.GameRequest, gameID uuid.UUID, season SeasonWithTeams) (db.Game, error)
	Delete(ctx context.Context, gameID uuid.UUID) error
}

// gameService is the concrete implementation backed by db_handler.DB.
type gameService struct {
	db db_handler.DB
}

func NewGameService(db db_handler.DB) GameService {
	return &gameService{db: db}
}

func (s *gameService) Create(ctx context.Context, req *api.GameRequest, season SeasonWithTeams) (db.Game, error) {
	var game db.Game

	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		game, txErr = createGame(ctx, queries, req, season)
		return txErr
	})
	if err != nil {
		return db.Game{}, err
	}

	return game, nil
}

func (s *gameService) GetAll(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error) {
	var games []db.Game

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		games, err = queries.GetGames(ctx, seasonID)
		return err
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting games")
	}

	return games, nil
}

func (s *gameService) Get(ctx context.Context, gameID uuid.UUID) (db.Game, error) {
	var game db.Game

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		game, err = queries.GetGame(ctx, gameID)
		return err
	})
	if err != nil {
		return db.Game{}, errors.Wrap(err, "failed to get game")
	}

	return game, nil
}

func (s *gameService) Update(ctx context.Context, req *api.GameRequest, gameID uuid.UUID, season SeasonWithTeams) (db.Game, error) {
	var game db.Game

	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		game, txErr = updateGame(ctx, queries, req, gameID, season)
		return txErr
	})
	if err != nil {
		return db.Game{}, err
	}

	return game, nil
}

func (s *gameService) Delete(ctx context.Context, gameID uuid.UUID) error {
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		return deleteGame(ctx, queries, gameID)
	})
	if err != nil {
		return err
	}
	return nil
}

func createGame(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.GameRequest,
	season SeasonWithTeams,
) (db.Game, error) {
	if err := validateGameRequest(req, season); err != nil {
		return db.Game{}, err
	}

	now := time.Now()

	// default status if not provided
	status := api.GameStatusScheduled
	if req.Status != "" {
		status = req.Status
	}

	createParams := db.CreateGameParams{
		ID:         uuid.New(),
		SeasonID:   season.ID,
		Round:      req.Round,
		Date:       req.Date,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		HomeScore:  toNullInt32(req.HomeScore),
		AwayScore:  toNullInt32(req.AwayScore),
		Status:     db.GameStatus(status),
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  sql.NullTime{Time: time.Time{}, Valid: false},
	}

	if err := queries.CreateGame(ctx, createParams); err != nil {
		return db.Game{}, errors.Wrap(err, "unable to create new game")
	}

	game, err := queries.GetGame(ctx, createParams.ID)
	if err != nil {
		return db.Game{}, errors.Wrap(err, "unable to get new game")
	}

	return game, nil
}

func updateGame(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.GameRequest,
	gameID uuid.UUID,
	season SeasonWithTeams,
) (db.Game, error) {
	if err := validateGameRequest(req, season); err != nil {
		return db.Game{}, err
	}

	now := time.Now()

	// default status if not provided
	status := api.GameStatusScheduled
	if req.Status != "" {
		status = req.Status
	}

	updateParams := db.UpdateGameParams{
		Round:      req.Round,
		Date:       req.Date,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		HomeScore:  toNullInt32(req.HomeScore),
		AwayScore:  toNullInt32(req.AwayScore),
		Status:     db.GameStatus(status),
		UpdatedAt:  now,
		ID:         gameID,
	}

	if err := queries.UpdateGame(ctx, updateParams); err != nil {
		return db.Game{}, errors.Wrap(err, "unable to update game")
	}

	updatedGame, err := queries.GetGame(ctx, gameID)
	if err != nil {
		return db.Game{}, errors.Wrap(err, "unable to get updated game")
	}

	return updatedGame, nil
}

func deleteGame(
	ctx context.Context,
	queries db_handler.Queries,
	gameID uuid.UUID,
) error {
	deleteParams := db.DeleteGameParams{
		ID:        gameID,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if err := queries.DeleteGame(ctx, deleteParams); err != nil {
		return errors.Wrap(err, "unable to delete game")
	}

	return nil
}

func validateGameRequest(req *api.GameRequest, season SeasonWithTeams) error {
	teamIDs := make(map[uuid.UUID]struct{}, len(season.Teams))
	for _, t := range season.Teams {
		teamIDs[t.ID] = struct{}{}
	}

	// Round in season bounds
	if req.Round < 1 || req.Round > season.Rounds {
		return errors.Errorf("round %d is out of bounds (1-%d)", req.Round, season.Rounds)
	}

	// Team in season
	if _, ok := teamIDs[req.HomeTeamID]; !ok {
		return errors.New("home team not in season")
	}
	if _, ok := teamIDs[req.AwayTeamID]; !ok {
		return errors.New("away team not in season")
	}

	// Date in season bounds
	if req.Date.Before(season.StartDate) || req.Date.After(season.EndDate) {
		return errors.Errorf("game date %s outside season bounds (%s - %s)", req.Date, season.StartDate, season.EndDate)
	}

	return nil
}

func toNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}
