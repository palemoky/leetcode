package longest_consecutive_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "empty input",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: 1,
		},
		{
			name:     "all duplicates",
			nums:     []int{5, 5, 5, 5},
			expected: 1,
		},
		{
			name:     "contains negatives",
			nums:     []int{-1, -2, -3, 10, 11},
			expected: 3,
		},
		{
			name:     "multiple runs choose longest",
			nums:     []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
			expected: 7,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"longestConsecutive": longestConsecutive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.nums)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
