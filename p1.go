package main

import (
	"flag"
	"math"
	"fmt"
	"log"
	"strconv"
)

var upper = flag.Int("upper", 1000, "upper bound for factor sums")

func main() {
	flag.Parse()

	upp := *upper - 1
	if float64(upp) > math.Sqrt(math.MaxInt64) {
		log.Fatal("overflow: upper bound is too large")
	}

	factors := []int{}
	for i, arg := range flag.Args() {
		fac, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("arg %v (value '%v') is invalid", i, arg)
		}
		factors = append(factors, fac)
	}

	tot := 0
	prod := 1
	for _, fac := range factors {
		prod *= fac
		count := upp / fac
		max := upp - (upp % fac)
		tot += max * (count + 1) / 2
	}

	// subtract common multiples of factors
	// prior to this, tot would, for example, contain 3*5 and 5*3
	count := upp / prod
	max := upp - (upp % prod)
	tot -= max * (count + 1) / 2

	fmt.Printf("Sum of all multiples of %#v up to %v: %v\n", factors, *upper, tot)
}
