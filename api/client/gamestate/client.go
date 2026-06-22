package gamestate

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gamestatev1 "github.com/bradley-adams/gainline/proto/gen/gamestate/v1"
)

type Client struct {
	client gamestatev1.GameStateServiceClient
}

func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		client: gamestatev1.NewGameStateServiceClient(conn),
	}, nil
}

func (c *Client) UpdateGameState(ctx context.Context, state *gamestatev1.GameState) error {
	_, err := c.client.UpdateGameState(ctx, &gamestatev1.UpdateGameStateRequest{
		State: state,
	})
	return err
}

func (c *Client) WatchGameState(ctx context.Context, gameID string) (<-chan *gamestatev1.GameState, error) {
	stream, err := c.client.WatchGameState(ctx, &gamestatev1.WatchGameStateRequest{
		GameId: gameID,
	})
	if err != nil {
		return nil, err
	}

	out := make(chan *gamestatev1.GameState)

	go func() {
		defer close(out)
		for {
			state, err := stream.Recv()
			if err != nil {
				return
			}
			select {
			case out <- state:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}
