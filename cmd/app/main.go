package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/clodoaldomarques/balances/configs"
	"github.com/clodoaldomarques/balances/internal/app/infra/rest/server"
	"github.com/clodoaldomarques/balances/pkg/logger"
)

func main() {
	c := configs.New(configs.WithAppPort(5000), configs.WithMysqlDBName("balances"))
	e := server.New()
	if err := e.Start(fmt.Sprintf(":%d", c.AppPort)); err != http.ErrServerClosed {
		logger.Fatal(context.Background(), err.Error(), logger.Fields{})
	}
}
