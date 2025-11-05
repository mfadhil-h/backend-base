# ─── Stage 1: Builder ───
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/air-verse/air@latest
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/app

# ─── Stage 2: Development ───
FROM builder AS dev
WORKDIR /app
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# ─── Stage 3: Production ───
FROM alpine:3.19 AS prod
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
