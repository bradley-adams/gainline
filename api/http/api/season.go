package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type SeasonRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Rounds    int32     `json:"rounds"`
}

type SeasonResponse struct {
	ID            uuid.UUID `json:"id"`
	CompetitionID uuid.UUID `json:"competition_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Rounds        int32     `json:"rounds"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     zero.Time `json:"deleted_at"`
}

func ToSeasonResponse(s db.Season) SeasonResponse {
	return SeasonResponse{
		ID:            s.ID,
		CompetitionID: s.CompetitionID,
		StartDate:     s.StartDate,
		EndDate:       s.EndDate,
		Rounds:        s.Rounds,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		DeletedAt:     zero.TimeFrom(s.DeletedAt.Time),
	}
}
