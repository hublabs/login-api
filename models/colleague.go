package models

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hublabs/common/api"

	"github.com/pangpanglabs/goutils/behaviorlog"
	"github.com/pangpanglabs/goutils/httpreq"
)

type Colleague struct{}

func (Colleague) AuthenticationByUsername(ctx context.Context, mode string, identiKey string, password string) (map[string]interface{}, error) {
	url := fmt.Sprintf(modelConfig.ColleagueApi + "/v1/login/token-detail")
	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	body := map[string]string{
		"mode":      mode,
		"identiKey": identiKey,
		"password":  password,
	}

	if statusCode, err := httpreq.New(http.MethodPost, url, body).WithBehaviorLogContext(behaviorlog.FromCtx(ctx)).
		Call(&v); err != nil {
		return nil, err
	} else if statusCode < 200 || statusCode >= 300 {
		return nil, api.ErrorRemoteService.New(err)
	}

	return v.Result, nil
}
