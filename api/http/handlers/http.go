package handlers

import (
	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/docs"
	"github.com/bradley-adams/gainline/http/middleware"
	"github.com/bradley-adams/gainline/service"
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

		// middleware
		seasonService := service.NewSeasonService(db)
		v1public.Use(middleware.CompetitionStructureValidator(logger, db, seasonService))

		// competitions
		competitionService := service.NewCompetitionService(db)
		v1public.POST("/competitions", handleCreateCompetition(logger, validate, competitionService))
		v1public.GET("/competitions", handleGetCompetitions(logger, competitionService))
		v1public.GET("/competitions/:competitionID", handleGetCompetition(logger, competitionService))
		v1public.PUT("/competitions/:competitionID", handleUpdateCompetition(logger, validate, competitionService))
		v1public.DELETE("/competitions/:competitionID", handleDeleteCompetition(logger, competitionService))

		// seasons
		v1public.POST("/competitions/:competitionID/seasons", handleCreateSeason(logger, validate, seasonService))
		v1public.GET("/competitions/:competitionID/seasons", handleGetSeasons(logger, seasonService))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID", handleGetSeason(logger))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID", handleUpdateSeason(logger, validate, seasonService))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID", handleDeleteSeason(logger, seasonService))

		// games
		v1public.POST("/competitions/:competitionID/seasons/:seasonID/games", handleCreateGame(logger, db, validate))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games", handleGetGames(logger, db))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleGetGame(logger, db))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleUpdateGame(logger, db, validate))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleDeleteGame(logger, db))

		// teams
		v1public.POST("/teams", handleCreateTeam(logger, db, validate))
		v1public.GET("/teams", handleGetTeams(logger, db))
		v1public.GET("/teams/:teamID", handleGetTeam(logger, db))
		v1public.PUT("/teams/:teamID", handleUpdateTeam(logger, db, validate))
		v1public.DELETE("/teams/:teamID", handleDeleteTeam(logger, db))

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
