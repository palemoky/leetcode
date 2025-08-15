package count_special_quadruplets

// 本题可以认为是三数之和，和的位置在i>=3，所以先锁定和的位置，然后在和的左侧搜索，查找是否有符合要求的结果，有就计数

// 解法一：迭代法
// Time: O(n^4), Space: O(1)
func countQuadruplets(nums []int) int {
	count := 0
	for d := len(nums) - 1; d >= 3; d-- {
		for c := d - 1; c >= 2; c-- {
			for b := c - 1; b >= 1; b-- {
				for a := b - 1; a >= 0; a-- {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}

	return count
}

// 解法二：哈希表
// 由 nums[a]+nums[b]+nums[c] == nums[d] 变形得 nums[a] + nums[b] == nums[d] - nums[c]
func countQuadrupletsHashMap(nums []int) int {
	n := len(nums)
	if n < 4 {
		return 0
	}
	ans := 0
	diff := map[int]int{} // diff = nums[d] - nums[c] 的计数

	// b 从 n-3 开始向左移动，确保存在 c=b+1 和 d>c
	for b := n - 3; b >= 1; b-- {
		// 把所有以 c = b+1 为左端的 (c,d) 对加入 diff
		c := b + 1
		for d := c + 1; d < n; d++ {
			diff[nums[d]-nums[c]]++
		}
		// 统计所有 a < b，使得 nums[a]+nums[b] 等于某个 diff
		for a := 0; a < b; a++ {
			ans += diff[nums[a]+nums[b]]
		}
	}
	return ans
}
