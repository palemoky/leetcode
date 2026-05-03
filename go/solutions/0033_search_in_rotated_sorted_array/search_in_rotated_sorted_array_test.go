package search_in_rotated_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "example 1: target in left rotated part",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "example 2: target not found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "example 3: array not rotated",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "target at beginning",
			nums:     []int{3, 1, 2},
			target:   3,
			expected: 0,
		},
		{
			name:     "target in right ordered part",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		{
			name:     "target at rotation point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   7,
			expected: 3,
		},
		{
			name:     "target at rotation point (right side)",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "single element array (found)",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element array (not found)",
			nums:     []int{1},
			target:   3,
			expected: -1,
		},
		{
			name:     "two elements (rotated)",
			nums:     []int{3, 1},
			target:   3,
			expected: 0,
		},
		{
			name:     "two elements (rotated, target in right)",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "ascending array (no rotation)",
			nums:     []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "ascending array (no rotation, not found)",
			nums:     []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "target smaller than all elements",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   -1,
			expected: -1,
		},
		{
			name:     "target larger than all elements",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   8,
			expected: -1,
		},
	}

	funcsToTest := map[string]func([]int, int) int{
		"search": search,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.nums, tc.target)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
