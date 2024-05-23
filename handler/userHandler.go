package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/domain"
	"github.com/xlund/chess-games-tracker/web/page"
)

func User() echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := session.Get("auth-session", c)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		profile, ok := session.Values["profile"].(domain.UserClaims)
		if !ok {
			return c.Redirect(http.StatusFound, "/")
		}
		t := page.User(profile)
		return t.Render(c.Request().Context(), c.Response().Writer)
	}
}
