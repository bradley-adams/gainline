package service

import (
	"context"

	"github.com/google/uuid"

	gamestateclient "github.com/bradley-adams/gainline/client/gamestate"
	gamestatev1 "github.com/bradley-adams/gainline/proto/gen/gamestate/v1"
)

// GameStateService defines the contract for live game state operations.
type GameStateService interface {
	UpdateGameState(ctx context.Context, gameID uuid.UUID, homeScore, awayScore int32, status string) error
	WatchGameState(ctx context.Context, gameID uuid.UUID) (<-chan *gamestatev1.GameState, error)
}

// gameStateService is the concrete implementation backed by the gamestate gRPC client.
type gameStateService struct {
	client *gamestateclient.Client
}

func NewGameStateService(client *gamestateclient.Client) GameStateService {
	return &gameStateService{client: client}
}

// UpdateGameState broadcasts the latest score and status for a game.
func (s *gameStateService) UpdateGameState(ctx context.Context, gameID uuid.UUID, homeScore, awayScore int32, status string) error {
	return s.client.UpdateGameState(ctx, &gamestatev1.GameState{
		GameId:    gameID.String(),
		HomeScore: homeScore,
		AwayScore: awayScore,
		Status:    status,
	})
}

// WatchGameState returns a channel of live state updates for a game.
func (s *gameStateService) WatchGameState(ctx context.Context, gameID uuid.UUID) (<-chan *gamestatev1.GameState, error) {
	return s.client.WatchGameState(ctx, gameID.String())
}
