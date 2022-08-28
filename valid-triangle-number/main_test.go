package main

import "testing"

func TestTriangleNumber(t *testing.T) {
	testCases := []*struct {
		name     string
		nums     []int
		expected int
	}{
		{"1", []int{2, 2, 3, 4}, 3},
		{"2", []int{4, 2, 3, 4}, 4},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := triangleNumber(tc.nums)
			if actual != tc.expected {
				t.Errorf("%v != %v", actual, tc.expected)
			}
		})
	}
}
