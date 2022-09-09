package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2}, 8},
		{"2", args{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leastInterval(tt.args.tasks, tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
