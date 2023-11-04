package main

import (
	"fmt"
	"recything/app/config"
	"recything/app/database"
	"recything/app/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var cfg = config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	route.NewRoute(e, dbMysql)
	database.InitMigrationMysql(dbMysql)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
