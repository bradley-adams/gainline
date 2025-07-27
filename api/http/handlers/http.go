package handlers

import (
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetupRouter(db db_handler.DB, logger zerolog.Logger) *gin.Engine {
	logger.Debug().Msg("setting up http router...")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/health", healthCheck(logger))

	v1public := router.Group("/v1").Use(cors.Default())
	{
		v1public.POST("/competitions", handleCreateCompetition(logger, db))
	}

	return router
}

func healthCheck(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "OK",
		})

	}
}
