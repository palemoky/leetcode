package leetcode

// Time: O(n), Space: O(n)
func isValid(s string) bool {
	braces := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			braces = append(braces, v)
		} else {
			if len(braces) == 0 || braces[len(braces)-1] != pairs[v] {
				return false
			}
			braces = braces[:len(braces)-1]
		}
	}
	return len(braces) == 0
}
