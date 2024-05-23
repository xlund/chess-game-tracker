package handler

import (
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		scheme := "http"
		if c.Request().TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + c.Request().Host)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
		logoutUrl.RawQuery = parameters.Encode()

		return c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
	}
}
