package jump_game_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: multiple jumps",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "example 2: forced path",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "single element",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "two elements",
			nums:     []int{1, 0},
			expected: 1,
		},
		{
			name:     "already at end",
			nums:     []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "large jumps",
			nums:     []int{5, 1, 1, 1, 1},
			expected: 1,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"jump": jump,
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
