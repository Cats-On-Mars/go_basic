package main

import "testing"

func BenchmarkFib1(b *testing.B){
	for i := 0; i < b.N; i++ {
		fib1(10)
	}
}

func BenchmarkFib2(b *testing.B){
	for i := 0; i < b.N; i++ {
		fib2(10)
	}
}