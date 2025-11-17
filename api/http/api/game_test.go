package api

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GameRequest validation", func() {
	var (
		validate *validator.Validate
		team1    uuid.UUID
		team2    uuid.UUID
		date1    time.Time
	)

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterStructValidation(ValidateGameRequest, GameRequest{})
		validate.RegisterValidation("game_status", ValidateGameStatus)

		team1 = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		team2 = uuid.MustParse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")
		date1 = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	})

	It("passes with valid scheduled game", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})

	It("passes with valid playing game with scores", func() {
		home := int32(3)
		away := int32(2)
		game := GameRequest{
			Round:      10,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			HomeScore:  &home,
			AwayScore:  &away,
			Status:     GameStatusPlaying,
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})

	It("passes with zero scores for playing/finished games", func() {
		home := int32(0)
		away := int32(0)
		game := GameRequest{
			Round:      5,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			HomeScore:  &home,
			AwayScore:  &away,
			Status:     GameStatusFinished,
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})

	It("passes with zero scores for canceled games", func() {
		game := GameRequest{
			Round:      5,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			HomeScore:  nil,
			AwayScore:  nil,
			Status:     GameStatusCancelled,
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})

	It("fails if AwayTeamID equals HomeTeamID", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team1,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if scheduled game has non-nil scores", func() {
		home := int32(1)
		away := int32(2)
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			HomeScore:  &home,
			AwayScore:  &away,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if playing game has nil scores", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusPlaying,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if finished game has nil scores", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusFinished,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if scores are negative", func() {
		home := int32(-1)
		away := int32(2)
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			HomeScore:  &home,
			AwayScore:  &away,
			Status:     GameStatusFinished,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if round is below minimum", func() {
		game := GameRequest{
			Round:      0,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails if round exceeds maximum", func() {
		game := GameRequest{
			Round:      53,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("passes at round boundaries", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())

		game.Round = 52
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})

	It("fails if date is missing", func() {
		game := GameRequest{
			Round:      1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     GameStatusScheduled,
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("fails with invalid game status", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     "invalid_status",
		}
		Expect(validate.Struct(&game)).To(HaveOccurred())
	})

	It("passes if Status is empty and game is scheduled (omitempty)", func() {
		game := GameRequest{
			Round:      1,
			Date:       date1,
			HomeTeamID: team1,
			AwayTeamID: team2,
			Status:     "",
		}
		Expect(validate.Struct(&game)).NotTo(HaveOccurred())
	})
})
