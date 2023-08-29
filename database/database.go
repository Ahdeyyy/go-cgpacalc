package database

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

	if err != nil {
		log.Panicln(err)
	}


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

	rows, er := c.Db.Query(`
		SELECT unit, grade
		FROM courses 
		WHERE session = ? ;`,course.Session )

	if er != nil {
		return er
	}

	defer rows.Close()

	totalPoints := 0
	totalUnits := 0

	for rows.Next() {
		var unit int
		var grade byte
		er = rows.Scan(&unit,&grade)
		if er != nil {
			return er
		}
		totalPoints += unit * GradeToPoint(grade)
		totalUnits += unit
	}

	gpa := float32(totalPoints/totalUnits)

	_, err = c.Db.Exec(`UPDATE semesters SET gpa = ? WHERE session = ?`, gpa, course.Session)

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
func (c CgpaRepo) DeleteCourse(course Course) error {
	stmt := `DELETE FROM courses WHERE code = ?`
	_, err := c.Db.Exec(stmt, course.Code)
	if err != nil {
		return err
	
	rows, er := c.Db.Query(`
		SELECT unit, grade
		FROM courses 
		WHERE session = ? ;`,course.Session )

	if er != nil {
		return er
	}

	defer rows.Close()

	totalPoints := 0
	totalUnits := 0

	for rows.Next() {
		var unit int
		var grade byte
		er = rows.Scan(&unit,&grade)
		if er != nil {
			return er
		}
		totalPoints += unit * GradeToPoint(grade)
		totalUnits += unit
	}

	gpa := float32(totalPoints/totalUnits)

	_, err = c.Db.Exec(`UPDATE semesters SET gpa = ? WHERE session = ?`, gpa, course.Session)

	if err != nil {
		return err
	}

}
	return nil
}

func (c CgpaRepo) GetSemesters() ([]Semester, error) {
	var semesters []Semester 
	stmt := "SELECT * FROM semesters"
	res, err := c.Db.Query(stmt)
	if err != nil {
		return semesters, err
	}
	defer res.Close()
	for res.Next() {
		sem := Semester{}
		err = res.Scan(&sem.Session,&sem.Gpa)
		if err != nil {
			continue
		}
		semesters =	append(semesters,sem)
	}

	return semesters, nil 
}




func (c CgpaRepo)GetCourses(semester Semester) ([]Course,error) {
	var courses []Course
	stmt :=	`SELECT * FROM courses WHERE session = ? `
	res, err := c.Db.Query(stmt,semester.Session)
	if err != nil {
		return courses, err
	}
	defer res.Close()

	for res.Next() {
		cour := Course{}
		err = res.Scan(&cour.Session, &cour.Name,&cour.Code, &cour.Unit,&cour.Grade)
		if err != nil {
			continue
		}
		courses = append(courses, cour)
	}

	return courses,nil

}

func (c CgpaRepo)GetCgpa() (float32,error) {
	stmt := `SELECT SUM(gpa) / COUNT(*) AS CGPA FROM semesters`
	var cgpa float32
	err := c.Db.QueryRow(stmt).Scan(&cgpa)
	if err != nil {
		return 0.0, err
	}
	return cgpa, nil
}


func GradeToPoint(grade byte) int {
	switch grade {
		case 'A', 'a':
			return 5
		case 'B', 'b':
			return 4
		case 'C','c' :
			return 3
		case 'D', 'd' :
			return 2
		case 'E', 'e': 
			return 1
		default: 
			return 0
	}
}

