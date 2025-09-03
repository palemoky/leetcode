package fibonacci_number

// 注意：斐波那契数有从 0 和从 1 开始两个版本，其判断条件也分别是 n>=2 和 n>=3
// Version 1: 0 1 1 2 3 5 8 13   n >= 2
// Version 2: 1 1 2 3 5 8 13 21  n >= 3
// 本题采用的是 Version 1

// 解法一：朴素递归法 (Naive Recursive)
// 思路直观，但大量重复计算导致效率极低
// Time: O(2^n), Space: O(1)
func fibRecursive(n int) int {
	switch {
	case n >= 2:
		return fibRecursive(n-1) + fibRecursive(n-2)
	case n >= 0:
		return n
	default: // n < 0
		return 0
	}
}

// 解法一优化：递归记忆优化
// Time: O(n), Space: O(n)
var memo = map[int]int{}

func fibRecursiveMemo(n int) int {
	switch {
	case n >= 2:
		if val, ok := memo[n]; ok {
			return val
		}

		memo[n] = fibRecursiveMemo(n-2) + fibRecursiveMemo(n-1)

		return memo[n]
	case n >= 0:
		return n
	default: // n < 0
		return 0
	}
}

// 解法二：递归求解
// Time: O(n), Space: O(n)
func fibDP(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1) // 注意 n 是从 0 开始，因此要 n+1
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 解法二优化：通过 3 个变量滚动来实现将空间复杂度降为 O(1)
// Time: O(n), Space: O(1)
func fibDPOptimized(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, 2) // 注意 n 是从 0 开始，因此要 n+1
	dp[0], dp[1] = 0, 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = dp[0] + dp[1]
		dp[0], dp[1] = dp[1], sum
	}

	return dp[1]
}

// 解法三：迭代法 (Iterative)
// 最高效、最推荐的解法
// Time: O(n), Space: O(1)
func fibIterative(n int) int {
	if n < 2 {
		return n
	}

	x, y := 0, 1
	for i := 2; i <= n; i++ { // 注意此处的条件是 <=
		// x, y 分别代表 f(i-2) 和 f(i-1)
		// 计算 f(i) 并更新 x 和 y
		x, y = y, x+y
	}

	return y
}
