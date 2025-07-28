package climbing_stairs

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

// 解法二：记忆优化递归
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

// 解法三：迭代求解（推荐）
// Time: O(n), Space: O(1)
func climbStairsIterative(n int) int {
	switch {
	case n > 2:
		x, y := 1, 2
		for i := 3; i <= n; i++ {
			x, y = y, x+y
		}

		return y

	case n > 0:
		return n

	default:
		return 0
	}
}
