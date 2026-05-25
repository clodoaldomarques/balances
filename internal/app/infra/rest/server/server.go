package server

import (
	"github.com/clodoaldomarques/balances-api/internal/app/infra/rest/accounts"
	"github.com/clodoaldomarques/core-sdk/pkg/logger"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Server struct {
	http *echo.Echo
}

func New() *echo.Echo {
	e := echo.New()
	routes(e)
	return e
}

func routes(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	// logger interceptor
	e.Use(logger.InterceptorWithConfig(logger.InterceptorConfig{
		MaxBodySize:     5 * 1024,
		LogRequestBody:  true,
		LogResponseBody: false, // ligue só para debug
		RedactFields:    []string{"password", "token", "credit_card"},
	}))

	// health check
	e.GET("/", HealthCheck)

	// Accounts handler
	e.POST("/accounts", accounts.CreateNewAccount)
	e.PUT("/accounts/:orgID/:accountID/limits", accounts.UpdateAccountLimits)
	e.PUT("/accounts/:orgID/:accountID/status", accounts.UpdateAccountStatus)
	e.POST("/accounts/entries", accounts.ProcessEntry)
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	r, ok := i.(accounts.EntityRequest)
	if !ok {
		return nil
	}
	if err := r.Validate(); err != nil {
		return err
	}
	return nil
}
