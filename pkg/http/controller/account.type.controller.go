package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (n NzediApiController) DeleteAccountType(ctx echo.Context) error {
	var state spec.AccountType
	logger.Log.Info("Account type receives deletion operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&state); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if state.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccountType, err := n.AccountTypeFactory.CreateAccountTypeFactory(state)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountTypeService.DeleteAccountType(*entityAccountType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountType(ctx echo.Context) error {
	logger.Log.Info("Account type receives getAll operation.")
	deleteResult, err := n.AccountTypeService.ReadAccountTypeAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchAccountType(ctx echo.Context) error {
	var state spec.AccountType
	logger.Log.Info("Account type receives update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&state); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if state.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccountType, err := n.AccountTypeFactory.CreateAccountTypeFactory(state)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountTypeService.UpdateAccountType(*entityAccountType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostAccountType(ctx echo.Context) error {
	var accountType spec.AccountType
	logger.Log.Info("Account type receives create operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&accountType); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if accountType.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccountType, err := n.AccountTypeFactory.CreateAccountTypeFactory(accountType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountTypeService.UpdateAccountType(*entityAccountType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountTypeAccountId(ctx echo.Context, accountId string) error {
	logger.Log.Info("Account type receives get account by accountId operation.")
	if accountId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing value"))
	}
	deleteResult, err := n.AccountTypeService.ReadAccountTypeWithAccountId(accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountTypeId(ctx echo.Context, accountId string) error {
	logger.Log.Info("Account type receives get operation.")
	if accountId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing value"))
	}

	deleteResult, err := n.AccountTypeService.ReadAccountType(accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
