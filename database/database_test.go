package database

import (
	"database/sql"
	"reflect"
	"testing"
)

var testDb CgpaRepo  = NewCgpaRepo(":memory:")

func TestGetAllCourses(t *testing.T) {
	s := NewSemester("2021/22")
	s2 := NewSemester("2020/21")

	sessions := []Semester{s, s2}
	c := NewCourse("2021/22","physics", "phy101",4,'A')
	c1 := NewCourse("2021/22","algebra", "mat101",4,'A')
	c2 := NewCourse("2021/22","calculus", "mat102",4,'A')
	c3 := NewCourse("2020/21","intro to programming", "csc101",4,'A')
	c4 := NewCourse("2020/21","physics", "phy103",4,'A')

	courses := []Course{c, c1, c2, c3, c4}

	var err error

	for _, session := range sessions {
		err = testDb.AddSemester(session)
		if err != nil {
			t.Errorf("error: %s",err)
		}
	}

	for _, course := range courses {

		err = testDb.AddCourse(course)
		if err != nil {
			t.Errorf("error: %s",err)
		}
	}

	courses2, err := testDb.GetAllCourses()

	if err != nil {
		t.Errorf("error: %s",err)
	}

	if !reflect.DeepEqual(courses, courses2) {
		t.Errorf("error: expected %v, got %v", courses, courses2)
	}

	for _, course := range courses {
		testDb.DeleteCourse(course)
	}

	for _, session := range sessions {
		testDb.DeleteSemester(session.Session)
	}

}

	



func TestGetCourses(t *testing.T) {
	
	s := NewSemester("2021/22")
	c := NewCourse("2021/22","physics", "phy101",4,'A')
	c1 := NewCourse("2021/22","algebra", "mat101",4,'A')
	c2 := NewCourse("2021/22","calculus", "mat102",4,'A')
	c3 := NewCourse("2021/22","intro to programming", "csc101",4,'A')

	courses := []Course{c, c1, c2, c3}
	e := testDb.AddSemester(s)

	if e != nil {
		t.Errorf("error: %s",e)
	}

	for _, course := range courses {
		e = testDb.AddCourse(course)
		if e != nil {
			t.Errorf("error: %s",e)
		}
	}


	courses2, err := testDb.GetCourses(s)

	if err != nil {
		t.Errorf("error: %s",e)
	}

	if !reflect.DeepEqual(courses, courses2) {
		t.Errorf("error: expected %v, got %v", courses, courses2)
	}
	testDb.DeleteSemester(s.Session)
	for _, course := range courses {
		testDb.DeleteCourse(course)
	}


}

func TestAddAndGetCourse(t *testing.T) {
	s := NewSemester("2021/22")
	c := NewCourse("2021/22","physics", "phy106",4,'A')
	_ = testDb.AddSemester(s)
	e := testDb.AddCourse(c)
	
	if e != nil {
		t.Errorf("error: %s",e)
	}
	c2, err := testDb.GetCourse(c.Code)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if !reflect.DeepEqual(c, c2) {
		t.Errorf("expected :%v, got %v", c, c2)
	}

	s,_ = testDb.GetSemester(s.Session)
	t.Logf("%v",s)

	testDb.DeleteCourse(c2)
	testDb.DeleteSemester(s.Session)

}

func TestDeleteCourse(t *testing.T) {

	c := NewCourse("2021/22","physics", "phy108",4,'A')
	testDb.AddCourse(c)
	testDb.DeleteCourse(c)
	c2, err := testDb.GetCourse(c.Code)
	if err != nil && err != sql.ErrNoRows  {
		t.Errorf("error %s", err)
	}

	if c2.Code != "" {
		t.Errorf("expected no value, got %s", c2.Code)
	}
}

func TestAddAndGetSemester(t *testing.T) {
	s := NewSemester("2021/22")
	testDb.AddSemester(s)
	s2, err := testDb.GetSemester(s.Session)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	if !reflect.DeepEqual(s, s2) {
		t.Errorf("expected: %v, got %v ", s, s2)
	}
	testDb.DeleteSemester(s.Session)

}


func TestDeleteSemester(t *testing.T) {
	s := NewSemester("2019/20")
	testDb.AddSemester(s)
	testDb.DeleteSemester(s.Session)
	s2,err := testDb.GetSemester(s.Session)

	
	if err != nil && err != sql.ErrNoRows  {
		t.Errorf("error %s", err)
	}

	if s2.Session != "" {
		t.Errorf("expected no value, got %s", s2.Session)
	}

}
