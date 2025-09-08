package api

import (
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ValidateTeamRequest", func() {
	var validate *validator.Validate

	BeforeEach(func() {
		validate = validator.New()
		validate.RegisterValidation("entity_name", validation.ValidateEntityName)
	})

	It("passes with valid team data", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "HIG",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).NotTo(HaveOccurred())
	})

	It("fails when name is empty", func() {
		team := &TeamRequest{
			Name:         "",
			Abbreviation: "HIG",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when name is too short (<3)", func() {
		team := &TeamRequest{
			Name:         "AB",
			Abbreviation: "HIG",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when name is too long (>100)", func() {
		longName := ""
		for i := 0; i < 101; i++ {
			longName += "A"
		}
		team := &TeamRequest{
			Name:         longName,
			Abbreviation: "HIG",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when abbreviation is missing", func() {
		team := &TeamRequest{
			Name:     "Highlanders",
			Location: "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when abbreviation is too short (<2)", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "H",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when abbreviation is too long (>4)", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "HIGHD",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when abbreviation contains numbers", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "H1G",
			Location:     "Dunedin",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("passes when location is empty", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "HIG",
			Location:     "",
		}
		Expect(validate.Struct(team)).NotTo(HaveOccurred())
	})

	It("fails when location is too short (<2)", func() {
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "HIG",
			Location:     "A",
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})

	It("fails when location is too long (>100)", func() {
		longLocation := ""
		for i := 0; i < 101; i++ {
			longLocation += "B"
		}
		team := &TeamRequest{
			Name:         "Highlanders",
			Abbreviation: "HIG",
			Location:     longLocation,
		}
		Expect(validate.Struct(team)).To(HaveOccurred())
	})
})
