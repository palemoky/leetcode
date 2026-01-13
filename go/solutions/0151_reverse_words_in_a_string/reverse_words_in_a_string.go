package reverse_words_in_a_string

// Solution 1: 提取单词 → 逆序遍历 → 拼接
// 本题与 186、1592 题相同思路解法
// Time: O(n), Space: O(n)
func reverseWords(s string) string {
	// 1. 找出单词加入数组
	word := []byte{}
	words := []string{}
	for i := range s {
		if s[i] != ' ' {
			word = append(word, s[i])
		} else if len(word) > 0 {
			words = append(words, string(word))
			word = []byte{}
		}
	}

	// 处理最后一个单词
	if len(word) > 0 {
		words = append(words, string(word))
	}

	// 边界情况：没有单词
	if len(words) == 0 {
		return ""
	}

	// 2. 逆序遍历单词数组构建结果，单词间用空格分隔
	ans := ""
	for i := len(words) - 1; i >= 0; i-- {
		ans += words[i] + " "
	}

	return ans[:len(ans)-1]
}
