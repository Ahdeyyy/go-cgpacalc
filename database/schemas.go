package database

import (
	"strconv"
)

type Semester struct {
	Session string // i.e 2021/22 First
	Gpa     float32
}

func (s Semester) Title() string {
	return s.Session
}

func (s Semester) Description() string {
	return "GPA: " + strconv.FormatFloat(float64(s.Gpa), 'f', 2, 32)
}

func (s Semester) FilterValue() string {
	return s.Session
}

type Course struct {
	Session string // i.e 2021/22 First
	Name    string
	Code    string
	Unit    int
	Grade   byte
}

func (c Course) Title() string {
	return c.Name
}

func (c Course) Description() string {
	return "Code: " + c.Code + " Unit: " + strconv.Itoa(c.Unit) + " Grade: " + string(c.Grade)
}

func (c Course) FilterValue() string {
	return c.Code
}


func NewCourse(Session string, Name string, Code string, Unit int, Grade byte) Course {
	return Course{
		Session: Session,
		Name:    Name,
		Code:    Code,
		Unit:    Unit,
		Grade:   Grade,
	}
}

func NewSemester(session string) Semester {
	return Semester{
		Session: session,
		Gpa:     0.0,
	}
}
