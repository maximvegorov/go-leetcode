package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findBall(t *testing.T) {
	tests := []*struct {
		name string
		arg  [][]int
		want []int
	}{
		{"1", [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}, []int{1, -1, -1, -1, -1}},
		{"2", [][]int{{-1}}, []int{-1}},
		{"3", [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}, []int{0, 1, 2, 3, 4, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBall(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
