package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	getHelloHandler()
}

func getHelloHandler() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", helloHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world.")
}
