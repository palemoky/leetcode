package reverse_words_in_a_string_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"Example 1", []byte("the sky is blue"), []byte("blue is sky the")},
		{"Example 2", []byte("a good   example"), []byte("example   good a")},
		{"Single word", []byte("hello"), []byte("hello")},
		{"Two words", []byte("hello world"), []byte("world hello")},
		{"Empty", []byte(""), []byte("")},
		{"Single char", []byte("a"), []byte("a")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Make a copy since the function modifies in-place
			s := make([]byte, len(tc.input))
			copy(s, tc.input)
			reverseWords(s)
			assert.Equal(t, tc.want, s)
		})
	}
}
