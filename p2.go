package main

import (
	"flag"
	"fmt"
	//"log"
)

var upper = flag.Int("upper", 4000000, "upper bound for factor sums")

func main() {
	flag.Parse()

	x1, x2 := 2, 8
	tot := x1
	for x2 <= *upper {
		tot += x2
		x1, x2 = x2, x2*4+x1
	}
	fmt.Printf("sum of even fibonacci up to %v: %v\n", *upper, tot)
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
