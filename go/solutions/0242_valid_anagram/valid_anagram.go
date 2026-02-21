package valid_anagram

import "slices"

// Solution 1: 既然是查找异位词，那么排序后就完全一样
// Time: O(nlogn), Space: O(n)
func isAnagramSorting(s string, t string) bool {
	sb, tb := []byte(s), []byte(t)
	slices.Sort(sb)
	slices.Sort(tb)

	return string(sb) == string(tb)
}

// Solution 2: 如果两个字符串是异位词，那么每个字符出现的总和为零
// Time: O(n), Space: O(1)
func isAnagramCounting(s string, t string) bool {
	count := [26]int{}
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range t {
		count[ch-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
