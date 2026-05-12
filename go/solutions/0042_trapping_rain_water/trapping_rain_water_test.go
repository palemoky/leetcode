package trapping_rain_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "example 1",
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		{
			name:  "example 2",
			input: []int{4, 2, 0, 3, 2, 5},
			want:  9,
		},
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "single bar",
			input: []int{5},
			want:  0,
		},
		{
			name:  "two bars",
			input: []int{5, 1},
			want:  0,
		},
		{
			name:  "all same height",
			input: []int{3, 3, 3, 3},
			want:  0,
		},
		{
			name:  "strictly increasing",
			input: []int{1, 2, 3, 4},
			want:  0,
		},
		{
			name:  "strictly decreasing",
			input: []int{4, 3, 2, 1},
			want:  0,
		},
		{
			name:  "simple valley",
			input: []int{2, 0, 2},
			want:  2,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"monotonic stack": trapMonotonicStack,
		"two pointers":    trapTwoPointers,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.input)
					assert.Equal(t, tc.want, result)
				})
			}
		})
	}
}
