package main

import (
	"testing"
)

const LargeSize = 10000000

type LargeStruct struct {
	Data [LargeSize]int
}

// 值方法
func (ls LargeStruct) ValueMethod() int {
	sum := 0
	for _, val := range ls.Data {
		sum += val
	}
	return sum
	// return 1
}

// 指针方法
func (ls *LargeStruct) PointerMethod() int {
	sum := 0
	for _, val := range ls.Data {
		sum += val
	}
	return sum
	// return 1
}

func BenchmarkLargeValueMethod(b *testing.B) {
	ls := LargeStruct{
		Data: [LargeSize]int{},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ls.ValueMethod()
	}
}

func BenchmarkLargePointerMethod(b *testing.B) {
	ls := &LargeStruct{
		Data: [LargeSize]int{},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ls.PointerMethod()
	}
}
