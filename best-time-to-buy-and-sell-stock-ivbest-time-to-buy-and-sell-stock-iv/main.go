package best_time_to_buy_and_sell_stock_iv

func maxProfit(k int, prices []int) int {
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
	for n := 2; n <= k; n++ {
		minPriceIndex := 2 * (n - 1)
		if minPriceIndex >= len(prices) {
			break
		}
		dp2[minPriceIndex] = 0
		for i := minPriceIndex + 1; i < len(prices); i++ {
			pi := prices[i]
			profit := 0
			for j := minPriceIndex; j < i; j++ {
				pj := pi - prices[j] + dp1[j-1]
				if pj > profit {
					profit = pj
				}
			}
			if prev := dp2[i-1]; profit < prev {
				profit = prev
			}
			dp2[i] = profit
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		dp1, dp2 = dp2, dp1
	}

	return maxProfit
}
