package search_suggestions_system

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	res := make([][]string, 0, len(searchWord))
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		pos := sort.SearchStrings(products, prefix)
		k := min(len(products), pos+3)
		var top3 []string
		for j := pos; j < k; j++ {
			product := products[j]
			if !strings.HasPrefix(product, prefix) {
				break
			}
			top3 = append(top3, product)
		}
		products = products[pos:]
		res = append(res, top3)
	}
	return res
}
