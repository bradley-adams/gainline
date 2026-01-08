package api

import (
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StageRequest validation", func() {
	var (
		validate *validator.Validate
		stage    StageRequest
	)

	BeforeEach(func() {
		validate = validator.New()

		// Register all custom validators used in StageRequest
		validate.RegisterValidation("entity_name", validation.ValidateEntityName)
		validate.RegisterValidation("stage_type", ValidateStageType)

		stage = StageRequest{
			Name:       "Group Stage",
			StageType:  StageTypeRegular,
			OrderIndex: 1,
		}
	})

	It("passes with valid stage data", func() {
		Expect(validate.Struct(stage)).To(Succeed())
	})

	It("fails when Name is missing", func() {
		stage.Name = ""
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
		validationErrors, ok := err.(validator.ValidationErrors)
		Expect(ok).To(BeTrue())
		Expect(validationErrors[0].Tag()).To(Equal("required"))
	})

	It("fails when Name is too short", func() {
		stage.Name = "AB"
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("fails when Name is too long", func() {
		longName := ""
		for i := 0; i < 101; i++ {
			longName += "A"
		}
		stage.Name = longName
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("fails when StageType is missing", func() {
		stage.StageType = ""
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("fails when StageType is invalid", func() {
		stage.StageType = "invalid_type"
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("fails when OrderIndex is missing or zero", func() {
		stage.OrderIndex = 0
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("fails when OrderIndex is over 100", func() {
		stage.OrderIndex = 101
		err := validate.Struct(stage)
		Expect(err).To(HaveOccurred())
	})

	It("passes with StageType finals", func() {
		stage.StageType = StageTypeFinals
		Expect(validate.Struct(stage)).To(Succeed())
	})
})
