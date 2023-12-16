package determine_if_two_strings_are_close

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloseStrings(t *testing.T) {
	testCases := []*struct {
		word1, word2 string
		expected     bool
	}{
		{"abc", "bac", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := closeStrings(tc.word1, tc.word2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
