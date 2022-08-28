package main

import "testing"

func Test_isNumber(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{"0", true},
		{"e", false},
		{".", false},
		{"0089", true},
		{"-0.1", true},
		{"+3.14", true},
		{"4.", true},
		{"-.9", true},
		{"2e10", true},
		{"-90E3", true},
		{"3e+7", true},
		{"+6e-1", true},
		{"53.5e93", true},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := isNumber(tt.arg); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
