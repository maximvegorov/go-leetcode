package longest_string_chain

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dp := make([]int, len(words))
	max := 0
	for i, w2 := range words {
		maxi := 0
		for j, w1 := range words[:i] {
			if isPredecessor([]byte(w1), []byte(w2)) && dp[j] > maxi {
				maxi = dp[j]
			}
		}
		maxi++
		dp[i] = maxi
		if maxi > max {
			max = maxi
		}
	}
	return max
}

func isPredecessor(w1, w2 []byte) bool {
	if len(w1)+1 != len(w2) {
		return false
	}
	i := 0
	for i < len(w1) && w1[i] == w2[i] {
		i++
	}
	j := i + 1
	for j < len(w2) {
		if w1[i] != w2[j] {
			return false
		}
		i++
		j++
	}
	return true
}
