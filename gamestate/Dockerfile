# Stage 1: Build the Go binary
FROM golang:1.25.0 AS builder

WORKDIR /app

# Copy proto module first (needed for local replace directive)
COPY proto/ ./proto/

# Copy gamestate module
COPY gamestate/ ./gamestate/

WORKDIR /app/gamestate
RUN go build -o app

# Stage 2: Create a lightweight image to run the app
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/gamestate/app .
COPY gamestate/.env .
EXPOSE 50051
CMD ["./app"]