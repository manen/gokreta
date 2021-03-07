package kreta

import "time"

type GradeType struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type GradeMode struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type SubjectCategory struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type Subject struct {
	ID       string           `json:"Uid"`
	Category *SubjectCategory `json:"Kategoria"`
	Title    string           `json:"Nev"`
}

type ClassGroup struct {
	ID string `json:"Uid"`
}

type Grade struct {
	ID               string      `json:"Uid"`
	TeacherName      string      `json:"ErtekeloTanarNeve"`
	GradeType        *GradeType  `json:"ErtekFajta"`
	Kind             string      `json:"Jelleg"`
	DateWritten      *time.Time  `json:"KeszitesDatuma"`
	DateSeen         *time.Time  `json:"LattamozasDatuma"`
	Mode             *GradeMode  `json:"Mod"`
	DateAdded        *time.Time  `json:"RogzitesDatuma"`
	Weight           int         `json:"SulySzazalekErteke"`
	Value            int         `json:"SzamErtek"`
	ValueString      string      `json:"SzovegesErtek"`
	ValueStringShort string      `json:"SzovegesErtekelesRovidNev"`
	Subject          *Subject    `json:"Tantargy"`
	Topic            string      `json:"Tema"`
	Type             *GradeType  `json:"Tipus"`
	ClassGroup       *ClassGroup `json:"OsztalyCsoport"`
}

type Absences struct{}

type Timetable struct{}

type Notes struct{}

type Notices struct{}

type Tests struct{}

type Groups struct{}

type UserData struct{}

type YearOrder struct{}

type Homework struct{}

type HomeworkComments struct{}
