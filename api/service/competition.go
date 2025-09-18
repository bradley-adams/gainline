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

// CompetitionService defines the contract for competition-related operations.
type CompetitionService interface {
	Create(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error)
	GetAll(ctx context.Context) ([]db.Competition, error)
	Get(ctx context.Context, id uuid.UUID) (db.Competition, error)
	Update(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// DefaultCompetitionService is the concrete implementation backed by db_handler.DB.
type DefaultCompetitionService struct {
	db db_handler.DB
}

// NewCompetitionService constructs a new CompetitionService.
func NewCompetitionService(db db_handler.DB) CompetitionService {
	return &DefaultCompetitionService{db: db}
}

// --- Service methods ---

func (s *DefaultCompetitionService) Create(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error) {
	req.Name = strings.TrimSpace(req.Name)

	var competition db.Competition
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var err error
		competition, err = createCompetition(ctx, queries, req)
		return err
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
}

func (s *DefaultCompetitionService) GetAll(ctx context.Context) ([]db.Competition, error) {
	var competitions []db.Competition

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
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

func (s *DefaultCompetitionService) Get(ctx context.Context, competitionID uuid.UUID) (db.Competition, error) {
	var competition db.Competition

	err := db_handler.Run(ctx, s.db, func(queries db_handler.Queries) error {
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

func (s *DefaultCompetitionService) Update(ctx context.Context, competitionID uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	req.Name = strings.TrimSpace(req.Name)

	var competition db.Competition
	err := db_handler.RunInTransaction(ctx, s.db, func(queries db_handler.Queries) error {
		var txErr error
		competition, txErr = updateCompetition(ctx, queries, competitionID, req)
		return txErr
	})
	if err != nil {
		return db.Competition{}, err
	}

	return competition, nil
}

func (s *DefaultCompetitionService) Delete(ctx context.Context, competitionID uuid.UUID) error {
	return db_handler.RunInTransaction(ctx, s.db, func(q db_handler.Queries) error {
		return deleteCompetition(ctx, q, competitionID)
	})
}

// --- Private helpers ---

func createCompetition(ctx context.Context, queries db_handler.Queries, req *api.CompetitionRequest) (db.Competition, error) {
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

func updateCompetition(ctx context.Context, queries db_handler.Queries, competitionID uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
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

func deleteCompetition(ctx context.Context, queries db_handler.Queries, competitionID uuid.UUID) error {
	now := time.Now()

	deleteGamesByCompetitionIDParams := db.DeleteGamesByCompetitionIDParams{
		DeletedAt:     sql.NullTime{Time: now, Valid: true},
		CompetitionID: competitionID,
	}
	if err := queries.DeleteGamesByCompetitionID(ctx, deleteGamesByCompetitionIDParams); err != nil {
		return errors.Wrap(err, "unable to delete games for competition")
	}

	deleteSeasonsByCompetitionIDParams := db.DeleteSeasonsByCompetitionIDParams{
		DeletedAt:     sql.NullTime{Time: now, Valid: true},
		CompetitionID: competitionID,
	}
	if err := queries.DeleteSeasonsByCompetitionID(ctx, deleteSeasonsByCompetitionIDParams); err != nil {
		return errors.Wrap(err, "unable to delete seasons for competition")
	}

	deleteCompetitionParams := db.DeleteCompetitionParams{
		DeletedAt: sql.NullTime{Time: now, Valid: true},
		ID:        competitionID,
	}
	if err := queries.DeleteCompetition(ctx, deleteCompetitionParams); err != nil {
		return errors.Wrap(err, "unable to delete competition")
	}

	return nil
}
