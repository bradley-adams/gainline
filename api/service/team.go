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

func CreateTeam(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.TeamRequest,
) (db.Team, error) {
	var team db.Team

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		team, err = createTeam(ctx, queries, req)
		return err
	})
	if err != nil {
		return db.Team{}, err
	}

	return team, nil
}

func createTeam(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.TeamRequest,
) (db.Team, error) {
	now := time.Now()
	createTeamParams := db.CreateTeamParams{
		ID:           uuid.New(),
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
		Location:     req.Location,
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	err := queries.CreateTeam(ctx, createTeamParams)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to create new team")
	}

	team, err := queries.GetTeam(ctx, createTeamParams.ID)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to get new team")
	}

	return team, nil
}

func GetTeams(
	ctx context.Context,
	dbHandler db_handler.DB,
) ([]db.Team, error) {
	var teams []db.Team

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		teams, err = queries.GetTeams(ctx)
		if err != nil {
			return errors.Wrap(err, "unable to get teams")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func GetTeam(
	ctx context.Context,
	dbHandler db_handler.DB,
	teamID uuid.UUID,
) (db.Team, error) {
	var team db.Team

	err := db_handler.Run(ctx, dbHandler, func(queries db_handler.Queries) error {
		var err error
		team, err = queries.GetTeam(ctx, teamID)
		if err != nil {
			return errors.Wrap(err, "unable to get team")
		}
		return nil
	})
	if err != nil {
		return db.Team{}, err
	}

	return team, nil
}

func UpdateTeam(
	ctx context.Context,
	dbHandler db_handler.DB,
	req *api.TeamRequest,
	teamID uuid.UUID,
) (db.Team, error) {
	var team db.Team

	err := db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		var txErr error
		team, txErr = updateTeam(ctx, queries, req, teamID)
		return txErr
	})
	if err != nil {
		return db.Team{}, err
	}

	return team, nil
}

func updateTeam(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.TeamRequest,
	teamID uuid.UUID,
) (db.Team, error) {
	now := time.Now()
	updateTeamParams := db.UpdateTeamParams{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
		Location:     req.Location,
		UpdatedAt:    now,
		ID:           teamID,
	}

	err := queries.UpdateTeam(ctx, updateTeamParams)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to update team")
	}

	updatedTeam, err := queries.GetTeam(ctx, teamID)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to get updated team")
	}

	return updatedTeam, nil
}

func DeleteTeam(
	ctx context.Context,
	dbHandler db_handler.DB,
	teamID uuid.UUID,
) error {
	return db_handler.RunInTransaction(ctx, dbHandler, func(queries db_handler.Queries) error {
		return deleteTeam(ctx, queries, teamID)
	})
}

func deleteTeam(
	ctx context.Context,
	queries db_handler.Queries,
	teamID uuid.UUID,
) error {
	deleteTeamParams := db.DeleteTeamParams{
		ID:        teamID,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := queries.DeleteTeam(ctx, deleteTeamParams)
	if err != nil {
		return errors.Wrap(err, "unable to delete team")
	}

	return nil
}
