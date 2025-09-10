package validation

import "github.com/go-playground/validator/v10"

func Register(v *validator.Validate) {
	v.RegisterValidation("entity_name", ValidateEntityName)
	v.RegisterValidation("unique_team_uuids", ValidateUniqueUUIDs)
}
