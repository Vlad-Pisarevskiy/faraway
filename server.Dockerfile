FROM golang:1.26.1 AS builder

WORKDIR /app

COPY go.sum ./
COPY go.mod ./
COPY ./config ./config

COPY ./internal/pow ./internal/pow
COPY ./internal/server ./internal/server
COPY ./internal/protocol ./internal/protocol
COPY ./internal/quotes ./internal/quotes
COPY ./cmd/server ./cmd/server

RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM alpine
WORKDIR /app

COPY .env ./.env
COPY --from=builder /app/server ./server

CMD ["./server"]