package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// handleWatchGame streams live game state updates to the client via SSE
//
//	@Summary	Watch live game state
//	@ID			watch-game
//	@Tags		Games
//	@Produce	text/event-stream
//	@Param		competitionID	path	string	true	"Competition ID"
//	@Param		seasonID		path	string	true	"Season ID"
//	@Param		gameID			path	string	true	"Game ID"
//	@Success	200
//	@Failure	400	{object}	response.ErrorResponse	"Invalid game ID"
//	@Failure	500	{object}	response.ErrorResponse	"Unable to watch game"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games/{gameID}/live [get]
func handleWatchGame(logger zerolog.Logger, gameStateService service.GameStateService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gameID, err := uuid.Parse(ctx.Param("gameID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid game ID")
			return
		}

		updates, err := gameStateService.WatchGameState(ctx.Request.Context(), gameID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to watch game")
			return
		}

		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")

		for state := range updates {
			data, err := json.Marshal(state)
			if err != nil {
				logger.Error().Err(err).Msg("failed to marshal game state")
				continue
			}
			ctx.SSEvent("update", string(data))
			ctx.Writer.Flush()
		}
	}
}
