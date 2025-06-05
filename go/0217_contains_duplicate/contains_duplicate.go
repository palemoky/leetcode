package leetcode

// Time: O(n), Space: O(n)
func containsDuplicateHashMap(nums []int) bool {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	for _, count := range counter {
		if count >= 2 {
			return true
		}
	}

	return false
}
