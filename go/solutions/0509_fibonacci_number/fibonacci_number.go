package fibonacci_number

// n 0 1 1 2 3 5 8 13
// i 0 1 2 3 4 5 6 7

// 解法一：朴素递归法 (Naive Recursive)
// 思路直观，但大量重复计算导致效率极低
// Time: O(2^n), Space: O(1)
func fibRecursive(n int) int {
	switch {
	case n > 1:
		return fibRecursive(n-1) + fibRecursive(n-2)
	case n >= 0:
		return n
	default: // n < 0
		return 0
	}
}

// 递归记忆优化
// Time: O(n), Space: O(n)
var memo = map[int]int{}

func fibRecursiveMemo(n int) int {
	switch {
	case n > 1:
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

// 解法二：迭代法 (Iterative)
// 最高效、最推荐的解法
// Time: O(n), Space: O(1)
func fibIterative(n int) int {
	if n < 0 {
		return 0
	} else if n <= 1 {
		return n
	}

	x, y := 0, 1
	for i := 2; i <= n; i++ { // 注意此处的条件是 <=
		// a, b 分别代表 f(i-2) 和 f(i-1)
		// 计算 f(i) 并更新 a 和 b
		x, y = y, x+y
	}

	return y
}
