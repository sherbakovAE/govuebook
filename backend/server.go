package main

import (
	echo "github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Static("/", "frontend/dist")
	e.Logger.Fatal(e.Start(":1323"))
}
