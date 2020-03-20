package config

import (
	"os"
)

var (
	AppEnv             string
	Httpport           string
	Service            string
	DataBaseDriver     string
	LoginApiConnection string
)

func Read() {
	if appEnv := os.Getenv("APP_ENV"); appEnv == "" {
		AppEnv = "test"
	} else {
		AppEnv = appEnv
	}
	if httpport := os.Getenv("HTTP_PORT"); httpport == "" {
		Httpport = "8002"
	} else {
		Httpport = httpport
	}
	if dataBaseDriver := os.Getenv("DATABASE_DRIVER"); dataBaseDriver == "" {
		DataBaseDriver = "sqlite3"
	} else {
		DataBaseDriver = dataBaseDriver
	}
	if loginApiConnection := os.Getenv("LOGIN_API_CONNECTION"); loginApiConnection == "" {
		LoginApiConnection = ":memory:"
	} else {
		LoginApiConnection = loginApiConnection
	}
	Service = "login-api"
}

func ReadForTest() {
	AppEnv = "test"
	Httpport = "8002"
	Service = "login-api"
	DataBaseDriver = "sqlite3"
	LoginApiConnection = ":memory:"
}
