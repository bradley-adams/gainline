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

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)
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

	validTeamIDs := []uuid.UUID{validTeamID, validTeamID2}

	validTimeNow := time.Now()

	validSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Rounds:    15,
		Teams:     validTeamIDs,
	}

	var validNilSeason db.Season

	validSeasonFromDB := db.Season{
		ID:            validSeasonID,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow,
		EndDate:       validTimeNow.AddDate(0, 5, 0),
		Rounds:        15,
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonFromDB2 := db.Season{
		ID:            validSeasonID2,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow,
		EndDate:       validTimeNow.AddDate(0, 5, 0),
		Rounds:        15,
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonsFromDB := []db.Season{
		validSeasonFromDB,
		validSeasonFromDB2,
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

	var validNilSeasonWithTeams SeasonWithTeams
	var validNilSeasonsWithTeams []SeasonWithTeams

	validSeasonWithTeams := SeasonWithTeams{
		ID:            validSeasonFromDB.ID,
		CompetitionID: validSeasonFromDB.CompetitionID,
		StartDate:     validSeasonFromDB.StartDate,
		EndDate:       validSeasonFromDB.EndDate,
		Rounds:        validSeasonFromDB.Rounds,
		Teams:         []db.Team{},
		CreatedAt:     validSeasonFromDB.CreatedAt,
		UpdatedAt:     validSeasonFromDB.UpdatedAt,
		DeletedAt:     zero.TimeFrom(validSeasonFromDB.DeletedAt.Time),
	}

	validSeasonWithTeams2 := SeasonWithTeams{
		ID:            validSeasonFromDB2.ID,
		CompetitionID: validSeasonFromDB2.CompetitionID,
		StartDate:     validSeasonFromDB2.StartDate,
		EndDate:       validSeasonFromDB2.EndDate,
		Rounds:        validSeasonFromDB2.Rounds,
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
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Rounds).To(Equal(validSeasonResponse.Rounds))
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

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed creating season: a valid testing error"))
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

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed creating season: unable to create season: a valid testing error"))
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

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(ContainSubstring("failed creating season: unable to get team ")))
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

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err).To(MatchError(ContainSubstring("failed creating season: unable to add team ")))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed creating season: unable to get season: a valid testing error"))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return([]db.GetSeasonTeamsRow{}, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed creating season: a valid testing error"))
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
			).Return(nil, nil)
			mockQueries.EXPECT().GetSeasonTeams(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, nil)

			seasons, err := GetSeasons(context.Background(), mockDB, validCompetitionID)
			Expect(err).NotTo(HaveOccurred())

			Expect(seasons[0].ID).To(Equal(validSeasonsResponse[0].ID))
			Expect(seasons[0].CompetitionID).To(Equal(validSeasonsResponse[0].CompetitionID))
			Expect(seasons[0].StartDate).To(Equal(validSeasonsResponse[0].StartDate))
			Expect(seasons[0].EndDate).To(Equal(validSeasonsResponse[0].EndDate))
			Expect(seasons[0].Rounds).To(Equal(validSeasonsResponse[0].Rounds))
			Expect(seasons[0].CreatedAt).To(Equal(validSeasonsResponse[0].CreatedAt))
			Expect(seasons[0].UpdatedAt).To(Equal(validSeasonsResponse[0].UpdatedAt))
			Expect(seasons[0].DeletedAt.Time).To(Equal(validSeasonsResponse[0].DeletedAt.Time))

			Expect(seasons[1].ID).To(Equal(validSeasonsResponse[1].ID))
			Expect(seasons[1].CompetitionID).To(Equal(validSeasonsResponse[1].CompetitionID))
			Expect(seasons[1].StartDate).To(Equal(validSeasonsResponse[1].StartDate))
			Expect(seasons[1].EndDate).To(Equal(validSeasonsResponse[1].EndDate))
			Expect(seasons[1].Rounds).To(Equal(validSeasonsResponse[1].Rounds))
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

			seasons, err := GetSeasons(context.Background(), mockDB, validCompetitionID)

			Expect(seasons).To(Equal(validNilSeasonsWithTeams))
			Expect(err.Error()).To(Equal("failed getting seasons: unable to get seasons: a valid testing error"))
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

			seasons, err := GetSeasons(context.Background(), mockDB, validCompetitionID)

			Expect(seasons).To(Equal(validNilSeasonsWithTeams))
			Expect(err.Error()).To(Equal("failed getting seasons: unable to get season teams: a valid testing error"))

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
			).Return([]db.GetSeasonTeamsRow{}, nil)

			season, err := GetSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Rounds).To(Equal(validSeasonResponse.Rounds))
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

			season, err := GetSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed getting season: unable to get season: a valid testing error"))
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

			season, err := GetSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed getting season: unable to get season teams: a valid testing error"))
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
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)
			Expect(err).NotTo(HaveOccurred())

			Expect(season.ID).To(Equal(validSeasonResponse.ID))
			Expect(season.CompetitionID).To(Equal(validSeasonResponse.CompetitionID))
			Expect(season.StartDate).To(Equal(validSeasonResponse.StartDate))
			Expect(season.EndDate).To(Equal(validSeasonResponse.EndDate))
			Expect(season.Rounds).To(Equal(validSeasonResponse.Rounds))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: a valid testing error"))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to update season: a valid testing error"))
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
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to get team 22222222-2222-4222-8222-222222222222: a valid testing error"))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to get season's existing teams: a valid testing error"))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to remove team 33333333-3333-4333-8333-333333333333 from season aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa: a valid testing error"))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to get season: a valid testing error"))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)
			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to get season teams: a valid testing error"))
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

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: unable to get team: a valid testing error"))
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
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				validCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: a valid testing error"))
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

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

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

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("failed deleting season: a valid testing error"))
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

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)
			Expect(err.Error()).To(Equal("failed deleting season: unable to delete games for season: a valid testing error"))
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

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)
			Expect(err.Error()).To(Equal("failed deleting season: unable to delete season teams for season: a valid testing error"))
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
			mockQueries.EXPECT().DeleteSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("failed deleting season: unable to delete season: a valid testing error"))
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

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("failed deleting season: a valid testing error"))
		})
	})
})
