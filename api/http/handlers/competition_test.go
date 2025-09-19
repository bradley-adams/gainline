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
	GetAllFn func(ctx context.Context) ([]db.Competition, error)
	GetFn    func(ctx context.Context, id uuid.UUID) (db.Competition, error)
	UpdateFn func(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error)
	DeleteFn func(ctx context.Context, id uuid.UUID) error
}

func (m *mockCompetitionService) Create(ctx context.Context, req *api.CompetitionRequest) (db.Competition, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, req)
	}
	return db.Competition{}, nil
}

func (m *mockCompetitionService) GetAll(ctx context.Context) ([]db.Competition, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn(ctx)
	}
	return nil, nil
}

func (m *mockCompetitionService) Get(ctx context.Context, id uuid.UUID) (db.Competition, error) {
	if m.GetFn != nil {
		return m.GetFn(ctx, id)
	}
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Update(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, id, req)
	}
	return db.Competition{}, nil
}

func (m *mockCompetitionService) Delete(ctx context.Context, id uuid.UUID) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

var _ = Describe("competition handlers", func() {
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
		router.GET("/competitions", handleGetCompetitions(logger, mockSvc))
		router.GET("/competitions/:competitionID", handleGetCompetition(logger, mockSvc))
		router.PUT("/competitions/:competitionID", handleUpdateCompetition(logger, validate, mockSvc))
		router.DELETE("/competitions/:competitionID", handleDeleteCompetition(logger, mockSvc))
	})

	Describe("create", func() {
		It("returns 201 for valid competition", func() {
			reqBody := `{"name":"Test Competition"}`
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("returns 400 for invalid JSON", func() {
			reqBody := `{"name":`
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 400 for validation failure", func() {
			reqBody := `{"name":""}`
			req := httptest.NewRequest(http.MethodPost, "/competitions", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
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

	Describe("get all competitions", func() {
		It("returns 200 and list of competitions", func() {
			mockSvc.GetAllFn = func(ctx context.Context) ([]db.Competition, error) {
				return []db.Competition{{ID: uuid.New(), Name: "Comp1"}}, nil
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 500 when service fails", func() {
			mockSvc.GetAllFn = func(ctx context.Context) ([]db.Competition, error) {
				return nil, fmt.Errorf("db failure")
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get one competition", func() {
		It("returns 200 for valid ID", func() {
			compID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, id uuid.UUID) (db.Competition, error) {
				return db.Competition{ID: id, Name: "Comp1"}, nil
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions/"+compID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid ID format", func() {
			req := httptest.NewRequest(http.MethodGet, "/competitions/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, id uuid.UUID) (db.Competition, error) {
				return db.Competition{}, fmt.Errorf("db failure")
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions/"+compID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("update competition", func() {
		It("returns 200 for valid update", func() {
			compID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
				return db.Competition{ID: id, Name: req.Name}, nil
			}

			reqBody := `{"name":"Updated"}`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid JSON", func() {
			compID := uuid.New()
			reqBody := `{"name":`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 400 for validation failure", func() {
			compID := uuid.New()
			reqBody := `{"name":""}`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, id uuid.UUID, req *api.CompetitionRequest) (db.Competition, error) {
				return db.Competition{}, fmt.Errorf("db failure")
			}

			reqBody := `{"name":"Fail"}`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("delete competition", func() {
		It("returns 204 for successful deletion", func() {
			compID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, id uuid.UUID) error { return nil }

			req := httptest.NewRequest(http.MethodDelete, "/competitions/"+compID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNoContent))
		})

		It("returns 400 for invalid ID format", func() {
			req := httptest.NewRequest(http.MethodDelete, "/competitions/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, id uuid.UUID) error { return fmt.Errorf("db failure") }

			req := httptest.NewRequest(http.MethodDelete, "/competitions/"+compID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
