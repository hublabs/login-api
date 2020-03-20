package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeApiController struct {
}

func (c HomeApiController) Init(e *echo.Echo) {
	e.GET("/ping", c.Ping)
}

func (c HomeApiController) Ping(ctx echo.Context) error {
	return ReturnResultApiSucc(ctx, http.StatusOK, "login-api-ping")
}
