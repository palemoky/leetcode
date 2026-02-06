package two_sum

// 注意本题返回的是数组的下标，因此不能排序后用双指针处理，这会导致索引丢失

// 解法一：暴力解法
// Time: O(n^2), Space: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 解法二：哈希表+一次遍历
// Time: O(N), Space: O(N)
func twoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{i, j}
		}

		seen[num] = i
	}

	return []int{}
}
