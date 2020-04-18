package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"govuebook/backend/backend/handlers"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Static("/", "frontend/dist")

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/books", handlers.Books)

	e.Logger.Fatal(e.Start(":1323"))
}
