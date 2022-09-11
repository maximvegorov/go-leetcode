package sum_closest

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := nums[0] + nums[1] + nums[2]
	if len(nums) == 3 {
		return sum
	}
	dist := abs(sum - target)
	for j := 1; j < len(nums)-1; j++ {
		nj := nums[j]
		l := j - 1
		r := j + 1
		for l >= 0 && r < len(nums) {
			s3 := nums[l] + nj + nums[r]
			d3 := abs(s3 - target)
			if d3 < dist {
				sum = s3
				dist = d3
			}
			if s3 <= target {
				r++
			} else {
				l--
			}
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
