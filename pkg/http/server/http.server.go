package server

import (
	"fmt"
	spec "github.com/ESPOIR-DITE/ditkay-eshop/api/server/ditkay-eshop-api"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/logger"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type HttpServer interface {
	Start() error
}

type HttpServerImpl struct {
	ServiceConfig config.AppConfig
	Controller    spec.ServerInterface
}

var _ HttpServer = &HttpServerImpl{}

func NewHttpServerImpl(
	serviceConfig config.AppConfig,
	whsJwoConnectorController spec.ServerInterface,
) *HttpServerImpl {
	return &HttpServerImpl{
		ServiceConfig: serviceConfig,
		Controller:    whsJwoConnectorController,
	}
}

func initMiddleware(e *echo.Echo) error {
	initLogger(e)
	initRecoverFunction(e)
	return initRequestValidator(e)
}

func initRecoverFunction(e *echo.Echo) {
	e.Use(echomiddleware.RecoverWithConfig(echomiddleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
}
func initLogger(e *echo.Echo) {
	e.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values echomiddleware.RequestLoggerValues) error {
			if values.URI != "/health" {
				logger.Log.WithFields(logrus.Fields{
					"URI":    values.URI,
					"status": values.Status,
				}).Info("request")
			}
			return nil
		},
	}))

	// Middleware for logging requests and responses
	e.Use(echomiddleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		requestUri := c.Request().RequestURI
		if requestUri != "/health" {
			logger.Log.WithFields(logrus.Fields{
				"method":   c.Request().Method,
				"request":  string(reqBody),
				"response": string(resBody),
			}).Debugf("request and response(%d) for URI=%s", c.Response().Status, requestUri)
		}

	}))
}

func initRequestValidator(e *echo.Echo) error {
	swagger, err := spec.GetSwagger()
	if err != nil {
		return fmt.Errorf("Error loading application swagger spec: %s", err.Error())
	}
	swagger.Servers = nil
	e.Use(middleware.OapiRequestValidator(swagger))
	return nil

}
func (whs *HttpServerImpl) Start() error {

	e := echo.New()
	spec.RegisterHandlers(e, whs.Controller)

	if err := initMiddleware(e); err != nil {
		return err
	}

	if err := e.Start(fmt.Sprintf(":%d", whs.ServiceConfig.Port())); err != http.ErrServerClosed {
		return err
	}
	return nil
}
