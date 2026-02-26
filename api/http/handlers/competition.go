package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	validate *validator.Validate,
	competitionService service.CompetitionService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &api.CompetitionRequest{}
		if err := ctx.ShouldBindJSON(req); err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		if err := validate.Struct(req); err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "invalid request")
			return
		}

		competition, err := competitionService.Create(ctx.Request.Context(), req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to add competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, api.ToCompetitionResponse(competition))
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
func handleGetCompetitions(logger zerolog.Logger, competitionService service.CompetitionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitions, err := competitionService.GetAll(ctx.Request.Context())
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get competitions")
			return
		}

		competitionResponse := make([]api.CompetitionResponse, 0, len(competitions))
		for _, competition := range competitions {
			competitionResponse = append(competitionResponse, api.ToCompetitionResponse(competition))
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, competitionResponse)
	}
}

// handleGetCompetitions2 retrieves competitions (paginated)
//
//	@Summary	Retrieve competitions 2
//	@ID			get-competitions2
//	@Tags		Competitions
//	@Produce	json
//	@Param		page		query		int	false	"Page number"		default(1)
//	@Param		page_size	query		int	false	"Items per page"	default(20)
//	@Success	200			{object}	api.PaginatedResponse[api.CompetitionResponse]
//	@Failure	500			{object}	response.ErrorResponse
//	@Router		/competitions2 [get]
func handleGetCompetitions2(
	logger zerolog.Logger,
	competitionService service.CompetitionService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, pageSize, offset := getPagination(ctx)

		competitions, total, err := competitionService.GetAllPaginated(
			ctx.Request.Context(),
			pageSize,
			offset,
		)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get competitions")
			return
		}

		data := make([]api.CompetitionResponse, 0, len(competitions))
		for _, competition := range competitions {
			data = append(data, api.ToCompetitionResponse(competition))
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		response.RespondSuccess(ctx, logger, http.StatusOK, api.PaginatedResponse[api.CompetitionResponse]{
			Data: data,
			Pagination: api.PaginationMeta{
				Page:       page,
				PageSize:   pageSize,
				Total:      total,
				TotalPages: totalPages,
			},
		})
	}
}

// handleGetCompetition retrieves a competition by ID
//
//	@Summary	Get a single competition by ID
//	@ID			get-competition
//	@Tags		Competitions
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Success	200				{object}	api.CompetitionResponse	"Competition found"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid competition ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [get]
func handleGetCompetition(logger zerolog.Logger, competitionService service.CompetitionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		competition, err := competitionService.Get(ctx.Request.Context(), competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to get competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToCompetitionResponse(competition))
	}
}

// handleUpdateCompetition updates a competition by ID
//
//	@Summary	Update an existing competition
//	@ID			update-competition
//	@Tags		Competitions
//	@Accept		json
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"	default(44dd315c-1abc-43aa-9843-642f920190d1)
//	@Param		competition		body		api.CompetitionRequest	true	"Competition details to update"
//	@Success	200				{object}	api.CompetitionResponse	"Competition updated"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid request"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [put]
func handleUpdateCompetition(
	logger zerolog.Logger,
	validate *validator.Validate,
	competitionService service.CompetitionService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &api.CompetitionRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "bad request")
			return
		}

		// Validate tags on CompetitionRequest struct
		err = validate.Struct(req)
		if err != nil {
			response.RespondError(ctx, logger, err, 400, "invalid request")
			return
		}

		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		competition, err := competitionService.Update(ctx.Request.Context(), competitionID, req)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to update competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusOK, api.ToCompetitionResponse(competition))
	}
}

// handleDeleteCompetition deletes a competition by ID
//
//	@Summary	Delete a competition
//	@ID			delete-competition
//	@Tags		Competitions
//	@Produce	json
//	@Param		competitionID	path		string					true	"UUID of the competition"	default(a973dd2c-ecd3-4578-b5c3-9022a3f0ecbd)
//	@Success	204				{string}	string					"Successfully deleted"
//	@Failure	400				{object}	response.ErrorResponse	"Invalid competition ID"
//	@Failure	500				{object}	response.ErrorResponse	"Internal server error"
//	@Router		/competitions/{competitionID} [delete]
func handleDeleteCompetition(logger zerolog.Logger, competitionService service.CompetitionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID, err := uuid.Parse(ctx.Param("competitionID"))
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID format")
			return
		}

		err = competitionService.Delete(ctx.Request.Context(), competitionID)
		if err != nil {
			response.RespondError(ctx, logger, err, http.StatusInternalServerError, "Unable to delete competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusNoContent, nil)
	}
}

func getPagination(ctx *gin.Context) (page, pageSize, offset int) {
	page = 1
	pageSize = 20

	if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
		page = p
	}

	if ps, err := strconv.Atoi(ctx.Query("page_size")); err == nil && ps > 0 && ps <= 100 {
		pageSize = ps
	}

	offset = (page - 1) * pageSize
	return
}
