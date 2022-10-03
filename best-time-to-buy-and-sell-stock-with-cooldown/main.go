package best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, len(prices)+2)
	for i := len(prices) - 1; i >= 0; i-- {
		pi := 0
		for j := i + 1; j < len(prices); j++ {
			pj := prices[j] - prices[i] + dp[j+2]
			if pj > pi {
				pi = pj
			}
		}
		if pi < dp[i+1] {
			pi = dp[i+1]
		}
		dp[i] = pi
	}
	return dp[0]
}
