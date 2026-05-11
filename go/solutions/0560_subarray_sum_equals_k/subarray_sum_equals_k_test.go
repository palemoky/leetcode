package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			name:     "contains negative numbers",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0},
			k:        0,
			expected: 6,
		},
		{
			name:     "single element equals k",
			nums:     []int{5},
			k:        5,
			expected: 1,
		},
		{
			name:     "single element not equals k",
			nums:     []int{5},
			k:        2,
			expected: 0,
		},
		{
			name:     "empty input",
			nums:     []int{},
			k:        0,
			expected: 0,
		},
	}

	funcsToTest := map[string]func([]int, int) int{
		"Brute Force": subarraySumBruteForce,
		"Prefix Sum":  subarraySumPrefixSum,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.nums, tc.k)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
