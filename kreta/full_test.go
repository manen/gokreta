package kreta_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/manen/gokreta/kreta"
)

func Test(t *testing.T) {
	s := kreta.NewSession("hu.ekreta.student/1.0.5/Android/0/0", os.Getenv("SCHOOL"), os.Getenv("STUDENT"), os.Getenv("PASSWORD"))

	g, err := s.Grades(time.Date(2020, 9, 01, 01, 01, 01, 0, time.Local), time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(g[0].Subject.Title)

	a, err := s.Absences(time.Date(2020, 9, 01, 01, 01, 01, 0, time.Local), time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(a[0].DateRecorded)
}
