package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/handler"
)

func main() {
	app := echo.New()
	index := 0

	h := handler.Homehandler{index}
	app.GET("/", h.HandleUserShow)
	app.POST("/carousel/next", handler.HtmxCarouselHandler{&index, 1}.HandlerCarouselUpdate)
	app.POST("/carousel/previous", handler.HtmxCarouselHandler{&index, -1}.HandlerCarouselUpdate)
	app.Static("/static", "assets")

	app.Start(":3030")
}
