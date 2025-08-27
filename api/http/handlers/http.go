package handlers

import (
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes and configures the HTTP router for handling incoming requests
//
//	@title			Gainline Api
//	@description	A set of endpoints for managing gainline tasks
//	@version		1.0
func SetupRouter(db db_handler.DB, logger zerolog.Logger, validate *validator.Validate) *gin.Engine {
	logger.Debug().Msg("setting up http router...")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/health", healthCheck(logger))

	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1public := router.Group("/v1").Use(cors.Default())
	{
		v1public.OPTIONS("/*path")

		//competitions
		v1public.POST("/competitions", handleCreateCompetition(logger, db, validate))
		v1public.GET("/competitions", handleGetCompetitions(logger, db))
		v1public.GET("/competitions/:competitionID", handleGetCompetition(logger, db))
		v1public.PUT("/competitions/:competitionID", handleUpdateCompetition(logger, db))
		v1public.DELETE("/competitions/:competitionID", handleDeleteCompetition(logger, db))

		//seasons
		v1public.POST("/competitions/:competitionID/seasons", handleCreateSeason(logger, db, validate))
		v1public.GET("/competitions/:competitionID/seasons", handleGetSeasons(logger, db))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID", handleGetSeason(logger, db))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID", handleUpdateSeason(logger, db))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID", handleDeleteSeason(logger, db))

		//teams
		v1public.POST("/teams", handleCreateTeam(logger, db))
		v1public.GET("/teams", handleGetTeams(logger, db))
		v1public.GET("/teams/:teamID", handleGetTeam(logger, db))
		v1public.PUT("/teams/:teamID", handleUpdateTeam(logger, db))
		v1public.DELETE("/teams/:teamID", handleDeleteTeam(logger, db))

		// games
		v1public.POST("/competitions/:competitionID/seasons/:seasonID/games", handleCreateGame(logger, db))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games", handleGetGames(logger, db))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleGetGame(logger, db))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleUpdateGame(logger, db))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleDeleteGame(logger, db))
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
