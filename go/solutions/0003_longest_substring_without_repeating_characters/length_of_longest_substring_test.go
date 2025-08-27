package length_of_longest_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		// --- 1. 基本情况 ---
		{name: "Empty String", input: "", want: 0},
		{name: "Single Character String", input: "a", want: 1},

		// --- 2. 典型情况 ---
		{name: "Typical repeat at start/end", input: "abcabcbb", want: 3}, // "abc"
		{name: "All same character", input: "bbbbb", want: 1},             // "b"
		{name: "Repeat with jump", input: "pwwkew", want: 3},              // "wke"
		{name: "All Unique Characters", input: "abcdefg", want: 7},

		// --- 3. 关键边界情况 ---
		{name: "Repetition at the end", input: "abcdeaf", want: 6}, // "bcdeaf"
		{
			name: "Repetition requires window to jump",
			// 当遇到第二个 'b' 时，左指针需要从 a 的位置跳到第一个 b 的下一个位置
			input: "tmmzuxt",
			want:  5, // "mzuxt"
		},
		{
			name: "The 'abba' case - critical for jump logic",
			// 当遇到第二个 'a' 时，左指针必须跳过 'b'，直接移动到第一个 'a' 的下一个位置
			input: "abba",
			want:  2, // "ab" or "ba"
		},
		{
			name:  "String with only spaces",
			input: "   ",
			want:  1, // " "
		},

		// --- 4. Unicode 字符 (检验算法对 rune 的处理是否正确) ---
		{name: "String with unicode characters", input: "你好世界你好", want: 4},          // "你好世界"
		{name: "String with repeating unicode characters", input: "我爱爱北京", want: 3}, // "我爱" or "爱北京"
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstringSlidingWindow(tc.input)
			// 使用 %q 可以让字符串在打印时带上引号，更清晰
			assert.Equal(t, tc.want, got, "lengthOfLongestSubstring(%q) failed", tc.input)
		})
	}
}
