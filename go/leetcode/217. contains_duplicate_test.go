package leetcode

import (
	"testing"
)

func TestContainsDuplicateHashMap(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{"Example 1", []int{1,2,3,1}, true},
		{"Example 2", []int{1,2,3,4}, false},
		{"Example 3", []int{1,1,1,3,3,4,3,2,4,2}, true},
		{"Example 4", []int{}, false},
		{"Example 5", []int{0}, false},
	}

	t.Run("containsDuplicateHashMap", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := containsDuplicateHashMap(tc.nums)
				if got != tc.want {
					t.Errorf("containsDuplicateHashMap(%q) = %v; want %v", tc.nums, got, tc.want)
				}
			})
		}
	})
}

func BenchmarkContainsDuplicateHashMap(b *testing.B) {
	for b.Loop() {
		containsDuplicateHashMap([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5})
	}
}
