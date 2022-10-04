package best_time_to_buy_and_sell_stock_iv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//		{"1", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
		{"2", args{2, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.args.k, tt.args.prices)
			assert.Equal(t, tt.want, got)
		})
	}
}
