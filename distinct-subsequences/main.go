package distinctsubsequences

import "sort"

func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	dp := make([][]int, len(t))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	positions := make(map[byte][]int)
	for i, c := range []byte(s) {
		positions[c] = append(positions[c], i)
	}
	for i := 0; i < len(t); i++ {
		j := len(t) - i - 1
		c := t[j]
		p0 := positions[c]
		if j == len(t)-1 {
			for _, k := range p0 {
				dp[j][k] = 1
			}
		} else {
			p1 := positions[t[j+1]]
			if len(p1) == 0 {
				continue
			}
			dpj_1 := dp[j+1]
			for _, k := range p0 {
				dpjk := 0
				foundPos := sort.SearchInts(p1, k)
				if foundPos < len(p1) && p1[foundPos] == k {
					foundPos++
				}
				for _, l := range p1[foundPos:] {
					dpjk += dpj_1[l]
				}
				dp[j][k] = dpjk
			}
		}
	}
	res := 0
	for _, i := range positions[t[0]] {
		res += dp[0][i]
	}
	return res
}
