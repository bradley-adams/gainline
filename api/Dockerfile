# Stage 1: Build the Go binary
FROM golang:1.24.1 AS builder

WORKDIR /app
COPY . .
RUN go build -o app

# Stage 2: Create a lightweight image to run the app
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
