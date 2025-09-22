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

var _ = Describe("team", func() {
	var ctrl *gomock.Controller
	var mockDB *mock_db.MockDB
	var mockQueries *mock_db.MockQueries
	var svc TeamService

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)
		svc = NewTeamService(mockDB)
	})

	validTeamID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	validTeamID2 := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	validTimeNow := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)

	validTeamRequest := &api.TeamRequest{
		Name:         "Team A",
		Abbreviation: "TA",
		Location:     "Somewhere",
	}

	var validNilTeam db.Team
	var validNilTeams []db.Team

	validTeamFromDB := db.Team{
		ID:           validTeamID,
		Name:         "Team A",
		Abbreviation: "TA",
		Location:     "Somewhere",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validUpdatedTeamFromDB := db.Team{
		ID:           validTeamID,
		Name:         "Team A Updated",
		Abbreviation: "TUA",
		Location:     "Elsewhere",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow.Add(1 * time.Hour),
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validTeamsFromDB := []db.Team{
		validTeamFromDB,
		{
			ID:           validTeamID2,
			Name:         "Team B",
			Abbreviation: "TB",
			Location:     "Anywhere",
			CreatedAt:    validTimeNow,
			UpdatedAt:    validTimeNow,
			DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
		},
	}

	validTeamResponse := api.ToTeamResponse(validTeamFromDB)
	validUpdatedTeamResponse := api.ToTeamResponse(validUpdatedTeamFromDB)
	validTeamsResponse := []api.TeamResponse{
		api.ToTeamResponse(validTeamFromDB),
		api.ToTeamResponse(validTeamsFromDB[1]),
	}

	validTestError := errors.New("a valid testing error")

	Describe("CreateTeam", func() {
		It("should create a new team without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Create(context.Background(), validTeamRequest)
			Expect(err).NotTo(HaveOccurred())

			Expect(team.ID).To(Equal(validTeamResponse.ID))
			Expect(team.Name).To(Equal(validTeamResponse.Name))
			Expect(team.Abbreviation).To(Equal(validTeamResponse.Abbreviation))
			Expect(team.Location).To(Equal(validTeamResponse.Location))
			Expect(team.CreatedAt).To(Equal(validTeamResponse.CreatedAt))
			Expect(team.UpdatedAt).To(Equal(validTeamResponse.UpdatedAt))
			Expect(team.DeletedAt.Time).To(Equal(validTeamResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Create(context.Background(), validTeamRequest)

			Expect(team).To(Equal(validNilTeam))
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
			mockQueries.EXPECT().CreateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			team, err := svc.Create(context.Background(), validTeamRequest)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal("unable to create new team: a valid testing error"))
		})

		It("should rollback and return formatted error on get new team failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilTeam, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			team, err := svc.Create(context.Background(), validTeamRequest)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal("unable to get new team: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Create(context.Background(), validTeamRequest)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("GetTeams", func() {
		It("should retrieve all teams without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetTeams(
				gomock.Any(),
			).Return(validTeamsFromDB, nil)

			teams, err := svc.GetAll(context.Background())
			Expect(err).NotTo(HaveOccurred())

			Expect(teams[0].ID).To(Equal(validTeamsResponse[0].ID))
			Expect(teams[0].Name).To(Equal(validTeamsResponse[0].Name))
			Expect(teams[0].Abbreviation).To(Equal(validTeamsResponse[0].Abbreviation))
			Expect(teams[0].Location).To(Equal(validTeamsResponse[0].Location))
			Expect(teams[0].CreatedAt).To(Equal(validTeamsResponse[0].CreatedAt))
			Expect(teams[0].UpdatedAt).To(Equal(validTeamsResponse[0].UpdatedAt))
			Expect(teams[0].DeletedAt.Time).To(Equal(validTeamsResponse[0].DeletedAt.Time))

			Expect(teams[1].ID).To(Equal(validTeamsResponse[1].ID))
			Expect(teams[1].Name).To(Equal(validTeamsResponse[1].Name))
			Expect(teams[1].Abbreviation).To(Equal(validTeamsResponse[1].Abbreviation))
			Expect(teams[1].Location).To(Equal(validTeamsResponse[1].Location))
			Expect(teams[1].CreatedAt).To(Equal(validTeamsResponse[1].CreatedAt))
			Expect(teams[1].UpdatedAt).To(Equal(validTeamsResponse[1].UpdatedAt))
			Expect(teams[1].DeletedAt.Time).To(Equal(validTeamsResponse[1].DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetTeams(
				gomock.Any(),
			).Return(nil, validTestError)

			teams, err := svc.GetAll(context.Background())

			Expect(teams).To(Equal(validNilTeams))
			Expect(err.Error()).To(Equal("failed getting teams: a valid testing error"))
		})
	})

	Describe("GetTeam", func() {
		It("should retrieve a team without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTeamFromDB, nil)

			team, err := svc.Get(context.Background(), validTeamID)
			Expect(err).NotTo(HaveOccurred())

			Expect(team.ID).To(Equal(validTeamResponse.ID))
			Expect(team.Name).To(Equal(validTeamResponse.Name))
			Expect(team.Abbreviation).To(Equal(validTeamResponse.Abbreviation))
			Expect(team.Location).To(Equal(validTeamResponse.Location))
			Expect(team.CreatedAt).To(Equal(validTeamResponse.CreatedAt))
			Expect(team.UpdatedAt).To(Equal(validTeamResponse.UpdatedAt))
			Expect(team.DeletedAt.Time).To(Equal(validTeamResponse.DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Team{}, validTestError)

			team, err := svc.Get(context.Background(), validTeamID)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal("failed to get team: a valid testing error"))
		})
	})

	Describe("UpdateTeam", func() {
		It("should update a team without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validUpdatedTeamFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Update(context.Background(), validTeamRequest, validTeamID)
			Expect(err).NotTo(HaveOccurred())

			Expect(team.ID).To(Equal(validUpdatedTeamResponse.ID))
			Expect(team.Name).To(Equal(validUpdatedTeamResponse.Name))
			Expect(team.Abbreviation).To(Equal(validUpdatedTeamResponse.Abbreviation))
			Expect(team.Location).To(Equal(validUpdatedTeamResponse.Location))
			Expect(team.CreatedAt).To(Equal(validUpdatedTeamResponse.CreatedAt))
			Expect(team.UpdatedAt).To(Equal(validUpdatedTeamResponse.UpdatedAt))
			Expect(team.DeletedAt.Time).To(Equal(validUpdatedTeamResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Update(context.Background(), validTeamRequest, validTeamID)

			Expect(team).To(Equal(validNilTeam))
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
			mockQueries.EXPECT().UpdateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			team, err := svc.Update(context.Background(), validTeamRequest, validTeamID)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal("unable to update team: a valid testing error"))
		})

		It("should rollback and return formatted error on getting updated team failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilTeam, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			team, err := svc.Update(context.Background(), validTeamRequest, validTeamID)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal("unable to get updated team: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validUpdatedTeamFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			team, err := svc.Update(context.Background(), validTeamRequest, validTeamID)

			Expect(team).To(Equal(validNilTeam))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("DeleteTeam", func() {
		It("should soft delete a team without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := svc.Delete(context.Background(), validTeamID)
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

			err := svc.Delete(context.Background(), validTeamID)

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})

		It("should rollback and return formatted error on delete failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			err := svc.Delete(context.Background(), validTeamID)

			Expect(err.Error()).To(Equal("unable to delete team: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteTeam(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := svc.Delete(context.Background(), validTeamID)

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})
})
