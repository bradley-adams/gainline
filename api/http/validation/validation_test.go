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
	RunSpecs(t, "validation Suite")
}

var _ = Describe("ValidateSeasonDates", func() {
	var validate *validator.Validate

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterStructValidation(ValidateSeasonDates, api.SeasonRequest{})
	})

	team1, _ := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	team2, _ := uuid.Parse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")

	It("should pass when start date is before end date", func() {
		season := &api.SeasonRequest{
			StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			Rounds:    1,
			Teams:     []uuid.UUID{team1, team2},
		}

		err := validate.Struct(season)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should fail when start date is after end date", func() {
		season := &api.SeasonRequest{
			StartDate: time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			EndDate:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			Rounds:    1,
			Teams:     []uuid.UUID{team1, team2},
		}

		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())

		validationErrors, ok := err.(validator.ValidationErrors)
		Expect(ok).To(BeTrue())

		Expect(validationErrors).To(ContainElement(
			WithTransform(func(e validator.FieldError) string { return e.Tag() }, Equal("start_must_be_before_end")),
		))
	})
})
