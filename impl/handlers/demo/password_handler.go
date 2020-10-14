package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlePasswordPage(c echo.Context) error {
	hpd := NewHomePageData("Resource Owner Password Demo")
	return c.Render(http.StatusOK, "demo_password.html", hpd)
}

func HandlePasswordPost(c echo.Context) error {
	hpd := NewHomePageData("Resource Owner Password Demo")

	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := getAuthCodeConfig().PasswordCredentialsToken(context.Background(), username, password)
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
