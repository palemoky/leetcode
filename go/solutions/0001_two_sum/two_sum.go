package two_sum

// 注意本题返回的是数组的下标，因此不能排序后用双指针处理，这会导致索引丢失

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

// Time: O(N), Space: O(N)
func twoSumHashMap(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}

	return []int{}
}
