package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func([]int, int) []int{
	"BruteForce": twoSumBruteForce,
	"HashMap":    twoSumHashMap,
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"LeetCode Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"LeetCode Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"LeetCode Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"No solution", []int{1, 2, 3}, 7, []int{}},
		{"Empty array", []int{}, 0, []int{}},
		{"Single element", []int{1}, 1, []int{}},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"Zero target", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"Multiple pairs", []int{1, 2, 3, 4, 4}, 8, []int{3, 4}},
		{"Duplicate numbers", []int{1, 5, 1, 5}, 10, []int{1, 3}},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := fn(tc.nums, tc.target)
					assert.ElementsMatch(t, tc.want, got, "%s: input=%v, target=%d", fnName, tc.nums, tc.target)
				})
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 1, 8, 3, 6, 4, 5, 9, 10}
	target := 19
	for fnName, fn := range funcsToTest {
		b.Run(fnName, func(b *testing.B) {
			for b.Loop() {
				fn(nums, target)
			}
		})
	}
}
