package kdiffpairsinanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 1, 4, 1, 5}, 2}, 2},
		{"2", args{[]int{1, 2, 3, 4, 5}, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPairs(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
