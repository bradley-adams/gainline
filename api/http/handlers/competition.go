package handlers

import (
	"net/http"

	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/guregu/null/zero"
	"github.com/rs/zerolog"
)

// handleCreateCompetition creates a new competition with the provided details
//
//	@Summary	Create a new competition
//	@ID			create-competition
//	@Tags		Competitions
//	@Accept		json
//	@Produce	json
//	@Param		competition	body		api.CompetitionRequest	true	"Competition details to create"
//	@Success	201			{object}	api.CompetitionResponse	"Successful operation"
//	@Failure	400			{object}	response.ErrorResponse	"Bad request"
//	@Failure	404			{object}	response.ErrorResponse	"Not found"
//	@Failure	500			{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions [post]
func handleCreateCompetition(
	logger zerolog.Logger,
	db db_handler.DB,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &api.CompetitionRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "bad request")
			return
		}

		competition, err := service.CreateCompetition(ctx, db, req)
		if err != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to add competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, api.CompetitionResponse{
			ID:        competition.ID,
			Name:      competition.Name,
			CreatedAt: competition.CreatedAt,
			UpdatedAt: competition.UpdatedAt,
			DeletedAt: zero.TimeFrom(competition.DeletedAt.Time),
		})
	}
}

// handleGetCompetitions retrieves all competitions
//
//	@Summary	Retrieve all competitions
//	@ID			get-competitions
//	@Tags		Competitions
//	@Produce	json
//	@Success	200	{array}		api.CompetitionResponse	"List of competitions"
//	@Failure	500	{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions [get]
func handleGetCompetitions(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitions, err := service.GetCompetitions(ctx, db)
		if err != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to get competitions")
			return
		}

		competitionsResponseList := []api.CompetitionResponse{}
		for _, competition := range competitions {
			response := api.CompetitionResponse{
				ID:        competition.ID,
				Name:      competition.Name,
				CreatedAt: competition.CreatedAt,
				UpdatedAt: competition.UpdatedAt,
				DeletedAt: zero.TimeFrom(competition.DeletedAt.Time),
			}
			competitionsResponseList = append(competitionsResponseList, response)
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, competitionsResponseList)
	}
}

// handleGetCompetition retrieves a competition by ID
//
//	@Summary	Get a single competition by ID
//	@ID			get-competition
//	@Tags		Competitions
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"
//	@Success	200				{object}	api.CompetitionResponse	"Competition found"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid competition ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [get]
func handleGetCompetition(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionIDParam := ctx.Param("competitionID")

		competitionID, err := uuid.Parse(competitionIDParam)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		competition, err := service.GetCompetition(ctx, db, competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to get competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.CompetitionResponse{
			ID:        competition.ID,
			Name:      competition.Name,
			CreatedAt: competition.CreatedAt,
			UpdatedAt: competition.UpdatedAt,
			DeletedAt: zero.TimeFrom(competition.DeletedAt.Time),
		})
	}
}

// handleUpdateCompetition updates a competition by ID
//
//	@Summary	Update an existing competition
//	@ID			update-competition
//	@Tags		Competitions
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"
//	@Param		competition		body		api.CompetitionRequest	true	"Competition details to update"
//	@Success	200				{object}	api.CompetitionResponse	"Competition updated"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [put]
func handleUpdateCompetition(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &api.CompetitionRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "bad request")
			return
		}

		competitionIDParam := ctx.Param("competitionID")

		competitionID, err := uuid.Parse(competitionIDParam)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		competition, err := service.UpdateCompetition(ctx, db, competitionID, req)
		if err != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to update competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.CompetitionResponse{
			ID:        competition.ID,
			Name:      competition.Name,
			CreatedAt: competition.CreatedAt,
			UpdatedAt: competition.UpdatedAt,
			DeletedAt: zero.TimeFrom(competition.DeletedAt.Time),
		})
	}
}

// handleDeleteCompetition deletes a competition by ID
//
//	@Summary	Delete a competition
//	@ID			delete-competition
//	@Tags		Competitions
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"
//	@Success	204				{string}	string					"Successfully deleted"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid competition ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [delete]
func handleDeleteCompetition(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionIDParam := ctx.Param("competitionID")

		competitionID, err := uuid.Parse(competitionIDParam)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		err = service.DeleteCompetition(ctx, db, competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to delete competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusNoContent, nil)
	}
}
