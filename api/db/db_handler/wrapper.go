package db_handler

import (
	"context"
	"database/sql"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error

	New(db db.DBTX) Queries

	HealthCheck() error

	db.DBTX
}

type Queries interface {
	//Competition
	CreateCompetition(ctx context.Context, arg db.CreateCompetitionParams) error
	GetCompetition(ctx context.Context, id uuid.UUID) (db.Competition, error)
	GetCompetitions(ctx context.Context) ([]db.Competition, error)
	UpdateCompetition(ctx context.Context, arg db.UpdateCompetitionParams) error
	DeleteCompetition(ctx context.Context, arg db.DeleteCompetitionParams) error

	//Season
	CreateSeason(ctx context.Context, arg db.CreateSeasonParams) error
	GetSeason(ctx context.Context, id uuid.UUID) (db.Season, error)
	GetSeasons(ctx context.Context, competitionID uuid.UUID) ([]db.Season, error)
	UpdateSeason(ctx context.Context, arg db.UpdateSeasonParams) error
	DeleteSeason(ctx context.Context, arg db.DeleteSeasonParams) error

	//Team
	CreateTeam(ctx context.Context, arg db.CreateTeamParams) error
	GetTeam(ctx context.Context, id uuid.UUID) (db.Team, error)
	GetTeams(ctx context.Context) ([]db.Team, error)
	UpdateTeam(ctx context.Context, arg db.UpdateTeamParams) error
	DeleteTeam(ctx context.Context, arg db.DeleteTeamParams) error
}

type DBWrapper struct {
	*sql.DB
}

func (d DBWrapper) HealthCheck() error {
	return d.Ping()
}

func (d DBWrapper) New(dbtx db.DBTX) Queries {
	return db.New(dbtx)
}

func (d DBWrapper) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (d DBWrapper) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func RunInTransaction(
	ctx context.Context,
	db DB,
	f func(q Queries) error,
) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(db.New(tx))
	if err != nil {
		return handleTxError(err, db, tx)
	}

	err = db.Commit(tx)
	if err != nil {
		return err
	}

	return nil
}

func Run(
	ctx context.Context,
	db DB,
	f func(q Queries) error,
) error {
	return f(db.New(db))
}

func handleTxError(err error, db DB, tx *sql.Tx) error {
	rollbackErr := db.Rollback(tx)
	if rollbackErr != nil {
		return errors.Wrap(err, rollbackErr.Error())
	}
	return err
}
