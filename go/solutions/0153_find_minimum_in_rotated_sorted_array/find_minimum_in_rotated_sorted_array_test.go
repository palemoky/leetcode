package find_minimum_in_rotated_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1: rotated array",
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "example 2: rotated array with more elements",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "example 3: not rotated (already sorted)",
			input:    []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "two elements rotated",
			input:    []int{2, 1},
			expected: 1,
		},
		{
			name:     "two elements not rotated",
			input:    []int{1, 2},
			expected: 1,
		},
		{
			name:     "minimum at the end",
			input:    []int{2, 3, 4, 5, 1},
			expected: 1,
		},
		{
			name:     "minimum at the beginning",
			input:    []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"findMin": findMin,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
