package main

/**
值方法和指针方法
**/
import (
	"fmt"
)

type Student struct {
	Age  int
	Name string
}

func (s *Student) SetPointer(age int, name string) {
	s.Age = age
	s.Name = name
}

func (s Student) SetValue(age int, name string) {
	s.Age = age
	s.Name = name
}

func main() {
	var s1 Student
	s1.SetValue(10, "aa")
	fmt.Println(s1)
	s1.SetPointer(20, "bb")
	fmt.Println(s1)
}
