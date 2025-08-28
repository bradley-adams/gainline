package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/bradley-adams/gainline/http/response"
	"github.com/bradley-adams/gainline/service"
)

func CompetitionStructureValidator(logger zerolog.Logger, db db_handler.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		competitionID := ctx.Param("competitionID")
		seasonID := ctx.Param("seasonID")
		gameID := ctx.Param("gameID")

		var compUUID, seasonUUID, gameUUID uuid.UUID
		var err error

		if competitionID != "" {
			compUUID, err = uuid.Parse(competitionID)
			if err != nil {
				response.RespondAbortError(ctx, logger, err, http.StatusBadRequest, "Invalid competition ID")
				return
			}
		}

		if seasonID != "" {
			seasonUUID, err = uuid.Parse(seasonID)
			if err != nil {
				response.RespondAbortError(ctx, logger, err, http.StatusBadRequest, "Invalid season ID")
				return
			}

			err := validateSeason(ctx.Request.Context(), db, seasonUUID, compUUID)
			if err != nil {
				response.RespondAbortError(ctx, logger, err, http.StatusForbidden, "Season does not belong to competition")
				return
			}
		}

		if gameID != "" {
			gameUUID, err = uuid.Parse(gameID)
			if err != nil {
				response.RespondAbortError(ctx, logger, err, http.StatusBadRequest, "Invalid game ID")
				return
			}

			err := validateGame(ctx.Request.Context(), db, gameUUID, seasonUUID)
			if err != nil {
				response.RespondAbortError(ctx, logger, err, http.StatusForbidden, "Game does not belong to season")
				return
			}
		}

		ctx.Next()
	}
}

func validateGame(ctx context.Context, db db_handler.DB, gameID, seasonID uuid.UUID) error {
	game, err := service.GetGame(ctx, db, gameID)
	if err != nil {
		return err
	}
	if game.SeasonID != seasonID {
		return errors.New("unauthorized: game does not belong to season")
	}
	return nil
}

func validateSeason(ctx context.Context, db db_handler.DB, seasonID, competitionID uuid.UUID) error {
	season, err := service.GetSeason(ctx, db, competitionID, seasonID)
	if err != nil {
		return err
	}
	if season.CompetitionID != competitionID {
		return errors.New("unauthorized: season does not belong to competition")
	}
	return nil
}
