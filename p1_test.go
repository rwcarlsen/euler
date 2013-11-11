package main

import (
	"testing"
)

var facs = []int{3, 5}

const up = 1000000

func BenchmarkAlgo1(b *testing.B) {
	tot := 0
	for i := 0; i < b.N; i++ {
		tot = 0
		for val := 0; val < up; val++ {
			for _, fac := range facs {
				if val % fac == 0 {
					tot += val
					break
				}
			}
		}
	}
	b.Logf("Sum of all multiples of %#v up to %v: %v\n", facs, up, tot)
}

func BenchmarkAlgo2(b *testing.B) {
	tot := 0
	for i := 0; i < b.N; i++ {
		tot = 0
		prod := 1
		for _, fac := range facs {
			prod *= fac
			for val := 0; val < up; val += fac {
				tot += val
			}
		}
		for val := 0; val < up; val += prod {
			tot -= val
		}
	}
	b.Logf("Sum of all multiples of %#v up to %v: %v\n", facs, up, tot)
}

func BenchmarkAlgo3(b *testing.B) {
	tot := 0
	upp := up - 1
	for i := 0; i < b.N; i++ {
		tot = 0
		prod := 1
		for _, fac := range facs {
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
	}
	b.Logf("Sum of all multiples of %#v up to %v: %v\n", facs, up, tot)
}

