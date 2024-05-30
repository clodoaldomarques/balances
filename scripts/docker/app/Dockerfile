## build
FROM golang:1.20 AS build

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
COPY .env /

#mysql access
ENV MYSQL_DB_HOST=localhost
ENV MYSQL_DB_PORT=3306
ENV MYSQL_DB_NOME=balances
ENV MYSQL_DB_USUARIO=admin
ENV MYSQL_DB_SENHA=b4l4nc3s

#AWS 
ENV AWS_ADDRESS="http://192.168.49.2:4566"
ENV AWS_REGION="us-east-1"
ENV AWS_PROFILE="localstack"
ENV AWS_ID="test"
ENV AWS_SECRET="test"

#SnS
ENV CREATE_ACCOUNT_SNS_TOPIC="arn:aws:sns:us-east-1:000000000000:create-account-sns-topic"
ENV CHANGE_AVAILABLE_SNS_TOPIC="arn:aws:sns:us-east-1:000000000000:change-available-sns-topic"

EXPOSE 5000

ENTRYPOINT [ "/balances-api" ]