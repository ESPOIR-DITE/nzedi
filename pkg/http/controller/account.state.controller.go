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

func (n NzediApiController) DeleteAccountState(ctx echo.Context) error {
	var state spec.AccountState
	logger.Log.Info("Account state received for deletion operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&state); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if state.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountStateFactory.CreateAccountState(state)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountStateService.DeleteAccountState(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountState(ctx echo.Context) error {
	logger.Log.Info("Account state received GetAll operation.")
	deleteResult, err := n.AccountStateService.ReadAccountStateAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchAccountState(ctx echo.Context) error {
	var state spec.AccountState
	logger.Log.Info("Account state received Update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&state); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if state.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountStateFactory.CreateAccountState(state)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountStateService.UpdateAccountState(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostAccountState(ctx echo.Context) error {
	var state spec.AccountState
	logger.Log.Info("Account state received create operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&state); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if state.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	entityAccount, err := n.AccountStateFactory.CreateAccountState(state)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountStateService.CreateAccountState(*entityAccount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetAccountStateId(ctx echo.Context, id string) error {
	logger.Log.Info("Account state received get operation.")
	accountStateId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.AccountStateService.ReadAccountState(accountStateId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
