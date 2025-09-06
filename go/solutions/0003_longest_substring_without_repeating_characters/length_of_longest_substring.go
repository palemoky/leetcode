package length_of_longest_substring

// 解法一：哈希表+滑动窗口
// 本题求解时需要注意两个 +1：
// 1. 当哈希表发现重复值时，左指针应移动到其首次出现的 +1 位置
// 2. 计算最大窗口时差值需要 +1
func lengthOfLongestSubstringSlidingWindow(s string) int {
	// 用哈希表记录所有出现过的字符及其索引
	seen := map[byte]int{}

	// 用左右指针作为滑动窗口来获取最长不重复的子字符串
	left, maxWin := 0, 0
	for right := 0; right < len(s); right++ {
		// 出现重复字符时，将左指针移动到重复字符索引的+1位置
		if i, ok := seen[s[right]]; ok && i >= left {
			left = i + 1
		}

		// 因为要记录每个字符最新出现的索引，所以要不断更新哈希表
		seen[s[right]] = right
		maxWin = max(maxWin, right-left+1)
	}

	return maxWin
}
