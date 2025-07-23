package main

import (
	"os"
	"time"

	"github.com/bradley-adams/gainline/http/handlers"
	"github.com/rs/zerolog"
)

const serviceName = "gainline-api"

func main() {

	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}

	logger := zerolog.New(writer).With().Timestamp().Logger()

	logger.Info().Msgf("%s starting", serviceName)

	logger.Debug().Msg("setting up router...")
	r := handlers.SetupRouter(logger)

	logger.Info().Msg(serviceName + " started")
	logger.Fatal().Err(r.Run(":8080")).Msg("failed to start server")
}
