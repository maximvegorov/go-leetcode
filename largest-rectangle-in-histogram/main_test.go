package largestrectangleinhistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
		{"2", args{[]int{2, 1, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.args.heights)
			assert.Equal(t, tt.want, got)
		})
	}
}
