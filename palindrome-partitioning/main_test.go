package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{"aab"}, [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"2", args{"a"}, [][]string{{"a"}}},
		{"3", args{"efe"}, [][]string{{"e", "f", "e"}, {"efe"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
