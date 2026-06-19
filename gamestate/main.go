package main

import (
	"net"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/bradley-adams/gainline/gamestate/redis"
	"github.com/bradley-adams/gainline/gamestate/server"
	gamestatev1 "github.com/bradley-adams/gainline/proto/gen/gamestate/v1"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const serviceName = "gainline-gamestate"

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

	redisClient := setupRedisClient(logger)
	gameStateServer := server.New(redisClient)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to listen")
	}

	grpcServer := grpc.NewServer()
	gamestatev1.RegisterGameStateServiceServer(grpcServer, gameStateServer)

	logger.Info().Msg(serviceName + " started")
	logger.Fatal().Err(grpcServer.Serve(lis)).Msg("failed to start server")
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

func setupRedisClient(logger zerolog.Logger) *redis.Client {
	logger.Info().Msg("setting up redis client...")

	addr := viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT")

	client, err := redis.New(addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to redis")
	}

	logger.Info().Msg("redis connection established")

	return client
}
