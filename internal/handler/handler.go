package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/mdparser"
	"github.com/lorenzhohermuth/portfolio/view/component"
	"github.com/lorenzhohermuth/portfolio/view/page"
)

type Homehandler struct {
	Index int
	Entrys *[]component.CarouselEntry
}

func(h Homehandler) HandleUserShow(ctx echo.Context) error {
	events := []component.Event{
		{time.Now(), time.Now(), "Title 1"},
		{time.Now(), time.Now(), "Title 2"},
		{time.Now(), time.Now(), "Title 3"},
		{time.Now(), time.Now(), "Title 4"},
		{time.Now(), time.Now(), "Title 5"},
	}

	*h.Entrys = mdparser.GetProjects()
	
	return render(ctx, page.ShowHome(*h.Entrys, h.Index, events))
}

type HtmxCarouselHandler struct {
	Index *int
	Direction int
	Entrys *[]component.CarouselEntry
}

func(h HtmxCarouselHandler) HandlerCarouselUpdate(ctx echo.Context) error {
	*h.Index += h.Direction
	return render(ctx, component.Carousel(*h.Entrys, int(*h.Index)))
}
