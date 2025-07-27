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
		return db.Competition{}, err
	}

	competition, err = queries.GetCompetition(ctx, createCompetitionParams.ID)
	if err != nil {
		return db.Competition{}, errors.Wrap(err, "unable to get new customer")
	}

	return competition, nil
}
