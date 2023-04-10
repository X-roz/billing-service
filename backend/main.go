package main

import (
	"billing-service/db"
	"billing-service/internals/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Enable Cross Origin request
	e.Use(middleware.CORS())

	db.ConnectToDb()

	route.Init(e.Group("/"))
	e.Start(":9000")
}
