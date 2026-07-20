FROM golang:1.26.1 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY ./internal/client ./internal/client
COPY ./internal/pow ./internal/pow
COPY ./internal/protocol ./internal/protocol
COPY ./cmd/client ./cmd/client


RUN go build -o client ./cmd/client

FROM alpine:latest

COPY --from=builder /app/client /client
CMD ["/client"]
