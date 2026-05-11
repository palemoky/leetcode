package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "strictly increasing",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "strictly decreasing",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6,
		},
		{
			name:     "all equal heights",
			height:   []int{3, 3, 3, 3},
			expected: 9,
		},
		{
			name:     "contains zeros",
			height:   []int{0, 2, 0, 4, 0},
			expected: 4,
		},
		{
			name:     "empty input",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "single bar",
			height:   []int{10},
			expected: 0,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"two pointers": maxArea,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.height)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
