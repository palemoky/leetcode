package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "example 1: normal case",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "example 2: nums1 has one element",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "example 3: nums1 is empty",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "nums2 is empty",
			nums1:    []int{1, 2, 3},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "all nums2 elements are smaller",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "all nums2 elements are larger",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "interleaved elements",
			nums1:    []int{1, 3, 5, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 4, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "duplicate elements",
			nums1:    []int{1, 2, 2, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 2, 3},
			n:        3,
			expected: []int{1, 2, 2, 2, 2, 3},
		},
		{
			name:     "negative numbers",
			nums1:    []int{-3, -1, 0, 0, 0},
			m:        2,
			nums2:    []int{-2, 0, 1},
			n:        3,
			expected: []int{-3, -2, -1, 0, 1},
		},
	}

	funcsToTest := map[string]func([]int, int, []int, int){
		"merge": merge,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					fn(tc.nums1, tc.m, tc.nums2, tc.n)
					assert.Equal(t, tc.expected, tc.nums1)
				})
			}
		})
	}
}
