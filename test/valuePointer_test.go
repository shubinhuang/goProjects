package main

import (
	"fmt"
	"testing"
)

type MyStruct struct {
	Value int
}

// 值方法
func (m MyStruct) ValueMethod() {
	m.Value = 5
}

// 指针方法
func (m *MyStruct) PointerMethod() {
	m.Value = 5
}

func TestValueMethod(t *testing.T) {
	myStruct := MyStruct{}
	myStruct.ValueMethod()
	fmt.Println(myStruct.Value)

	if myStruct.Value != 5 {
		t.Errorf("Expected result to be 5, got %d", myStruct.Value)
	}
}

func TestPointerMethod(t *testing.T) {
	myStruct := &MyStruct{}
	myStruct.ValueMethod()
	fmt.Println(myStruct)

	if myStruct.Value != 5 {
		t.Errorf("Expected result to be 5, got %d", myStruct.Value)
	}
}

func BenchmarkValueMethod(b *testing.B) {
	myStruct := MyStruct{Value: 5}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		myStruct.ValueMethod()
	}
}

func BenchmarkPointerMethod(b *testing.B) {
	myStruct := &MyStruct{Value: 5}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		myStruct.PointerMethod()
	}
}

// test Student

func TestStudentValueMethod(t *testing.T) {
	sv1 := Student{}
	sv1.SetValue(20, "aa")
	if sv1.Age != 20 {
		t.Errorf("在方法中被操作的是原对象的副本，不会影响原对象，结果是%v", sv1)
	}
}

func TestStudentPointerMethod(t *testing.T) {
	var sv1 Student
	// var sp1 *Student
	// sv1.SetPointer(20, "aa")
	sv1.SetPointer(21, "bb")
	// fmt.Println(sv1)
	fmt.Println(sv1)
}

func TestStudentMethod(t *testing.T) {
	var sv1 = Student{}
	var sp1 = &Student{}
	var sv2 = Student{}
	var sp2 = &Student{}

	fmt.Printf("初始为: %v, %v\n", sv1, sp1)

	sv1.SetValue(10, "aa")
	sp1.SetPointer(11, "bb")
	sv2.SetPointer(12, "cc")
	sp2.SetValue(13, "cc")

	fmt.Printf("值调用值方法, result = %v\n", sv1)
	fmt.Printf("指针调用指针方法, result = %v\n", sp1)
	fmt.Printf("值调用指针方法, result = %v\n", sv2)
	fmt.Printf("指针调用值方法, result = %v\n", sp2)
}
