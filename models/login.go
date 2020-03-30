package models

import (
	"context"
	"errors"
	"net/http"

	"github.com/hublabs/login-api/jwtauth"
	"github.com/labstack/echo"
)

type Login struct{}

func (Login) LoginByUserName(ctx context.Context, userName string, password string) (map[string]interface{}, error) {
	colleagueId, isLoginSuccess, err := Colleague{}.AuthenticationByUserName(ctx, userName, password)
	if err != nil {
		return nil, err
	}

	if isLoginSuccess == false || colleagueId == 0 {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: errors.New("login failed.")}
	}

	tokenInfo := map[string]interface{}{"colleagueId": colleagueId}

	token, err := jwtauth.NewToken(tokenInfo, "colleague")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"token": token}, nil
}
