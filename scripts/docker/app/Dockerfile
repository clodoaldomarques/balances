## build
FROM golang:1.26.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY cmd/app/main.go .
RUN go build -o /balances-api

## deploy
FROM ubuntu AS base

WORKDIR /

COPY --from=build /balances-api /balances-api

EXPOSE 5000

ENTRYPOINT [ "/balances-api" ]