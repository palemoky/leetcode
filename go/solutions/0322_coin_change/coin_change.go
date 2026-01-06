package coin_change

import "sort"

// coinChangeErr 使用贪心算法（错误示例）
// 贪心策略：总是优先选择面值最大的硬币
// 问题：贪心算法不能保证得到最优解
// 反例：coins=[1,3,4], amount=6
//   - 贪心：4+1+1=3枚
//   - 最优：3+3=2枚
//
// Time: O(n log n + n), Space: O(1)
func coinChangeErr(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 排序，从大到小选择硬币
	sort.Ints(coins)
	nums, remain := 0, 0
	for i := len(coins) - 1; i >= 0; i-- {
		if amount >= coins[i] {
			remain = amount % coins[i]
			nums += amount / coins[i]
		}

		if remain == 0 {
			break
		}
	}

	if nums == 0 {
		return -1
	}
	return nums
}

// coinChange 使用动态规划（正确解法）
// dp[i] 表示凑成金额 i 所需的最少硬币数
// 状态转移方程：dp[i] = min(dp[i], dp[i-coin]+1)
// 含义：要凑成金额 i，可以先凑成 i-coin，再加一枚面值为 coin 的硬币
// Time: O(amount × len(coins)), Space: O(amount)
func coinChange(coins []int, amount int) int {
	// dp[i] 表示凑成金额 i 所需的最少硬币数
	dp := make([]int, amount+1)

	// 初始化：dp[0]=0（凑成0元需要0枚硬币）
	// 其他位置初始化为 amount+1（表示不可能，因为最多需要 amount 枚面值为1的硬币）
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1

		// 尝试使用每一种硬币
		for _, coin := range coins {
			if i >= coin {
				// 状态转移：dp[i] = min(dp[i], dp[i-coin]+1)
				// 如果使用面值为 coin 的硬币，需要 dp[i-coin]+1 枚
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 如果 dp[amount] 仍然大于 amount，说明无法凑成
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
