package kth_largest_element_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"distinct", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"leetcode_example", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"single", []int{1}, 1, 1},
		{"duplicates", []int{2, 1, 2}, 2, 2},
		{"k_equals_len", []int{2, 1}, 2, 1},
		{"negatives", []int{-1, -1, -2, -3}, 1, -1},
	}

	funcsToTest := map[string]func(nums []int, k int) int{
		"SortedArray": findKthLargestSortedArray,
		"Heap":        findKthLargestHeap,
		"QuickSelect": findKthLargestQuickSelect,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(append([]int{}, tc.nums...), tc.k)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
