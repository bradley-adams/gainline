package api

import "github.com/go-playground/validator/v10"

func Register(v *validator.Validate) {
	v.RegisterValidation("game_status", ValidateGameStatus)
	v.RegisterValidation("stage_type", ValidateStageType)

	v.RegisterStructValidation(ValidateGameRequest, GameRequest{})
	v.RegisterStructValidation(ValidateSeasonRequest, SeasonRequest{})

}
