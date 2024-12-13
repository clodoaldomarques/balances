package main

import (
	"balances/configs"
	"balances/internal/app/infra/rest/server"
	"balances/pkg/logger"
	"context"
	"fmt"
	"net/http"
)

func main() {
	c := configs.New(configs.WithAppPort(5000))
	e := server.New()
	if err := e.Start(fmt.Sprintf(":%d", c.AppPort)); err != http.ErrServerClosed {
		logger.Fatal(context.Background(), err.Error(), logger.Fields{})
	}
}
