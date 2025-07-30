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

func CreateSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
) (db.Season, error) {
	var season db.Season

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		season, err = createSeason(ctx, queries, req, competitionID)
		return err
	})
	if err != nil {
		return db.Season{}, err
	}

	return season, nil
}

func createSeason(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
) (db.Season, error) {
	now := time.Now()
	createSeasonParams := db.CreateSeasonParams{
		ID:            uuid.New(),
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
		return db.Season{}, errors.Wrap(err, "unable to create new season")
	}

	season, err := queries.GetSeason(ctx, createSeasonParams.ID)
	if err != nil {
		return db.Season{}, errors.Wrap(err, "unable to get new season")
	}

	return season, nil
}

func GetSeasons(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
) ([]db.Season, error) {
	var seasons []db.Season

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		seasons, err = queries.GetSeasons(ctx, competitionID)
		if err != nil {
			return errors.Wrap(err, "unable to get seasons")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return seasons, nil
}

func GetSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (db.Season, error) {
	var season db.Season

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		season, err := queries.GetSeason(ctx, seasonID)
		if err != nil {
			return errors.Wrap(err, "unable to get season")
		}
		if season.CompetitionID != competitionID {
			return errors.New("season does not belong to the specified competition")
		}
		return nil
	})
	if err != nil {
		return db.Season{}, err
	}

	return season, nil
}

func UpdateSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (db.Season, error) {
	var season db.Season

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		season, txErr = updateSeason(ctx, queries, req, competitionID, seasonID)
		return txErr
	})
	if err != nil {
		return db.Season{}, err
	}

	return season, nil
}

func updateSeason(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.SeasonRequest,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) (db.Season, error) {
	season, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return db.Season{}, errors.Wrap(err, "unable to get season for update")
	}

	if season.CompetitionID != competitionID {
		return db.Season{}, errors.New("season does not belong to the specified competition")
	}

	now := time.Now()
	updateSeasonParams := db.UpdateSeasonParams{
		CompetitionID: competitionID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Rounds:        req.Rounds,
		UpdatedAt:     now,
		ID:            seasonID,
	}

	err = queries.UpdateSeason(ctx, updateSeasonParams)
	if err != nil {
		return db.Season{}, errors.Wrap(err, "unable to update season")
	}

	updatedSeason, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return db.Season{}, errors.Wrap(err, "unable to get updated season")
	}

	return updatedSeason, nil
}

func DeleteSeason(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) error {
	return db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		return deleteSeason(ctx, queries, competitionID, seasonID)
	})
}

func deleteSeason(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
	seasonID uuid.UUID,
) error {
	season, err := queries.GetSeason(ctx, seasonID)
	if err != nil {
		return errors.Wrap(err, "unable to get season for deletion")
	}

	if season.CompetitionID != competitionID {
		return errors.New("season does not belong to the specified competition")
	}

	deleteSeasonParams := db.DeleteSeasonParams{
		ID:        season.ID,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err = queries.DeleteSeason(ctx, deleteSeasonParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete season")
	}

	return nil
}
