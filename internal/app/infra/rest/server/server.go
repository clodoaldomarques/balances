package server

import (
	"balances/internal/app/infra/rest/accounts"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	fillRoutes(e)
	return e
}

func fillRoutes(e *echo.Echo) {

	// health check
	e.GET("/", HealthCheck)

	// Accounts handler
	e.POST("/accounts", accounts.CreateNewAccount)
	e.PUT("/accounts/:accountID/limits", accounts.UpdateAccountLimits)
	e.PUT("/accounts/:accountID/status", accounts.UpdateAccountStatus)
	e.POST("/accounts/entries", accounts.ProcessEntry)

}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
