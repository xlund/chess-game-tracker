package api

import (
	"embed"
	"encoding/gob"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/domain"
	"github.com/xlund/chess-games-tracker/handler"
)

func New(auth *domain.Authenticator, fs embed.FS) *echo.Echo {
	e := echo.New()
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.StaticFS("/public", fs)

	e.GET("/", handler.Index())
	e.GET("/login", handler.Login(*auth))
	e.GET("/callback", handler.Callback(*auth))
	return e
}
