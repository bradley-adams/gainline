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
	var validate *validator.Validate

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("competition_name", ValidateCompetitionName)
		validate.RegisterValidation("unique_team_uuids", ValidateUniqueUUIDs)
	})

	Describe("ValidateCompetitionName", func() {
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
			Expect(validationErrors[0].Tag()).To(Equal("competition_name"))
		})

		It("should fail with empty name as its too short", func() {
			comp := &api.CompetitionRequest{Name: ""}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("ValidateSeasonRequest", func() {
		var team1, team2 uuid.UUID

		BeforeEach(func() {
			team1 = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
			team2 = uuid.MustParse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")
		})

		It("should pass with valid season data", func() {
			season := &api.SeasonRequest{
				StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
				Rounds:    10,
				Teams:     []uuid.UUID{team1, team2},
			}
			err := validate.Struct(season)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail when end date is before start date", func() {
			season := &api.SeasonRequest{
				StartDate: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
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
				StartDate: time.Now(),
				EndDate:   time.Now().Add(24 * time.Hour),
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
				StartDate: time.Now(),
				EndDate:   time.Now().Add(24 * time.Hour),
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
				StartDate: time.Now(),
				EndDate:   time.Now().Add(24 * time.Hour),
				Rounds:    5,
				Teams:     []uuid.UUID{team1, team1},
			}
			err := validate.Struct(season)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("unique_team_uuids"))
		})
	})
})
