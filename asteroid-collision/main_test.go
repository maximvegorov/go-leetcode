package asteroid_collision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 10, -5}}, []int{5, 10}},
		{"2", args{[]int{-2, -2, 1, -2}}, []int{-2, -2, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.args.asteroids)
			assert.Equal(t, tt.want, got)
		})
	}
}
