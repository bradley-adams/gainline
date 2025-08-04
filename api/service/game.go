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

func CreateGame(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.GameRequest,
	seasonID uuid.UUID,
) (db.Game, error) {
	var game db.Game

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		game, txErr = createGame(ctx, queries, req, seasonID)
		return txErr
	})
	if err != nil {
		return db.Game{}, err
	}

	return game, nil
}

func createGame(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.GameRequest,
	seasonID uuid.UUID,
) (db.Game, error) {
	now := time.Now()

	// default status if not provided
	status := api.GameStatusScheduled
	if req.Status != "" {
		status = req.Status
	}

	createParams := db.CreateGameParams{
		ID:         uuid.New(),
		SeasonID:   seasonID,
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

func toNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func GetGames(
	ctx context.Context,
	dbHandler db_handler.DB,
	seasonID uuid.UUID,
) ([]db.Game, error) {
	var games []db.Game

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		games, err = queries.GetGames(ctx, seasonID)
		if err != nil {
			return errors.Wrap(err, "unable to get games")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return games, nil
}

func GetGame(
	ctx context.Context,
	dbHandler db_handler.DB,
	gameID uuid.UUID,
) (db.Game, error) {
	var game db.Game

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		game, err = queries.GetGame(ctx, gameID)
		if err != nil {
			return errors.Wrap(err, "unable to get game")
		}
		return nil
	})
	if err != nil {
		return db.Game{}, err
	}

	return game, nil
}

func UpdateGame(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.GameRequest,
	gameID uuid.UUID,
) (db.Game, error) {
	var game db.Game

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		game, txErr = updateGame(ctx, queries, req, gameID)
		return txErr
	})
	if err != nil {
		return db.Game{}, err
	}

	return game, nil
}

func updateGame(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.GameRequest,
	gameID uuid.UUID,
) (db.Game, error) {
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

func DeleteGame(
	ctx context.Context,
	dbHandler db_handler.DB,
	gameID uuid.UUID,
) error {
	return db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		return deleteGame(ctx, queries, gameID)
	})
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
