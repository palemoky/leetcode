package find_first_palindromic_string_in_the_array

// 解法一：使用对撞指针判断回文字符串
// Time: O(n), Space: O(1)
func firstPalindrome(words []string) string {
	ans := ""
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ans
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
