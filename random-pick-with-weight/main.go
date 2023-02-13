package randompickwithweight

import (
	"math/rand"
	"sort"
)

type Solution struct {
	prefixSum []int
	n         int
	r         *rand.Rand
}

func Constructor(weights []int) Solution {
	prefixSum := make([]int, len(weights)+1)
	for i, weight := range weights {
		prefixSum[i+1] = prefixSum[i] + weight
	}
	return Solution{prefixSum: prefixSum, n: prefixSum[len(prefixSum)-1], r: rand.New(rand.NewSource(0))}
}

func (this *Solution) PickIndex() int {
	point := this.r.Intn(this.n)
	return sort.Search(len(this.prefixSum), func(i int) bool {
		return this.prefixSum[i] > point
	}) - 1
}
