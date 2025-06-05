package solution

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
