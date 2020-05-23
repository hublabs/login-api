package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"runtime"

	"github.com/hublabs/login-api/models"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/kafka"
)

var (
	appEnv           = ""
	ctx              context.Context
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
	xormEngine       *xorm.Engine
)

func init() {
	runtime.GOMAXPROCS(1)
	var err error
	xormEngine, err = xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	models.SetModelConfig(&models.ModelConfig{
		AppEnv:       "test",
		ColleagueApi: "http://localhost:80/colleague-api",
	})

	echoApp = echo.New()
	handleWithFilter = func(handlerFunc echo.HandlerFunc, echoContext echo.Context) error {
		return echomiddleware.ContextDB("login-api", xormEngine, kafka.Config{})(handlerFunc)(echoContext)
	}
	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}

func SetContext(req *http.Request) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, xormEngine.NewSession())))

	return c, rec
}
func SetContextWithSession(req *http.Request, session *xorm.Session) (echo.Context, *httptest.ResponseRecorder) {

	rec := httptest.NewRecorder()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, session)))

	return c, rec
}
