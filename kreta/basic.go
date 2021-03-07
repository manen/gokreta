package kreta

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/manen/gokreta/endpoints"
)

func (s *Session) Grades(from, to time.Time) ([]*Grade, error) {
	vals := &url.Values{}

	vals.Add("datumTol", from.Format("2020-09-01T00-00-00"))
	vals.Add("datumIg", to.Format("2020-09-01T00-00-00"))

	res, err := s.restGet(endpoints.Grades, vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var grades []*Grade
	err = json.NewDecoder(strings.NewReader(res)).Decode(&grades)
	if err != nil {
		return nil, err
	}

	return grades, nil
}

func (s *Session) Absences(from, to time.Time) ([]*Absence, error) {
	vals := &url.Values{}

	vals.Add("datumTol", from.Format("2020-09-01T00-00-00"))
	vals.Add("datumIg", to.Format("2020-09-01T00-00-00"))

	res, err := s.restGet(endpoints.Absences, vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var absences []*Absence
	err = json.NewDecoder(strings.NewReader(res)).Decode(&absences)
	if err != nil {
		return nil, err
	}

	return absences, nil
}

func (s *Session) Timetable() (*Timetable, error) {
	return nil, nil
}

func (s *Session) Notes() (*Notes, error) {
	return nil, nil
}

func (s *Session) Notices() (*Notices, error) {
	return nil, nil
}

func (s *Session) Tests() (*Tests, error) {
	return nil, nil
}

func (s *Session) Groups() (*Groups, error) {
	return nil, nil
}

func (s *Session) UserData() (*UserData, error) {
	return nil, nil
}

func (s *Session) YearOrder() (*YearOrder, error) {
	return nil, nil
}

func (s *Session) Homework() (*Homework, error) {
	return nil, nil
}

func (s *Session) HomeworkComments(id string) (*HomeworkComments, error) {
	return nil, nil
}

func (s *Session) HomeworkComment(id, comment string) error {
	return nil
}
