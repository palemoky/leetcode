package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "example 1",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "example 2",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "no zeros",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "zeros at beginning",
			input:    []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "zeros at end",
			input:    []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "alternating zeros",
			input:    []int{1, 0, 2, 0, 3, 0, 4},
			expected: []int{1, 2, 3, 4, 0, 0, 0},
		},
	}

	funcsToTest := map[string]func([]int){
		"moveZeroes": moveZeroes,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// Make a copy since the function modifies in-place
					nums := make([]int, len(tc.input))
					copy(nums, tc.input)
					fn(nums)
					assert.Equal(t, tc.expected, nums)
				})
			}
		})
	}
}
