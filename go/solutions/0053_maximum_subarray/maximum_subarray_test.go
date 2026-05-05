package maximum_subarray

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
			name:     "example 1: mixed positive and negative numbers",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "example 2: single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "example 3: all positive numbers",
			input:    []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "all negative numbers",
			input:    []int{-8, -3, -6, -2, -5, -4},
			expected: -2,
		},
		{
			name:     "best subarray in the middle",
			input:    []int{-1, -2, 10, -1, 2, 3, -20, 4},
			expected: 14,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"Kadane": maxSubArray,
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
