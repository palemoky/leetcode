package longest_consecutive_sequence

// Solution 1: 哈希表
// Time: O(n), Space: O(n)
func longestConsecutive(nums []int) int {
	has := map[int]bool{} // 直接初始化为 bool 可以简化哈希表判断
	for _, num := range nums {
		has[num] = true
	}

	ans := 0
	for num := range has {
		prev, next := num-1, num+1
		if has[prev] { // 上一个数存在于哈希表，说明 num 不是序列的起点
			continue
		}

		for has[next] { // 不断查找下一个序列是否在哈希表中
			next++
		}

		ans = max(ans, next-num)
	}

	return ans
}
