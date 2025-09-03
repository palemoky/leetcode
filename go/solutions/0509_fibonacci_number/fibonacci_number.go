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

// 递归记忆优化
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

// 解法二：迭代法 (Iterative)
// 最高效、最推荐的解法
// Time: O(n), Space: O(1)
func fibIterative(n int) int {
	switch {
	case n >= 2:
		x, y := 0, 1
		for i := 2; i <= n; i++ { // 注意此处的条件是 <=
			// x, y 分别代表 f(i-2) 和 f(i-1)
			// 计算 f(i) 并更新  和 y
			x, y = y, x+y
		}

		return y

	case n > 0:
		return n

	default:
		return 0
	}
}
