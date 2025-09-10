package api

import "github.com/go-playground/validator/v10"

func Register(v *validator.Validate) {
	v.RegisterValidation("game_status", ValidateGameStatus)
	v.RegisterStructValidation(ValidateGameRequest, GameRequest{})
}
