package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Target exists",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "Insert at beginning",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "Insert in middle",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "Insert at end",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "Empty array",
			nums:   []int{},
			target: 1,
			want:   0,
		},
		{
			name:   "Single element, insert before",
			nums:   []int{3},
			target: 2,
			want:   0,
		},
		{
			name:   "Single element, insert after",
			nums:   []int{3},
			target: 4,
			want:   1,
		},
		{
			name:   "Single element, target exists",
			nums:   []int{3},
			target: 3,
			want:   0,
		},
	}

	funcsToTest := map[string]func(nums []int, target int) int{
		"Naive":      searchInsert,
		"LowerBound": searchInsertLowerBound,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(append([]int{}, tc.nums...), tc.target)
					assert.Equal(t, tc.want, got, "nums=%v, target=%d", tc.nums, tc.target)
				})
			}
		})
	}
}
