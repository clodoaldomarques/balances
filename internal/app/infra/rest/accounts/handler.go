package accounts

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/app/infra/database/mysqldb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateNewAccount(c echo.Context) error {
	ctx := c.Request().Context()

	a := new(PostAccountRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	acc := a.ToEntity()

	r := mysqldb.NewRepository()
	s := accounts.NewService(r)
	newAcc, err := s.CreateNewAccount(ctx, acc)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPostAccountResponse(newAcc)
	return c.JSON(http.StatusCreated, resp)
}

func UpdateAccountLimits(c echo.Context) error {
	ctx := c.Request().Context()
	accID, err := strconv.ParseInt(c.Param("accountID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	a := new(PutAccountRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	r := mysqldb.NewRepository()
	s := accounts.NewService(r)
	acc, err := s.UpdateAccountLimits(ctx, accID, a.Limits)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPutAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}

func UpdateAccountStatus(c echo.Context) error {
	ctx := c.Request().Context()
	accID, err := strconv.ParseInt(c.Param("accountID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	a := new(PutAccountRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	r := mysqldb.NewRepository()
	s := accounts.NewService(r)
	acc, err := s.UpdateAccountStatus(ctx, accID, accounts.Status(a.Status))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPutAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}

func ProcessEntry(c echo.Context) error {
	ctx := c.Request().Context()

	e := new(PostEntryRequest)
	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	r := mysqldb.NewRepository()
	s := accounts.NewService(r)
	acc, err := s.ProcessEntry(ctx, e.ToEntity())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPostAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}
