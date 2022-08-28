package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	testCases := []*struct {
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{3, [][]int{{2, 0}, {2, 1}}, []int{1, 0, 2}},
		{3, [][]int{{0, 2}, {1, 2}, {2, 0}}, []int{}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
		{7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}, []int{6, 5, 4, 2, 3, 0, 1}},
		{8, [][]int{{1, 0}, {2, 6}, {1, 7}, {5, 1}, {6, 4}, {7, 0}, {0, 5}}, []int{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			actual := findOrder(tc.numCourses, tc.prerequisites)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
