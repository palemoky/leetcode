package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		nums   []int
		target int
		want   int // 期望的索引，-1 表示未找到
	}{
		// --- 1. 基本情况 ---
		{
			name:   "Empty Array",
			nums:   []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "Single Element Array - Found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "Single Element Array - Not Found",
			nums:   []int{5},
			target: 3,
			want:   -1,
		},

		// --- 2. 典型情况 ---
		{
			name:   "Target in the middle (Odd length)",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 5,
			want:   3,
		},
		{
			name:   "Target in the left half",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 0,
			want:   1,
		},
		{
			name:   "Target in the right half",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 12,
			want:   5,
		},
		{
			name:   "Target in the middle (Even length)",
			nums:   []int{2, 5, 7, 8, 11, 12},
			target: 7, // or 8, depending on mid calculation
			want:   2,
		},

		// --- 3. 关键边界情况 ---
		{
			name:   "Target is the first element",
			nums:   []int{5, 7, 8, 9, 10},
			target: 5,
			want:   0,
		},
		{
			name:   "Target is the last element",
			nums:   []int{5, 7, 8, 9, 10},
			target: 10,
			want:   4,
		},
		{
			name:   "Target smaller than all elements",
			nums:   []int{5, 7, 8, 9, 10},
			target: 3,
			want:   -1,
		},
		{
			name:   "Target larger than all elements",
			nums:   []int{5, 7, 8, 9, 10},
			target: 11,
			want:   -1,
		},
		{
			name:   "Target not in array, but within range",
			nums:   []int{5, 7, 9, 10},
			target: 8,
			want:   -1,
		},
	}

	funcsToTest := map[string]func(nums []int, target int) int{
		"Iterative": searchIterative,
		"Recursive": searchRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.nums, tc.target)
					assert.Equal(t, tc.want, got, "Search failed for nums=%v, target=%d", tc.nums, tc.target)
				})
			}
		})
	}
}
