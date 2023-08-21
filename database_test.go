package main

import (
	"database/sql"
	"reflect"
	"testing"
)

var testDb CgpaRepo  = NewCgpaRepo("test.db")


func TestAddSemester(t *testing.T) {
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

func TestGetSemester(t *testing.T) {
	s := NewSemester("2022/23")
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
