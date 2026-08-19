# Stage 1: Build the Go binary
FROM golang:1.25.0 AS builder

WORKDIR /app

# Copy proto module first (needed for local replace directive)
COPY proto/ ./proto/

# Copy api module
COPY api/ ./api/

WORKDIR /app/api
RUN go build -o app

# Stage 2: Create a lightweight image to run the app
FROM debian:bookworm-slim

# Install CA certificates so Go can make outbound HTTPS calls to Auth0
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/api/app .
COPY api/.env .
EXPOSE 8080
CMD ["./app"]