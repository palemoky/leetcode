package reverse_words_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{"Example 1", "the sky is blue", "blue is sky the"},
		{"Example 2", "  hello world  ", "world hello"},
		{"Example 3", "a good   example", "example good a"},
		{"Single word", "hello", "hello"},
		{"Two words", "hello world", "world hello"},
		{"Multiple spaces", "  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Only spaces", "   ", ""},
		{"Empty string", "", ""},
		{"Leading spaces", "  hello", "hello"},
		{"Trailing spaces", "hello  ", "hello"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := reverseWords(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}
