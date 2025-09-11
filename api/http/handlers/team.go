package handlers

import (
	"net/http"

	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// handleCreateTeam creates a new team with the provided details
//
//	@Summary	Create a new team
//	@ID			create-team
//	@Tags		Teams
//	@Accept		json
//	@Produce	json
//	@Param		team	body		api.TeamRequest			true	"Team details to create"
//	@Success	201		{object}	api.TeamResponse		"Successful operation"
//	@Failure	400		{object}	response.ErrorResponse	"Bad request"
//	@Failure	500		{object}	response.ErrorResponse	"Internal server error"
//	@Router		/teams [post]
func handleCreateTeam(
	logger zerolog.Logger,
	db db_handler.DB,
	validate *validator.Validate,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &api.TeamRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		// Validate tags on TeamRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		team, err := service.CreateTeam(ctx, db, req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to add team")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, api.ToTeamResponse(team))
	}
}

// handleGetTeams retrieves all teams
//
//	@Summary	Retrieve all teams
//	@ID			get-teams
//	@Tags		Teams
//	@Produce	json
//	@Success	200	{array}		api.TeamResponse		"List of teams"
//	@Failure	500	{object}	response.ErrorResponse	"Internal server error"
//	@Router		/teams [get]
func handleGetTeams(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		teams, err := service.GetTeams(ctx, db)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get teams")
			return
		}

		teamsResponse := make([]api.TeamResponse, 0, len(teams))
		for _, team := range teams {
			teamsResponse = append(teamsResponse, api.ToTeamResponse(team))
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, teamsResponse)
	}
}

// handleGetTeam retrieves a team by ID
//
//	@Summary	Get a single team by ID
//	@ID			get-team
//	@Tags		Teams
//	@Produce	json
//	@Param		teamID	path		string					true	"Team ID"	default(013952a5-87e1-4d26-a312-09b2aff54241)
//	@Success	200		{object}	api.TeamResponse		"Team found"
//	@Failure	400		{object}	response.ErrorResponse	"Invalid team ID"
//	@Failure	500		{object}	response.ErrorResponse	"Internal server error"
//	@Router		/teams/{teamID} [get]
func handleGetTeam(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		teamID, err := uuid.Parse(ctx.Param("teamID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid team ID")
			return
		}

		team, err := service.GetTeam(ctx, db, teamID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get team")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToTeamResponse(team))
	}
}

// handleUpdateTeam updates an existing team
//
//	@Summary	Update an existing team
//	@ID			update-team
//	@Tags		Teams
//	@Accept		json
//	@Produce	json
//	@Param		teamID	path		string					true	"Team ID"	default(013952a5-87e1-4d26-a312-09b2aff54241)
//	@Param		team	body		api.TeamRequest			true	"Team details to update"
//	@Success	200		{object}	api.TeamResponse		"Team updated"
//	@Failure	400		{object}	response.ErrorResponse	"Bad request"
//	@Failure	500		{object}	response.ErrorResponse	"Internal server error"
//	@Router		/teams/{teamID} [put]
func handleUpdateTeam(
	logger zerolog.Logger,
	db db_handler.DB,
	validate *validator.Validate,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		teamID, err := uuid.Parse(ctx.Param("teamID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid team ID")
			return
		}

		req := &api.TeamRequest{}
		err = ctx.BindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		// Validate tags on TeamRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		team, err := service.UpdateTeam(ctx, db, req, teamID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to update team")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToTeamResponse(team))
	}
}

// handleDeleteTeam deletes a team by ID
//
//	@Summary	Delete a team by ID
//	@ID			delete-team
//	@Tags		Teams
//	@Produce	json
//	@Param		teamID	path			string	true	"Team ID"	default(013952a5-87e1-4d26-a312-09b2aff54241)
//	@Success	204		"No Content"	"Team deleted successfully"
//	@Failure	400		{object}		response.ErrorResponse	"Invalid team ID"
//	@Failure	500		{object}		response.ErrorResponse	"Internal server error"
//	@Router		/teams/{teamID} [delete]
func handleDeleteTeam(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		teamID, err := uuid.Parse(ctx.Param("teamID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid team ID")
			return
		}

		err = service.DeleteTeam(ctx, db, teamID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to delete team")
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
