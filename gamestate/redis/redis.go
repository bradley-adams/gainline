package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	gamestatev1 "github.com/bradley-adams/gainline/proto/gen/gamestate/v1"
)

type Client struct {
	rdb *redis.Client
}

func New(addr string) *Client {
	return &Client{
		rdb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func gameKey(gameID string) string {
	return fmt.Sprintf("game:%s", gameID)
}

func gameChannel(gameID string) string {
	return fmt.Sprintf("game:%s:updates", gameID)
}

// SetGameState stores state as JSON and publishes it to the game's channel.
func (c *Client) SetGameState(ctx context.Context, state *gamestatev1.GameState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("marshal game state: %w", err)
	}

	if err := c.rdb.Set(ctx, gameKey(state.GetGameId()), data, 0).Err(); err != nil {
		return fmt.Errorf("set game state: %w", err)
	}

	if err := c.rdb.Publish(ctx, gameChannel(state.GetGameId()), data).Err(); err != nil {
		return fmt.Errorf("publish game state: %w", err)
	}

	return nil
}

// SubscribeGameState returns a channel of state updates for gameID.
// The channel closes when ctx is cancelled.
func (c *Client) SubscribeGameState(ctx context.Context, gameID string) (<-chan *gamestatev1.GameState, error) {
	sub := c.rdb.Subscribe(ctx, gameChannel(gameID))

	out := make(chan *gamestatev1.GameState)

	go func() {
		defer close(out)
		defer sub.Close()

		ch := sub.Channel()
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-ch:
				if !ok {
					return
				}
				var state gamestatev1.GameState
				if err := json.Unmarshal([]byte(msg.Payload), &state); err != nil {
					continue
				}
				select {
				case out <- &state:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return out, nil
}
