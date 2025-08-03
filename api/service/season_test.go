package service

import (
	"context"
	"database/sql"
	"fmt"
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

	validSeasonID := uuid.New()
	validSeasonID2 := uuid.New()

	validCompetitionID := uuid.New()
	invalidCompetitionID := uuid.New()

	validSeasonTeamID := uuid.New()
	validSeasonTeamID2 := uuid.New()
	validSeasonTeamID3 := uuid.New()

	validTeamID := uuid.New()
	validTeamID2 := uuid.New()
	validTeamID3 := uuid.New()

	validTeamIDs := []uuid.UUID{validTeamID, validTeamID2}

	validTimeNow := time.Now()

	validSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Rounds:    15,
		TeamIDs:   validTeamIDs,
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
		It("should create a new season with no errors", func() {
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

		It("transaction begin error should return formatted", func() {
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

		It("insert season error should return formatted and then rollback", func() {
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

		It("get team error should return formatted and then rollback", func() {
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

		It("create season teams error should return formatted and then rollback", func() {
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

		It("get season error should return formatted and then rollback", func() {
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

		It("a commit error should return formatted", func() {
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
		It("should get all seasons with no errors", func() {
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

		It("get seasons error should return formatted", func() {
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

		It("get season teams error should return formatted", func() {
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
		It("should get a season with no errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("get season error should return formatted", func() {
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

		It("get season teams error should return formatted", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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
		It("should update a season with no errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("transaction begin error should return formatted", func() {
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

		It("get season error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
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

		It("season must belong to competition error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := UpdateSeason(
				context.Background(),
				mockDB,
				validSeasonRequest,
				invalidCompetitionID,
				validSeasonID,
			)

			Expect(season).To(Equal(validNilSeasonWithTeams))
			Expect(err.Error()).To(Equal("failed updating season: season does not belong to the specified competition"))
		})

		It("update season error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("get team error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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
			expected := fmt.Sprintf(
				"failed updating season: unable to get team %s: a valid testing error",
				validTeamID2.String(),
			)
			Expect(err).To(MatchError(expected))
		})

		It("get season team error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("delete season team error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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
			expected := fmt.Sprintf(
				"failed updating season: unable to remove team %s from season %s: a valid testing error",
				validTeamID3.String(),
				validSeasonID.String(),
			)
			Expect(err).To(MatchError(expected))
		})

		It("get season after update error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("get season team after update error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("get team after update error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

		It("get team after update error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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
		It("should soft delete a season with no errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
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
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
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

		It("transaction begin error should return formatted", func() {
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

		It("get season error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("failed deleting season: unable to get season for deletion: a valid testing error"))
		})

		It("get season teams error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
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
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("failed deleting season: unable to get season teams for deletion: a valid testing error"))
		})

		It("delete season team error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
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
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			expected := fmt.Sprintf(
				"failed deleting season: unable to remove team %s from season %s: a valid testing error",
				validTeamID.String(),
				validSeasonID.String(),
			)
			Expect(err).To(MatchError(expected))
		})

		It("delete season error should return formatted and then rollback", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
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
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
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

		It("a commit error should return formatted", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
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
			mockQueries.EXPECT().DeleteSeasonTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonTeam(
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
