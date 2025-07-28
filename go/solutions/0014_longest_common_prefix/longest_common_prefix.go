package longest_common_prefix

import "strings"

// Time: O(S), Space: O(1)
func longestCommonPrefixVerticalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	baseStr := strs[0]
	// 虽然这里是两个for循环嵌套，但最坏的情况也只是遍历了一遍所有字符串，所以时间复杂度是O(N)
	for i := 0; i < len(baseStr); i++ {
		for j := 1; j < len(strs); j++ {
			// 结束条件：
			// 1. 待匹配字符与基准字符不同
			// 2. 基准字符串就是最短长度，则其为最长公共前缀
			if i == len(strs[j]) || baseStr[i] != strs[j][i] {
				return baseStr[:i]
			}
		}
	}

	return baseStr
}

// Time: O(S), Space: O(1)
func longestCommonPrefixHorizontalScanningBuiltin(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 1. 假设第一个字符串就是公共前缀
	prefix := strs[0]
	// 2. 遍历从第二个开始的每一个字符串
	for i := 1; i < len(strs); i++ {
		// 3. 不断缩短 prefix，直到当前字符串 (strs[i]) 以它为前缀为止
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			break
		}
	}

	return prefix
}

// Time: O(S), Space: O(1)
func longestCommonPrefixHorizontalScanningByIndex(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	// 遍历从第二个开始的每一个字符串
	for i := 1; i < len(strs); i++ {
		currentStr := strs[i]
		// 找到当前 prefix 和 currentStr 的公共前缀长度
		j := 0
		for j < len(prefix) && j < len(currentStr) && prefix[j] == currentStr[j] {
			j++
		}

		// 根据公共长度来更新 prefix
		prefix = prefix[:j]

		if prefix == "" {
			return ""
		}
	}

	return prefix
}
