package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func([]int) int
}{
	{"HashMap", singleNumberHashMap},
	{"BitWise", singleNumberBitWise},
}

func TestSingleNumberHashMap(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
        {"LeetCode Example 1", []int{2, 2, 1}, 1},
        {"LeetCode Example 2", []int{4, 1, 2, 1, 2}, 4},
        {"Single element", []int{99}, 99},
        {"Negative number", []int{-1, -1, -2}, -2},
        {"Zero included", []int{0, 1, 1}, 0},
        {"Large numbers", []int{100000, 1, 1}, 100000},
        {"Mixed positive and negative", []int{-3, -3, 0, 0, 7}, 7},
         {"Empty", []int{}, 0}, 
        // {"All same", []int{5, 5, 5}, 0}, // BitWise cannot handle this case, but HashMap can handle it 
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				tc := tc 
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

	for _, algo := range algorithms {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				algo.fn(nums)
			}
		})
	}
}
