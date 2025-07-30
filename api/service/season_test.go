package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/bradley-adams/gainline/db/db"
	mock_db "github.com/bradley-adams/gainline/db/db_handler/mock"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/google/uuid"
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

	validTimeNow := time.Now()

	validSeasonRequest := &api.SeasonRequest{
		StartDate: validTimeNow,
		EndDate:   validTimeNow.AddDate(0, 5, 0),
		Rounds:    15,
	}

	var validNilSeason db.Season
	var validNilSeasons []db.Season

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

	validSeasonResponse := api.ToSeasonResponse(validSeasonFromDB)

	validSeasonResponse2 := api.ToSeasonResponse(validSeasonFromDB2)

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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal(validTestError.Error()))
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to create new season: a valid testing error"))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilSeason, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to get new season: a valid testing error"))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			season, err := CreateSeason(context.Background(), mockDB, validSeasonRequest, validCompetitionID)

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal(validTestError.Error()))
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

			Expect(seasons).To(Equal(validNilSeasons))
			Expect(err.Error()).To(Equal("unable to get seasons: a valid testing error"))
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to get season: a valid testing error"))
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
			mockQueries.EXPECT().GetSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validSeasonFromDB, nil)
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal(validTestError.Error()))
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to get season for update: a valid testing error"))
		})

		It("season must being to competition error should return formatted and then rollback", func() {
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("season does not belong to the specified competition"))
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to update season: a valid testing error"))
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
			).Return(validSeasonFromDB, nil)
			mockQueries.EXPECT().UpdateSeason(
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

			Expect(season).To(Equal(validNilSeason))
			Expect(err.Error()).To(Equal("unable to get updated season: a valid testing error"))
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

			Expect(err.Error()).To(Equal(validTestError.Error()))
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

			Expect(err.Error()).To(Equal("unable to get season for deletion: a valid testing error"))
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
			mockQueries.EXPECT().DeleteSeason(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(1)

			err := DeleteSeason(context.Background(), mockDB, validCompetitionID, validSeasonID)

			Expect(err.Error()).To(Equal("unable to delete season: a valid testing error"))
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

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})
})
