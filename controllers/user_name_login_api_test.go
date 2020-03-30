package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/hublabs/login-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_UserNameApiController_LoginByUserName(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/v1/logins/user-name",
		strings.NewReader(`{"userName":"xiao.ming", "password":"1111"}`))

	c, rec := SetContext(req)

	dbSession := factory.DB(c.Request().Context()).(*xorm.Session)
	dbSession.Begin()
	defer func() {
		dbSession.Close()
		dbSession.Rollback()
	}()

	test.Ok(t, UserNameLoginApiController{}.LoginByUserName(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, len(v.Result["token"].(string)) > 10, true)
}
