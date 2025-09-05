package contains_duplicate

// 解法一：使用哈希表计数，>= 2即为存在重复元素
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

// 解法一优化(推荐)：只要已经在哈希表中存在，即可认为存在重复元素
// Time: O(n), Space: O(n)
func containsDuplicateHashMapOptimized(nums []int) bool {
	scaned := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := scaned[num]; ok {
			return true
		}
		scaned[num] = struct{}{}
	}

	return false
}
