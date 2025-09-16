package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// Manual mock for CompetitionService
type mockCompetitionService struct {
	CreateFn func(ctx context.Context, dbHandlerDB db_handler.DB, req *api.CompetitionRequest) (db.Competition, error)
}

func (m *mockCompetitionService) Create(ctx context.Context, dbHandlerDB db_handler.DB, req *api.CompetitionRequest) (db.Competition, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, dbHandlerDB, req)
	}
	return db.Competition{}, nil
}

func (m *mockCompetitionService) GetAll(ctx context.Context, dbHandlerDB db_handler.DB) ([]db.Competition, error) {
	return nil, nil
}

func (m *mockCompetitionService) Get(ctx context.Context, dbHandlerDB db_handler.DB, id uuid.UUID) (db.Competition, error) {
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Update(ctx context.Context, dbHandlerDB db_handler.DB, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Delete(ctx context.Context, dbHandlerDB db_handler.DB, id uuid.UUID) error {
	return nil
}

func TestHandleCreateCompetition(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	// Validator and logger
	validate := validator.New()
	_ = validate.RegisterValidation("entity_name", func(fl validator.FieldLevel) bool {
		return fl.Field().String() != ""
	})
	logger := zerolog.Nop()

	// Manual mock service
	mockSvc := &mockCompetitionService{
		CreateFn: func(ctx context.Context, dbHandlerDB db_handler.DB, req *api.CompetitionRequest) (db.Competition, error) {
			return db.Competition{
				ID:   uuid.New(),
				Name: req.Name,
			}, nil
		},
	}

	// Setup Gin router with the handler
	router := gin.New()
	router.POST("/competitions", handleCreateCompetition(logger, nil, validate, mockSvc))

	// Prepare a valid JSON request
	reqBody := `{"name":"Test Competition"}`
	req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}
