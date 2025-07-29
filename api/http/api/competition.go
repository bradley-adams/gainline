package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type CompetitionRequest struct {
	Name string `json:"name"`
}

type CompetitionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt zero.Time `json:"deleted_at"`
}
