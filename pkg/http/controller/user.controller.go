package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (n NzediApiController) DeleteUser(ctx echo.Context) error {
	var userObject spec.User
	logger.Log.Info("User receives deletion operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userObject); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userObject.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUser, err := n.UserFactory.CreateUser(userObject)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserService.DeleteUser(*createUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetUser(ctx echo.Context) error {
	logger.Log.Info("User receives getAll operation.")

	deleteResult, err := n.UserService.ReadUserAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchUser(ctx echo.Context) error {
	var userObject spec.User
	logger.Log.Info("User receives update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userObject); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userObject.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUser, err := n.UserFactory.CreateUser(userObject)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserService.UpdateUser(*createUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostUser(ctx echo.Context) error {
	var userObject spec.User
	logger.Log.Info("User receives create operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userObject); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userObject.Id < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUser, err := n.UserFactory.CreateUser(userObject)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserService.CreateUser(*createUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
func (n NzediApiController) GetUserAccountId(ctx echo.Context, accountId int) error {
	logger.Log.Info("User receives get with account id operation.")
	if accountId < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing value"))
	}

	deleteResult, err := n.UserService.ReadUserWithAccountId(accountId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetUsersUserId(ctx echo.Context, userId int) error {
	logger.Log.Info("User receives get operation.")
	if userId < 0 {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing value"))
	}
	deleteResult, err := n.UserService.ReadUserWithAccountId(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
