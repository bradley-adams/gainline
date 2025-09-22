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

// TeamService defines the contract for team-related operations.
type TeamService interface {
	Create(ctx context.Context, req *api.TeamRequest) (db.Team, error)
	GetAll(ctx context.Context) ([]db.Team, error)
	Get(ctx context.Context, teamID uuid.UUID) (db.Team, error)
	Update(ctx context.Context, req *api.TeamRequest, teamID uuid.UUID) (db.Team, error)
	Delete(ctx context.Context, teamID uuid.UUID) error
}

// teamService is the concrete implementation backed by db_handler.DB.
type teamService struct {
	db db_handler.DB
}

// NewTeamService returns a new TeamService backed by db_handler.DB.
func NewTeamService(db db_handler.DB) TeamService {
	return &teamService{db: db}
}

func (s *teamService) Create(ctx context.Context, req *api.TeamRequest) (db.Team, error) {
	var team db.Team

	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		team, txErr = createTeam(ctx, queries, req)
		return txErr
	})
	if err != nil {
		return db.Team{}, err
	}

	return team, nil
}

func (s *teamService) GetAll(ctx context.Context) ([]db.Team, error) {
	var teams []db.Team

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		teams, err = queries.GetTeams(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting teams")
	}

	return teams, nil
}

func (s *teamService) Get(ctx context.Context, teamID uuid.UUID) (db.Team, error) {
	var team db.Team

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		team, err = queries.GetTeam(ctx, teamID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return db.Team{}, errors.Wrap(err, "failed to get team")
	}

	return team, nil
}

func (s *teamService) Update(ctx context.Context, req *api.TeamRequest, teamID uuid.UUID) (db.Team, error) {
	var team db.Team

	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		team, txErr = updateTeam(ctx, queries, req, teamID)
		return txErr
	})
	if err != nil {
		return db.Team{}, err
	}

	return team, nil
}

func (s *teamService) Delete(ctx context.Context, teamID uuid.UUID) error {
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		return deleteTeam(ctx, queries, teamID)
	})
	if err != nil {
		return err
	}
	return nil
}

// --- Internal helpers ---

func createTeam(
	ctx context.Context,
	queries db_handler.Queries,
	req *api.TeamRequest,
) (db.Team, error) {
	now := time.Now()
	params := db.CreateTeamParams{
		ID:           uuid.New(),
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
		Location:     req.Location,
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	if err := queries.CreateTeam(ctx, params); err != nil {
		return db.Team{}, errors.Wrap(err, "unable to create new team")
	}

	team, err := queries.GetTeam(ctx, params.ID)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to get new team")
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
	params := db.UpdateTeamParams{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
		Location:     req.Location,
		UpdatedAt:    now,
		ID:           teamID,
	}

	if err := queries.UpdateTeam(ctx, params); err != nil {
		return db.Team{}, errors.Wrap(err, "unable to update team")
	}

	updatedTeam, err := queries.GetTeam(ctx, teamID)
	if err != nil {
		return db.Team{}, errors.Wrap(err, "unable to get updated team")
	}

	return updatedTeam, nil
}

func deleteTeam(
	ctx context.Context,
	queries db_handler.Queries,
	teamID uuid.UUID,
) error {
	params := db.DeleteTeamParams{
		ID:        teamID,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if err := queries.DeleteTeam(ctx, params); err != nil {
		return errors.Wrap(err, "unable to delete team")
	}

	return nil
}
