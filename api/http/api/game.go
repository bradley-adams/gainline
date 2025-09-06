package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type GameStatus string

const (
	GameStatusScheduled GameStatus = "scheduled"
	GameStatusPlaying   GameStatus = "playing"
	GameStatusFinished  GameStatus = "finished"
)

func (gs GameStatus) String() string {
	return string(gs)
}

type GameRequest struct {
	Round      int32      `json:"round" validate:"required,min=1,max=52"`
	Date       time.Time  `json:"date" validate:"required" example:"2025-01-01T00:00:00Z"`
	HomeTeamID uuid.UUID  `json:"home_team_id" validate:"required,uuid" swaggertype:"string" example:"013952a5-87e1-4d26-a312-09b2aff54241"`
	AwayTeamID uuid.UUID  `json:"away_team_id" validate:"required,uuid" swaggertype:"string" example:"7b6cdb33-3bc6-4b0c-bac2-82d2a6bc6a97"`
	HomeScore  *int32     `json:"home_score,omitempty" validate:"omitempty,min=0"`
	AwayScore  *int32     `json:"away_score,omitempty" validate:"omitempty,min=0"`
	Status     GameStatus `json:"status,omitempty" validate:"omitempty,game_status" example:"playing"`
}

type GameResponse struct {
	ID         uuid.UUID  `json:"id"`
	SeasonID   uuid.UUID  `json:"season_id"`
	Round      int32      `json:"round"`
	Date       time.Time  `json:"date"`
	HomeTeamID uuid.UUID  `json:"home_team_id"`
	AwayTeamID uuid.UUID  `json:"away_team_id"`
	HomeScore  *int32     `json:"home_score,omitempty"`
	AwayScore  *int32     `json:"away_score,omitempty"`
	Status     GameStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  zero.Time  `json:"deleted_at"`
}

func ToGameResponse(g db.Game) GameResponse {
	var homeScore *int32
	if g.HomeScore.Valid {
		h := int32(g.HomeScore.Int32)
		homeScore = &h
	}

	var awayScore *int32
	if g.AwayScore.Valid {
		a := int32(g.AwayScore.Int32)
		awayScore = &a
	}

	status := GameStatus(g.Status)

	return GameResponse{
		ID:         g.ID,
		SeasonID:   g.SeasonID,
		Round:      g.Round,
		Date:       g.Date,
		HomeTeamID: g.HomeTeamID,
		AwayTeamID: g.AwayTeamID,
		HomeScore:  homeScore,
		AwayScore:  awayScore,
		Status:     status,
		CreatedAt:  g.CreatedAt,
		UpdatedAt:  g.UpdatedAt,
		DeletedAt:  zero.TimeFrom(g.DeletedAt.Time),
	}
}
