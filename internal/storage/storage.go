package storage

import "github.com/shivam-02/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentById(id int64, name string, email string, age int) (int64, error)
	DeleteById(id int64) (int64, error)
}
