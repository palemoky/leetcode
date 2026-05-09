package non_overlapping_intervals

import (
	"cmp"
	"math"
	"slices"
)

// Solution 1: 排序查找
// 转换思路，求最多不重叠区间
// Time: O(nlogn), Space: O(logn)
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		// 不重叠区间的核心就是前一个的右边界和下一个的左边界不重叠，所以要以右边界排序
		return cmp.Compare(a[1], b[1])
	})

	ans := 0

	// 题目中的索引可能是负数
	prevEnd := math.MinInt
	for _, interval := range intervals {
		// 比较是否重叠
		if interval[0] >= prevEnd {
			ans++
			prevEnd = interval[1]
		}
	}

	return len(intervals) - ans
}
