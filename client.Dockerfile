FROM golang:1.26.1 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./config ./config

COPY ./internal/client ./internal/client
COPY ./internal/pow ./internal/pow
COPY ./internal/protocol ./internal/protocol
COPY ./cmd/client ./cmd/client


RUN CGO_ENABLED=0 go build -o client ./cmd/client

FROM alpine:latest
WORKDIR /app

COPY .env ./.env
COPY --from=builder /app/client ./client

CMD ["./client"]
