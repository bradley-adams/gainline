package api

import (
	"time"

	"github.com/bradley-adams/gainline/http/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SeasonRequest validation", func() {
	var (
		validate *validator.Validate
		team1    uuid.UUID
		team2    uuid.UUID
		date1    time.Time
		date2    time.Time
	)

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("unique_team_uuids", validation.ValidateUniqueUUIDs)

		team1 = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		team2 = uuid.MustParse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")

		date1 = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		date2 = time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	})

	It("passes with valid season data", func() {
		season := &SeasonRequest{
			StartDate: date1,
			EndDate:   date2,
			Teams:     []uuid.UUID{team1, team2},
		}
		Expect(validate.Struct(season)).To(Succeed())
	})

	It("fails when StartDate is missing", func() {
		season := &SeasonRequest{
			EndDate: date2,
			Teams:   []uuid.UUID{team1, team2},
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
		validationErrors, ok := err.(validator.ValidationErrors)
		Expect(ok).To(BeTrue())
		Expect(validationErrors[0].Tag()).To(Equal("required"))
	})

	It("fails when EndDate is missing", func() {
		season := &SeasonRequest{
			StartDate: date1,
			Teams:     []uuid.UUID{team1, team2},
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
		validationErrors, ok := err.(validator.ValidationErrors)
		Expect(ok).To(BeTrue())
		Expect(validationErrors[0].Tag()).To(Equal("required"))
	})

	It("fails when EndDate is before StartDate", func() {
		season := &SeasonRequest{
			StartDate: date2,
			EndDate:   date1,
			Teams:     []uuid.UUID{team1, team2},
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
		validationErrors, ok := err.(validator.ValidationErrors)
		Expect(ok).To(BeTrue())
		Expect(validationErrors[0].Tag()).To(Equal("gtfield"))
	})

	It("fails when Teams < 2", func() {
		season := &SeasonRequest{
			StartDate: date1,
			EndDate:   date2,
			Teams:     []uuid.UUID{team1},
		}
		Expect(validate.Struct(season)).To(HaveOccurred())
	})

	It("fails when Teams > 100", func() {
		teams := make([]uuid.UUID, 101)
		for i := 0; i < 101; i++ {
			teams[i] = uuid.New()
		}
		season := &SeasonRequest{
			StartDate: date1,
			EndDate:   date2,
			Teams:     teams,
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
	})

	It("fails when Teams contain duplicates", func() {
		season := &SeasonRequest{
			StartDate: date1,
			EndDate:   date2,
			Teams:     []uuid.UUID{team1, team1},
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
	})

	It("fails when Teams contain nil UUID", func() {
		season := &SeasonRequest{
			StartDate: date1,
			EndDate:   date2,
			Teams:     []uuid.UUID{team1, uuid.Nil},
		}
		err := validate.Struct(season)
		Expect(err).To(HaveOccurred())
	})
})
