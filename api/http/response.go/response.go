package response

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RespondError(ctx *gin.Context, logger zerolog.Logger, err error, statusCode int, errMessage string) {
	logger.Error().Msg(err.Error())

	ctx.JSON(statusCode, gin.H{
		"message": errMessage,
	})
}

func RespondSuccess(ctx *gin.Context, logger zerolog.Logger, status int, response interface{}) {
	logger.Debug().Msgf("processed successfully")

	ctx.JSON(status, response)
}
