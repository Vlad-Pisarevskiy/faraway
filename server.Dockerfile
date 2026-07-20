FROM golang:1.26.1 AS builder

WORKDIR /app

COPY go.sum ./
COPY go.mod ./
COPY .env ./env

COPY ./internal/server ./internal/server
COPY ./internal/protocol ./internal/protocol
COPY ./internal/quotes ./internal/quotes
COPY ./cmd/sever ./cmd/server

RUN go build -o server ./cmd/server

FROM alpine

copy --from=builder ./app/server /server

CMD["/server"]