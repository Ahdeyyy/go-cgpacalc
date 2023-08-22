package main 

type Semester struct {
  Session string // i.e 2021/22
  Gpa float32
}

type Course struct {
	Session string // i.e 2021/22
	Name string
	Code string 
	Unit int
	Grade byte
}

func NewCourse(Session string, Name string, Code string, Unit int, Grade byte ) Course {
	return Course{
		Session: Session,
		Name: Name,
		Code: Code,
		Unit: Unit,
		Grade: Grade,
	}
}


func NewSemester (session string) Semester {
  return Semester{
    Session: session,
    Gpa: 0.0, 
  }
}

