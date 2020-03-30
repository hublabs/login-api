package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type LoginApiController struct {
}

func (c LoginApiController) Init(e *echo.Echo) {
	e.GET("/v1/logins/ping", c.ping)
}

func (c LoginApiController) ping(ctx echo.Context) error {
	return renderSucc(ctx, http.StatusOK, "ping")
}
