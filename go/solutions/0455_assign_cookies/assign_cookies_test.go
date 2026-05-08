package assign_cookies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		g        []int
		s        []int
		expected int
	}{
		{
			name:     "example 1",
			g:        []int{1, 2, 3},
			s:        []int{1, 1},
			expected: 1,
		},
		{
			name:     "example 2",
			g:        []int{1, 2},
			s:        []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "no cookies",
			g:        []int{1, 2, 3},
			s:        []int{},
			expected: 0,
		},
		{
			name:     "no children",
			g:        []int{},
			s:        []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "all cookies too small",
			g:        []int{5, 10},
			s:        []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "all children satisfied",
			g:        []int{1, 2, 3},
			s:        []int{3, 3, 3},
			expected: 3,
		},
	}

	funcsToTest := map[string]func([]int, []int) int{
		"findContentChildren": findContentChildren,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.g, tc.s)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
