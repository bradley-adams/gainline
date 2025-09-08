package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var competitionNameRegex = regexp.MustCompile(`^[A-Za-z0-9 .,'-]+$`)

func ValidateEntityName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return competitionNameRegex.MatchString(name)
}

func ValidateUniqueUUIDs(fl validator.FieldLevel) bool {
	teams, ok := fl.Field().Interface().([]uuid.UUID)
	if !ok {
		return false
	}

	seen := make(map[uuid.UUID]struct{})
	for _, t := range teams {
		if _, exists := seen[t]; exists {
			return false
		}
		seen[t] = struct{}{}
	}

	return true
}
