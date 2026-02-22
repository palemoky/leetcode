package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func(string) string{
	"CenterExpansion": longestPalindrome,
}

func TestLongestPalindrome(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		s    string
		want []string // 可能有多个合法答案
	}{
		{"Example 1", "babad", []string{"bab", "aba"}},
		{"Example 2", "cbbd", []string{"bb"}},
		{"Single char", "a", []string{"a"}},
		{"All same", "aaaa", []string{"aaaa"}},
		{"Even palindrome", "abba", []string{"abba"}},
		{"Entire string", "racecar", []string{"racecar"}},
		{"No palindrome > 1", "abcde", []string{"a", "b", "c", "d", "e"}},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					assert.Contains(t, tc.want, fn(tc.s))
				})
			}
		})
	}
}
