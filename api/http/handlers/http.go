package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetupRouter(logger zerolog.Logger) *gin.Engine {
	logger.Debug().Msg("setting up http router...")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", healthCheck(logger))

	return r
}

func healthCheck(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "OK",
		})

	}
}
