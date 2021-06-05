package main

import (
	"fmt"
	"github.com/huf0813/rentgo_backend_api/infra/app_driver"
	"github.com/huf0813/rentgo_backend_api/infra/mysql_driver"
	"github.com/huf0813/rentgo_backend_api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := mysql_driver.NewMysqlDriver()
	if err != nil {
		panic(err)
	}

	routes.NewRoutes(e, db)

	appDriver, err := app_driver.NewAppDriver()
	if err != nil {
		panic(err)
	}
	port := fmt.Sprintf(":%s", appDriver.AppPort)

	e.Logger.Fatal(e.Start(port))
}
