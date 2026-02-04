package minimum_size_subarray_sum

import "math"

// Solution 1:
// 本题是 76 题的基础
// Time: O(n), Space: O(1)
func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, math.MaxInt

	for right := range nums {
		sum += nums[right]

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}
