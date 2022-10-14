package increasingtripletsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{nums: []int{2, 1, 5, 0, 4, 6}}, true},
		{"2", args{nums: []int{20, 100, 10, 12, 5, 13}}, true},
		{"3", args{nums: []int{0, 4, 1, -1, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingTriplet(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
