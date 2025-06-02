package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func([]int) bool
}{
	{"HashMap", containsDuplicateHashMap},
}

func TestContainsDuplicateHashMap(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, true},
		{"Example 2", []int{1, 2, 3, 4}, false},
		{"Example 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"Example 4", []int{}, false},
		{"Example 5", []int{0}, false},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := algo.fn(tc.nums)
					assert.Equal(t, tc.want, got, "%s: input=%v", algo.name, tc.nums)
				})
			}
		})
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	for _, algo := range algorithms {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				algo.fn(nums)
			}
		})
	}
}
