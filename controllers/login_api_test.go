package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/login-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_LoginApiController_GetLoginById(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/logins/ping", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		dbSession.Close()
		dbSession.Rollback()
	}()

	test.Ok(t, LoginApiController{}.ping(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  string                 `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result, "ping")
}
