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

// handleCreateSeason creates a new season with the provided details
//
//	@Summary	Create a new season
//	@ID			create-season
//	@Tags		Seasons
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		season			body		api.SeasonRequest		true	"Season details to create"
//	@Success	201				{object}	api.SeasonResponse		"Successful operation"
//	@Failure	400				{object}	response.ErrorResponse	"Bad request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons [post]
func handleCreateSeason(
	logger zerolog.Logger,
	validate *validator.Validate,
	seasonService service.SeasonService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID")
			return
		}

		req := &api.SeasonRequest{}
		err = ctx.ShouldBindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		// Validate tags on SeasonRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		season, err := seasonService.Create(ctx.Request.Context(), req, competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to add season")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, service.ToSeasonResponse(season))
	}
}

// handleGetSeasons retrieves all seasons for a given competition
//
//	@Summary	Retrieve all seasons for a competition
//	@ID			get-seasons
//	@Tags		Seasons
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Success	200				{array}		api.SeasonResponse		"List of seasons"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons [get]
func handleGetSeasons(logger zerolog.Logger, seasonService service.SeasonService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID")
			return
		}

		seasons, err := seasonService.GetAll(ctx.Request.Context(), competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get seasons")
			return
		}

		seasonsResponse := make([]api.SeasonResponse, 0, len(seasons))
		for _, season := range seasons {
			seasonsResponse = append(seasonsResponse, service.ToSeasonResponse(season))
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, seasonsResponse)
	}
}

// handleGetSeason retrieves a season by ID
//
//	@Summary	Get a single season by ID
//	@ID			get-season
//	@Tags		Seasons
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Success	200				{object}	api.SeasonResponse		"Season found"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid season ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID} [get]
func handleGetSeason(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		season := ctx.MustGet("season").(service.SeasonAggregate)

		response.RespondSuccess(ctx, logger, http.StatusOK, service.ToSeasonResponse(season))
	}
}

// handleUpdateSeason updates an existing season
//
//	@Summary	Update an existing season
//	@ID			update-season
//	@Tags		Seasons
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path		string					true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Param		season			body		api.SeasonRequest		true	"Season details to update"
//	@Success	200				{object}	api.SeasonResponse		"Season updated"
//	@Failure	400				{object}	response.ErrorResponse	"Bad request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID} [put]
func handleUpdateSeason(
	logger zerolog.Logger,
	validate *validator.Validate,
	seasonService service.SeasonService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID")
			return
		}

		seasonID, err := uuid.Parse(ctx.Param("seasonID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid season ID")
			return
		}

		req := &api.SeasonRequest{}
		err = ctx.ShouldBindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		// Validate tags on SeasonRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		season, err := seasonService.Update(ctx.Request.Context(), req, competitionID, seasonID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to update season")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, service.ToSeasonResponse(season))
	}
}

// handleDeleteSeason deletes a season by ID
//
//	@Summary	Delete a season by ID
//	@ID			delete-season
//	@Tags		Seasons
//	@Produce	json
//	@Param		competitionID	path			string	true	"Competition ID"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		seasonID		path			string	true	"Season ID"			default(9300778f-cce0-4efe-af6c-e399d8170315)
//	@Success	204				"No Content"	"Season deleted successfully"
//	@Failure	400				{object}		response.ErrorResponse	"Invalid season ID"
//	@Failure	500				{object}		response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID}/seasons/{seasonID} [delete]
func handleDeleteSeason(logger zerolog.Logger, seasonService service.SeasonService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		seasonID, err := uuid.Parse(ctx.Param("seasonID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid season ID")
			return
		}

		err = seasonService.Delete(ctx.Request.Context(), seasonID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to delete season")
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
