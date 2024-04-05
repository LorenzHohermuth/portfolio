package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/view/component"
	"github.com/lorenzhohermuth/portfolio/view/page"
)

type Homehandler struct {
	Index int
}

func(h Homehandler) HandleUserShow(ctx echo.Context) error {
	entrys := []component.CarouselEntry{
		{"/static/test.jpg", "This is a Tree" , "I like Trees"},
		{"/static/flower.jpeg", "This is a Flower" , "I like Flowers"},
		{"/static/paris.jpg", "This is the Eiffel Tower" , "I not like Franc"},
	}
	return render(ctx, page.ShowHome(entrys, h.Index))
}

type HtmxCarouselHandler struct {
	Index *int
	Direction int
}

func(h HtmxCarouselHandler) HandlerCarouselUpdate(ctx echo.Context) error {
	entrys := []component.CarouselEntry{
		{"/static/test.jpg", "This is a Tree" , "I like Trees"},
		{"/static/flower.jpeg", "This is a Flower" , "I like Flowers"},
		{"/static/paris.jpg", "This is the Eiffel Tower" , "I not like Franc"},
	}
	*h.Index += h.Direction
	return render(ctx, component.Carousel(entrys, int(*h.Index)))
}
