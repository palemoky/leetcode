package group_anagrams

import "slices"

// 异位词：字母相同但顺序不同的字符串
// 本题是 242 题的升级版，图解可参看 drawio 文件和 core/checklist/go/group_anagrams 中的图解

// Solution 1: 异位词排序后就完全一样，用排序后的字符串作为 key
// Time: O(m*nlogn), Space: O(m*n)
func groupAnagramsSorting(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		slices.Sort(bytes)
		key := string(bytes)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Solution 2: 统计每个字符出现的次数作为 key
// Time: O(m*n), Space: O(m*n)
func groupAnagramsCounting(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, str := range strs {
		var count [26]int
		for _, ch := range str {
			count[ch-'a']++
		}
		groups[count] = append(groups[count], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
