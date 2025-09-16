package service

import (
	"context"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
)

// DefaultCompetitionService wraps the existing functions to implement CompetitionService interface
type DefaultCompetitionService struct{}

func (s DefaultCompetitionService) Create(ctx context.Context, db db_handler.DB, req *api.CompetitionRequest) (db.Competition, error) {
	return CreateCompetition(ctx, db, req)
}

func (s DefaultCompetitionService) GetAll(ctx context.Context, db db_handler.DB) ([]db.Competition, error) {
	return GetCompetitions(ctx, db)
}

func (s DefaultCompetitionService) Get(ctx context.Context, db db_handler.DB, id uuid.UUID) (db.Competition, error) {
	return GetCompetition(ctx, db, id)
}

func (s DefaultCompetitionService) Update(ctx context.Context, db db_handler.DB, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	return UpdateCompetition(ctx, db, id, req)
}

func (s DefaultCompetitionService) Delete(ctx context.Context, db db_handler.DB, id uuid.UUID) error {
	return DeleteCompetition(ctx, db, id)
}
