package main

import (
	"testing"
)

const up = 4000000

func BenchmarkAlgo1(b *testing.B) {
	tot := 0
	for i := 0; i < b.N; i++ {
		tot = 0
		for fib, f := NewFibonacci(); f <= up; f = fib.Next() {
			if (f & 1) == 0 {
				tot += fib.Val()
			}
		}
	}
	b.Log("total: ", tot)
}

func BenchmarkAlgo2(b *testing.B) {
	tot := 0
	for i := 0; i < b.N; i++ {
		tot = 0
		fib, _ := NewFibonacci()
		fib.Next()
		for fib.Next() <= up {
			tot += fib.Val()
			fib.Next()
			fib.Next()
		}
	}
	b.Log("total: ", tot)
}

func BenchmarkAlgo3(b *testing.B) {
	tot := 0
	for i := 0; i < b.N; i++ {
		x1, x2 := 2, 8
		tot = x1
		for x2 <= up {
			tot += x2
			x1, x2 = x2, x2*4+x1
		}
	}
	b.Log("total: ", tot)
}
