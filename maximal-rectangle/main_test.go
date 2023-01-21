package maximalrectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]byte{{'0', '0', '1'}, {'1', '1', '1'}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximalRectangle(tt.args.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}
