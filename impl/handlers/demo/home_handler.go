package demo

import (
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func HandleHome(c echo.Context) error {
	hpd := NewHomePageData("Demo Home")
	return c.Render(http.StatusOK, "demo_home.html", hpd)
}

type HomePageData struct {
	Type            string
	AuthCodeFlowURL string
	ImplicitFlowURL string
	HomeURL         string
	HybridFlowURL   string
	AccessCode      string
	AccessToken     string
	RefreshToken    string
	IDToken         string
	State           string
	CodeVerifier    string
}

func NewHomePageData(pageType string) *HomePageData {
	h := &HomePageData{Type: pageType}
	h.State = "RandomIdString0wefwefwefwef"
	h.AuthCodeFlowURL = getAuthCodeConfig().AuthCodeURL(h.State)
	impUrl, _ := url.Parse(h.AuthCodeFlowURL)
	query := impUrl.Query()
	query.Set("response_type", "token id_token")
	query.Set("nonce", "a-nonce-value")
	impUrl.RawQuery = query.Encode()
	h.ImplicitFlowURL = impUrl.String()
	query.Set("response_type", "code token")
	impUrl.RawQuery = query.Encode()
	h.HybridFlowURL = impUrl.String()
	h.HomeURL = config.NewSDKConfig().Issuer
	return h
}

const ValidIdChars = "1234567890abcdefghijklmnopqrstuvwxyz"

func RandomIdString(length int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = ValidIdChars[seededRand.Intn(len(ValidIdChars))]
	}
	return string(b)
}
