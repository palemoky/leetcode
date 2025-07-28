package climbing_stairs

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

func climbStairsIterative(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	x, y := 1, 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}
