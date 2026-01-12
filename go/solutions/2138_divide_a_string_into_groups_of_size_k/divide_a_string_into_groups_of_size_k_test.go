package divide_a_string_into_groups_of_size_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		s    string
		k    int
		fill byte
		want []string
	}{
		{"Example 1", "abcdefghi", 3, 'x', []string{"abc", "def", "ghi"}},
		{"Example 2 - need fill", "abcdefghij", 3, 'x', []string{"abc", "def", "ghi", "jxx"}},
		{"Single group", "abc", 3, 'x', []string{"abc"}},
		{"Need padding", "ab", 3, 'z', []string{"abz"}},
		{"k equals 1", "hello", 1, 'x', []string{"h", "e", "l", "l", "o"}},
		{"Empty string", "", 3, 'x', []string{}},
		{"Single char need fill", "a", 3, 'y', []string{"ayy"}},
		{"Multiple fills", "a", 5, '*', []string{"a****"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := divideString(tc.s, tc.k, tc.fill)
			assert.Equal(t, tc.want, got)
		})
	}
}
