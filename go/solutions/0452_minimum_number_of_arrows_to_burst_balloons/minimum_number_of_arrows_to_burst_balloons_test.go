package minimum_number_of_arrows_to_burst_balloons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "example 1",
			points:   [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			expected: 2,
		},
		{
			name:     "example 2",
			points:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			expected: 4,
		},
		{
			name:     "example 3",
			points:   [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			expected: 2,
		},
		{
			name:     "single balloon",
			points:   [][]int{{-5, 5}},
			expected: 1,
		},
		{
			name:     "nested balloons",
			points:   [][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			expected: 1,
		},
		{
			name:     "touching endpoints",
			points:   [][]int{{1, 2}, {2, 3}, {3, 4}},
			expected: 2,
		},
		{
			name:     "empty input",
			points:   [][]int{},
			expected: 0,
		},
	}

	funcsToTest := map[string]func([][]int) int{
		"findMinArrowShots": findMinArrowShots,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					points := clone2D(tc.points)
					result := fn(points)
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
