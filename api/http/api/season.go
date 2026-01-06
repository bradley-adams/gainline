package api

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type SeasonRequest struct {
	StartDate time.Time      `json:"start_date" validate:"required" example:"2025-01-01T00:00:00Z"`
	EndDate   time.Time      `json:"end_date" validate:"required,gtfield=StartDate" example:"2025-12-31T23:59:59Z"`
	Stages    []StageRequest `json:"stages" validate:"required,min=1,max=50,dive"`
	Teams     []uuid.UUID    `json:"teams" validate:"required,min=2,max=100,unique_team_uuids,dive,required" swaggertype:"array,string" example:"013952a5-87e1-4d26-a312-09b2aff54241,7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"`
}

type SeasonResponse struct {
	ID            uuid.UUID       `json:"id"`
	CompetitionID uuid.UUID       `json:"competition_id"`
	StartDate     time.Time       `json:"start_date"`
	EndDate       time.Time       `json:"end_date"`
	Stages        []StageResponse `json:"stages"`
	Teams         []TeamResponse  `json:"teams"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     zero.Time       `json:"deleted_at"`
}

func ValidateSeasonRequest(sl validator.StructLevel) {
	season := sl.Current().Interface().(SeasonRequest)

	stages := season.Stages
	seen := make(map[int32]bool, len(stages))

	for _, s := range stages {
		if seen[s.OrderIndex] {
			sl.ReportError(
				stages,
				"stages",
				"Stages",
				"duplicate_order",
				"duplicate stage order_index",
			)
			return
		}
		seen[s.OrderIndex] = true
	}

	for i := 1; i <= len(stages); i++ {
		if !seen[int32(i)] {
			sl.ReportError(
				stages,
				"stages",
				"Stages",
				"non_contiguous_order",
				"stage order_index must be contiguous starting at 1",
			)
			return
		}
	}
}
