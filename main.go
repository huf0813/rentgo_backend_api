package main

import (
	"github.com/huf0813/rentgo_backend_api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.NewRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
