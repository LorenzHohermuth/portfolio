package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lorenzhohermuth/portfolio/view/page"
)

type Homehandler struct {}

func(h Homehandler) HandleUserShow(ctx echo.Context) error {
	return render(ctx, page.ShowHome())
}
