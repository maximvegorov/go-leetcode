package minimumremovetomakevalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"lee(t(c)o)de)"}, "lee(t(c)o)de"},
		{"2", args{"a)b(c)d"}, "ab(c)d"},
		{"3", args{"))(("}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minRemoveToMakeValid(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
