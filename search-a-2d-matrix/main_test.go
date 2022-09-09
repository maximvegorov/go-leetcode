package search_a_2d_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []*struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.args.matrix, tt.args.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
