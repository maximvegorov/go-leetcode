package hindex

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	invH := sort.Search(len(citations), func(i int) bool {
		return citations[i] >= len(citations)-i
	})
	return len(citations) - invH
}
