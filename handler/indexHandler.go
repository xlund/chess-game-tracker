package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/web/page"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := page.Home()
		return t.Render(c.Request().Context(), c.Response().Writer)
	}
}
