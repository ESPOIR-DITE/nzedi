package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (n NzediApiController) DeleteUserType(ctx echo.Context) error {
	var userType spec.UserType
	logger.Log.Info("User type controller receives deletion operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userType); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userType.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUserType, err := n.UserTypeFactory.CreateUserType(userType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserTypeService.DeleteUserType(*createUserType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetUserType(ctx echo.Context) error {
	logger.Log.Info("User type controller receives getAll request.")

	deleteResult, err := n.UserTypeService.ReadUserTypeAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchUserType(ctx echo.Context) error {
	var userType spec.UserType
	logger.Log.Info("User type controller receives update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userType); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userType.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUserType, err := n.UserTypeFactory.CreateUserType(userType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserTypeService.UpdateUserType(*createUserType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostUserType(ctx echo.Context) error {
	var userType spec.UserType
	logger.Log.Info("User type controller receives create request.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&userType); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if userType.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createUserType, err := n.UserTypeFactory.CreateUserType(userType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.UserTypeService.CreateUserType(*createUserType)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetUserTypeId(ctx echo.Context, id string) error {
	logger.Log.Info("User type controller receives get operation.")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing value"))
	}
	deleteResult, err := n.UserTypeService.ReadUserType(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
