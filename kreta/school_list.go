package kreta

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/manen/gokreta/endpoints"
)

func (s *Session) SchoolsOfficial(apiKey string) ([]ListSchool, error) {
	vals := url.Values{}
	vals.Add("apiKey", apiKey)

	res, err := s.restGet(endpoints.OfficialSchools, vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var schools []ListSchool
	err = json.NewDecoder(strings.NewReader(res)).Decode(&schools)
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *Session) Schools() ([]ListSchool, error) {
	res, err := http.Get(endpoints.Schools)

	var schools []ListSchool
	err = json.NewDecoder(res.Body).Decode(&schools)
	if err != nil {
		return nil, err
	}

	return schools, nil
}
