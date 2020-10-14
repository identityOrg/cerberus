package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

func HandleCallback(c echo.Context) error {
	hpd := NewHomePageData("Authorization Code Flow Response")
	hpd.AccessCode = c.QueryParam("code")
	hpd.AccessToken = c.QueryParam("token")
	hpd.IDToken = c.QueryParam("id_token")
	if hpd.AccessCode != "" {
		token, err := getAuthCodeConfig().Exchange(context.Background(), hpd.AccessCode,
			oauth2.SetAuthURLParam("state", hpd.State),
			oauth2.SetAuthURLParam("scope", strings.Join(getAuthCodeConfig().Scopes, " ")),
		)
		if err != nil {
			hpd.Type = err.Error()
			return c.Render(http.StatusOK, "demo_home.html", hpd)
		} else {
			hpd.AccessToken = token.AccessToken
			hpd.RefreshToken = token.RefreshToken
			idToken := token.Extra("id_token")
			if idToken != nil {
				hpd.IDToken = idToken.(string)
			}
		}
	}
	return c.Render(http.StatusOK, "demo_home.html", hpd)
}
