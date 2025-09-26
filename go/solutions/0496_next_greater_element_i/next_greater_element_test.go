package next_greater_element_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "example 1",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			name:  "example 2",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		{
			name:  "no greater element",
			nums1: []int{3, 2, 1},
			nums2: []int{3, 2, 1},
			want:  []int{-1, -1, -1},
		},
		{
			name:  "all have greater element",
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4},
		},
		{
			name:  "single element",
			nums1: []int{1},
			nums2: []int{1, 2},
			want:  []int{2},
		},
		{
			name:  "single element no greater",
			nums1: []int{2},
			nums2: []int{1, 2},
			want:  []int{-1},
		},
		{
			name:  "empty nums1",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  []int{},
		},
		{
			name:  "descending order",
			nums1: []int{5, 4, 3},
			nums2: []int{5, 4, 3, 2, 1},
			want:  []int{-1, -1, -1},
		},
		{
			name:  "ascending order",
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 3, 4, 5},
			want:  []int{2, 3, 4},
		},
	}

	funcToTest := map[string]func(nums1, nums2 []int) []int{
		"LTR": nextGreaterElementStackLeftToRight,
		"RTL": nextGreaterElementStackRightToLeft,
	}

	for fnName, fn := range funcToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					t.Run(tc.name, func(t *testing.T) {
						got := fn(tc.nums1, tc.nums2)
						assert.Equal(t, tc.want, got)
					})
				})
			}
		})
	}
}
