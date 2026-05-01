package length_of_longest_substring

// 解法一：哈希表 + 滑动窗口
// Time: O(n), Space: O(n)
func lengthOfLongestSubstringSlidingWindow(s string) int {
	lastSeen := map[rune]int{} // 反向映射哈希表

	left, maxLen := 0, 0
	for right, char := range s {
		// 注意检查 prev >= left，防止左指针回退
		if prev, seen := lastSeen[char]; seen && prev >= left {
			left = prev + 1 // 出现重复字符时，将左指针移动到重复字符索引的 +1 位置
		}

		lastSeen[char] = right

		maxLen = max(maxLen, right-left+1) // 差值计算的是区间，注意计数需 +1
	}

	return maxLen
}
