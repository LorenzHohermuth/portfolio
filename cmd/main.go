package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/handler"
	"github.com/lorenzhohermuth/portfolio/internal/mdparser"
	"github.com/gofor-little/env"
)

func main() {
	envErr := env.Load(".env")
	if envErr != nil {
		panic(envErr)
	}
	app := echo.New()
	index := 0

	projects := mdparser.GetProjects()
	h := handler.Homehandler{index, projects}
	app.GET("/", h.HandleUserShow)
	app.POST("/carousel/next", handler.HtmxCarouselHandler{&index, 1, projects}.HandlerCarouselUpdate)
	app.POST("/carousel/previous", handler.HtmxCarouselHandler{&index, -1, projects}.HandlerCarouselUpdate)
	app.Static("/static", "assets")

	app.Start(":3030")
}
