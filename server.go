package main

import (
	"echo-learn/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.RouteHandlers(e)
	e.Logger.Fatal(e.Start(":1323"))
}
