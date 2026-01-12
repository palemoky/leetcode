package divide_a_string_into_groups_of_size_k

// Solution 1:
// Time: O(n), Space: O(n)
func divideString(s string, k int, fill byte) []string {
	ans := []string{}
	chars := []byte(s)

	for i := 0; i < len(s); i += k {
		// 使用 min 避免越界
		end := min(i+k, len(s))
		group := chars[i:end]

		// 如果当前组不足 k 个字符，填充
		currentLen := len(group)
		if currentLen < k {
			for range k - currentLen {
				group = append(group, fill)
			}
		}

		ans = append(ans, string(group))
	}

	return ans
}
