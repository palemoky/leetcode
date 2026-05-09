package minimum_number_of_arrows_to_burst_balloons

import (
	"cmp"
	"math"
	"slices"
)

// Solution 1: 排序后比较
// Time: O(nlogn), Space: O(logn)
func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int { return cmp.Compare(a[1], b[1]) })

	count := 0
	prevEnd := math.MinInt
	for _, point := range points {
		if point[0] > prevEnd {
			count++
			prevEnd = point[1]
		}
	}

	return count
}
