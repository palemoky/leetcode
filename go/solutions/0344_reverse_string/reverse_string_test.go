package reverse_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Standard case with odd length",
			input:    []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:     "Standard case with even length",
			input:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name:     "Edge case: empty slice",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "Edge case: single element slice",
			input:    []byte{'A'},
			expected: []byte{'A'},
		},
		{
			name:     "Edge case: two elements slice",
			input:    []byte{'a', 'b'},
			expected: []byte{'b', 'a'},
		},
		{
			name:     "Slice with numbers and symbols",
			input:    []byte{'1', ' ', 'a', '&'},
			expected: []byte{'&', 'a', ' ', '1'},
		},
	}

	functionsToTest := map[string]func([]byte){
		"Iterative":   reverseStringIterative,
		"TwoPointers": reverseStringTwoPoints,
		"Recursive":   reverseStringRecursive,
	}

	for fnName, fn := range functionsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// 注意此处是原地修改，这会导致原始数据被污染，因此需使用副本以避免数据竞争
					input := append([]byte{}, tc.input...)
					fn(input)
					assert.Equal(t, tc.expected, input)
				})
			}
		})
	}
}
