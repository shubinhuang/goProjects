package main

import "testing"

func benchFib(b *testing.B, num int) {
	for n := 0; n < b.N; n++ {
		Fib(num)
	}
}

func BenchmarkFib_10(b *testing.B) {
	benchFib(b, 10)
}

func BenchmarkFib_30(b *testing.B) {
	benchFib(b, 30)
}

func BenchmarkFib_40(b *testing.B) {
	benchFib(b, 40)
}
