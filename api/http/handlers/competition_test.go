package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Manual mock for CompetitionService
type mockCompetitionService struct {
	CreateFn func(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error)
}

func (m *mockCompetitionService) Create(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, req)
	}
	return db.Competition{}, nil
}

func (m *mockCompetitionService) GetAll(ctx context.Context) ([]db.Competition, error) {
	return nil, nil
}

func (m *mockCompetitionService) Get(ctx context.Context, id uuid.UUID) (db.Competition, error) {
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Update(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Delete(ctx context.Context, id uuid.UUID) error { return nil }

var _ = Describe("competition", func() {
	var (
		router   *gin.Engine
		validate *validator.Validate
		logger   zerolog.Logger
		mockSvc  *mockCompetitionService
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)

		validate = validator.New()
		_ = validate.RegisterValidation("entity_name", func(fl validator.FieldLevel) bool {
			return fl.Field().String() != ""
		})
		logger = zerolog.Nop()

		mockSvc = &mockCompetitionService{
			CreateFn: func(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error) {
				return db.Competition{
					ID:   uuid.New(),
					Name: req.Name,
				}, nil
			},
		}

		router = gin.New()
		router.POST("/competitions", handleCreateCompetition(logger, validate, mockSvc))
	})

	Describe("create", func() {
		It("should return 201 when creating a valid competition", func() {
			reqBody := `{"name":"Test Competition"}`
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("should return 400 for invalid JSON", func() {
			reqBody := `{"name":` // malformed JSON
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return 400 for validation failure", func() {
			reqBody := `{"name":""}` // fails entity_name validation
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return 500 when service fails", func() {
			// simulate error
			mockSvc.CreateFn = func(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error) {
				return db.Competition{}, fmt.Errorf("db failure")
			}

			reqBody := `{"name":"Test Competition"}`
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
