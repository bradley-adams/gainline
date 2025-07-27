package handlers

import (
	"net/http"

	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/response.go"
	"github.com/bradley-adams/gainline/service"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
	"github.com/rs/zerolog"
)

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

		competition, serviceErr := service.CreateCompetition(ctx, db, req)
		if serviceErr != nil {
			response.RespondError(ctx, logger, err, 500, "Unable to add competition")
			return
		}

		response.RespondSuccess(ctx, logger, http.StatusCreated, api.CompetitionResponse{
			ID:        competition.ID.String(),
			Name:      competition.Name,
			CreatedAt: competition.CreatedAt,
			UpdatedAt: competition.UpdatedAt,
			DeletedAt: zero.TimeFrom(competition.DeletedAt.Time),
		})
	}
}
