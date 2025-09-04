package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type CompetitionRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100,competition_name"`
}

type CompetitionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt zero.Time `json:"deleted_at"`
}

func ToCompetitionResponse(c db.Competition) CompetitionResponse {
	return CompetitionResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: zero.TimeFrom(c.DeletedAt.Time),
	}
}
