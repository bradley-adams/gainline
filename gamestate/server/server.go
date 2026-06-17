package server

import (
	"context"

	gamestatev1 "github.com/bradley-adams/gainline/proto/gen/gamestate/v1"
)

type store interface {
	SetGameState(ctx context.Context, state *gamestatev1.GameState) error
	SubscribeGameState(ctx context.Context, gameID string) (<-chan *gamestatev1.GameState, error)
}

type Server struct {
	gamestatev1.UnimplementedGameStateServiceServer

	store store
}

func New(store store) *Server {
	return &Server{store: store}
}

// UpdateGameState writes the latest game state and broadcasts it to watchers.
func (s *Server) UpdateGameState(ctx context.Context, req *gamestatev1.UpdateGameStateRequest) (*gamestatev1.UpdateGameStateResponse, error) {
	if err := s.store.SetGameState(ctx, req.GetState()); err != nil {
		return nil, err
	}
	return &gamestatev1.UpdateGameStateResponse{}, nil
}

// WatchGameState streams state updates for a game until the client disconnects.
func (s *Server) WatchGameState(req *gamestatev1.WatchGameStateRequest, stream gamestatev1.GameStateService_WatchGameStateServer) error {
	ctx := stream.Context()

	updates, err := s.store.SubscribeGameState(ctx, req.GetGameId())
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case state, ok := <-updates:
			if !ok {
				return nil
			}
			if err := stream.Send(state); err != nil {
				return err
			}
		}
	}
}
