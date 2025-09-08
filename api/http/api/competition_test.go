package api

import (
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("validation", func() {
	var (
		validate *validator.Validate
	)

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("entity_name", validation.ValidateEntityName)
	})

	Describe("ValidateCompetitionRequest", func() {
		It("should pass with a valid name within min and max length", func() {
			comp := &CompetitionRequest{Name: "Super Rugby Pacific"}
			err := validate.Struct(comp)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail when name is empty (required)", func() {
			comp := &CompetitionRequest{Name: ""}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("required"))
		})

		It("should fail when name is too short (less than 3 characters)", func() {
			comp := &CompetitionRequest{Name: "AB"}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("min"))
		})

		It("should fail when name is too long (more than 100 characters)", func() {
			longName := make([]byte, 101)
			for i := range longName {
				longName[i] = 'A'
			}
			comp := &CompetitionRequest{Name: string(longName)}
			err := validate.Struct(comp)
			Expect(err).To(HaveOccurred())
			validationErrors, ok := err.(validator.ValidationErrors)
			Expect(ok).To(BeTrue())
			Expect(validationErrors[0].Tag()).To(Equal("max"))
		})
	})
})
