package leetcode

import (
	"testing"
)

func TestSingleNumberHashMap(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{2, 2, 1}, 1},
		{"Example 2", []int{4, 1, 2, 1, 2}, 4},
		{"Example 3", []int{1}, 1},
		{"Example 4", []int{}, 0},
	}

	t.Run("singleNumberHashMap", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := singleNumberHashMap(tc.nums)
				if got != tc.want {
					t.Errorf("singleNumberHashMap(%q) = %v; want %v", tc.nums, got, tc.want)
				}
			})
		}
	})
}

func BenchmarkSingleNumberHashMap(b *testing.B) {
	for b.Loop() {
		singleNumberHashMap([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5})
	}
}
