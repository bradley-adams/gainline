package validation

import (
	"testing"
	"time"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation Suite")
}

var _ = Describe("validation", func() {
	var (
		validate *validator.Validate

		team1 uuid.UUID
		team2 uuid.UUID

		date1 time.Time
		date2 time.Time
	)

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("entity_name", ValidateEntityName)
		validate.RegisterValidation("unique_team_uuids", ValidateUniqueUUIDs)
		validate.RegisterValidation("game_status", ValidateGameStatus)

		validate.RegisterStructValidation(ValidateGameRequest, api.GameRequest{})

		team1 = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		team2 = uuid.MustParse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")

		date1 = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		date2 = time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	})

	Describe("ValidateCompetitionRequest", func() {
		It("should pass with a valid name", func() {
			comp := &api.CompetitionRequest{Name: "Super Rugby Pacific"}
			err := validate.Struct(comp)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should allow numbers, spaces, and punctuation", func() {
			comp := &api.CompetitionRequest{Name: "Division 1 - Men's Cup, 2025"}
			err := validate.Struct(comp)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail with invalid characters", func() {
			comp := &api.CompetitionRequest{Name: "Super Rugby!!!"}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("entity_name"))
		})

		It("should fail with empty name as it's too short", func() {
			comp := &api.CompetitionRequest{Name: ""}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("ValidateSeasonRequest", func() {
		It("should pass with valid season data", func() {
			season := &api.SeasonRequest{
				StartDate: date1,
				EndDate:   date2,
				Rounds:    10,
				Teams:     []uuid.UUID{team1, team2},
			}
			err := validate.Struct(season)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail when end date is before start date", func() {
			season := &api.SeasonRequest{
				StartDate: date2,
				EndDate:   date1,
				Rounds:    10,
				Teams:     []uuid.UUID{team1, team2},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("gtfield"))
		})

		It("should fail when there are fewer than 2 teams", func() {
			season := &api.SeasonRequest{
				StartDate: date1,
				EndDate:   date2,
				Rounds:    5,
				Teams:     []uuid.UUID{team1},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("min"))
		})

		It("should fail when rounds exceed 52", func() {
			season := &api.SeasonRequest{
				StartDate: date1,
				EndDate:   date2,
				Rounds:    53,
				Teams:     []uuid.UUID{team1, team2},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("max"))
		})

		It("should fail when teams contain duplicates", func() {
			season := &api.SeasonRequest{
				StartDate: date1,
				EndDate:   date2,
				Rounds:    5,
				Teams:     []uuid.UUID{team1, team1},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("unique_team_uuids"))
		})

		It("should fail when teams contain a nil UUID", func() {
			season := &api.SeasonRequest{
				StartDate: date1,
				EndDate:   date2,
				Rounds:    5,
				Teams:     []uuid.UUID{team1, uuid.Nil},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("required"))
		})
	})

	Describe("ValidateGameRequest", func() {
		It("should pass with valid scheduled game", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should pass with valid playing game with scores", func() {
			home := int32(3)
			away := int32(2)
			game := &api.GameRequest{
				Round:      52,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				HomeScore:  &home,
				AwayScore:  &away,
				Status:     api.GameStatusPlaying,
			}
			err := validate.Struct(game)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should pass when scheduled game has nil scores", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail if AwayTeamID equals HomeTeamID", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team1,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail when HomeTeamID is invalid", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: uuid.Nil,
				AwayTeamID: team2,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail when AwayTeamID is invalid", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: uuid.Nil,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail when round is out of range", func() {
			game := &api.GameRequest{
				Round:      53,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail when date is missing", func() {
			game := &api.GameRequest{
				Round:      1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail with invalid game status", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     "invalid_status",
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail if scheduled game has non-nil scores", func() {
			home := int32(1)
			away := int32(2)
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				HomeScore:  &home,
				AwayScore:  &away,
				Status:     api.GameStatusScheduled,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail if playing game has nil scores", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusPlaying,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail if finished game has nil scores", func() {
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				Status:     api.GameStatusFinished,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})

		It("should fail if scores are negative", func() {
			home := int32(-1)
			away := int32(2)
			game := &api.GameRequest{
				Round:      1,
				Date:       date1,
				HomeTeamID: team1,
				AwayTeamID: team2,
				HomeScore:  &home,
				AwayScore:  &away,
				Status:     api.GameStatusFinished,
			}
			err := validate.Struct(game)
			Expect(err).To(HaveOccurred())
		})
	})
})
