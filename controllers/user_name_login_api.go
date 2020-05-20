package controllers

import (
	"net/http"

	"github.com/hublabs/common/api"
	"github.com/hublabs/login-api/models"
	"github.com/labstack/echo"
)

type UsernameLoginApiController struct {
}

func (c UsernameLoginApiController) Init(g *echo.Echo) {
	//账号密码登录
	g.POST("/v1/logins/user-name", c.LoginByUsername)
}

//用账号密码登陆
func (c UsernameLoginApiController) LoginByUsername(ctx echo.Context) error {
	var loginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&loginUser); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	mode, err := GetUsernameLoginMode(loginUser.Username)
	if err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function LoginByUsername <=======================*/
	tokens, err := models.Login{}.LoginByUsername(ctx.Request().Context(), mode, loginUser.Username, loginUser.Password)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, tokens)
}
