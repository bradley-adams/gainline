package handlers

import (
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes and configures the HTTP router for handling incoming requests
//
//	@title			Gainline Api
//	@description	A set of endpoints for managing gainline tasks
//	@version		1.0
func SetupRouter(db db_handler.DB, logger zerolog.Logger) *gin.Engine {
	logger.Debug().Msg("setting up http router...")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/health", healthCheck(logger))

	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1public := router.Group("/v1").Use(cors.Default())
	{
		v1public.POST("/competitions", handleCreateCompetition(logger, db))
		v1public.GET("/competitions", handleGetCompetitions(logger, db))
		v1public.GET("/competitions/:competitionID", handleGetCompetition(logger, db))
		v1public.PUT("/competitions/:competitionID", handleUpdateCompetition(logger, db))
		v1public.DELETE("/competitions/:competitionID", handleDeleteCompetition(logger, db))
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
