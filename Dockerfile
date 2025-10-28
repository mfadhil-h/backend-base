# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build binary
RUN go build -o server ./cmd/app/main.go

# --- Runtime stage ---
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
