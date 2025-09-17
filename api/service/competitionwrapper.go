package service

import (
	"context"
	"strings"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
)

type DefaultCompetitionService struct {
	db db_handler.DB
}

func NewCompetitionService(db db_handler.DB) CompetitionService {
	return &DefaultCompetitionService{db: db}
}

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
	return GetCompetitions(ctx, s.db)
}

func (s *DefaultCompetitionService) Get(ctx context.Context, id uuid.UUID) (db.Competition, error) {
	return GetCompetition(ctx, s.db, id)
}

func (s *DefaultCompetitionService) Update(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	return UpdateCompetition(ctx, s.db, id, req)
}

func (s *DefaultCompetitionService) Delete(ctx context.Context, id uuid.UUID) error {
	return DeleteCompetition(ctx, s.db, id)
}
