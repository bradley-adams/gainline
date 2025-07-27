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
	CreateCompetition(ctx context.Context, arg db.CreateCompetitionParams) error
	GetCompetition(ctx context.Context, id uuid.UUID) (db.Competition, error)
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

func handleTxError(err error, db DB, tx *sql.Tx) error {
	rollbackErr := db.Rollback(tx)
	if rollbackErr != nil {
		return errors.Wrap(err, rollbackErr.Error())
	}
	return err
}
