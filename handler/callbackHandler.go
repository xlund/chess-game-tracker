package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/domain"
)

// Handler for our callback.
func Callback(auth domain.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := session.Get("auth-session", c)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		code := c.QueryParam("code")

		token, profile, err := auth.GetProfile(code, c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		session.Values["access_token"] = token
		session.Values["profile"] = profile
		if err := session.Save(c.Request(), c.Response()); err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusFound, "/user")

	}
}
