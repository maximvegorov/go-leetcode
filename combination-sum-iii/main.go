package combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	var res collector
	var digits []byte
	for d := 1; d <= 9; d++ {
		digits = append(digits, byte(d))
	}
	appendCombinations(k, n, digits, &res)
	return res.combs
}

func appendCombinations(k int, n int, digits []byte, dest *collector) {
	if k == 0 {
		if n == 0 {
			dest.append([]int{})
		}
		return
	}
	for i, d := range digits {
		var res collector
		appendCombinations(k-1, n-int(d), digits[i+1:], &res)
		for _, tail := range res.combs {
			comb := make([]int, len(tail)+1)
			comb[0] = int(d)
			copy(comb[1:], tail)
			dest.append(comb)
		}
	}
}

type collector struct {
	combs [][]int
}

func (c *collector) append(comb []int) {
	c.combs = append(c.combs, comb)
}
