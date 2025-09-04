package reverse_string

// 解法一：迭代
// Time: O(n), Space: O(1)
func reverseStringIterative(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// 解法二（推荐）：双指针
// Time: O(n), Space: O(1)
func reverseStringTwoPoints(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 解法三：递归
func reverseStringRecursive(s []byte) {
	reverse(s, 0, len(s)-1)
}

func reverse(s []byte, left, right int) {
	if left >= right {
		return
	}
	s[left], s[right] = s[right], s[left]

	reverse(s, left+1, right-1)
}
