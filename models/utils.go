package models

import "github.com/hublabs/login-api/jwtauth"

type ModelConfig struct {
	AppEnv       string
	ColleagueApi string
}

var modelConfig *ModelConfig

func SetModelConfig(m *ModelConfig) {
	modelConfig = m
}

//TODO 환경셋팅이 되면 그떄는 사용하지 않음.
func GetTempTokenForLogin() (string, error) {
	tokenDetail := make(map[string]interface{})
	tokenDetail["tenantCode"] = "hublabs"
	tokenDetail["colleagueId"] = int64(1)
	token, err := jwtauth.NewToken(nil, "colleague")
	if err != nil {
		return "", err
	}
	return token, nil
}
