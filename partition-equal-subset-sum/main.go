package partition_equal_subset_sum

import "sort"

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	target := sum / 2
	target -= nums[0]
	if target < 0 {
		return false
	}
	s := newSolver(nums[1:], target)
	return s.solve(0, target)
}

type Status byte

const (
	None Status = iota
	False
	True
)

type solver struct {
	nums  []int
	cache [][]Status
}

func newSolver(nums []int, target int) *solver {
	cache := make([][]Status, len(nums)+1)
	for i := range cache {
		cache[i] = make([]Status, target+1)
	}
	return &solver{
		nums:  nums,
		cache: cache,
	}
}

func (s *solver) solve(i int, target int) bool {
	if target < 0 {
		return false
	} else if target == 0 {
		return true
	}
	if i >= len(s.nums) {
		return false
	}
	if status := s.cache[i][target]; status != None {
		return status == True
	}
	res := False
	if s.solve(i+1, target-s.nums[i]) || s.solve(i+1, target) {
		res = True
	}
	s.cache[i][target] = res
	return res == True
}
