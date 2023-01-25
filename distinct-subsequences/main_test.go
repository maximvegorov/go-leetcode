package distinctsubsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"rabbbit", "rabbit"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDistinct(tt.args.s, tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
