package next_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1: [1,2,3] → [1,3,2]",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "example 2: [3,2,1] → [1,2,3]",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "example 3: [1,1,5] → [1,5,1]",
			input:    []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two elements ascending: [1,2] → [2,1]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "two elements descending: [2,1] → [1,2]",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "with duplicates: [2,3,1,3,3] → [2,3,3,1,3]",
			input:    []int{2, 3, 1, 3, 3},
			expected: []int{2, 3, 3, 1, 3},
		},
	}

	funcsToTest := map[string]func([]int){
		"nextPermutation": nextPermutation,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					fn(tc.input)
					assert.Equal(t, tc.expected, tc.input)
				})
			}
		})
	}
}
