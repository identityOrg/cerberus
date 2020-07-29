package demo

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func HandleHome(c echo.Context) error {
	hpd := NewHomePageData("Demo Home")
	hpd.SetStateCookie(c)
	return c.Render(http.StatusOK, "demo_home.html", hpd)
}

type HomePageData struct {
	Type            string
	AuthCodeFlowURL string
	ImplicitFlowURL string
	AccessCode      string
	AccessToken     string
	RefreshToken    string
	State           string
	OldState        string
	CodeVerifier    string
}

func (hpd HomePageData) SetStateCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "state"
	cookie.Value = hpd.State
	c.SetCookie(cookie)
}

func (hpd *HomePageData) ReadStateCookie(c echo.Context) {
	h, err := c.Cookie("state")
	if err == nil {
		hpd.OldState = h.Value
	}
}

func NewHomePageData(pageType string) *HomePageData {
	h := &HomePageData{Type: pageType}
	h.State = RandomIdString(20)
	h.AuthCodeFlowURL = conf.AuthCodeURL(h.State)
	impUrl, _ := url.Parse(h.AuthCodeFlowURL)
	query := impUrl.Query()
	query.Set("response_type", "token")
	impUrl.RawQuery = query.Encode()
	h.ImplicitFlowURL = impUrl.String()
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
