package domain

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type Authenticator interface {
	VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
	VerifyState(state string, session *sessions.Session) bool
	GetLoginURL(c echo.Context, session *sessions.Session) (string, error)
	GetClaims(code string, ctx echo.Context) (*oauth2.Token, *UserClaims, error)
}
