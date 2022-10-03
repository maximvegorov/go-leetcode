package best_time_to_buy_and_sell_stock_with_transaction_fee

/*
O(len(prices)^2)
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, len(prices)+1)
	for i := len(prices) - 1; i >= 0; i-- {
		pi := 0
		for j := i + 1; j < len(prices); j++ {
			pj := prices[j] - prices[i] + dp[j+1]
			if pj > pi {
				pi = pj
			}
		}
		pi -= fee
		if pi < dp[i+1] {
			pi = dp[i+1]
		}
		dp[i] = pi
	}
	return dp[0]
}
