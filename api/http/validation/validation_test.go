package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validation Suite")
}

var _ = Describe("field-level validators", func() {
	var validate *validator.Validate

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("entity_name", ValidateEntityName)
		validate.RegisterValidation("unique_team_uuids", ValidateUniqueUUIDs)
	})

	Describe("ValidateEntityName", func() {
		It("should pass with valid names", func() {
			Expect(validate.Var("Super Rugby Pacific", "entity_name")).To(Succeed())
			Expect(validate.Var("Division 1 - Men's Cup, 2025", "entity_name")).To(Succeed())
		})

		It("should fail with invalid characters", func() {
			Expect(validate.Var("Super Rugby!!!", "entity_name")).To(HaveOccurred())
		})

		It("should fail with empty string", func() {
			Expect(validate.Var("", "entity_name")).To(HaveOccurred())
		})
	})

	Describe("ValidateUniqueUUIDs", func() {
		var team1, team2 uuid.UUID

		BeforeEach(func() {
			team1 = uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
			team2 = uuid.MustParse("7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97")
		})

		It("should pass with unique UUIDs", func() {
			Expect(validate.Var([]uuid.UUID{team1, team2}, "unique_team_uuids")).To(Succeed())
		})

		It("should fail with duplicates", func() {
			Expect(validate.Var([]uuid.UUID{team1, team1}, "unique_team_uuids")).To(HaveOccurred())
		})

		It("should pass with empty slice", func() {
			Expect(validate.Var([]uuid.UUID{}, "unique_team_uuids")).To(Succeed())
		})
	})
})
