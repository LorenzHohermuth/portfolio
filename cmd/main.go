package main

import (
	"github.com/gofor-little/env"
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/handler"
	"github.com/lorenzhohermuth/portfolio/view/component"
)

func main() {
	envErr := env.Load(".env")
	if envErr != nil {
		panic(envErr)
	}
	app := echo.New()
	index := 0

	var projects []component.CarouselEntry
	h := handler.Homehandler{Index: index, Entrys:  &projects}
	app.GET("/", h.HandleUserShow)
	app.POST("/carousel/next", handler.HtmxCarouselHandler{Index: &index,Direction:  1,Entrys:  &projects}.HandlerCarouselUpdate)
	app.POST("/carousel/previous", handler.HtmxCarouselHandler{Index: &index, Direction:   -1, Entrys:  &projects}.HandlerCarouselUpdate)
	app.Static("/static", "assets")

	app.Start(":3030")
}
