package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/clodoaldomarques/balances-api/configs"
	"github.com/clodoaldomarques/balances-api/internal/worker/application/consumer"
	"github.com/clodoaldomarques/balances-api/internal/worker/infra/rest/server"
	"github.com/clodoaldomarques/balances-api/pkg/logger"
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
