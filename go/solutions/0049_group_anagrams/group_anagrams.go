package group_anagrams

import "slices"

// Solution 1: 排序
// Time: O(N*KlogK), Space: O(N*K)
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

// Solution 2: 计数
// Time: O(N*K), Space: O(N*K)
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
