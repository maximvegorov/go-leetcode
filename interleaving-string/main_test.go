package interleaving_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		//{"2", args{"aabcc", "dbbca", "aadbbbaccc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3)
			assert.Equal(t, tt.want, got)
		})
	}
}
