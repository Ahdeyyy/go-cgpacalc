package main

import (
	"log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type CgpaRepo struct {
	Db *sql.DB
}

func NewCgpaRepo(path string) CgpaRepo {
	return CgpaRepo{
		Db: OpenDb(path),
	}
}

func OpenDb(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Panic("could not open database")
	}

	_, err = db.Exec(`
    PRAGMA JOURNAL_MODE = WAL;
    PRAGMA FOREIGN_KEYS = ON;
    PRAGMA BUSY_TIMEOUT = 500;
    `)

	if err != nil {
		log.Printf("err: %s ", err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS semesters (
      session TEXT UNIQUE, 
      gpa FLOAT
    );
    `)

	if err != nil {
		log.Panicln(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS courses (
			session TEXT,
			name TEXT,
			code TEXT UNIQUE,
			unit INT,
			grade CHARACTER(1)
		);
	`)

	return db
}

func (c CgpaRepo) AddSemester(semester Semester) error {
	stmt := `
	INSERT INTO semesters ( session, gpa ) VALUES ( ?, ?);
        `
	_, err := c.Db.Exec(stmt, semester.Session, semester.Gpa)

	if err != nil {
		log.Printf("error: %s", err)
		return err
	}
	return nil

}

func (c CgpaRepo) GetSemester(session string) (Semester, error) {
	stmt := `
	SELECT * FROM semesters WHERE session = ?
	`
	var s Semester
	result := c.Db.QueryRow(stmt, session)

	err := result.Scan(&s.Session, &s.Gpa)
	if err != nil {
		return s, err
	}
	return s, nil

}

func (c CgpaRepo) DeleteSemester(session string) error {
	stmt := `DELETE FROM semesters WHERE session = ?`
	_, err := c.Db.Exec(stmt, session)
	if err != nil {
		return err
	}
	return nil
}

func (c CgpaRepo) AddCourse(course Course) error {
	stmt := `
		INSERT INTO courses ( session, name, code, unit, grade) VALUES (?, ?, ?, ?, ?)
	`
	_, err := c.Db.Exec(stmt, course.Session, course.Name, course.Code, course.Unit, course.Grade)

	if err != nil {
		return err
	}

	return nil
}

func (c CgpaRepo) GetCourse(code string) (Course, error) {
	stmt := `
		SELECT * FROM courses WHERE code = ? 
		`
	var course Course

	result := c.Db.QueryRow(stmt, code)

	err := result.Scan(&course.Session, &course.Name, &course.Code, &course.Unit, &course.Grade)

	if err != nil {
		return course, err
	}
	return course, nil
}
