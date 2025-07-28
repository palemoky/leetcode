package reverse_string

// 解法一：迭代
// Time: O(n), Space: O(1)
func reverseStringIterative(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// 解法二：双指针（推荐）
// Time: O(n), Space: O(1)
func reverseStringTwoPoints(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
