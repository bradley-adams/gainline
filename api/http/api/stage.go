package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type StageType string

const (
	StageTypeRegular StageType = "regular"
	StageTypeFinals  StageType = "finals"
)

type StageRequest struct {
	Name       string    `json:"name" validate:"required,min=3,max=100,entity_name"`
	StageType  StageType `json:"stage_type" validate:"required,stage_type"`
	OrderIndex int32     `json:"order_index" validate:"required,min=1"`
}

type StageResponse struct {
	ID         uuid.UUID `json:"id"`
	SeasonID   uuid.UUID `json:"season_id"`
	Name       string    `json:"name"`
	StageType  StageType `json:"stage_type"`
	OrderIndex int32     `json:"order_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  zero.Time `json:"deleted_at"`
}

func ToStageResponse(s db.Stage) StageResponse {
	stageType := StageType(s.StageType)

	return StageResponse{
		ID:         s.ID,
		SeasonID:   s.SeasonID,
		Name:       s.Name,
		StageType:  stageType,
		OrderIndex: s.OrderIndex,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
		DeletedAt:  zero.TimeFrom(s.DeletedAt.Time),
	}
}

func ValidateStageType(fl validator.FieldLevel) bool {
	stageType := fl.Field().String()
	return stageType == "regular" || stageType == "finals"
}
