package valid_parentheses

// 解法一：使用栈+哈希表
// Time: O(n), Space: O(n)
func isValid(s string) bool {
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
