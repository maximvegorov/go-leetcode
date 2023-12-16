package merge_strings_alternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	} else if len(word2) == 0 {
		return word1
	}

	var b strings.Builder
	b.Grow(len(word1) + len(word2))
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		if i == j {
			b.WriteByte(word1[i])
			i++
		} else {
			b.WriteByte(word2[j])
			j++
		}
	}

	if i < len(word1) {
		b.WriteString(word1[i:])
	}

	if j < len(word2) {
		b.WriteString(word2[j:])
	}

	return b.String()
}
