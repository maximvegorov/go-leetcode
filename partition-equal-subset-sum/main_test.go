package partition_equal_subset_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 2, 3, 5}}, false},
		{"2", args{[]int{1, 3, 5, 5, 5, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartition(tt.args.nums)
			assert.Equal(t, got, tt.want)
		})
	}
}
