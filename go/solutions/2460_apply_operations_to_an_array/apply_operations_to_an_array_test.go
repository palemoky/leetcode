package apply_operations_to_an_array

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
			name:     "example 1",
			input:    []int{1, 2, 2, 1, 1, 0},
			expected: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name:     "example 2",
			input:    []int{0, 1},
			expected: []int{1, 0},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "no adjacent equal",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all equal",
			input:    []int{2, 2, 2, 2},
			expected: []int{4, 4, 0, 0},
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "zeros at beginning",
			input:    []int{0, 0, 1, 2, 2},
			expected: []int{1, 4, 0, 0, 0},
		},
		{
			name:     "multiple adjacent pairs",
			input:    []int{3, 3, 5, 5, 1, 1},
			expected: []int{6, 10, 2, 0, 0, 0},
		},
	}

	funcsToTest := map[string]func([]int) []int{
		"applyOperations":                  applyOperations,
		"applyOperationsWithSlidingWindow": applyOperationsWithSlidingWindow,
		"applyOperationsOnePass":           applyOperationsOnePass,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// 复制输入以避免测试间相互影响
					inputCopy := make([]int, len(tc.input))
					copy(inputCopy, tc.input)
					result := fn(inputCopy)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
