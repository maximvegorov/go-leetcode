package best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, len(prices))
	max := -prices[0] - fee
	for i := 1; i < len(dp); i++ {
		pi := prices[i]
		dpi := pi + max
		prev := dp[i-1]
		if dpi < prev {
			dpi = prev
		}
		dp[i] = dpi
		if maxi := prev - pi - fee; maxi > max {
			max = maxi
		}
	}
	return dp[len(dp)-1]
}
