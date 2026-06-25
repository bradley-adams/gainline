package handlers

import (
	"net/http"

	"github.com/bradley-adams/gainline/client/gamestate"
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

type RouterConfig struct {
	DB              db_handler.DB
	Logger          zerolog.Logger
	Validate        *validator.Validate
	GameStateClient *gamestate.Client
}

// SetupRouter initializes and configures the HTTP router for handling incoming requests
//
//	@title			Gainline Api
//	@description	A set of endpoints for managing gainline tasks
//	@version		1.0
func SetupRouter(cfg RouterConfig) *gin.Engine {
	cfg.Logger.Debug().Msg("setting up http router...")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/health", healthCheck(cfg.DB, cfg.Logger))

	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1public := router.Group("/v1").Use(cors.Default())
	{
		v1public.OPTIONS("/*path")

		// middleware
		seasonService := service.NewSeasonService(cfg.DB)
		gameService := service.NewGameService(cfg.DB)
		gameStateService := service.NewGameStateService(cfg.GameStateClient)
		v1public.Use(middleware.CompetitionStructureValidator(cfg.Logger, seasonService, gameService))

		// competitions
		competitionService := service.NewCompetitionService(cfg.DB)
		v1public.POST("/competitions", handleCreateCompetition(cfg.Logger, cfg.Validate, competitionService))
		v1public.GET("/competitions", handleGetCompetitions(cfg.Logger, cfg.Validate, competitionService))
		v1public.GET("/competitions/:competitionID", handleGetCompetition(cfg.Logger, competitionService))
		v1public.PUT("/competitions/:competitionID", handleUpdateCompetition(cfg.Logger, cfg.Validate, competitionService))
		v1public.DELETE("/competitions/:competitionID", handleDeleteCompetition(cfg.Logger, competitionService))

		// seasons
		v1public.POST("/competitions/:competitionID/seasons", handleCreateSeason(cfg.Logger, cfg.Validate, seasonService))
		v1public.GET("/competitions/:competitionID/seasons", handleGetSeasons(cfg.Logger, cfg.Validate, seasonService))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID", handleGetSeason(cfg.Logger))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID", handleUpdateSeason(cfg.Logger, cfg.Validate, seasonService))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID", handleDeleteSeason(cfg.Logger, seasonService))

		// games
		v1public.POST("/competitions/:competitionID/seasons/:seasonID/games", handleCreateGame(cfg.Logger, cfg.Validate, gameService))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/stages/:stageID/games", handleGetGames(cfg.Logger, gameService))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleGetGame(cfg.Logger, gameService))
		v1public.PUT("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleUpdateGame(cfg.Logger, gameService, cfg.Validate))
		v1public.DELETE("/competitions/:competitionID/seasons/:seasonID/games/:gameID", handleDeleteGame(cfg.Logger, gameService))
		v1public.GET("/competitions/:competitionID/seasons/:seasonID/games/:gameID/live", handleWatchGame(cfg.Logger, gameStateService))

		// teams
		teamService := service.NewTeamService(cfg.DB)
		v1public.POST("/teams", handleCreateTeam(cfg.Logger, cfg.Validate, teamService))
		v1public.GET("/teams", handleGetTeams(cfg.Logger, cfg.Validate, teamService))
		v1public.GET("/teams/:teamID", handleGetTeam(cfg.Logger, teamService))
		v1public.PUT("/teams/:teamID", handleUpdateTeam(cfg.Logger, cfg.Validate, teamService))
		v1public.DELETE("/teams/:teamID", handleDeleteTeam(cfg.Logger, teamService))
	}

	return router
}

func healthCheck(db db_handler.DB, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := db.HealthCheck(); err != nil {
			logger.Error().
				Err(err).
				Msg("database health check failed")

			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	}
}
