package interfaces_imp

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks int
}

type StudentIterface interface {
	IsPass() bool
	GetAge() int
	PrintDetails()
}

func (s Student) IsPass() bool {
	return s.Marks >= 33
}

func (s Student) GetAge() int {
	return s.Age
}

func (s Student) PrintDetails() {
	fmt.Printf("Name : %v\nAge: %v\nMarks: %v\n", s.Name, s.Age, s.Marks)
}
