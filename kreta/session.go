package kreta

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/manen/gokreta/endpoints"
)

// Session represents a Kréta session
type Session struct {
	userAgent string

	schoolID, studentID string

	token        string `json:"access_token"`
	refreshToken string `json:"refresh_token"`
}

// NewSession initializes a new Kréta session
func NewSession(userAgent, schoolID, studentID, password string) *Session {
	body := url.Values{}

	body.Set("userName", studentID)
	body.Set("password", password)
	body.Set("institute_code", schoolID)
	body.Set("grant_type", "password")
	body.Set("client_id", "kreta-ellenorzo-mobile")

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoints.Token, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", userAgent)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	s := &Session{
		userAgent:    userAgent,
		schoolID:     schoolID,
		studentID:    studentID,
		token:        "",
		refreshToken: "",
	}

	fuckGo := &struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(buf.String()), &fuckGo)

	s.token = fuckGo.AccessToken
	s.refreshToken = fuckGo.RefreshToken

	return s
}
