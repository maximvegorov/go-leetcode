package substringwithconcatenationofallwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"2", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
		{"3", args{"ababababab", []string{"ababa", "babab"}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.args.s, tt.args.words)
			assert.Equal(t, tt.want, got)
		})
	}
}
