package length_of_longest_substring

// 解法一：哈希表 + 滑动窗口
// Time: O(n), Space: O(n)
func lengthOfLongestSubstringSlidingWindow(s string) int {
	charIndex := map[byte]int{}

	left, maxLen := 0, 0
	for right := range s {
		// 出现重复字符时，将左指针移动到重复字符索引的 +1 位置
		// 注意检查 prevIndex >= left，防止左指针回退
		if prevIndex, exists := charIndex[s[right]]; exists && prevIndex >= left {
			left = prevIndex + 1
		}

		// 记录字符最新出现的索引
		charIndex[s[right]] = right

		// 注意窗口长度需要 +1
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
