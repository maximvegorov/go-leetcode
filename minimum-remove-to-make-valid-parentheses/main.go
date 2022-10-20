package minimumremovetomakevalidparentheses

import "bytes"

func minRemoveToMakeValid(s string) string {
	var segments [][]byte
	var stack []int
	i := 0
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c == '(' {
			if i < j {
				segments = append(segments, []byte(s[i:j]))
			}
			stack = append(stack, len(segments))
			segments = append(segments, []byte("("))
			i = j + 1
		} else if c == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				if i < j {
					segments = append(segments, []byte(s[i:j]))
				}
				i = j + 1
			}
		}
	}
	if i < len(s) {
		segments = append(segments, []byte(s[i:]))
	}
	for _, j := range stack {
		segments[j] = []byte("")
	}
	return string(bytes.Join(segments, []byte("")))
}
