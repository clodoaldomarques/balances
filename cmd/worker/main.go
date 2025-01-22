package main

import (
	"balances/configs"
	"balances/internal/worker/application/consumer"
	"balances/internal/worker/infra/rest/server"
	"balances/pkg/logger"
	"context"
	"fmt"
	"net/http"
)

func main() {
	ctx := context.Background()
	c := configs.New(configs.WithAppPort(5001), configs.WithMysqlDBName("worker"))
	setupSQSConsumer(ctx)
	setupHttpServer(ctx, c)
}

func setupHttpServer(ctx context.Context, c *configs.Config) {
	e := server.New()
	if err := e.Start(fmt.Sprintf(":%d", c.AppPort)); err != http.ErrServerClosed {
		logger.Fatal(ctx, err.Error(), logger.Fields{})
	}
}

func setupSQSConsumer(ctx context.Context) {
	go consumer.New(ctx).Start()
}
