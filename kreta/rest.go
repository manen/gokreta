package kreta

import (
	"io"
	"net/http"
	"strings"

	"github.com/manen/gokreta/endpoints"
)

func (s *Session) restGet(endpoint string, headers map[string]string) (string, error) {
	return s.rest("GET", endpoint, "", headers)
}

func (s *Session) restPost(endpoint, body string, headers map[string]string) (string, error) {
	return s.rest("POST", endpoint, body, headers)
}

func (s *Session) rest(httpType, endpoint, body string, headers map[string]string) (string, error) {
	req, err := http.NewRequest(httpType, endpoints.Token, strings.NewReader(body))
	req.Header.Add("User-Agent", s.userAgent)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		return "", nil
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}

	b := &strings.Builder{}
	_, err = io.Copy(b, res.Body)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
