package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/bradley-adams/gainline/db/db"
	mock_db "github.com/bradley-adams/gainline/db/db_handler/mock"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"go.uber.org/mock/gomock"
)

var _ = Describe("season", func() {
	var ctrl *gomock.Controller
	var mockDB *mock_db.MockDB
	var mockQueries *mock_db.MockQueries
	var svc SeasonService

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)
		svc = NewSeasonService(mockDB)
	})

	validSeasonID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	validSeasonID2 := uuid.MustParse("bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbbb")

	validCompetitionID := uuid.MustParse("cccccccc-cccc-4ccc-8ccc-cccccccccccc")

	validSeasonTeamID := uuid.MustParse("eeeeeeee-eeee-4eee-8eee-eeeeeeeeeeee")
	validSeasonTeamID2 := uuid.MustParse("ffffffff-ffff-4fff-8fff-ffffffffffff")
	validSeasonTeamID3 := uuid.MustParse("99999999-9999-4999-8999-999999999999")

	validTeamID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	validTeamID2 := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	validTeamID3 := uuid.MustParse("33333333-3333-4333-8333-333333333333")

	validStageID := uuid.MustParse("44444444-4444-4444-8444-444444444444")
	validStageID2 := uuid.MustParse("55555555-5555-4555-8555-555555555555")
	validStageID3 := uuid.MustParse("66666666-6666-4666-8666-666666666666")

	validTeamIDs := []uuid.UUID{validTeamID, validTeamID2}

	validTimeNow := time.Now()

	validStage := api.StageRequest{
		Name:       "Regular Season",
		StageType:  api.StageTypeRegular,
		OrderIndex: 1,
	}

	validCreateSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Stages:    []api.StageRequest{validStage},
		Teams:     validTeamIDs,
	}

	validStage2 := api.StageRequest{
		Name:       "Regular Season 2",
		StageType:  api.StageTypeRegular,
		OrderIndex: 2,
	}

	validStageUpdate := api.StageRequest{
		ID:         &validStageID2,
		Name:       "Semi Finals",
		StageType:  api.StageTypeFinals,
		OrderIndex: 3,
	}

	validUpdateSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Stages:    []api.StageRequest{validStage2, validStageUpdate},
		Teams:     validTeamIDs,
	}

	validStage3 := api.StageRequest{
		ID:         &validStageID3,
		Name:       "Regular Season 3",
		StageType:  api.StageTypeRegular,
		OrderIndex: 2,
	}

	invalidUpdateSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Stages:    []api.StageRequest{validStage3, validStageUpdate},
		Teams:     validTeamIDs,
	}

	var validNilSeason db.Season

	validSeasonFromDB := db.Season{
		ID:            validSeasonID,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow,
		EndDate:       validTimeNow.AddDate(0, 5, 0),
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonFromDB2 := db.Season{
		ID:            validSeasonID2,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow,
		EndDate:       validTimeNow.AddDate(0, 5, 0),
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonsFromDB := []db.Season{
		validSeasonFromDB,
		validSeasonFromDB2,
	}

	validRegularStageFromDB := db.Stage{
		ID:         validStageID,
		SeasonID:   validSeasonID,
		Name:       "Regular",
		OrderIndex: 1,
		CreatedAt:  validTimeNow,
		UpdatedAt:  validTimeNow,
	}

	validFinalStagesFromDB := db.Stage{
		ID:         validStageID2,
		SeasonID:   validSeasonID,
		Name:       "Semi Finals",
		OrderIndex: 2,
		CreatedAt:  validTimeNow,
		UpdatedAt:  validTimeNow,
	}

	validFinalStagesFromDB2 := db.Stage{
		ID:         uuid.MustParse("d40b4854-06e8-43ec-be4c-748f9a58c16f"),
		SeasonID:   validSeasonID,
		Name:       "Final",
		OrderIndex: 3,
		CreatedAt:  validTimeNow,
		UpdatedAt:  validTimeNow,
	}

	validStagesFromDB := []db.Stage{
		validRegularStageFromDB,
		validFinalStagesFromDB,
	}

	validStagesFromDB2 := []db.Stage{
		validRegularStageFromDB,
		validFinalStagesFromDB2,
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

	validTeamFromDB3 := db.Team{
		ID:           validTeamID3,
		Name:         "Test Team3",
		Abbreviation: "TT3",
		Location:     "Testville 33",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
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

	validSeasonTeamFromDB3 := db.GetSeasonTeamsRow{
		ID:        validSeasonTeamID3,
		TeamID:    validTeamID3,
		SeasonID:  validSeasonID,
		CreatedAt: validTimeNow,
		UpdatedAt: validTimeNow,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonTeamsFromDB := []db.GetSeasonTeamsRow{
		validSeasonTeamFromDB,
		validSeasonTeamFromDB2,
	}

	validSeasonTeamsFromDB2 := []db.GetSeasonTeamsRow{
		validSeasonTeamFromDB,
		validSeasonTeamFromDB3,
	}

	var validNilSeasonWithTeams SeasonAggregate
	var validNilSeasonsWithTeams []SeasonAggregate

	validSeasonWithTeams := SeasonAggregate{
		ID:            validSeasonFromDB.ID,
		CompetitionID: validSeasonFromDB.CompetitionID,
		StartDate:     validSeasonFromDB.StartDate,
		EndDate:       validSeasonFromDB.EndDate,
		Teams:         []db.Team{},
		CreatedAt:     validSeasonFromDB.CreatedAt,
		UpdatedAt:     validSeasonFromDB.UpdatedAt,
		DeletedAt:     zero.TimeFrom(validSeasonFromDB.DeletedAt.Time),
	}

	validSeasonWithTeams2 := SeasonAggregate{
		ID:            validSeasonFromDB2.ID,
		CompetitionID: validSeasonFromDB2.CompetitionID,
		StartDate:     validSeasonFromDB2.StartDate,
		EndDate:       validSeasonFromDB2.EndDate,
		Teams:         []db.Team{},
		CreatedAt:     validSeasonFromDB2.CreatedAt,
		UpdatedAt:     validSeasonFromDB2.UpdatedAt,
		DeletedAt:     zero.TimeFrom(validSeasonFromDB2.DeletedAt.Time),
	}

	validSeasonResponse := ToSeasonResponse(validSeasonWithTeams)

	validSeasonResponse2 := ToSeasonResponse(validSeasonWithTeams2)

	validSeasonsResponse := []api.SeasonResponse{
		validSeasonResponse,
		validSeasonResponse2,
	}

	validTestError := errors.New("a valid testing error")

	Describe("CreateSeason", func() {
		It("should create a new season without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Teams).To(HaveLen(2))
			Expect(season.Teams).To(ConsistOf(
				validTeamFromDB,
				validTeamFromDB2,
			))
			Expect(season.Stages).To(HaveLen(2))
			Expect(season.Stages).To(ConsistOf(
				validRegularStageFromDB,
				validFinalStagesFromDB,
			))
			Expect(season.CreatedAt).To(Equal(validSeasonResponse.CreatedAt))
			Expect(season.UpdatedAt).To(Equal(validSeasonResponse.UpdatedAt))
			Expect(season.DeletedAt.Time).To(Equal(validSeasonResponse.DeletedAt.Time))
		})

		It("should return a formatted error when transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("a valid testing error"))
		})

		It("should rollback and return a formatted error when inserting season fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to create season: a valid testing error"))
		})

		It("should rollback and return a formatted error when fetching a team fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(ContainSubstring("teams do not all exist: unable to get team ")))
			Expect(err.Error()).To(ContainSubstring(": a valid testing error"))
		})

		It("should rollback and return a formatted error when creating season teams fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(ContainSubstring("unable to add team ")))
			Expect(err.Error()).To(ContainSubstring(": a valid testing error"))
		})

		It("should rollback and return a formatted error when creating stage fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(ContainSubstring("unable to create stage")))
			Expect(err.Error()).To(ContainSubstring(": a valid testing error"))
		})

		It("should rollback and return a formatted error when fetching the season fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season: a valid testing error"))
		})

		It("should rollback and return a formatted error when fetching stages by seasonID fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.Stage{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season stages: a valid testing error"))
		})

		It("should return a formatted error when commit fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.GetSeasonTeamsRow{}, nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Create(context.Background(), validCreateSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("a valid testing error"))
		})
	})

	Describe("GetSeasons", func() {
		It("should get all seasons without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeasons(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonsFromDB, nil)
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
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB3, nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB2, nil)

			seasons, err := svc.GetAll(context.Background(), validCompetitionID)
			Expect(err).NotTo(HaveOccurred())

			Expect(seasons[0].ID).To(Equal(validSeasonsResponse[0].ID))
			Expect(seasons[0].CompetitionID).To(Equal(validSeasonsResponse[0].CompetitionID))
			Expect(seasons[0].StartDate).To(Equal(validSeasonsResponse[0].StartDate))
			Expect(seasons[0].EndDate).To(Equal(validSeasonsResponse[0].EndDate))
			Expect(seasons[0].Teams).To(HaveLen(2))
			Expect(seasons[0].Teams).To(ConsistOf(
				validTeamFromDB,
				validTeamFromDB2,
			))
			Expect(seasons[0].Stages).To(HaveLen(2))
			Expect(seasons[0].Stages).To(ConsistOf(
				validRegularStageFromDB,
				validFinalStagesFromDB,
			))
			Expect(seasons[0].CreatedAt).To(Equal(validSeasonsResponse[0].CreatedAt))
			Expect(seasons[0].UpdatedAt).To(Equal(validSeasonsResponse[0].UpdatedAt))
			Expect(seasons[0].DeletedAt.Time).To(Equal(validSeasonsResponse[0].DeletedAt.Time))

			Expect(seasons[1].ID).To(Equal(validSeasonsResponse[1].ID))
			Expect(seasons[1].CompetitionID).To(Equal(validSeasonsResponse[1].CompetitionID))
			Expect(seasons[1].StartDate).To(Equal(validSeasonsResponse[1].StartDate))
			Expect(seasons[1].EndDate).To(Equal(validSeasonsResponse[1].EndDate))
			Expect(seasons[1].Teams).To(HaveLen(2))
			Expect(seasons[1].Teams).To(ConsistOf(
				validTeamFromDB,
				validTeamFromDB3,
			))
			Expect(seasons[1].Stages).To(HaveLen(2))
			Expect(seasons[1].Stages).To(ConsistOf(
				validRegularStageFromDB,
				validFinalStagesFromDB2,
			))
			Expect(seasons[1].CreatedAt).To(Equal(validSeasonsResponse[1].CreatedAt))
			Expect(seasons[1].UpdatedAt).To(Equal(validSeasonsResponse[1].UpdatedAt))
			Expect(seasons[1].DeletedAt.Time).To(Equal(validSeasonsResponse[1].DeletedAt.Time))
		})

		It("should return a formatted error when getting seasons fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeasons(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)

			seasons, err := svc.GetAll(context.Background(), validCompetitionID)

			Expect(seasons).To(Equal(validNilSeasonsWithTeams))
			Expect(err.Error()).To(Equal("unable to get seasons: a valid testing error"))
		})

		It("should return a formatted error when getting season teams fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeasons(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonsFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)

			seasons, err := svc.GetAll(context.Background(), validCompetitionID)

			Expect(seasons).To(Equal(validNilSeasonsWithTeams))
			Expect(err.Error()).To(Equal("unable to get season teams: a valid testing error"))

		})
	})

	Describe("GetSeason", func() {
		It("should get a season without errors", func() {
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
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)

			season, err := svc.Get(context.Background(), validCompetitionID, validSeasonID)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Teams).To(HaveLen(2))
			Expect(season.Teams).To(ConsistOf(
				validTeamFromDB,
				validTeamFromDB2,
			))
			Expect(season.Stages).To(HaveLen(2))
			Expect(season.Stages).To(ConsistOf(
				validRegularStageFromDB,
				validFinalStagesFromDB,
			))
			Expect(season.CreatedAt).To(Equal(validSeasonResponse.CreatedAt))
			Expect(season.UpdatedAt).To(Equal(validSeasonResponse.UpdatedAt))
			Expect(season.DeletedAt.Time).To(Equal(validSeasonResponse.DeletedAt.Time))
		})

		It("should return a formatted error when getting a season fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Season{}, validTestError)

			season, err := svc.Get(context.Background(), validCompetitionID, validSeasonID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season: a valid testing error"))
		})

		It("should return a formatted error when getting season teams fails", func() {
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
			).Return([]db.GetSeasonTeamsRow{}, validTestError)

			season, err := svc.Get(context.Background(), validCompetitionID, validSeasonID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season teams: a valid testing error"))
		})
	})

	Describe("UpdateSeason", func() {
		It("should update a season without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)

			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Teams).To(HaveLen(2))
			Expect(season.Teams).To(ConsistOf(
				validTeamFromDB,
				validTeamFromDB2,
			))
			Expect(season.Stages).To(HaveLen(2))
			Expect(season.Stages).To(ConsistOf(
				validRegularStageFromDB,
				validFinalStagesFromDB,
			))
			Expect(season.CreatedAt).To(Equal(validSeasonResponse.CreatedAt))
			Expect(season.UpdatedAt).To(Equal(validSeasonResponse.UpdatedAt))
			Expect(season.DeletedAt.Time).To(Equal(validSeasonResponse.DeletedAt.Time))
		})

		It("should return a formatted error when transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("a valid testing error"))
		})

		It("should rollback and return a formatted error when updating the season fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to update season: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting a team fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				validTeamID,
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				validTeamID2,
			).Return(db.Team{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season teams: unable to get team 22222222-2222-4222-8222-222222222222: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting existing season teams fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season teams: unable to get season's existing teams: a valid testing error"))
		})

		It("should rollback and return a formatted error when deleting a season team fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season teams: unable to remove team 33333333-3333-4333-8333-333333333333 from season aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa: a valid testing error"))
		})

		It("should rollback and return an error when updating a stage that does not belong to the season", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.Stage{
				validRegularStageFromDB,
			}, nil)

			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				invalidUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(
				"unable to sync season stages: unable to update stage 66666666-6666-4666-8666-666666666666 for season aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa",
			))
		})

		It("should rollback and return a formatted error when getting the stages fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.Stage{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season stages: unable to get existing stages: a valid testing error"))
		})

		It("should rollback and return a formatted error when creating a new stage fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season stages: unable to create stage \"Regular Season 2\": a valid testing error"))
		})

		It("should rollback and return a formatted error when updating an existing stage fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season stages: unable to update stage 55555555-5555-4555-8555-555555555555: a valid testing error"))
		})

		It("should rollback and return a formatted error when deleteing an existing not needed stage fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to sync season stages: unable to delete stage 44444444-4444-4444-8444-444444444444: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting the season after update fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting season teams after update fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)
			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season teams: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting a team after update fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			).Return(db.Team{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get team: a valid testing error"))
		})

		It("should rollback and return a formatted error when getting a stage after update fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			).Return(db.Team{}, nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.Stage{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("unable to get season stages: a valid testing error"))
		})

		It("should return a formatted error when commit fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB2, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonTeamsFromDB2, nil)
			mockQueries.EXPECT().CreateSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockQueries.EXPECT().CreateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().UpdateStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStage(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
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
			mockQueries.EXPECT().GetStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validStagesFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := svc.Update(
				context.Background(),
				validUpdateSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("a valid testing error"))
		})
	})

	Describe("DeleteSeason", func() {
		It("should soft delete a season without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeamsBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := svc.Delete(context.Background(), validCompetitionID)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a formatted error when transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := svc.Delete(context.Background(), validCompetitionID)

			Expect(err.Error()).To(Equal("a valid testing error"))
		})

		It("should rollback and return a formatted error when deleting games fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := svc.Delete(context.Background(), validCompetitionID)
			Expect(err.Error()).To(Equal("unable to delete games for season: a valid testing error"))
		})

		It("should rollback and return a formatted error when deleting season teams fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeamsBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := svc.Delete(context.Background(), validCompetitionID)
			Expect(err.Error()).To(Equal("unable to delete season teams for season: a valid testing error"))
		})

		It("should rollback and return a formatted error when deleting the season stages fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeamsBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := svc.Delete(context.Background(), validCompetitionID)

			Expect(err.Error()).To(Equal("unable to delete stages for season: a valid testing error"))
		})

		It("should rollback and return a formatted error when deleting the season fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeamsBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := svc.Delete(context.Background(), validCompetitionID)

			Expect(err.Error()).To(Equal("unable to delete season: a valid testing error"))
		})

		It("should return a formatted error when commit fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeamsBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteStagesBySeasonID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := svc.Delete(context.Background(), validCompetitionID)

			Expect(err.Error()).To(Equal("a valid testing error"))
		})
	})
})
