package main

import (
	"flag"
	"fmt"
	//"log"
)

var upper = flag.Int("upper", 1000, "upper bound for factor sums")

func main() {
	flag.Parse()

	tot := 2
	fmt.Printf("fib %v: %v\n", 2, 2)
	fib, _ := NewFibonacci()
	fib.Next()
	fib.Next()
	for fib.Next()+fib.Next() <= *upper {
		tot += fib.Next()
		fmt.Printf("fib %v: %v\n", fib.Count(), fib.Val())
	}
}

type Fibonacci struct {
	a, b int
	i    int
}

func NewFibonacci() (Fibonacci, int) {
	return Fibonacci{1, 1, 0}, 1
}

func (f *Fibonacci) Count() int {
	return f.i
}

func (f Fibonacci) Prev() int {
	return f.a
}

func (f Fibonacci) Val() int {
	return f.b
}

func (f *Fibonacci) Next() int {
	if f.i >= 1 {
		f.a, f.b = f.b, f.a+f.b
	}
	f.i++
	return f.b
}
