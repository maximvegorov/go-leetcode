package best_time_to_buy_and_sell_stock_iii

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0

	dp1 := make([]int, len(prices))
	{
		minPrice := prices[0]
		for i := 1; i < len(prices); i++ {
			pi := prices[i]
			profit := pi - minPrice
			if prev := dp1[i-1]; profit < prev {
				profit = prev
			}
			dp1[i] = profit
			if pi < minPrice {
				minPrice = pi
			}
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	dp2 := make([]int, len(prices))
	{
		maxPrice := prices[len(prices)-1]
		for i := len(prices) - 2; i >= 0; i-- {
			pi := prices[i]
			profit := maxPrice - pi
			if prev := dp2[i+1]; profit < prev {
				profit = prev
			}
			dp2[i] = profit
			if pi > maxPrice {
				maxPrice = pi
			}
		}
	}

	for i := 1; i < len(prices)-1; i++ {
		pi := dp1[i] + dp2[i+1]
		if pi > maxProfit {
			maxProfit = pi
		}
	}
	return maxProfit
}
