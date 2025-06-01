package leetcode

// Time: O(n), Space: O(n)
func singleNumberHashMap(nums []int) int {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if count == 1 {
			return num
		}
	}

	return 0
}
