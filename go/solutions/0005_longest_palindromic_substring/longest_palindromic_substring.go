package longest_palindromic_substring

// 回文不满足单调性，滑动窗口不适用；应利用回文围绕中心对称的特征，从中心向两侧扩展

// Solution 1: 中心扩展法
// Time: O(n²), Space: O(1)
func longestPalindrome(s string) string {
	start, maxLen := 0, 0

	// 从中心向两侧扩展，返回回文长度
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		// 此时 [l+1, r-1] 是回文
		if r-l-1 > maxLen {
			start = l + 1
			maxLen = r - l - 1
		}
	}

	for i := range s {
		// 无法预知以某个位置为中心的最长回文是奇数还是偶数长度，因此需要同时遍历两种情况，取最长的那个
		expand(i, i)   // 奇数长度回文（如 "aba"）
		expand(i, i+1) // 偶数长度回文（如 "abba"）
	}

	return s[start : start+maxLen]
}
