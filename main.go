package main

import (
	"flag"
	"os"
	"sort"
	"time"

	"github.com/hublabs/common/api"
	"github.com/hublabs/login-api/config"
	"github.com/hublabs/login-api/controllers"
	"github.com/hublabs/login-api/factory"
	"github.com/hublabs/login-api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/kafka"
	"github.com/urfave/cli/v2"
)

var (
	appEnv = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
)

func main() {
	c := config.Init(*appEnv)
	api.SetErrorMessagePrefix(c.ServiceName)

	models.SetModelConfig(&models.ModelConfig{
		AppEnv:       *appEnv,
		ColleagueApi: c.ColleagueApi,
	})

	xormEngine := initXormEngine(c.Database.Driver, c.Database.Connection)
	factory.InitDB(xormEngine)
	defer xormEngine.Close()

	app := cli.NewApp()
	app.Name = "login"
	app.Commands = []*cli.Command{
		{
			Name:  "api-server",
			Usage: "run api server",
			Action: func(cliContext *cli.Context) error {
				if err := initEchoApp(xormEngine, c.ServiceName).Start(":" + c.HttpPort); err != nil {
					return err
				}
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}

func initEchoApp(xormEngine *xorm.Engine, serviceName string) *echo.Echo {
	xormEngine.SetMaxOpenConns(50)
	xormEngine.SetMaxIdleConns(50)
	xormEngine.SetConnMaxLifetime(60 * time.Second)

	e := echo.New()

	InitControllers(e)

	e.Static("/static", "static")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	e.Use(echomiddleware.ContextDB(serviceName, xormEngine, kafka.Config{}))

	// 초기에 token 인증을 처리하지 않고 후에는 처리 되여야 함.
	// e.Use(auth.UserClaimMiddelware())

	return e
}

func InitControllers(e *echo.Echo) {
	controllers.HomeApiController{}.Init(e)
	controllers.LoginApiController{}.Init(e)
	controllers.UsernameLoginApiController{}.Init(e)
}

func initXormEngine(driver, connection string) *xorm.Engine {
	xormEngine, err := xorm.NewEngine(driver, connection)
	if err != nil {
		panic(err)
	}
	xormEngine.SetMaxIdleConns(5)
	xormEngine.SetMaxOpenConns(20)
	xormEngine.SetConnMaxLifetime(time.Minute * 10)
	//xormEngine.ShowSQL()

	return xormEngine
}
