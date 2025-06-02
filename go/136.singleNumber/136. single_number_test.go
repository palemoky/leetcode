package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"HashMap", singleNumberHashMap},
		{"BitWise", singleNumberBitWise},
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

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"HashMap", singleNumberHashMap},
		{"BitWise", singleNumberBitWise},
	}

	for _, algo := range algorithms {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				algo.fn(nums)
			}
		})
	}
}
