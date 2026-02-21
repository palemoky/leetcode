package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func(string, string) bool{
	"Sorting":  isAnagramSorting,
	"Counting": isAnagramCounting,
}

func TestIsAnagram(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"Example 1", "anagram", "nagaram", true},
		{"Example 2", "rat", "car", false},
		{"Empty strings", "", "", true},
		{"Different lengths", "ab", "a", false},
		{"Single char same", "a", "a", true},
		{"Single char different", "a", "b", false},
		{"Repeated chars", "aacc", "ccac", false},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.want, fn(tc.s, tc.t))
				})
			}
		})
	}
}
