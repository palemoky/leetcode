package best_time_to_buy_and_sell_stock_ii

func maxProfitNaive(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

func maxProfitOptimized(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit += max(0, prices[i]-prices[i-1])
	}

	return maxProfit
}
