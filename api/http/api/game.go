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
	Round      int32      `json:"round"`
	Date       time.Time  `json:"date"`
	HomeTeamID uuid.UUID  `json:"home_team_id"`
	AwayTeamID uuid.UUID  `json:"away_team_id"`
	HomeScore  *int32     `json:"home_score,omitempty"`
	AwayScore  *int32     `json:"away_score,omitempty"`
	Status     GameStatus `json:"status,omitempty"`
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
