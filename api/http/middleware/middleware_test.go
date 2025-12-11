package middleware

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"

	"github.com/bradley-adams/gainline/db/db"
	mock_db "github.com/bradley-adams/gainline/db/db_handler/mock"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/service"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "middleware Suite")
}

var _ = Describe("middleware", func() {
	var (
		ctrl        *gomock.Controller
		mockDB      *mock_db.MockDB
		mockQueries *mock_db.MockQueries

		logBuffer      *bytes.Buffer
		logger         zerolog.Logger
		router         *gin.Engine
		createRecorder func() *httptest.ResponseRecorder

		seasonService service.SeasonService
		gameService   service.GameService
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)

		logBuffer = new(bytes.Buffer)
		logger = zerolog.New(io.MultiWriter(GinkgoWriter, logBuffer)).With().Str("testing", "testing").Logger()
		router = gin.Default()

		seasonService = service.NewSeasonService(mockDB)
		gameService = service.NewGameService(mockDB)

		createRecorder = func() *httptest.ResponseRecorder {
			return httptest.NewRecorder()
		}

		router.Use(CompetitionStructureValidator(logger, seasonService, gameService))

		router.GET("/test/competitions/:competitionID/seasons/:seasonID/games/:gameID", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

	})

	validCompetitionID := uuid.MustParse("58068a80-5b92-4cd4-b16c-fd869758e9e4")
	invalidCompetitionID := uuid.MustParse("19f0facb-ed06-4fba-80ad-6db58729956e")

	validSeasonID := uuid.MustParse("1de75aa2-cde6-45de-a9a0-100afac15e80")
	invalidSeasonID := uuid.MustParse("dd385348-d031-438e-bc26-2a70f6c04323")

	validSeasonTeamID := uuid.MustParse("eeeeeeee-eeee-4eee-8eee-eeeeeeeeeeee")
	validSeasonTeamID2 := uuid.MustParse("ffffffff-ffff-4fff-8fff-ffffffffffff")

	validTeamID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	validTeamID2 := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	validGameID := uuid.MustParse("68bf7fb7-89cb-4c71-8e80-c34b8b7fdcde")

	validStageID := uuid.MustParse("33333333-3333-4333-8333-333333333333")

	validHomeTeamID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	validAwayTeamID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	validTimeNow := time.Now()

	validSeasonFromDB := db.Season{
		ID:            validSeasonID,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow,
		EndDate:       validTimeNow.AddDate(0, 5, 0),
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonTeamFromDB := db.GetSeasonTeamsRow{
		ID:        validSeasonTeamID,
		TeamID:    validTeamID,
		SeasonID:  validSeasonID,
		CreatedAt: validTimeNow,
		UpdatedAt: validTimeNow,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonTeamFromDB2 := db.GetSeasonTeamsRow{
		ID:        validSeasonTeamID2,
		TeamID:    validTeamID2,
		SeasonID:  validSeasonID,
		CreatedAt: validTimeNow,
		UpdatedAt: validTimeNow,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonTeamsFromDB := []db.GetSeasonTeamsRow{
		validSeasonTeamFromDB,
		validSeasonTeamFromDB2,
	}

	validTeamFromDB := db.Team{
		ID:           validTeamID,
		Name:         "Test Team",
		Abbreviation: "TT",
		Location:     "Testville",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validTeamFromDB2 := db.Team{
		ID:           validTeamID2,
		Name:         "Test Team2",
		Abbreviation: "TT2",
		Location:     "Testville 22",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validGameFromDB := db.Game{
		ID:         validGameID,
		SeasonID:   validSeasonID,
		StageID:    validStageID,
		Date:       validTimeNow,
		HomeTeamID: validHomeTeamID,
		AwayTeamID: validAwayTeamID,
		HomeScore:  sql.NullInt32{Int32: 2, Valid: true},
		AwayScore:  sql.NullInt32{Int32: 1, Valid: true},
		Status:     db.GameStatus(api.GameStatusPlaying),
		CreatedAt:  validTimeNow,
		UpdatedAt:  validTimeNow,
		DeletedAt:  sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validTestError := errors.New("a valid testing error")

	Describe("CompetitionStructureValidator", func() {
		It("should verify season and game belong to the given competition", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)

			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validGameFromDB, nil)

			recorder := createRecorder()
			req, _ := http.NewRequest("GET", "/test/competitions/"+validCompetitionID.String()+"/seasons/"+validSeasonID.String()+"/games/"+validGameID.String(), nil)
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("success"))
		})

		It("should return error when GetSeason returns an error", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Season{}, validTestError)

			recorder := createRecorder()
			req, _ := http.NewRequest("GET", "/test/competitions/"+validCompetitionID.String()+"/seasons/"+validSeasonID.String()+"/games/"+validGameID.String(), nil)
			router.ServeHTTP(recorder, req)

			logContent := logBuffer.String()
			Expect(logContent).To(ContainSubstring(
				"unable to get season: a valid testing error",
			))

			Expect(recorder.Code).To(Equal(http.StatusForbidden))
			Expect(recorder.Body.String()).To(ContainSubstring("Season does not belong to competition"))
		})

		It("should return error when GetGame returns an error", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)

			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Game{}, validTestError)

			recorder := createRecorder()
			req, _ := http.NewRequest("GET", "/test/competitions/"+validCompetitionID.String()+"/seasons/"+validSeasonID.String()+"/games/"+validGameID.String(), nil)
			router.ServeHTTP(recorder, req)

			logContent := logBuffer.String()
			Expect(logContent).To(ContainSubstring(
				"unable to get game: a valid testing error",
			))

			Expect(recorder.Code).To(Equal(http.StatusForbidden))
			Expect(recorder.Body.String()).To(ContainSubstring("Game does not belong to season"))
		})

		It("should return error if season does not belong to competition", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)

			recorder := createRecorder()
			req, _ := http.NewRequest("GET", "/test/competitions/"+invalidCompetitionID.String()+"/seasons/"+validSeasonID.String()+"/games/"+validGameID.String(), nil)
			router.ServeHTTP(recorder, req)

			logContent := logBuffer.String()
			Expect(logContent).To(ContainSubstring("season does not belong to competition"))

			Expect(recorder.Code).To(Equal(http.StatusForbidden))
			Expect(recorder.Body.String()).To(ContainSubstring("Season does not belong to competition"))
		})

		It("should return error if game does not belong to season", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)

			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validGameFromDB, nil)

			recorder := createRecorder()
			req, _ := http.NewRequest("GET", "/test/competitions/"+validCompetitionID.String()+"/seasons/"+invalidSeasonID.String()+"/games/"+validGameID.String(), nil)
			router.ServeHTTP(recorder, req)

			logContent := logBuffer.String()
			Expect(logContent).To(ContainSubstring("unauthorized: game does not belong to season"))

			Expect(recorder.Code).To(Equal(http.StatusForbidden))
			Expect(recorder.Body.String()).To(ContainSubstring("Game does not belong to season"))
		})
	})
})
