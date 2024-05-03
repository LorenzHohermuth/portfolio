package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/internal/mdparser"
	"github.com/lorenzhohermuth/portfolio/view/component"
	"github.com/lorenzhohermuth/portfolio/view/page"
)

type Homehandler struct {
	Index int
	Entrys *[]component.CarouselEntry
}

func(h Homehandler) HandleShowHome(ctx echo.Context) error {
	events := mdparser.GetWork()
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
