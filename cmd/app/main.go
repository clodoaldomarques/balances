package main

import (
	"balances/configs"
	"balances/internal/app/infra/rest/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := configs.New()
	e := server.New()
	if err := e.Start(fmt.Sprintf(":%d", c.AppPort)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
