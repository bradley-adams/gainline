package validation

import (
	"regexp"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var competitionNameRegex = regexp.MustCompile(`^[A-Za-z0-9 .,'-]+$`)

func ValidateEntityName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return competitionNameRegex.MatchString(name)
}

func ValidateUniqueUUIDs(fl validator.FieldLevel) bool {
	teams, ok := fl.Field().Interface().([]uuid.UUID)
	if !ok {
		return false
	}

	seen := make(map[uuid.UUID]struct{})
	for _, t := range teams {
		if _, exists := seen[t]; exists {
			return false
		}
		seen[t] = struct{}{}
	}

	return true
}

var validStatuses = map[api.GameStatus]struct{}{
	api.GameStatusScheduled: {},
	api.GameStatusPlaying:   {},
	api.GameStatusFinished:  {},
	"":                      {},
}

func ValidateGameStatus(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(api.GameStatus)
	if !ok {
		return false
	}
	_, valid := validStatuses[status]
	return valid
}

func ValidateGameRequest(sl validator.StructLevel) {
	game := sl.Current().Interface().(api.GameRequest)

	// Teams must be different
	if game.HomeTeamID == game.AwayTeamID {
		sl.ReportError(game.HomeTeamID, "HomeTeamID", "home_team_id", "home_and_away_teams_must_differ", "")
		sl.ReportError(game.AwayTeamID, "AwayTeamID", "away_team_id", "home_and_away_teams_must_differ", "")
	}

	// Scheduled games must NOT have scores
	if game.Status == api.GameStatusScheduled && (game.HomeScore != nil || game.AwayScore != nil) {
		sl.ReportError(game.HomeScore, "HomeScore", "home_score", "no_scores_for_scheduled_games", "")
		sl.ReportError(game.AwayScore, "AwayScore", "away_score", "no_scores_for_scheduled_games", "")
	}

	// Playing or Finished games must have scores
	if (game.Status == api.GameStatusPlaying || game.Status == api.GameStatusFinished) &&
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
