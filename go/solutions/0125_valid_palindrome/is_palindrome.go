package valid_palindrome

// 本题是非常好的练习 ASCII 字符处理的题目

// 解法一：双指针对撞
// Time: O(n), Space: O(1)
func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A') // or + 32
	}
	return char
}

func isPalindromeTwoPointers(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// 跳过非字母数字字符
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
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
