package determine_if_two_strings_are_close

import (
	"maps"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freqs1 := freqs(word1)
	freqs2 := freqs(word2)

	counts1 := make(map[int]int, len(freqs1))
	for ch, cnt := range freqs1 {
		if freqs2[ch] == 0 {
			return false
		}
		counts1[cnt] += 1
	}

	counts2 := make(map[int]int, len(freqs2))
	for ch, cnt := range freqs2 {
		if freqs1[ch] == 0 {
			return false
		}
		counts2[cnt] += 1
	}

	return maps.Equal(counts1, counts2)
}

func freqs(word string) map[byte]int {
	res := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		res[word[i]] += 1
	}
	return res
}
