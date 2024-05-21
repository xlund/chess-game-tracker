package handler

import (
	"github.com/labstack/echo/v4"
	web "github.com/xlund/chess-games-tracker/web/template"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := web.Home()
		return t.Render(c.Request().Context(), c.Response().Writer)
	}
}
