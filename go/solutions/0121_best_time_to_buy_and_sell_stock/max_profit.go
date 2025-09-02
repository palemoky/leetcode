package best_time_to_buy_and_sell_stock

import "math"

func maxProfitGreedy(prices []int) int {
	minProfit, maxProfit := prices[0], 0
	for _, price := range prices {
		if price < minProfit {
			minProfit = price
		}
		if price-minProfit > maxProfit {
			maxProfit = price - minProfit
		}
	}

	return maxProfit
}

func maxProfitGreedyOptimized(prices []int) int {
	minPrice, maxProfit := math.MaxInt64, 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

func maxProfitDP(prices []int) int {
	n := len(prices)
	dp := [2]int{}
	dp[0] = 0
	dp[1] = -prices[0]
	for i := 1; i < n; i++ {
		v := prices[i]
		dp[0] = max(dp[0], dp[1]+v)
		dp[1] = max(dp[1], -v)
	}
	return dp[0]
}
