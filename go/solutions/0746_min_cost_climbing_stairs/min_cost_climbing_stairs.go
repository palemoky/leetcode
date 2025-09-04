package min_cost_climbing_stairs

// 请注意，
// 1. 本题的成本是指所在i阶向上爬的门票，而不是两个台阶间的成本，也不是到达i阶的成本
// 2. 台阶共有 len(cost)+1 阶
// Time: O(n), Space: O(n)
func minCostClimbingStairsDP(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		// dp[i-1]代表累计花费，cost[i-1]代表当前门票
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

// Time: O(n), Space: O(1)
func minCostClimbingStairsIterative(cost []int) int {
	x, y := 0, 0 // 由于可以从0阶或1阶开始，因此cost均为0
	for i := 2; i <= len(cost); i++ {
		x, y = y, min(y+cost[i-1], x+cost[i-2])
	}

	return y
}
