package main

import (
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/routes"
	"final-project-prakerja-golang-batch-11/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	utils.LoadEnv()
	configs.InitDatabase()

	e := echo.New()
	routes.InitRoutes(e)

	e.Start(":8000")
}
