package merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "example 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 4}},
			expected:  [][]int{{1, 4}},
		},
		{
			name:      "already disjoint",
			intervals: [][]int{{1, 2}, {4, 5}, {7, 9}},
			expected:  [][]int{{1, 2}, {4, 5}, {7, 9}},
		},
		{
			name:      "nested intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 8}},
			expected:  [][]int{{1, 10}},
		},
		{
			name:      "unsorted input",
			intervals: [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "empty input",
			intervals: [][]int{},
			expected:  [][]int{},
		},
	}

	funcsToTest := map[string]func([][]int) [][]int{
		"merge": merge,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					intervals := clone2D(tc.intervals)
					result := fn(intervals)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}

func clone2D(src [][]int) [][]int {
	if src == nil {
		return nil
	}

	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}

	return dst
}
