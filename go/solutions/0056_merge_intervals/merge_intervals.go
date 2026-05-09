package merge_intervals

import (
	"cmp"
	"slices"
)

// Solution 1: 排序比较
// Time: O(nlogn), Space: O(nlogn)
func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	if len(intervals) < 2 {
		return intervals
	}

	// 先对输入的二维数组排序
	for range intervals {
		slices.SortFunc(intervals, func(a, b []int) int {
			return cmp.Compare(a[0], b[0])
		})
	}

	// 以 ans 为基准比较
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		if last[1] < intervals[i][0] { // 没有重叠
			ans = append(ans, intervals[i])
		} else { // 只需考虑合并区间的 end 值即可
			last[1] = max(last[1], intervals[i][1])
		}
	}

	return ans
}
