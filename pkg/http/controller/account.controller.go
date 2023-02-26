package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (n NzediApiController) DeleteAccount(ctx echo.Context) error {
	var account spec.Account
	logger.Log.Info("Account received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&account); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if account.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountFactory.CreateAccount(account)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountService.DeleteAccount(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccount(ctx echo.Context) error {
	deleteResult, err := n.AccountService.ReadAccountAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchAccount(ctx echo.Context) error {
	var account spec.Account
	logger.Log.Info("Account received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&account); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if account.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountFactory.CreateAccount(account)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountService.UpdateAccount(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostAccount(ctx echo.Context) error {
	var account spec.Account
	logger.Log.Info("Account received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&account); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if account.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountFactory.CreateAccount(account)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountService.CreateAccount(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PutAccount(ctx echo.Context) error {
	var account spec.Account
	logger.Log.Info("Account received for deletion.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&account); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if account.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountFactory.CreateAccount(account)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountService.UserLogin(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountAccountId(ctx echo.Context, accountId string) error {
	if accountId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing data"))
	}
	id, err := strconv.Atoi(accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountService.ReadAccount(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
