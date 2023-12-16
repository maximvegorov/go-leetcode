package search_suggestions_system

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	testCases := []*struct {
		products   []string
		searchWord string
		expected   [][]string
	}{
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"}}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := suggestedProducts(tc.products, tc.searchWord)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
