package main

import (
	"testing"
)

func init() {
	*upper = 40000000
}

func BenchmarkAlgo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tot := 0
		for fib, f := NewFibonacci(); f <= *upper; f = fib.Next() {
			if (f & 1) == 0 {
				tot += fib.Next()
			}
		}
		b.Log("total: ", tot)
	}
}

func BenchmarkAlgo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tot := 2
		fib, _ := NewFibonacci()
		fib.Next()
		fib.Next()
		for fib.Next()+fib.Next() <= *upper {
			tot += fib.Next()
		}
		b.Log("total: ", tot)
	}
}

func BenchmarkAlgo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x1, x2 := 2, 8
		tot := x1
		for x2 <= *upper {
			tot += x2
			x1, x2 = x2, x2*4+x1
		}
		b.Log("total: ", tot)
	}
}
