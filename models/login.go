package models

import (
	"context"
	"errors"

	"github.com/hublabs/login-api/jwtauth"
)

type Login struct{}

func (Login) LoginByUserName(ctx context.Context, mode string, userName string, password string) (map[string]interface{}, error) {
	tokenDetail, err := Colleague{}.AuthenticationByUserName(ctx, mode, userName, password)
	if err != nil {
		return nil, err
	}

	if !IsValidTokenDetail(tokenDetail) {
		return nil, errors.New("login failed.")
	}

	token, err := jwtauth.NewToken(tokenDetail, "colleague")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"token": token}, nil
}

func IsValidTokenDetail(tokenDetail map[string]interface{}) bool {
	if tokenDetail == nil {
		return false
	}

	if colleagueId, ok := tokenDetail["colleagueId"]; ok {
		if id, ok := colleagueId.(float64); ok && id > 0 {
			return true
		}
	}

	return false
}
