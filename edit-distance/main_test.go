package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	testCases := []*struct {
		s1, s2   string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q - %q", tc.s1, tc.s2), func(t *testing.T) {
			actual := minDistance(tc.s1, tc.s2)
			assert.Equal(t, tc.expected, actual)

		})
	}
}
