package main

import (
	"database/sql"
	"reflect"
	"testing"
)

var testDb CgpaRepo  = NewCgpaRepo(":memory:")

func TestAddAndGetCourse(t *testing.T) {
	s := NewSemester("2021/22")
	c := NewCourse("2021/22","physics", "phy101",4,'A')
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

	c := NewCourse("2021/22","physics", "phy101",4,'A')
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
