package api

import (
	"time"

	"github.com/bradley-adams/gainline/db/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
)

type GameStatus string

const (
	GameStatusScheduled GameStatus = "scheduled"
	GameStatusPlaying   GameStatus = "playing"
	GameStatusFinished  GameStatus = "finished"
	GameStatusCancelled GameStatus = "cancelled"
)

func (gs GameStatus) String() string {
	return string(gs)
}

type GameRequest struct {
	Round      int32      `json:"round" validate:"required,min=1,max=52" example:"10"`
	Date       time.Time  `json:"date" validate:"required" example:"2025-08-02T00:00:00Z"`
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

func ValidateGameStatus(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(GameStatus)
	if !ok {
		return false
	}

	switch status {
	case GameStatusScheduled, GameStatusPlaying, GameStatusFinished, GameStatusCancelled:
		return true
	default:
		return false
	}
}

func ValidateGameRequest(sl validator.StructLevel) {
	game := sl.Current().Interface().(GameRequest)

	// Teams must be different
	if game.HomeTeamID == game.AwayTeamID {
		sl.ReportError(game.HomeTeamID, "HomeTeamID", "home_team_id", "home_and_away_teams_must_differ", "")
		sl.ReportError(game.AwayTeamID, "AwayTeamID", "away_team_id", "home_and_away_teams_must_differ", "")
	}

	// Scheduled games must NOT have scores
	if game.Status == GameStatusScheduled && (game.HomeScore != nil || game.AwayScore != nil) {
		sl.ReportError(game.HomeScore, "HomeScore", "home_score", "no_scores_for_scheduled_games", "")
		sl.ReportError(game.AwayScore, "AwayScore", "away_score", "no_scores_for_scheduled_games", "")
	}

	// Cancelled games must NOT have scores
	if game.Status == GameStatusCancelled && (game.HomeScore != nil || game.AwayScore != nil) {
		sl.ReportError(game.HomeScore, "HomeScore", "home_score", "no_scores_for_cancelled_games", "")
		sl.ReportError(game.AwayScore, "AwayScore", "away_score", "no_scores_for_cancelled_games", "")
	}

	// Playing or Finished games must have scores
	if (game.Status == GameStatusPlaying || game.Status == GameStatusFinished) &&
		(game.HomeScore == nil || game.AwayScore == nil) {
		sl.ReportError(game.HomeScore, "HomeScore", "home_score", "scores_required_for_playing_or_finished_games", "")
		sl.ReportError(game.AwayScore, "AwayScore", "away_score", "scores_required_for_playing_or_finished_games", "")
	}

	// Scores cannot be negative
	if game.HomeScore != nil && *game.HomeScore < 0 {
		sl.ReportError(game.HomeScore, "HomeScore", "home_score", "scores_cannot_be_negative", "")
	}
	if game.AwayScore != nil && *game.AwayScore < 0 {
		sl.ReportError(game.AwayScore, "AwayScore", "away_score", "scores_cannot_be_negative", "")
	}
}
