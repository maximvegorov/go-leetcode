package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	res := 0
	for i, vi := range nums {
		if vi <= 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			vj := nums[j]
			if vj <= 0 {
				continue
			}

			remain := nums[j+1:]

			max := vi + vj
			k1 := sort.Search(len(remain), func(i int) bool {
				return remain[i] < max
			})

			min := vi - vj
			if min < 0 {
				min = -min
			}
			if min >= max {
				continue
			}
			k2 := sort.Search(len(remain), func(i int) bool {
				return remain[i] <= min
			})
			count := k2 - k1
			if count <= 0 {
				continue
			}
			res += count
		}
	}
	return res
}
