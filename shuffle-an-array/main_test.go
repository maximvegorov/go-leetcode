package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const sampleSize = 1_000_000
	s := Constructor([]int{1, 2, 3})
	freqs := map[[3]int]int{}
	for i := 0; i < sampleSize; i++ {
		var perm [3]int
		b := s.Shuffle()
		copy(perm[:], b)
		freqs[perm]++
	}
	for k, v := range freqs {
		fmt.Printf("%q = %v\n", k, float64(v)/float64(sampleSize))
	}
}
