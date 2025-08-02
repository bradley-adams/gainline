package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type SeasonRequest struct {
	StartDate time.Time   `json:"start_date"`
	EndDate   time.Time   `json:"end_date"`
	Rounds    int32       `json:"rounds"`
	TeamIDs   []uuid.UUID `json:"TeamIDs"`
}

type SeasonResponse struct {
	ID            uuid.UUID      `json:"id"`
	CompetitionID uuid.UUID      `json:"competition_id"`
	StartDate     time.Time      `json:"start_date"`
	EndDate       time.Time      `json:"end_date"`
	Rounds        int32          `json:"rounds"`
	Teams         []TeamResponse `json:"teams"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     zero.Time      `json:"deleted_at"`
}
