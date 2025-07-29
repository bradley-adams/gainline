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

func CreateCompetition(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.CompetitionRequest,
) (db.Competition, error) {
	var competition db.Competition

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		competition, txErr = createCompetition(ctx, queries, req)
		return txErr
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
}

func createCompetition(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.CompetitionRequest,
) (db.Competition, error) {
	var competition db.Competition

	now := time.Now()
	createCompetitionParams := db.CreateCompetitionParams{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	err := queries.CreateCompetition(ctx, createCompetitionParams)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to create new competition")
	}

	competition, err = queries.GetCompetition(ctx, createCompetitionParams.ID)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to get new competition")
	}

	return competition, nil
}

func GetCompetitions(
	ctx context.Context,
	dbHandler db_handler.DB,
) ([]db.Competition, error) {
	var competitions []db.Competition

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		competitions, txErr = getCompetitions(ctx, queries)
		return txErr
	})
	if err != nil {
		return []db.Competition{}, err
	}

	return competitions, nil
}

func getCompetitions(
	ctx context.Context,
	queries db_handler.Queries,
) ([]db.Competition, error) {
	competitions, err := queries.GetCompetitions(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get competitions")
	}

	return competitions, nil
}

func GetCompetition(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
) (db.Competition, error) {
	var competition db.Competition

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		competition, txErr = getCompetition(ctx, queries, competitionID)
		return txErr
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
}

func getCompetition(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
) (db.Competition, error) {
	competitions, err := queries.GetCompetition(ctx, competitionID)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to get competition")
	}

	return competitions, nil
}

func UpdateCompetition(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
	req *api.CompetitionRequest,
) (db.Competition, error) {
	var competition db.Competition

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		competition, txErr = updateCompetition(ctx, queries, competitionID, req)
		return txErr
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
}

func updateCompetition(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
	req *api.CompetitionRequest,
) (db.Competition, error) {
	var competition db.Competition

	updateCompetitionParams := db.UpdateCompetitionParams{
		Name: req.Name,
		ID:   competitionID,
	}

	err := queries.UpdateCompetition(ctx, updateCompetitionParams)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to update competition")
	}

	competition, err = queries.GetCompetition(ctx, competitionID)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to get updated competition")
	}

	return competition, nil
}

func DeleteCompetition(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
) error {
	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		txErr := deleteCompetition(ctx, queries, competitionID)
		return txErr
	})
	if err != nil {
		return err
	}

	return nil
}

func deleteCompetition(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
) error {
	deleteCompetitionParams := db.DeleteCompetitionParams{
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:        competitionID,
	}

	err := queries.DeleteCompetition(ctx, deleteCompetitionParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete competition")
	}

	return nil
}
