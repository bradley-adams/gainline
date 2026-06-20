# Gainline gamestate

## Getting Started

### Start Redis:

```bash
make redis-up
```

### Run the App Locally:

```bash
go run .
```

## Testing with grpcurl

### Watch a game's live state:

```bash
grpcurl -plaintext -d '{"game_id":"test"}' localhost:50051 gamestate.v1.GameStateService/WatchGameState
```

This will hang, waiting for updates.

### Send a state update (in a separate terminal):

```bash
grpcurl -plaintext -d '{"state":{"game_id":"test","home_score":1,"away_score":0}}' localhost:50051 gamestate.v1.GameStateService/UpdateGameState
```

The update should appear in the watching terminal.

### Check what's stored in Redis:

```bash
make redis-cli
GET game:test
```
