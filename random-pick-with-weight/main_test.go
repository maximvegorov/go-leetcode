package randompickwithweight

import (
	"fmt"
	"testing"
)

func TestSolution_PickIndex(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		want    int
	}{
		{"1", []int{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.weights)
			for i := 0; i < len(tt.weights); i++ {
				index := this.PickIndex()
				fmt.Println(index)
			}
		})
	}
}
