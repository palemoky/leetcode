package valid_palindrome

// 本题是非常好的练习 ASCII 字符处理的题目

// 解法一：预处理 + 双指针
// Time: O(n), Space: O(n)
func isPalindrome(s string) bool {
	caseOffset := 'a' - 'A'

	filtered := []byte{}
	for _, ch := range []byte(s) {
		switch {
		case ch >= '0' && ch <= '9':
			filtered = append(filtered, ch)
		case ch >= 'a' && ch <= 'z':
			filtered = append(filtered, ch)
		case ch >= 'A' && ch <= 'Z': // 大写转小写
			filtered = append(filtered, ch+byte(caseOffset))
		}
	}

	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 解法二：原地双指针（推荐）
// Time: O(n), Space: O(1)
func isPalindromeInPlace(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// 跳过非字母数字字符
		// 跳过左侧非字母数字字符
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// 跳过右侧非字母数字字符
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// 转小写后对比左右两侧的值
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A') // or + 32
	}
	return char
}
