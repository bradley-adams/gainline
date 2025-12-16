package handlers

import (
	"net/http"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// handleCreateGame creates a new game for a season
//
//	@Summary	Create a new game
//	@ID			create-game
//	@Tags		Games
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Param		game			body		api.GameRequest			true	"Game details to create"
//	@Success	201				{object}	api.GameResponse		"Successful operation"
//	@Failure	400				{object}	response.ErrorResponse	"Bad request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games [post]
func handleCreateGame(
	logger zerolog.Logger,
	validate *validator.Validate,
	gameService service.GameService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		season := ctx.MustGet("season").(service.SeasonAggregate)

		req := &api.GameRequest{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Bad request")
			return
		}

		err := validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		game, err := gameService.Create(ctx.Request.Context(), req, season)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to create game")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, api.ToGameResponse(game))
	}
}

// handleGetGames retrieves all games for a season
//
//	@Summary	Get all games for a season
//	@ID			get-games
//	@Tags		Games
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Success	200				{array}		api.GameResponse		"List of games"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games [get]
func handleGetGames(logger zerolog.Logger, gameService service.GameService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		seasonID, err := uuid.Parse(ctx.Param("seasonID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid season ID")
			return
		}

		games, err := gameService.GetAll(ctx.Request.Context(), seasonID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get games")
			return
		}

		resp := make([]api.GameResponse, 0, len(games))
		for _, g := range games {
			resp = append(resp, api.ToGameResponse(g))
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, resp)
	}
}

// handleGetGame retrieves a single game by ID for a season
//
//	@Summary	Get a single game by ID
//	@ID			get-game
//	@Tags		Games
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Param		gameID			path		string					true	"Game ID"			default(4019a7f3-7741-4d8f-b3e0-1c7f3a0a1a01)
//	@Success	200				{object}	api.GameResponse		"Game found"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games/{gameID} [get]
func handleGetGame(logger zerolog.Logger, gameService service.GameService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gameID, err := uuid.Parse(ctx.Param("gameID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid game ID")
			return
		}

		game, err := gameService.Get(ctx.Request.Context(), gameID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get game")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToGameResponse(game))
	}
}

// handleUpdateGame updates a game by ID for a season
//
//	@Summary	Update a game
//	@ID			update-game
//	@Tags		Games
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Param		gameID			path		string					true	"Game ID"			default(4019a7f3-7741-4d8f-b3e0-1c7f3a0a1a01)
//	@Param		game			body		api.GameRequest			true	"Game details to update"
//	@Success	200				{object}	api.GameResponse		"Game updated"
//	@Failure	400				{object}	response.ErrorResponse	"Bad request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games/{gameID} [put]
func handleUpdateGame(
	logger zerolog.Logger,
	gameService service.GameService,
	validate *validator.Validate,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		season := ctx.MustGet("season").(service.SeasonAggregate)

		gameID, err := uuid.Parse(ctx.Param("gameID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid game ID")
			return
		}

		req := &api.GameRequest{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Bad request")
			return
		}

		// Validate tags on GameRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		game, err := gameService.Update(ctx.Request.Context(), req, gameID, season)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to update game")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToGameResponse(game))
	}
}

// handleDeleteGame deletes a game by ID for a season
//
//	@Summary	Delete a game by ID
//	@ID			delete-game
//	@Tags		Games
//	@Produce	json
//	@Param		competitionID	path			string	true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path			string	true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Param		gameID			path			string	true	"Game ID"			default(4019a7f3-7741-4d8f-b3e0-1c7f3a0a1a01)
//	@Success	204				"No Content"	"Game deleted successfully"
//	@Failure	400				{object}		response.ErrorResponse	"Invalid game ID"
//	@Failure	500				{object}		response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID}/games/{gameID} [delete]
func handleDeleteGame(logger zerolog.Logger, gameService service.GameService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gameID, err := uuid.Parse(ctx.Param("gameID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid game ID")
			return
		}

		err = gameService.Delete(ctx.Request.Context(), gameID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to delete game")
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
