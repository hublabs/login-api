package controllers

import (
	"net/http"

	"github.com/hublabs/common/api"
	"github.com/hublabs/login-api/models"
	"github.com/labstack/echo"
)

type UserNameLoginApiController struct {
}

func (c UserNameLoginApiController) Init(g *echo.Echo) {
	//账号密码登录
	g.POST("/v1/logins/user-name", c.LoginByUserName)
}

//用账号密码登陆
func (c UserNameLoginApiController) LoginByUserName(ctx echo.Context) error {
	var loginUser struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&loginUser); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function LoginByUserName <=======================*/
	tokens, err := models.Login{}.LoginByUserName(ctx.Request().Context(), loginUser.UserName, loginUser.Password)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, tokens)
}
