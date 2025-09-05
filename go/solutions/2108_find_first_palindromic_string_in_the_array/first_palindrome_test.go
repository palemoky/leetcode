package find_first_palindromic_string_in_the_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPalindrome(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		words []string
		want  string
	}{
		{"leetcode_example1", []string{"abc", "car", "ada", "racecar", "cool"}, "ada"},
		{"leetcode_example2", []string{"notapalindrome", "racecar"}, "racecar"},
		{"no_palindrome", []string{"abc", "def", "ghi"}, ""},
		{"first_is_palindrome", []string{"a", "b", "c"}, "a"},
		{"empty_array", []string{}, ""},
		{"single_palindrome", []string{"level"}, "level"},
		{"mixed_case", []string{"DeeD", "noon", "civic"}, "DeeD"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := firstPalindrome(append([]string{}, tc.words...))
			assert.Equal(t, tc.want, got)
		})
	}
}
