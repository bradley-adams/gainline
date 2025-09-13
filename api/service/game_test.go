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

var _ = Describe("game", func() {
	var ctrl *gomock.Controller
	var mockDB *mock_db.MockDB
	var mockQueries *mock_db.MockQueries

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)
		mockQueries = mock_db.NewMockQueries(ctrl)
	})

	validSeasonID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	validGameID := uuid.MustParse("bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbbb")
	validHomeTeamID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	validAwayTeamID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

	validTimeNow := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)

	homeScoreVal := int32(2)
	awayScoreVal := int32(1)

	validGameRequest := &api.GameRequest{
		Round:      3,
		Date:       validTimeNow,
		HomeTeamID: validHomeTeamID,
		AwayTeamID: validAwayTeamID,
		HomeScore:  &homeScoreVal,
		AwayScore:  &awayScoreVal,
		Status:     api.GameStatusPlaying,
	}

	var validNilGame db.Game
	var validNilGames []db.Game

	validGameFromDB := db.Game{
		ID:         validGameID,
		SeasonID:   validSeasonID,
		Round:      3,
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

	validUpdatedGameFromDB := db.Game{
		ID:         validGameID,
		SeasonID:   validSeasonID,
		Round:      4,
		Date:       validTimeNow.Add(2 * time.Hour),
		HomeTeamID: validHomeTeamID,
		AwayTeamID: validAwayTeamID,
		HomeScore:  sql.NullInt32{Int32: 5, Valid: true},
		AwayScore:  sql.NullInt32{Int32: 3, Valid: true},
		Status:     db.GameStatus(api.GameStatusFinished),
		CreatedAt:  validTimeNow,
		UpdatedAt:  validTimeNow.Add(2 * time.Hour),
		DeletedAt:  sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validGamesFromDB := []db.Game{
		validGameFromDB,
		{
			ID:         uuid.MustParse("736521aa-5332-405f-95ab-cc6beca13f95"),
			SeasonID:   validSeasonID,
			Round:      1,
			Date:       validTimeNow,
			HomeTeamID: validHomeTeamID,
			AwayTeamID: validAwayTeamID,
			HomeScore:  sql.NullInt32{Int32: 0, Valid: true},
			AwayScore:  sql.NullInt32{Int32: 0, Valid: true},
			Status:     db.GameStatus(api.GameStatusScheduled),
			CreatedAt:  validTimeNow,
			UpdatedAt:  validTimeNow,
			DeletedAt:  sql.NullTime{Time: time.Time{}, Valid: false},
		},
	}

	validGameResponse := api.ToGameResponse(validGameFromDB)
	validUpdatedGameResponse := api.ToGameResponse(validUpdatedGameFromDB)
	validGamesResponse := []api.GameResponse{
		api.ToGameResponse(validGameFromDB),
		api.ToGameResponse(validGamesFromDB[1]),
	}

	validCompetitionID := uuid.MustParse("cccccccc-cccc-4ccc-8ccc-cccccccccccc")

	validTeamFromDB := db.Team{
		ID:           validHomeTeamID,
		Name:         "Test Team",
		Abbreviation: "TT",
		Location:     "Testville",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validTeamFromDB2 := db.Team{
		ID:           validAwayTeamID,
		Name:         "Test Team2",
		Abbreviation: "TT2",
		Location:     "Testville 22",
		CreatedAt:    validTimeNow,
		UpdatedAt:    validTimeNow,
		DeletedAt:    sql.NullTime{Time: time.Time{}, Valid: false},
	}

	validSeasonWithTeams := SeasonWithTeams{
		ID:            validSeasonID,
		CompetitionID: validCompetitionID,
		StartDate:     validTimeNow.AddDate(0, -1, 0),
		EndDate:       validTimeNow.AddDate(0, 10, 0),
		Rounds:        15,
		Teams:         []db.Team{validTeamFromDB, validTeamFromDB2},
		CreatedAt:     validTimeNow,
		UpdatedAt:     validTimeNow,
		DeletedAt:     zero.Time{},
	}

	validTestError := errors.New("a valid testing error")

	Describe("CreateGame", func() {
		It("should create a new game without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validGameFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := CreateGame(context.Background(), mockDB, validGameRequest, validSeasonWithTeams)
			Expect(err).NotTo(HaveOccurred())

			Expect(game.ID).To(Equal(validGameResponse.ID))
			Expect(game.SeasonID).To(Equal(validGameResponse.SeasonID))
			Expect(game.Round).To(Equal(validGameResponse.Round))
			Expect(game.Date).To(Equal(validGameResponse.Date))
			Expect(game.HomeTeamID).To(Equal(validGameResponse.HomeTeamID))
			Expect(game.AwayTeamID).To(Equal(validGameResponse.AwayTeamID))

			Expect(game.HomeScore.Valid).To(BeTrue())
			Expect(game.HomeScore.Int32).To(Equal(*validGameResponse.HomeScore))
			Expect(game.AwayScore.Valid).To(BeTrue())
			Expect(game.AwayScore.Int32).To(Equal(*validGameResponse.AwayScore))

			Expect(string(game.Status)).To(Equal(string(validGameResponse.Status)))
			Expect(game.CreatedAt).To(Equal(validGameResponse.CreatedAt))
			Expect(game.UpdatedAt).To(Equal(validGameResponse.UpdatedAt))
			Expect(game.DeletedAt.Time).To(Equal(validGameResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := CreateGame(context.Background(), mockDB, validGameRequest, validSeasonWithTeams)

			Expect(game).To(Equal(validNilGame))
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
			mockQueries.EXPECT().CreateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			game, err := CreateGame(context.Background(), mockDB, validGameRequest, validSeasonWithTeams)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal("unable to create new game: a valid testing error"))
		})

		It("should rollback and return formatted error on get new game failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilGame, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			game, err := CreateGame(context.Background(), mockDB, validGameRequest, validSeasonWithTeams)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal("unable to get new game: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().CreateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validGameFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := CreateGame(context.Background(), mockDB, validGameRequest, validSeasonWithTeams)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("GetGames", func() {
		It("should retrieve all games without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGames(
				gomock.Any(),
				gomock.Any(),
			).Return(validGamesFromDB, nil)

			games, err := GetGames(context.Background(), mockDB, validSeasonID)
			Expect(err).NotTo(HaveOccurred())

			Expect(games[0].ID).To(Equal(validGamesResponse[0].ID))
			Expect(games[0].SeasonID).To(Equal(validGamesResponse[0].SeasonID))
			Expect(games[0].Round).To(Equal(validGamesResponse[0].Round))
			Expect(games[0].Date).To(Equal(validGamesResponse[0].Date))
			Expect(games[0].HomeTeamID).To(Equal(validGamesResponse[0].HomeTeamID))
			Expect(games[0].AwayTeamID).To(Equal(validGamesResponse[0].AwayTeamID))
			Expect(games[0].HomeScore.Valid).To(BeTrue())
			Expect(games[0].HomeScore.Int32).To(Equal(*validGamesResponse[0].HomeScore))
			Expect(games[0].AwayScore.Valid).To(BeTrue())
			Expect(games[0].AwayScore.Int32).To(Equal(*validGamesResponse[0].AwayScore))
			Expect(string(games[0].Status)).To(Equal(string(validGamesResponse[0].Status)))
			Expect(games[0].CreatedAt).To(Equal(validGamesResponse[0].CreatedAt))
			Expect(games[0].UpdatedAt).To(Equal(validGamesResponse[0].UpdatedAt))
			Expect(games[0].DeletedAt.Time).To(Equal(validGamesResponse[0].DeletedAt.Time))

			Expect(games[1].ID).To(Equal(validGamesResponse[1].ID))
			Expect(games[1].SeasonID).To(Equal(validGamesResponse[1].SeasonID))
			Expect(games[1].Round).To(Equal(validGamesResponse[1].Round))
			Expect(games[1].Date).To(Equal(validGamesResponse[1].Date))
			Expect(games[1].HomeTeamID).To(Equal(validGamesResponse[1].HomeTeamID))
			Expect(games[1].AwayTeamID).To(Equal(validGamesResponse[1].AwayTeamID))
			Expect(games[1].HomeScore.Valid).To(BeTrue())
			Expect(games[1].HomeScore.Int32).To(Equal(*validGamesResponse[1].HomeScore))
			Expect(games[1].AwayScore.Valid).To(BeTrue())
			Expect(games[1].AwayScore.Int32).To(Equal(*validGamesResponse[1].AwayScore))
			Expect(string(games[1].Status)).To(Equal(string(validGamesResponse[1].Status)))
			Expect(games[1].CreatedAt).To(Equal(validGamesResponse[1].CreatedAt))
			Expect(games[1].UpdatedAt).To(Equal(validGamesResponse[1].UpdatedAt))
			Expect(games[1].DeletedAt.Time).To(Equal(validGamesResponse[1].DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGames(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)

			games, err := GetGames(context.Background(), mockDB, validSeasonID)

			Expect(games).To(Equal(validNilGames))
			Expect(err.Error()).To(Equal("unable to get games: a valid testing error"))
		})
	})

	Describe("GetGame", func() {
		It("should retrieve a game without errors", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validGameFromDB, nil)

			game, err := GetGame(context.Background(), mockDB, validGameID)
			Expect(err).NotTo(HaveOccurred())

			Expect(game.ID).To(Equal(validGameResponse.ID))
			Expect(game.SeasonID).To(Equal(validGameResponse.SeasonID))
			Expect(game.Round).To(Equal(validGameResponse.Round))
			Expect(game.Date).To(Equal(validGameResponse.Date))
			Expect(game.HomeTeamID).To(Equal(validGameResponse.HomeTeamID))
			Expect(game.AwayTeamID).To(Equal(validGameResponse.AwayTeamID))

			Expect(game.HomeScore.Valid).To(BeTrue())
			Expect(game.HomeScore.Int32).To(Equal(*validGameResponse.HomeScore))
			Expect(game.AwayScore.Valid).To(BeTrue())
			Expect(game.AwayScore.Int32).To(Equal(*validGameResponse.AwayScore))

			Expect(string(game.Status)).To(Equal(string(validGameResponse.Status)))
			Expect(game.CreatedAt).To(Equal(validGameResponse.CreatedAt))
			Expect(game.UpdatedAt).To(Equal(validGameResponse.UpdatedAt))
			Expect(game.DeletedAt.Time).To(Equal(validGameResponse.DeletedAt.Time))
		})

		It("should return formatted error when retrieval fails", func() {
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(db.Game{}, validTestError)

			game, err := GetGame(context.Background(), mockDB, validGameID)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal("unable to get game: a valid testing error"))
		})
	})

	Describe("UpdateGame", func() {
		It("should update a game without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validUpdatedGameFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := UpdateGame(context.Background(), mockDB, validGameRequest, validGameID)
			Expect(err).NotTo(HaveOccurred())

			Expect(game.ID).To(Equal(validUpdatedGameResponse.ID))
			Expect(game.SeasonID).To(Equal(validUpdatedGameResponse.SeasonID))
			Expect(game.Round).To(Equal(validUpdatedGameResponse.Round))
			Expect(game.Date).To(Equal(validUpdatedGameResponse.Date))
			Expect(game.HomeTeamID).To(Equal(validUpdatedGameResponse.HomeTeamID))
			Expect(game.AwayTeamID).To(Equal(validUpdatedGameResponse.AwayTeamID))

			Expect(game.HomeScore.Valid).To(BeTrue())
			Expect(game.HomeScore.Int32).To(Equal(*validUpdatedGameResponse.HomeScore))
			Expect(game.AwayScore.Valid).To(BeTrue())
			Expect(game.AwayScore.Int32).To(Equal(*validUpdatedGameResponse.AwayScore))

			Expect(string(game.Status)).To(Equal(string(validUpdatedGameResponse.Status)))
			Expect(game.CreatedAt).To(Equal(validUpdatedGameResponse.CreatedAt))
			Expect(game.UpdatedAt).To(Equal(validUpdatedGameResponse.UpdatedAt))
			Expect(game.DeletedAt.Time).To(Equal(validUpdatedGameResponse.DeletedAt.Time))
		})

		It("should return formatted error if transaction begin fails", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			).Return(nil, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := UpdateGame(context.Background(), mockDB, validGameRequest, validGameID)

			Expect(game).To(Equal(validNilGame))
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
			mockQueries.EXPECT().UpdateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			game, err := UpdateGame(context.Background(), mockDB, validGameRequest, validGameID)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal("unable to update game: a valid testing error"))
		})

		It("should rollback and return formatted error on getting updated game failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validNilGame, validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			game, err := UpdateGame(context.Background(), mockDB, validGameRequest, validGameID)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal("unable to get updated game: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().UpdateGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockQueries.EXPECT().GetGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validUpdatedGameFromDB, nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			game, err := UpdateGame(context.Background(), mockDB, validGameRequest, validGameID)

			Expect(game).To(Equal(validNilGame))
			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("DeleteGame", func() {
		It("should soft delete a game without errors", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := DeleteGame(context.Background(), mockDB, validGameID)
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

			err := DeleteGame(context.Background(), mockDB, validGameID)
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
			mockQueries.EXPECT().DeleteGame(
				gomock.Any(),
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).AnyTimes()

			err := DeleteGame(context.Background(), mockDB, validGameID)
			Expect(err.Error()).To(Equal("unable to delete game: a valid testing error"))
		})

		It("should return formatted error on commit failure", func() {
			mockDB.EXPECT().BeginTx(
				gomock.Any(),
				gomock.Any(),
			)
			mockDB.EXPECT().New(
				gomock.Any(),
			).Return(mockQueries)
			mockQueries.EXPECT().DeleteGame(
				gomock.Any(),
				gomock.Any(),
			).Return(nil)
			mockDB.EXPECT().Commit(
				gomock.Any(),
			).Return(validTestError)
			mockDB.EXPECT().Rollback(
				gomock.Any(),
			).Times(0)

			err := DeleteGame(context.Background(), mockDB, validGameID)

			Expect(err.Error()).To(Equal(validTestError.Error()))
		})
	})

	Describe("validateGameRequest", func() {
		It("should allow a valid game request", func() {
			err := validateGameRequest(validGameRequest, validSeasonWithTeams)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should reject if round is greater than season rounds", func() {
			badReq := *validGameRequest
			badReq.Round = 16
			err := validateGameRequest(&badReq, validSeasonWithTeams)
			Expect(err).To(MatchError("round 16 is out of bounds (1-15)"))
		})

		It("should reject if home team not in season", func() {
			badReq := *validGameRequest
			badReq.HomeTeamID = uuid.MustParse("33333333-3333-4333-8333-333333333333")
			err := validateGameRequest(&badReq, validSeasonWithTeams)
			Expect(err).To(MatchError("home team not in season"))
		})

		It("should reject if away team not in season", func() {
			badReq := *validGameRequest
			badReq.AwayTeamID = uuid.MustParse("44444444-4444-4444-8444-444444444444")
			err := validateGameRequest(&badReq, validSeasonWithTeams)
			Expect(err).To(MatchError("away team not in season"))
		})

		It("should reject if date is before season start", func() {
			badReq := *validGameRequest
			badReq.Date = validSeasonWithTeams.StartDate.Add(-24 * time.Hour)
			err := validateGameRequest(&badReq, validSeasonWithTeams)
			Expect(err.Error()).To(ContainSubstring("outside season bounds"))
		})

		It("should reject if date is after season end", func() {
			badReq := *validGameRequest
			badReq.Date = validSeasonWithTeams.EndDate.Add(24 * time.Hour)
			err := validateGameRequest(&badReq, validSeasonWithTeams)
			Expect(err.Error()).To(ContainSubstring("outside season bounds"))
		})
	})
})
