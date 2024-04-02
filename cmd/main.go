package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/handler"
)

func main() {
	app := echo.New()

	h := handler.Homehandler{}
	app.GET("/", h.HandleUserShow)
	app.Static("/static", "assets")

	app.Start(":3030")
}
