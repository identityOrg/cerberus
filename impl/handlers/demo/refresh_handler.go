package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

func HandleRefresh(c echo.Context) error {
	hpd := NewHomePageData("Refresh Token Example")

	refreshToken := c.FormValue("refresh_token")
	accessToken := c.FormValue("access_token")
	oldToken := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		RefreshToken: refreshToken,
		Expiry:       time.Now().Add(-1 * time.Hour),
	}

	source := getAuthCodeConfig().TokenSource(context.Background(), oldToken)
	token, err := source.Token()
	if err != nil {
		hpd.Type = err.Error()
	} else {
		hpd.AccessToken = token.AccessToken
		hpd.RefreshToken = token.RefreshToken
		idToken := token.Extra("id_token")
		if idToken != nil {
			hpd.IDToken = idToken.(string)
		}
	}

	return c.Render(http.StatusOK, "demo_home.html", hpd)
}
