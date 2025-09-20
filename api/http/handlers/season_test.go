package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Manual mock for SeasonService
type mockSeasonService struct {
	CreateFn func(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (service.SeasonWithTeams, error)
	GetAllFn func(ctx context.Context, competitionID uuid.UUID) ([]service.SeasonWithTeams, error)
	GetFn    func(ctx context.Context, competitionID, seasonID uuid.UUID) (service.SeasonWithTeams, error)
	UpdateFn func(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (service.SeasonWithTeams, error)
	DeleteFn func(ctx context.Context, seasonID uuid.UUID) error
}

func (m *mockSeasonService) Create(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (service.SeasonWithTeams, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, req, competitionID)
	}
	return service.SeasonWithTeams{}, nil
}

func (m *mockSeasonService) GetAll(ctx context.Context, competitionID uuid.UUID) ([]service.SeasonWithTeams, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn(ctx, competitionID)
	}
	return nil, nil
}

func (m *mockSeasonService) Get(ctx context.Context, competitionID, seasonID uuid.UUID) (service.SeasonWithTeams, error) {
	if m.GetFn != nil {
		return m.GetFn(ctx, competitionID, seasonID)
	}
	return service.SeasonWithTeams{}, nil
}

func (m *mockSeasonService) Update(ctx context.Context, req *api.SeasonRequest, competitionID, seasonID uuid.UUID) (service.SeasonWithTeams, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, req, competitionID, seasonID)
	}
	return service.SeasonWithTeams{}, nil
}

func (m *mockSeasonService) Delete(ctx context.Context, seasonID uuid.UUID) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, seasonID)
	}
	return nil
}

var _ = Describe("season handlers", func() {
	var (
		router   *gin.Engine
		validate *validator.Validate
		logger   zerolog.Logger
		mockSvc  *mockSeasonService
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)

		validate = validator.New()
		validation.Register(validate)
		logger = zerolog.Nop()

		mockSvc = &mockSeasonService{}
		router = gin.New()

		router.POST("/competitions/:competitionID/seasons", handleCreateSeason(logger, validate, mockSvc))
		router.GET("/competitions/:competitionID/seasons", handleGetSeasons(logger, mockSvc))
		router.GET("/competitions/:competitionID/seasons/:seasonID", handleGetSeason(logger))
		router.PUT("/competitions/:competitionID/seasons/:seasonID", handleUpdateSeason(logger, validate, mockSvc))
		router.DELETE("/competitions/:competitionID/seasons/:seasonID", handleDeleteSeason(logger, mockSvc))
	})

	Describe("create season", func() {
		It("returns 201 for valid request", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (service.SeasonWithTeams, error) {
				return service.SeasonWithTeams{ID: uuid.New(), Rounds: req.Rounds}, nil
			}

			compID := uuid.New()
			reqBody := `{
				"start_date":"2025-01-01T00:00:00Z",
				"end_date":"2025-12-31T23:59:59Z",
				"rounds":10,
				"teams":["013952a5-87e1-4d26-a312-09b2aff54241","7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"]
			}`
			req := httptest.NewRequest(http.MethodPost, "/competitions/"+compID.String()+"/seasons", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("returns 400 for invalid JSON", func() {
			compID := uuid.New()
			reqBody := `{"rounds":`
			req := httptest.NewRequest(http.MethodPost, "/competitions/"+compID.String()+"/seasons", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 400 for validation failure", func() {
			compID := uuid.New()
			reqBody := `{"rounds":0}`
			req := httptest.NewRequest(http.MethodPost, "/competitions/"+compID.String()+"/seasons", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.SeasonRequest, competitionID uuid.UUID) (service.SeasonWithTeams, error) {
				return service.SeasonWithTeams{}, fmt.Errorf("db failure")
			}

			compID := uuid.New()
			reqBody := `{
				"start_date":"2025-01-01T00:00:00Z",
				"end_date":"2025-12-31T23:59:59Z",
				"rounds":10,
				"teams":["013952a5-87e1-4d26-a312-09b2aff54241","7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"]
			}`
			req := httptest.NewRequest(http.MethodPost, "/competitions/"+compID.String()+"/seasons", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get all seasons", func() {
		It("returns 200 and list of seasons", func() {
			compID := uuid.New()
			mockSvc.GetAllFn = func(ctx context.Context, competitionID uuid.UUID) ([]service.SeasonWithTeams, error) {
				return []service.SeasonWithTeams{{ID: uuid.New(), Rounds: 5}}, nil
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions/"+compID.String()+"/seasons", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			mockSvc.GetAllFn = func(ctx context.Context, competitionID uuid.UUID) ([]service.SeasonWithTeams, error) {
				return nil, fmt.Errorf("db failure")
			}

			req := httptest.NewRequest(http.MethodGet, "/competitions/"+compID.String()+"/seasons", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("update season", func() {
		It("returns 200 for valid update", func() {
			compID := uuid.New()
			seasonID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.SeasonRequest, competitionID, sID uuid.UUID) (service.SeasonWithTeams, error) {
				return service.SeasonWithTeams{ID: sID, Rounds: req.Rounds}, nil
			}

			reqBody := `{
				"start_date":"2025-01-01T00:00:00Z",
				"end_date":"2025-12-31T23:59:59Z",
				"rounds":12,
				"teams":["013952a5-87e1-4d26-a312-09b2aff54241","7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"]
			}`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String()+"/seasons/"+seasonID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid JSON", func() {
			compID := uuid.New()
			seasonID := uuid.New()
			reqBody := `{"rounds":`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String()+"/seasons/"+seasonID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			seasonID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.SeasonRequest, competitionID, sID uuid.UUID) (service.SeasonWithTeams, error) {
				return service.SeasonWithTeams{}, fmt.Errorf("db failure")
			}

			reqBody := `{
				"start_date":"2025-01-01T00:00:00Z",
				"end_date":"2025-12-31T23:59:59Z",
				"rounds":12,
				"teams":["013952a5-87e1-4d26-a312-09b2aff54241","7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"]
			}`
			req := httptest.NewRequest(http.MethodPut, "/competitions/"+compID.String()+"/seasons/"+seasonID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("delete season", func() {
		It("returns 204 for successful deletion", func() {
			compID := uuid.New()
			seasonID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, seasonID uuid.UUID) error { return nil }

			req := httptest.NewRequest(http.MethodDelete, "/competitions/"+compID.String()+"/seasons/"+seasonID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNoContent))
		})

		It("returns 400 for invalid ID format", func() {
			compID := uuid.New()
			req := httptest.NewRequest(http.MethodDelete, "/competitions/"+compID.String()+"/seasons/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			compID := uuid.New()
			seasonID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, seasonID uuid.UUID) error { return fmt.Errorf("db failure") }

			req := httptest.NewRequest(http.MethodDelete, "/competitions/"+compID.String()+"/seasons/"+seasonID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
