package kreta

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/manen/gokreta/endpoints"
)

// Session represents a Kréta session
type Session struct {
	userAgent string
	client    *http.Client

	token        string `json:"accessToken"`
	refreshToken string
}

type customTransport struct {
	UserAgent string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", t.UserAgent)
	return http.DefaultTransport.RoundTrip(req)
}

// NewSession initializes a new Kréta session
func NewSession(userAgent, schoolID, studentID, password string) *Session {
	body := url.Values{}

	body.Set("userName", studentID)
	body.Set("password", password)
	body.Set("institute_code", schoolID)
	body.Set("grant_type", "password")
	body.Set("client_id", "kreta-ellenorzo-mobile")

	client := &http.Client{Transport: &customTransport{UserAgent: userAgent}}

	req, err := http.NewRequest("POST", endpoints.Token, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	s := &Session{
		userAgent:    userAgent,
		client:       client,
		token:        "",
		refreshToken: "",
	}

	err = json.NewDecoder(res.Body).Decode(s)
	if err != nil {
		panic(err)
	}

	return s
}
