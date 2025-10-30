# ─── Dockerfile (production) ───
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -buildvcs=false -o server ./cmd/app

# ─── Runtime stage ───
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
