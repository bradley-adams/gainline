package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Manual mock for TeamService
type mockTeamService struct {
	CreateFn func(ctx context.Context, req *api.TeamRequest) (db.Team, error)
	GetAllFn func(ctx context.Context) ([]db.Team, error)
	GetFn    func(ctx context.Context, teamID uuid.UUID) (db.Team, error)
	UpdateFn func(ctx context.Context, req *api.TeamRequest, teamID uuid.UUID) (db.Team, error)
	DeleteFn func(ctx context.Context, teamID uuid.UUID) error
}

func (m *mockTeamService) Create(ctx context.Context, req *api.TeamRequest) (db.Team, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, req)
	}
	return db.Team{}, nil
}

func (m *mockTeamService) GetAll(ctx context.Context) ([]db.Team, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn(ctx)
	}
	return nil, nil
}

func (m *mockTeamService) Get(ctx context.Context, teamID uuid.UUID) (db.Team, error) {
	if m.GetFn != nil {
		return m.GetFn(ctx, teamID)
	}
	return db.Team{}, nil
}

func (m *mockTeamService) Update(ctx context.Context, req *api.TeamRequest, teamID uuid.UUID) (db.Team, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, req, teamID)
	}
	return db.Team{}, nil
}

func (m *mockTeamService) Delete(ctx context.Context, teamID uuid.UUID) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, teamID)
	}
	return nil
}

var _ = Describe("team handlers", func() {
	var (
		router   *gin.Engine
		validate *validator.Validate
		logger   zerolog.Logger
		mockSvc  *mockTeamService
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)

		validate = validator.New()
		validation.Register(validate)
		logger = zerolog.Nop()

		mockSvc = &mockTeamService{}
		router = gin.New()

		router.POST("/teams", handleCreateTeam(logger, validate, mockSvc))
		router.GET("/teams", handleGetTeams(logger, mockSvc))
		router.GET("/teams/:teamID", handleGetTeam(logger, mockSvc))
		router.PUT("/teams/:teamID", handleUpdateTeam(logger, validate, mockSvc))
		router.DELETE("/teams/:teamID", handleDeleteTeam(logger, mockSvc))
	})

	Describe("create team", func() {
		It("returns 201 for valid request", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.TeamRequest) (db.Team, error) {
				return db.Team{
					ID:           uuid.New(),
					Name:         req.Name,
					Abbreviation: req.Abbreviation,
					Location:     req.Location,
				}, nil
			}

			reqBody := `{"name":"Highlanders","abbreviation":"HIG","location":"Dunedin"}`
			req := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("returns 400 for invalid JSON", func() {
			req := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBufferString(`{"name":`))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 400 for validation failure", func() {
			req := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBufferString(`{"name":""}`))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.TeamRequest) (db.Team, error) {
				return db.Team{}, fmt.Errorf("db failure")
			}

			reqBody := `{"name":"Crusaders","abbreviation":"CRU","location":"Christchurch"}`
			req := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get all teams", func() {
		It("returns 200 and list of teams", func() {
			mockSvc.GetAllFn = func(ctx context.Context) ([]db.Team, error) {
				return []db.Team{
					{ID: uuid.New(), Name: "Hurricanes", Abbreviation: "HUR", Location: "Wellington"},
				}, nil
			}

			req := httptest.NewRequest(http.MethodGet, "/teams", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 500 when service fails", func() {
			mockSvc.GetAllFn = func(ctx context.Context) ([]db.Team, error) {
				return nil, fmt.Errorf("db failure")
			}

			req := httptest.NewRequest(http.MethodGet, "/teams", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get team by ID", func() {
		It("returns 200 for valid team", func() {
			teamID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, tID uuid.UUID) (db.Team, error) {
				return db.Team{ID: tID, Name: "Chiefs", Abbreviation: "CHI", Location: "Hamilton"}, nil
			}

			req := httptest.NewRequest(http.MethodGet, "/teams/"+teamID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid UUID", func() {
			req := httptest.NewRequest(http.MethodGet, "/teams/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			teamID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, tID uuid.UUID) (db.Team, error) {
				return db.Team{}, fmt.Errorf("db failure")
			}

			req := httptest.NewRequest(http.MethodGet, "/teams/"+teamID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("update team", func() {
		It("returns 200 for valid update", func() {
			teamID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.TeamRequest, tID uuid.UUID) (db.Team, error) {
				return db.Team{ID: tID, Name: req.Name, Abbreviation: req.Abbreviation, Location: req.Location}, nil
			}

			reqBody := `{"name":"Blues","abbreviation":"BLU","location":"Auckland"}`
			req := httptest.NewRequest(http.MethodPut, "/teams/"+teamID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid JSON", func() {
			teamID := uuid.New()
			req := httptest.NewRequest(http.MethodPut, "/teams/"+teamID.String(), bytes.NewBufferString(`{"name":`))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			teamID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.TeamRequest, tID uuid.UUID) (db.Team, error) {
				return db.Team{}, fmt.Errorf("db failure")
			}

			reqBody := `{"name":"Moana Pasifika","abbreviation":"MOA","location":"Auckland"}`
			req := httptest.NewRequest(http.MethodPut, "/teams/"+teamID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("delete team", func() {
		It("returns 204 for successful deletion", func() {
			teamID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, tID uuid.UUID) error { return nil }

			req := httptest.NewRequest(http.MethodDelete, "/teams/"+teamID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNoContent))
		})

		It("returns 400 for invalid UUID", func() {
			req := httptest.NewRequest(http.MethodDelete, "/teams/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			teamID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, tID uuid.UUID) error { return fmt.Errorf("db failure") }

			req := httptest.NewRequest(http.MethodDelete, "/teams/"+teamID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
