## build
FROM golang:1.22 AS build

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

#mysql access
ENV MYSQL_HOST=192.168.49.2
ENV MYSQL_PORT=30306
ENV MYSQL_DATABASE=balances
ENV MYSQL_USER=admin
ENV MYSQL_PASS=b4l4nc3s

#AWS 
ENV AWS_ADDRESS="http://192.168.49.2:4566"
ENV AWS_REGION="us-east-1"
ENV AWS_ACCESS_KEY_ID="test"
ENV AWS_SECRET_ACCESS_KEY="test"

#SnS
ENV BALANCES_SNS_TOPIC="arn:aws:sns:us-east-1:000000000000:balances-sns-topic"

EXPOSE 5000

ENTRYPOINT [ "/balances-api" ]