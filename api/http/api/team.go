package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type TeamRequest struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Location     string `json:"location"`
}

type TeamResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	Location     string    `json:"location"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    zero.Time `json:"deleted_at"`
}

func ToTeamResponse(t db.Team) TeamResponse {
	return TeamResponse{
		ID:           t.ID,
		Name:         t.Name,
		Abbreviation: t.Abbreviation,
		Location:     t.Location,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		DeletedAt:    zero.TimeFrom(t.DeletedAt.Time),
	}
}
