package controller

import (
	"encoding/json"
	"errors"
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (n NzediApiController) DeleteCompany(ctx echo.Context) error {
	var company spec.Company
	logger.Log.Info("Company receives deletion operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&company); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if company.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createCompany, err := n.CompanyFactory.CreateCompany(company)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.CompanyService.DeleteCompany(*createCompany)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetCompany(ctx echo.Context) error {
	logger.Log.Info("Company receives getAll operation.")

	deleteResult, err := n.CompanyService.ReadCompanyAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PatchCompany(ctx echo.Context) error {
	var company spec.Company
	logger.Log.Info("Company receives update operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&company); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if company.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createCompany, err := n.CompanyFactory.CreateCompany(company)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	deleteResult, err := n.CompanyService.UpdateCompany(*createCompany)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) PostCompany(ctx echo.Context) error {
	var company spec.Company
	logger.Log.Info("Company receives create operation.")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&company); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if company.Id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("missing required value"))
	}
	createCompany, err := n.CompanyFactory.CreateCompany(company)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	createResult, err := n.CompanyService.CreateCompany(*createCompany)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, createResult)
}

func (n NzediApiController) GetCompanyUserId(ctx echo.Context, userId string) error {
	logger.Log.Info("Company receives get by user/manager operation.")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("Error, missing value"))
	}
	deleteResult, err := n.CompanyService.ReadCompanyWithUserId(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}

func (n NzediApiController) GetCompanyId(ctx echo.Context, id string) error {
	logger.Log.Info("Company receives get by user/manager operation.")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, errors.New("Error, missing value"))
	}

	deleteResult, err := n.CompanyService.ReadCompany(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, deleteResult)
}
