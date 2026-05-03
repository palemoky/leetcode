package sqrtx

// Solution 1: 二分查找
// Time: O(logn), Space: O(1)
func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x { // 注意被截断的小数也符合要求
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}
