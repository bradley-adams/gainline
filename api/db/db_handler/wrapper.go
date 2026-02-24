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
	GetCompetitionsPaginated(ctx context.Context, arg db.GetCompetitionsPaginatedParams) ([]db.Competition, error)
	CountCompetitions(ctx context.Context) (int64, error)
	UpdateCompetition(ctx context.Context, arg db.UpdateCompetitionParams) error
	DeleteCompetition(ctx context.Context, arg db.DeleteCompetitionParams) error

	//Season
	CreateSeason(ctx context.Context, arg db.CreateSeasonParams) error
	GetSeason(ctx context.Context, id uuid.UUID) (db.Season, error)
	GetSeasons(ctx context.Context, competitionID uuid.UUID) ([]db.Season, error)
	UpdateSeason(ctx context.Context, arg db.UpdateSeasonParams) error
	DeleteSeason(ctx context.Context, arg db.DeleteSeasonParams) error
	DeleteSeasonsByCompetitionID(ctx context.Context, arg db.DeleteSeasonsByCompetitionIDParams) error

	//Stage
	CreateStage(ctx context.Context, arg db.CreateStageParams) error
	GetStagesBySeasonID(ctx context.Context, id uuid.UUID) ([]db.Stage, error)
	UpdateStage(ctx context.Context, arg db.UpdateStageParams) error
	DeleteStage(ctx context.Context, arg db.DeleteStageParams) error
	DeleteStagesBySeasonID(ctx context.Context, arg db.DeleteStagesBySeasonIDParams) error
	DeleteStagesByCompetitionID(ctx context.Context, arg db.DeleteStagesByCompetitionIDParams) error

	//Team
	CreateTeam(ctx context.Context, arg db.CreateTeamParams) error
	GetTeam(ctx context.Context, id uuid.UUID) (db.Team, error)
	GetTeams(ctx context.Context) ([]db.Team, error)
	UpdateTeam(ctx context.Context, arg db.UpdateTeamParams) error
	DeleteTeam(ctx context.Context, arg db.DeleteTeamParams) error

	//SeasonTeams
	CreateSeasonTeams(ctx context.Context, arg db.CreateSeasonTeamsParams) error
	GetSeasonTeams(ctx context.Context, seasonID uuid.UUID) ([]db.GetSeasonTeamsRow, error)
	DeleteSeasonTeam(ctx context.Context, arg db.DeleteSeasonTeamParams) error
	DeleteSeasonTeamsBySeasonID(ctx context.Context, arg db.DeleteSeasonTeamsBySeasonIDParams) error
	DeleteSeasonTeamsByCompetitionID(ctx context.Context, arg db.DeleteSeasonTeamsByCompetitionIDParams) error

	//Game
	CreateGame(ctx context.Context, arg db.CreateGameParams) error
	GetGame(ctx context.Context, id uuid.UUID) (db.Game, error)
	GetGames(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error)
	UpdateGame(ctx context.Context, arg db.UpdateGameParams) error
	DeleteGame(ctx context.Context, arg db.DeleteGameParams) error
	DeleteGamesByCompetitionID(ctx context.Context, arg db.DeleteGamesByCompetitionIDParams) error
	DeleteGamesBySeasonID(ctx context.Context, arg db.DeleteGamesBySeasonIDParams) error
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
