package climbing_stairs

// 本题的本质就是509题的斐波那契数

// 解法一：朴素递归
// Time: O(2^n), Space: O(1)
func climbStairsRecursive(n int) int {
	switch {
	case n > 2:
		return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
	case n > 0:
		return n
	default:
		return 0
	}
}

// 解法一优化：记忆优化递归
// Time: O(n), Space: O(n)
var memo = map[int]int{}

func climbStairsRecursiveMemo(n int) int {
	switch {
	case n > 2:
		if val, ok := memo[n]; ok {
			return val
		}

		memo[n] = climbStairsRecursiveMemo(n-1) + climbStairsRecursiveMemo(n-2)
		return memo[n]
	case n > 0:
		return n
	default:
		return 0
	}
}

// 解法二：DP
// Time: O(n), Space: O(n)
func climbStairsDP(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1) // 当输入为1时，n+1可以避免dp[2]的越界
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 解法三（推荐）：迭代求解，可将解法二的空间复杂度优化为O(1)
// Time: O(n), Space: O(1)
func climbStairsIterative(n int) int {
	if n <= 2 {
		return n
	}

	x, y := 1, 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}
