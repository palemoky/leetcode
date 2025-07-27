package single_number

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

// Time: O(n), Space: O(1)
func singleNumberBitWise(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	return ans
}
