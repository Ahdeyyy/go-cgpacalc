package main

import (

	_ "github.com/mattn/go-sqlite3"
)

type Semester struct {
  Session string // i.e 2021/22
  Gpa float32
}

func NewSemester (session string) Semester {
  return Semester{
    Session: session,
    Gpa: 0.0, 
  }
}

