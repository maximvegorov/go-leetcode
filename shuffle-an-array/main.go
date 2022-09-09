package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	r    *rand.Rand
	orig []int
}

func Constructor(nums []int) Solution {
	return Solution{r: rand.New(rand.NewSource(time.Now().UnixNano())), orig: nums}
}

func (s *Solution) Reset() []int {
	res := make([]int, len(s.orig))
	copy(res, s.orig)
	return res
}

func (s *Solution) Shuffle() []int {
	res := make([]int, len(s.orig))
	copy(res, s.orig)
	for i := 0; i < len(res)-1; i++ {
		k := s.r.Intn(len(res) - i)
		j := i + k
		if i != j {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}
