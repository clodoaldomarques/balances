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
	r := mysqldb.NewRepository(ctx)
	defer r.Close()
	s := accounts.NewService(r)

	a := new(PostAccountRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(a); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	acc := a.ToEntity()

	newAcc, err := s.CreateNewAccount(ctx, acc)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPostAccountResponse(newAcc)
	return c.JSON(http.StatusCreated, resp)
}

func UpdateAccountLimits(c echo.Context) error {
	ctx := c.Request().Context()
	r := mysqldb.NewRepository(ctx)
	defer r.Close()
	s := accounts.NewService(r)

	accID, err := strconv.ParseInt(c.Param("accountID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	a := new(PutAccountLimitsRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(a); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	acc, err := s.UpdateAccountLimits(ctx, accID, a.Limits)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPutAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}

func UpdateAccountStatus(c echo.Context) error {
	ctx := c.Request().Context()
	r := mysqldb.NewRepository(ctx)
	defer r.Close()
	s := accounts.NewService(r)

	accID, err := strconv.ParseInt(c.Param("accountID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	a := new(PutAccountStatusRequest)
	if err := c.Bind(a); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(a); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	acc, err := s.UpdateAccountStatus(ctx, accID, accounts.Status(a.Status))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPutAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}

func ProcessEntry(c echo.Context) error {
	ctx := c.Request().Context()
	r := mysqldb.NewRepository(ctx)
	defer r.Close()
	s := accounts.NewService(r)

	e := new(PostEntryRequest)
	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(e); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	acc, err := s.ProcessEntry(ctx, e.ToEntity())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := AccountToPostAccountResponse(acc)

	return c.JSON(http.StatusOK, resp)
}
