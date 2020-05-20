package controllers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/hublabs/common/api"
	"github.com/pangpanglabs/goutils/behaviorlog"

	"github.com/labstack/echo"
)

const (
	EmailMode  string = "email"
	MobileMode string = "mobile"
)

func renderFail(c echo.Context, err error) error {
	if err == nil {
		err = api.ErrorUnknown.New(nil)
	}
	behaviorlog.FromCtx(c.Request().Context()).WithError(err)
	var apiError api.Error
	if ok := errors.As(err, &apiError); ok {
		return c.JSON(apiError.Status(), api.Result{
			Success: false,
			Error:   apiError,
		})
	}
	return err
}

func renderSuccArray(c echo.Context, withHasMore, hasMore bool, totalCount int64, result interface{}) error {
	if withHasMore {
		return renderSucc(c, http.StatusOK, api.ArrayResultMore{
			HasMore: hasMore,
			Items:   result,
		})
	} else {
		return renderSucc(c, http.StatusOK, api.ArrayResult{
			TotalCount: totalCount,
			Items:      result,
		})
	}
}

func renderSucc(c echo.Context, status int, result interface{}) error {
	return c.JSON(status, api.Result{
		Success: true,
		Result:  result,
	})
}

//email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func GetUsernameLoginMode(username string) (string, error) {
	if VerifyEmailFormat(username) {
		return EmailMode, nil
	}
	if VerifyMobileFormat(username) {
		return MobileMode, nil
	}
	return "", errors.New("invalid account.")
}
