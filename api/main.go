package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/bradley-adams/gainline/docs"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"github.com/bradley-adams/gainline/db/db_handler"
	"github.com/bradley-adams/gainline/http/api"
	"github.com/bradley-adams/gainline/http/handlers"
	"github.com/bradley-adams/gainline/http/validation"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const serviceName = "gainline-api"

func main() {
	time.Local = time.UTC
	err := setUpEnvVars()
	if err != nil {
		panic(err)
	}

	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}

	port := viper.GetString("PORT")

	logger := zerolog.New(writer).With().Str("service", serviceName).Str("port", port).Timestamp().Logger()

	logger.Info().Msgf("%s starting", serviceName)

	db := setupWrapperDB(logger)
	validate, err := setUpValidator(logger)
	if err != nil {
		panic(err)
	}

	logger.Debug().Msg("setting up router...")
	r := handlers.SetupRouter(db, logger, validate)

	logger.Info().Msg(serviceName + " started")
	logger.Fatal().Err(r.Run(":8080")).Msg("failed to start server")
}

func setUpEnvVars() error {
	viper.AllowEmptyEnv(true)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func setupWrapperDB(logger zerolog.Logger) *db_handler.DBWrapper {
	logger.Info().Msg("setting up WrapperDB...")

	return &db_handler.DBWrapper{
		DB: setupDB(logger).DB,
	}
}

func setupDB(logger zerolog.Logger) *sqlx.DB {
	logger.Info().Msg("starting database connection...")

	dbURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"),
		viper.GetString("DB_SSL_MODE"),
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
	}

	return db
}

func setUpValidator(logger zerolog.Logger) (*validator.Validate, error) {
	logger.Info().Msg("setting up validator...")

	validate := validator.New()

	// Field-level validators
	validate.RegisterValidation("entity_name", validation.ValidateEntityName)
	validate.RegisterValidation("unique_team_uuids", validation.ValidateUniqueUUIDs)
	validate.RegisterValidation("game_status", validation.ValidateGameStatus)

	// Struct-level validators
	validate.RegisterStructValidation(validation.ValidateGameRequest, api.GameRequest{})

	return validate, nil
}
