package reverse_words_in_a_string_ii

// Solution 1: 两次翻转（原地算法）
// 核心思想：先整体翻转，再逐个翻转每个单词，可参照 189 题的图解 reverse.png
// Time: O(n), Space: O(1)
func reverseWords(s []byte) {
	n := len(s)
	if n == 0 {
		return
	}

	// 1. 翻转整个数组
	reverse(s, 0, n-1)

	// 2. 翻转每个单词（用双指针找单词边界）
	start := 0
	// 因为右指针指向单词末尾的下一个位置，所以需要遍历到 n 以处理最后一个单词
	for i := 0; i <= n; i++ {
		// 遇到空格或到达末尾，说明找到一个单词
		if i == n || s[i] == ' ' {
			// 翻转 [start, i-1] 范围的单词
			reverse(s, start, i-1)
			// 下一个单词的起始位置
			start = i + 1
		}
	}
}

// 翻转数组的 [left, right] 范围
func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
