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

// 解法一：使用暴力递归（自顶向下）
// 思路：对于每个金额，尝试使用每一种硬币，递归求解剩余金额
// 问题：存在大量重复计算，时间复杂度指数级
// Time: O(len(coins)^amount), Space: O(amount) 递归栈深度
func coinChangeBruteForce(coins []int, amount int) int {
	// 基础情况
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	minCoins := amount + 1 // 初始化为不可能的大值

	// 尝试使用每一种硬币
	for _, coin := range coins {
		subResult := coinChangeBruteForce(coins, amount-coin)
		// 如果子问题有解，更新最小值
		if subResult >= 0 {
			minCoins = min(minCoins, subResult+1)
		}
	}

	// 如果没有找到解决方案
	if minCoins > amount {
		return -1
	}

	return minCoins
}

// 解法一优化：使用记忆化递归（自顶向下优化）
// 思路：在暴力递归的基础上，使用 memo 缓存已计算的结果
// Time: O(amount × len(coins)), Space: O(amount)
func coinChangeMemo(coins []int, amount int) int {
	memo := make(map[int]int)
	return coinChangeMemoHelper(coins, amount, memo)
}

func coinChangeMemoHelper(coins []int, amount int, memo map[int]int) int {
	// 基础情况
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	// 检查缓存
	if val, ok := memo[amount]; ok {
		return val
	}

	minCoins := amount + 1 // 初始化为不可能的大值

	// 尝试使用每一种硬币
	for _, coin := range coins {
		subResult := coinChangeMemoHelper(coins, amount-coin, memo)
		// 如果子问题有解，更新最小值
		if subResult >= 0 {
			minCoins = min(minCoins, subResult+1)
		}
	}

	// 缓存结果
	if minCoins > amount {
		memo[amount] = -1
	} else {
		memo[amount] = minCoins
	}

	return memo[amount]
}

// 解法二：使用动态规划（自底向上）
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
