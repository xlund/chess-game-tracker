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
	GetLoginURL(c echo.Context, session *sessions.Session) (string, error)
	VerifyState(state string, session *sessions.Session) bool
	GetProfile(code string, session context.Context) (*oauth2.Token, *UserClaims, error)
}
