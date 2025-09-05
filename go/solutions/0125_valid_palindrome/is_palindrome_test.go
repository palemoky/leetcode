package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeImplementations(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		s    string
		want bool
	}{
		{"simple_palindrome", "aba", true},
		{"with_spaces", "A man a plan a canal Panama", true},
		{"with_punctuations", "A man, a plan, a canal: Panama", true},
		{"not_palindrome", "race a car", false},
		{"empty", "", true},
		{"single_char", "a", true},
		{"only_symbols", "!!!", true},
		{"numbers", "12321", true},
		{"mixed_case", "AbBa", true},
		{"mixed_case_not", "AbBc", false},
	}

	funcs := map[string]func(string) bool{
		"TwoPointers": isPalindromeTwoPointers,
		"Optimized":   isPalindromeTwoPointersOptimized,
	}

	for fnName, fn := range funcs {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.s)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
