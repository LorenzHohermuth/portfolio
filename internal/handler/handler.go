package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/view/component"
	"github.com/lorenzhohermuth/portfolio/view/page"
)

type Homehandler struct {}

func(h Homehandler) HandleUserShow(ctx echo.Context) error {
	entrys := []component.CarouselEntry{
		{"/static/test.jpg", "This is a Tree" , "I like Trees"},
		{"/static/flower.jpeg", "This is a Tree" , "I like Trees"},
		{"/static/paris.jpg", "This is a Tree" , "I like Trees"},
	}
	return render(ctx, page.ShowHome(entrys))
}
