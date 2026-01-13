package rearrange_spaces_between_words

// Solution 1: 统计空格和单词，然后重新分配
// Time: O(n), Space: O(n)
func rearrangeSpaces(text string) string {
	// 1. 统计空格总数
	totalSpaces := 0
	for _, ch := range text {
		if ch == ' ' {
			totalSpaces++
		}
	}

	// 2. 提取所有单词（使用 strings.Fields 自动处理多个连续空格）
	words := []string{}
	word := []byte{}
	for i := range text {
		if text[i] != ' ' { // 把单词逐个添加到 word 中
			word = append(word, text[i])
		} else if len(word) > 0 { // 遇到空格，说明一个单词结束
			words = append(words, string(word))
			word = []byte{} // 重置 word
		}
	}
	// 处理最后一个单词
	if len(word) > 0 {
		words = append(words, string(word))
	}

	wordCount := len(words)

	// 3. 边界情况：只有一个单词
	if wordCount == 1 {
		spaces := ""
		for range totalSpaces {
			spaces += " "
		}
		return words[0] + spaces
	}

	// 4. 计算每个间隙的空格数和剩余空格数
	gapSpaces := totalSpaces / (wordCount - 1)
	tailSpaces := totalSpaces % (wordCount - 1)

	// 5. 构建结果
	result := []byte{}
	for i, w := range words {
		// 先放单词
		result = append(result, []byte(w)...)

		// 最后一个单词后不加间隙空格，只加剩余空格
		if i < wordCount-1 {
			// 添加间隙空格
			for range gapSpaces {
				result = append(result, ' ')
			}
		}
	}

	// 添加剩余空格到末尾
	for range tailSpaces {
		result = append(result, ' ')
	}

	return string(result)
}
