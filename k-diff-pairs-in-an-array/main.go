package kdiffpairsinanarray

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}
	res := 0
	for numi, freqi := range m {
		numj := numi - k
		if k == 0 {
			if freqi > 1 {
				res += 1
			}
		} else {
			if m[numj] > 0 {
				res += 1
			}
		}
	}
	return res
}
