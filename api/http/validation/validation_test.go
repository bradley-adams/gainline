package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
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
	})

	Describe("ValidateCompetitionName", func() {
		type Competition struct {
			Name string `validate:"competition_name"`
		}

		It("should pass with a valid name", func() {
			comp := &Competition{Name: "Super Rugby Pacific"}
			err := validate.Struct(comp)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should allow numbers, spaces, and punctuation", func() {
			comp := &Competition{Name: "Division 1 - Men's Cup, 2025"}
			err := validate.Struct(comp)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail with invalid characters", func() {
			comp := &Competition{Name: "Super Rugby!!!"}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())

			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("competition_name"))
		})

		It("should fail with empty name as its too short", func() {
			comp := &Competition{Name: ""}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
		})
	})
})
