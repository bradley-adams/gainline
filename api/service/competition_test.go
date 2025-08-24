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

var _ = Describe("competition", func() {
	var ctrl *gomock.Controller
	var mockDB *mock_db.MockDB
	var mockQueries *mock_db.MockQueries

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)
	})

	validCompetitionID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	validCompetitionID2 := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	validTimeNow := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)

	validCompetitionRequest := &api.CompetitionRequest{
		Name: "Test Competition",
	}

	var validNilCompetition db.Competition
	var validNilCompetitions []db.Competition

	validCompetitionFromDB := db.Competition{
		ID:        validCompetitionID,
		Name:      "Test Competition",
		CreatedAt: validTimeNow,
		UpdatedAt: validTimeNow,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validCompetitionFromDB2 := db.Competition{
		ID:        validCompetitionID2,
		Name:      "Test Competition 2",
		CreatedAt: validTimeNow,
		UpdatedAt: validTimeNow,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validCompetitionsFromDB := []db.Competition{
		validCompetitionFromDB,
		validCompetitionFromDB2,
	}

	validCompetitionResponse := api.ToCompetitionResponse(validCompetitionFromDB)

	validCompetitionResponse2 := api.ToCompetitionResponse(validCompetitionFromDB2)

	validCompetitionsResponse := []api.CompetitionResponse{
		validCompetitionResponse,
		validCompetitionResponse2,
	}

	validTestError := errors.New("a valid testing error")

	Describe("CreateCompetition", func() {
		It("should create a new competition without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validCompetitionFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := CreateCompetition(context.Background(), mockDB, validCompetitionRequest)
			Expect(err).NotTo(HaveOccurred())

			Expect(competition.ID).To(Equal(validCompetitionResponse.ID))
			Expect(competition.Name).To(Equal(validCompetitionResponse.Name))
			Expect(competition.CreatedAt).To(Equal(validCompetitionResponse.CreatedAt))
			Expect(competition.UpdatedAt).To(Equal(validCompetitionResponse.UpdatedAt))
			Expect(competition.DeletedAt.Time).To(Equal(validCompetitionResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := CreateCompetition(context.Background(), mockDB, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})

		It("should rollback and return formatted error on insert failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			competition, err := CreateCompetition(context.Background(), mockDB, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal("unable to create new competition: a valid testing error"))
		})

		It("should rollback and return formatted error on get failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilCompetition, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			competition, err := CreateCompetition(context.Background(), mockDB, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal("unable to get new competition: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validCompetitionFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := CreateCompetition(context.Background(), mockDB, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("GetCompetitions", func() {
		It("should retrieve all competitions without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetCompetitions(
				gomock.Any(),
			).Return(validCompetitionsFromDB, nil)

			competitionsResult, err := GetCompetitions(context.Background(), mockDB)
			Expect(err).NotTo(HaveOccurred())

			Expect(competitionsResult[0].ID).To(Equal(validCompetitionsResponse[0].ID))
			Expect(competitionsResult[0].Name).To(Equal(validCompetitionsResponse[0].Name))
			Expect(competitionsResult[0].CreatedAt).To(Equal(validCompetitionsResponse[0].CreatedAt))
			Expect(competitionsResult[0].UpdatedAt).To(Equal(validCompetitionsResponse[0].UpdatedAt))
			Expect(competitionsResult[0].DeletedAt.Time).To(Equal(validCompetitionsResponse[0].DeletedAt.Time))

			Expect(competitionsResult[1].ID).To(Equal(validCompetitionsResponse[1].ID))
			Expect(competitionsResult[1].Name).To(Equal(validCompetitionsResponse[1].Name))
			Expect(competitionsResult[1].CreatedAt).To(Equal(validCompetitionsResponse[1].CreatedAt))
			Expect(competitionsResult[1].UpdatedAt).To(Equal(validCompetitionsResponse[1].UpdatedAt))
			Expect(competitionsResult[1].DeletedAt.Time).To(Equal(validCompetitionsResponse[1].DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetCompetitions(
				gomock.Any(),
			).Return(nil, validTestError)

			competitions, err := GetCompetitions(context.Background(), mockDB)

			Expect(competitions).To(Equal(validNilCompetitions))
			Expect(err.Error()).To(Equal("unable to get competitions: a valid testing error"))
		})
	})

	Describe("GetCompetition", func() {
		It("should retrieve a competition without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validCompetitionFromDB, nil)

			competition, err := GetCompetition(context.Background(), mockDB, validCompetitionID)
			Expect(err).NotTo(HaveOccurred())

			Expect(competition.ID).To(Equal(validCompetitionResponse.ID))
			Expect(competition.Name).To(Equal(validCompetitionResponse.Name))
			Expect(competition.CreatedAt).To(Equal(validCompetitionResponse.CreatedAt))
			Expect(competition.UpdatedAt).To(Equal(validCompetitionResponse.UpdatedAt))
			Expect(competition.DeletedAt.Time).To(Equal(validCompetitionResponse.DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Competition{}, validTestError)

			competition, err := GetCompetition(context.Background(), mockDB, validCompetitionID)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal("unable to get competition: a valid testing error"))
		})
	})

	Describe("UpdateCompetition", func() {
		It("should update a competition without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validCompetitionFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := UpdateCompetition(
				context.Background(),
				mockDB,
				validCompetitionID,
				validCompetitionRequest,
			)
			Expect(err).NotTo(HaveOccurred())

			Expect(competition.ID).To(Equal(validCompetitionResponse.ID))
			Expect(competition.Name).To(Equal(validCompetitionResponse.Name))
			Expect(competition.CreatedAt).To(Equal(validCompetitionResponse.CreatedAt))
			Expect(competition.UpdatedAt).To(Equal(validCompetitionResponse.UpdatedAt))
			Expect(competition.DeletedAt.Time).To(Equal(validCompetitionResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := UpdateCompetition(
				context.Background(),
				mockDB,
				validCompetitionID,
				validCompetitionRequest,
			)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})

		It("should rollback and return formatted error on update failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			competition, err := UpdateCompetition(context.Background(), mockDB, validCompetitionID, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal("unable to update competition: a valid testing error"))
		})

		It("should rollback and return formatted error on get updated competition failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilCompetition, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			competition, err := UpdateCompetition(context.Background(), mockDB, validCompetitionID, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal("unable to get updated competition: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validCompetitionFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			competition, err := UpdateCompetition(context.Background(), mockDB, validCompetitionID, validCompetitionRequest)

			Expect(competition).To(Equal(validNilCompetition))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("DeleteCompetition", func() {
		It("should soft delete a competition without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)

			mockQueries.EXPECT().DeleteSeasonsByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})

		It("should rollback and return error if deleting games fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)
			Expect(err.Error()).To(Equal("unable to delete games for competition: a valid testing error"))
		})

		It("should rollback and return error if deleting seasons fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonsByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)
			Expect(err.Error()).To(Equal("unable to delete seasons for competition: a valid testing error"))
		})

		It("should rollback and return formatted error on delete failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonsByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)

			Expect(err.Error()).To(Equal("unable to delete competition: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGamesByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteSeasonsByCompetitionID(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().DeleteCompetition(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := DeleteCompetition(context.Background(), mockDB, validCompetitionID)

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})
})
