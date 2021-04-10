package kreta_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/manen/gokreta/kreta"
)

var s *kreta.Session

func TestAuth(t *testing.T) {
	s = kreta.NewSession(
		"hu.ekreta.student/1.0.5/Android/0/0",
		os.Getenv("SCHOOL"),
		os.Getenv("STUDENT"),
		os.Getenv("PASSWORD"),
		true,
	)
}

func TestGrades(t *testing.T) {
	g, err := s.Grades(time.Date(2020, 9, 01, 01, 01, 01, 0, time.Local), time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(g[0].Subject.Title)
}

func TestAbsences(t *testing.T) {
	a, err := s.Absences(time.Date(2020, 9, 01, 01, 01, 01, 0, time.Local), time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(a[0].DateRecorded)
}

func TestTimetable(t *testing.T) {
	ti, err := s.Timetable(time.Date(2020, 9, 01, 01, 01, 01, 0, time.Local), time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(ti[0].Subject.Title)
}

func TestUserData(t *testing.T) {
	ud, err := s.UserData()
	if err != nil {
		panic(err)
	}

	log.Println(ud.Guardians[0].TelephoneNumber)
}
