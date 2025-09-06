package valid_parentheses

// 解法一：使用栈+哈希表
// Time: O(n), Space: O(n)
func isValidIfElse(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ { // 注意 range 返回的是 rune，这个返回的是 byte
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			n := len(stack)
			if n == 0 || stack[n-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:n-1]
		}
	}

	return len(stack) == 0
}

// 解法一的switch case写法
// Time: O(n), Space: O(n)
func isValidSwitchCase(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) > 0 && stack[len(stack)-1] == pairs[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
