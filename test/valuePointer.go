package main

/**
值方法和指针方法
**/

type Student struct {
	Age  int
	Name string
}

// 指针方法
func (s *Student) SetPointer(age int, name string) {
	s.Age = age
	s.Name = name
}

// 值方法
func (s Student) SetValue(age int, name string) {
	s.Age = age
	s.Name = name
}

func main() {
	// var s1 Student
	// s1.SetValue(10, "aa")
	// fmt.Println(s1)
	// s1.SetPointer(20, "bb")
	// fmt.Println(s1)
	// s2 := &Student{0, "00"}
	// s2.SetValue(10, "aa")
	// fmt.Println(s2)

}
