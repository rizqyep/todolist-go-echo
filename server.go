package main

import (
	"echo-learn/routes"
	"echo-learn/utils"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()
	envMap := utils.GetEnv()
	routes.RouteHandlers(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", envMap["PORT"])))

}
