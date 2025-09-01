package service

import (
	"context"
	"database/sql"
	"strings"
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
	req.Name = strings.TrimSpace(req.Name)

	var competition db.Competition
	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		competition, err = createCompetition(ctx, queries, req)
		return err
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

	competition, err := queries.GetCompetition(ctx, createCompetitionParams.ID)
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

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		competitions, err = queries.GetCompetitions(ctx)
		if err != nil {
			return errors.Wrap(err, "unable to get competitions")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

func GetCompetition(
	ctx context.Context,
	dbHandler db_handler.DB,
	competitionID uuid.UUID,
) (db.Competition, error) {
	var competition db.Competition

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		competition, err = queries.GetCompetition(ctx, competitionID)
		if err != nil {
			return errors.Wrap(err, "unable to get competition")
		}
		return nil
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
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
	updateCompetitionParams := db.UpdateCompetitionParams{
		Name: req.Name,
		ID:   competitionID,
	}

	err := queries.UpdateCompetition(ctx, updateCompetitionParams)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to update competition")
	}

	competition, err := queries.GetCompetition(ctx, competitionID)
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
	return db_handler.RunInTransaction(ctx, dbHandler, func(q db_handler.Queries) error {
		return deleteCompetition(ctx, q, competitionID)
	})
}

func deleteCompetition(
	ctx context.Context,
	queries db_handler.Queries,
	competitionID uuid.UUID,
) error {
	now := time.Now()

	deleteGamesByCompetitionIDParams := db.DeleteGamesByCompetitionIDParams{
		DeletedAt:     sql.NullTime{Time: now, Valid: true},
		CompetitionID: competitionID,
	}

	err := queries.DeleteGamesByCompetitionID(ctx, deleteGamesByCompetitionIDParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete games for competition")
	}

	deleteSeasonsByCompetitionIDParams := db.DeleteSeasonsByCompetitionIDParams{
		DeletedAt:     sql.NullTime{Time: now, Valid: true},
		CompetitionID: competitionID,
	}

	err = queries.DeleteSeasonsByCompetitionID(ctx, deleteSeasonsByCompetitionIDParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete seasons for competition")
	}

	deleteCompetitionParams := db.DeleteCompetitionParams{
		DeletedAt: sql.NullTime{Time: now, Valid: true},
		ID:        competitionID,
	}

	err = queries.DeleteCompetition(ctx, deleteCompetitionParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete competition")
	}

	return nil
}
