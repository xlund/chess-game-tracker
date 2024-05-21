package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/domain"
)

// Handler for our login.
func Login(a domain.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Save the state inside the session.
		session, err := session.Get("auth-session", c)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		url, err := a.GetLoginURL(c, session)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
