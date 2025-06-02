package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"Example 4", []int{3, 3}, 5, []int{}},
	}

	algorithms := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"BruteForce", twoSumBruteForce},
		{"HashMap", twoSumHashMap},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := algo.fn(tc.nums, tc.target)
					assert.ElementsMatch(t, tc.want, got, "%s: input=%v, target=%d", algo.name, tc.nums, tc.target)
				})
			}
		})
	}
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	nums := []int{2, 7, 11, 15, 1, 8, 3, 6, 4, 5, 9, 10}
	target := 19
	for b.Loop() {
		twoSumBruteForce(nums, target)
	}
}

func BenchmarkTwoSumHashMap(b *testing.B) {
	nums := []int{2, 7, 11, 15, 1, 8, 3, 6, 4, 5, 9, 10}
	target := 19
	for b.Loop() {
		twoSumHashMap(nums, target)
	}
}
