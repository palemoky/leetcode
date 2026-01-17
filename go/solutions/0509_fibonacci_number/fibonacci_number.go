package fibonacci_number

// 注意：斐波那契数有从 0 和从 1 开始两个版本，其判断条件也分别是 n>=2 和 n>=3
// Version 1: 0 1 1 2 3 5 8 13   n >= 2
// Version 2: 1 1 2 3 5 8 13 21  n >= 3
// 本题采用的是 Version 1

// 解法一：朴素递归法 (Naive Recursive)
// 思路直观，但大量重复计算导致效率极低
// Time: O(2^n), Space: O(1)
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

// 解法一优化：递归记忆优化
// Time: O(n), Space: O(n)
func fibRecursiveMemo(n int) int {
	memo := make(map[int]int)
	return fibMemoHelper(n, memo)
}

func fibMemoHelper(n int, memo map[int]int) int {
	// 基础情况
	if n < 2 {
		return n
	}

	// 检查缓存
	if val, ok := memo[n]; ok {
		return val
	}

	// 先计算 n-1 再计算 n-2，更符合缓存局部性原理
	memo[n] = fibMemoHelper(n-1, memo) + fibMemoHelper(n-2, memo)
	return memo[n]
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

// 解法三：迭代法 (Iterative)，可将解法二的空间复杂度优化为O(1)
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
