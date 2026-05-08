package assign_cookies

import "sort"

// Solution 1: 双指针
// Time: O(mlogm+nlogn), Space: O(logm+logn)
func findContentChildren(g, s []int) int {
	ans := 0

	sort.Ints(g)
	sort.Ints(s)

	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		// 遍历饼干，直到满足胃口
		for j < n && g[i] > s[j] {
			j++
		}

		// 分配饼干
		if j < n {
			ans++
			j++
		}
	}

	return ans
}
