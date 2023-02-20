package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/roviery/catetin-api/routes"
)

func main() {
	e := routes.Init()
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start("localhost:8080"))
}
