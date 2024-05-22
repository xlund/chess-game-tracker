package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/xlund/chess-games-tracker/domain"
	"golang.org/x/oauth2"
)

type Auth0Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func New() (domain.Authenticator, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://"+os.Getenv("AUTH0_DOMAIN")+"/")
	if err != nil {
		return nil, err
	}
	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Auth0Authenticator{provider, conf}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken
func (a *Auth0Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}
	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Auth0Authenticator) GetLoginURL(c echo.Context, session *sessions.Session) (string, error) {
	state, err := generateRandomState()
	if err != nil {
		return "", err
	}
	// Save the state inside the session.
	session.Values["state"] = state
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return "", err
	}
	return a.AuthCodeURL(state), nil
}

func (a *Auth0Authenticator) VerifyState(state string, session *sessions.Session) bool {
	storedState, ok := session.Values["state"].(string)
	if !ok {
		return false
	}
	return state == storedState
}

func (a *Auth0Authenticator) GetClaims(code string, ctx echo.Context) (*oauth2.Token, *domain.UserClaims, error) {
	token, err := a.Exchange(ctx.Request().Context(), code)
	if err != nil {
		return nil, nil, err
	}

	idToken, err := a.VerifyIDToken(ctx.Request().Context(), token)
	if err != nil {
		return nil, nil, err
	}

	var profile domain.UserClaims
	if err := idToken.Claims(&profile); err != nil {
		return nil, nil, err
	}

	return token, &profile, nil
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
