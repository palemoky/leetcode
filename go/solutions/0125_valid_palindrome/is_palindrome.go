package valid_palindrome

// 本题是非常好的练习 ASCII 字符处理的题目

// 解法一：使用对撞指针
// Time: O(n), Space: O(1)
func isPalindromeTwoPointers(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		leftLetter, rightLetter := s[left], s[right]
		// 跳过非字母字符
		isLowerLeft := leftLetter >= 'a' && leftLetter <= 'z'
		isUpperLeft := leftLetter >= 'A' && leftLetter <= 'Z'
		isNumberLeft := leftLetter >= '0' && leftLetter <= '9'
		if !isLowerLeft && !isUpperLeft && !isNumberLeft {
			left++
			continue
		}

		isLowerRight := rightLetter >= 'a' && rightLetter <= 'z'
		isUpperRight := rightLetter >= 'A' && rightLetter <= 'Z'
		isNumberRight := rightLetter >= '0' && rightLetter <= '9'
		if !isLowerRight && !isUpperRight && !isNumberRight {
			right--
			continue
		}

		// 所有大写转小写
		leftChar, rightChar := leftLetter, rightLetter // 默认为小写
		if leftLetter >= 'A' && leftLetter <= 'Z' {
			leftChar = +leftLetter + 32
		}
		if rightLetter >= 'A' && rightLetter <= 'Z' {
			rightChar = rightLetter + 32
		}

		// 对比左右两侧的值
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}

	return true
}

// 解法一优化（推荐）
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

func isPalindromeTwoPointersOptimized(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// 跳过非字母字符
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		// 所有大写转小写
		leftChar, rightChar := toLower(s[left]), toLower(s[right])

		// 对比左右两侧的值
		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}
