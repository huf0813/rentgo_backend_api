package main

import (
	"fmt"
	"github.com/huf0813/rentgo_backend_api/infra/app_driver"
	"github.com/huf0813/rentgo_backend_api/infra/auth"
	"github.com/huf0813/rentgo_backend_api/infra/mysql_driver"
	"github.com/huf0813/rentgo_backend_api/routes"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	e := echo.New()

	db, err := mysql_driver.NewMysqlDriver()
	if err != nil {
		panic(err)
	}

	authMiddleware, err := auth.NewAuthMiddleware()
	if err != nil {
		panic(err)
	}

	timeOut := 10 * time.Second

	routes.NewRoutes(e, db, timeOut, authMiddleware)

	appDriver, err := app_driver.NewAppDriver()
	if err != nil {
		panic(err)
	}
	port := fmt.Sprintf(":%s", appDriver.AppPort)

	e.Logger.Fatal(e.Start(port))
	//waw := domain.Product{}
	//db.Preload(clause.Associations).
	//	Preload("InvoiceReviews.").
	//	Preload("ProductImages").
	//	First(&waw, 2)
	//fmt.Println(waw.Name)
	//fmt.Println(len(waw.ProductImages))
}
