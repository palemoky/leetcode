package non_overlapping_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "example 1",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			expected:  1,
		},
		{
			name:      "example 2",
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			expected:  2,
		},
		{
			name:      "example 3",
			intervals: [][]int{{1, 2}, {2, 3}},
			expected:  0,
		},
		{
			name:      "empty input",
			intervals: [][]int{},
			expected:  0,
		},
		{
			name:      "single interval",
			intervals: [][]int{{-10, -1}},
			expected:  0,
		},
		{
			name:      "fully nested intervals",
			intervals: [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			expected:  3,
		},
		{
			name:      "unsorted mixed intervals",
			intervals: [][]int{{5, 7}, {1, 3}, {2, 4}, {6, 8}},
			expected:  2,
		},
	}

	funcsToTest := map[string]func([][]int) int{
		"eraseOverlapIntervals": eraseOverlapIntervals,
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
