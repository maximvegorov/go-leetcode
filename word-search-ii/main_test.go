package wordsearchii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"a", args{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}}, []string{"eat", "oath"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.args.board, tt.args.words)
			assert.Equal(t, tt.want, got)
		})
	}
}
