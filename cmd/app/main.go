package main

import (
	"balances/internal/app/infra/rest/server"
	"log"
	"net/http"
)

func main() {
	e := server.New()

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
