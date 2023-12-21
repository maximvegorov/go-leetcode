package evaluatedivision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_calcEquation(b *testing.B) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	for i := 0; i < b.N; i++ {
		calcEquation(equations, values, queries)
	}
}
