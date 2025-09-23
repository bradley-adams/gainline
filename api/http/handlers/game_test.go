package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/bradley-adams/gainline/db/db"
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

// Manual mock for GameService
type mockGameService struct {
	CreateFn func(ctx context.Context, req *api.GameRequest, season service.SeasonWithTeams) (db.Game, error)
	GetAllFn func(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error)
	GetFn    func(ctx context.Context, gameID uuid.UUID) (db.Game, error)
	UpdateFn func(ctx context.Context, req *api.GameRequest, gameID uuid.UUID, season service.SeasonWithTeams) (db.Game, error)
	DeleteFn func(ctx context.Context, gameID uuid.UUID) error
}

func (m *mockGameService) Create(ctx context.Context, req *api.GameRequest, season service.SeasonWithTeams) (db.Game, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, req, season)
	}
	return db.Game{}, nil
}
func (m *mockGameService) GetAll(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn(ctx, seasonID)
	}
	return nil, nil
}
func (m *mockGameService) Get(ctx context.Context, gameID uuid.UUID) (db.Game, error) {
	if m.GetFn != nil {
		return m.GetFn(ctx, gameID)
	}
	return db.Game{}, nil
}
func (m *mockGameService) Update(ctx context.Context, req *api.GameRequest, gameID uuid.UUID, season service.SeasonWithTeams) (db.Game, error) {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, req, gameID, season)
	}
	return db.Game{}, nil
}
func (m *mockGameService) Delete(ctx context.Context, gameID uuid.UUID) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, gameID)
	}
	return nil
}

var _ = Describe("game handlers", func() {
	var (
		router   *gin.Engine
		validate *validator.Validate
		logger   zerolog.Logger
		mockSvc  *mockGameService
		season   service.SeasonWithTeams
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)

		validate = validator.New()
		validation.Register(validate)
		api.Register(validate)

		logger = zerolog.Nop()

		mockSvc = &mockGameService{}
		router = gin.New()

		season = service.SeasonWithTeams{
			ID:        uuid.New(),
			Rounds:    3,
			StartDate: time.Now().Add(-time.Hour * 24),
			EndDate:   time.Now().Add(time.Hour * 24 * 30),
			Teams: []db.Team{
				{ID: uuid.New(), Name: "Home"},
				{ID: uuid.New(), Name: "Away"},
			},
		}

		// mount routes
		router.POST("/seasons/:seasonID/games", func(c *gin.Context) {
			c.Set("season", season)
			handleCreateGame(logger, validate, mockSvc)(c)
		})
		router.GET("/seasons/:seasonID/games", handleGetGames(logger, mockSvc))
		router.GET("/seasons/:seasonID/games/:gameID", handleGetGame(logger, mockSvc))
		router.PUT("/seasons/:seasonID/games/:gameID", func(c *gin.Context) {
			c.Set("season", season)
			handleUpdateGame(logger, mockSvc, validate)(c)
		})
		router.DELETE("/seasons/:seasonID/games/:gameID", handleDeleteGame(logger, mockSvc))
	})

	Describe("create game", func() {
		It("returns 201 for valid request", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.GameRequest, s service.SeasonWithTeams) (db.Game, error) {
				return db.Game{ID: uuid.New(), SeasonID: s.ID, Round: req.Round}, nil
			}

			reqBody := fmt.Sprintf(`{"round":1,"date":"%s","home_team_id":"%s","away_team_id":"%s"}`,
				time.Now().Format(time.RFC3339), season.Teams[0].ID, season.Teams[1].ID)

			req := httptest.NewRequest(http.MethodPost, "/seasons/"+season.ID.String()+"/games", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("returns 400 for invalid JSON", func() {
			req := httptest.NewRequest(http.MethodPost, "/seasons/"+season.ID.String()+"/games", bytes.NewBufferString(`{"round":`))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			mockSvc.CreateFn = func(ctx context.Context, req *api.GameRequest, s service.SeasonWithTeams) (db.Game, error) {
				return db.Game{}, fmt.Errorf("db failure")
			}

			reqBody := fmt.Sprintf(`{"round":1,"date":"%s","home_team_id":"%s","away_team_id":"%s"}`,
				time.Now().Format(time.RFC3339), season.Teams[0].ID, season.Teams[1].ID)

			req := httptest.NewRequest(http.MethodPost, "/seasons/"+season.ID.String()+"/games", bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get all games", func() {
		It("returns 200 with games", func() {
			mockSvc.GetAllFn = func(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error) {
				return []db.Game{{ID: uuid.New(), SeasonID: seasonID}}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/seasons/"+season.ID.String()+"/games", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 500 when service fails", func() {
			mockSvc.GetAllFn = func(ctx context.Context, seasonID uuid.UUID) ([]db.Game, error) {
				return nil, fmt.Errorf("db failure")
			}
			req := httptest.NewRequest(http.MethodGet, "/seasons/"+season.ID.String()+"/games", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("get game by ID", func() {
		It("returns 200 for valid game", func() {
			gameID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, gID uuid.UUID) (db.Game, error) {
				return db.Game{ID: gID, SeasonID: season.ID}, nil
			}
			req := httptest.NewRequest(http.MethodGet, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid UUID", func() {
			req := httptest.NewRequest(http.MethodGet, "/seasons/"+season.ID.String()+"/games/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			gameID := uuid.New()
			mockSvc.GetFn = func(ctx context.Context, gID uuid.UUID) (db.Game, error) {
				return db.Game{}, fmt.Errorf("db failure")
			}
			req := httptest.NewRequest(http.MethodGet, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("update game", func() {
		It("returns 200 for valid update", func() {
			gameID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.GameRequest, gID uuid.UUID, s service.SeasonWithTeams) (db.Game, error) {
				return db.Game{ID: gID, Round: req.Round, SeasonID: s.ID}, nil
			}

			reqBody := fmt.Sprintf(`{"round":2,"date":"%s","home_team_id":"%s","away_team_id":"%s"}`,
				time.Now().Format(time.RFC3339), season.Teams[0].ID, season.Teams[1].ID)

			req := httptest.NewRequest(http.MethodPut, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("returns 400 for invalid JSON", func() {
			gameID := uuid.New()
			req := httptest.NewRequest(http.MethodPut, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), bytes.NewBufferString(`{"round":`))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			gameID := uuid.New()
			mockSvc.UpdateFn = func(ctx context.Context, req *api.GameRequest, gID uuid.UUID, s service.SeasonWithTeams) (db.Game, error) {
				return db.Game{}, fmt.Errorf("db failure")
			}

			reqBody := fmt.Sprintf(`{"round":2,"date":"%s","home_team_id":"%s","away_team_id":"%s"}`,
				time.Now().Format(time.RFC3339), season.Teams[0].ID, season.Teams[1].ID)

			req := httptest.NewRequest(http.MethodPut, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), bytes.NewBufferString(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("delete game", func() {
		It("returns 204 for successful deletion", func() {
			gameID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, gID uuid.UUID) error { return nil }

			req := httptest.NewRequest(http.MethodDelete, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNoContent))
		})

		It("returns 400 for invalid UUID", func() {
			req := httptest.NewRequest(http.MethodDelete, "/seasons/"+season.ID.String()+"/games/invalid", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns 500 when service fails", func() {
			gameID := uuid.New()
			mockSvc.DeleteFn = func(ctx context.Context, gID uuid.UUID) error { return fmt.Errorf("db failure") }

			req := httptest.NewRequest(http.MethodDelete, "/seasons/"+season.ID.String()+"/games/"+gameID.String(), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
