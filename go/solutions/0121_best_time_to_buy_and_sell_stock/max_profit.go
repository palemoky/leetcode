package best_time_to_buy_and_sell_stock

import "math"

// 解法一：将第一个价格作为最低价，不断找到最低价后再求得最高利润
// Time: O(n), Space: O(1)
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

// 解法二（推荐）：本题的核心思路在于通过 min() 和 max() 不断维护最低价格和最大利润
// Time: O(n), Space: O(1)
func maxProfitGreedyOptimized(prices []int) int {
	minPrice, maxProfit := math.MaxInt64, 0
	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

// 解法三：
func maxProfitDP(prices []int) int {
	n := len(prices)
	dp := [2]int{}
	dp[0] = 0
	dp[1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}

	return dp[0]
}
