package length_of_longest_substring

func lengthOfLongestSubstringSlidingWindow(s string) int {
	// 用哈希表记录所有出现过的字符及其索引
	existed := map[rune]int{}

	// 用左右指针作为滑动窗口来获取最长不重复的子字符串
	str := []rune(s)
	left, maxLen := 0, 0
	for right := range str {
		// 出现重复字符时，将左指针移动到重复字符索引的+1位置
		if idx, ok := existed[str[right]]; ok && idx >= left {
			left = idx + 1
		}

		// 因为要记录每个字符最新出现的索引，所以要不断更新哈希表
		existed[str[right]] = right

		winSize := right - left + 1
		if winSize > maxLen {
			maxLen = winSize
		}
	}

	return maxLen
}
