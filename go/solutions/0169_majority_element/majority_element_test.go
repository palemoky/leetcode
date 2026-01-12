package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTest = map[string]func([]int) int{
	"Sorting":    majorityElementSorting,
	"HashMap":    majorityElementHashMap,
	"BoyerMoore": majorityElementBoyerMoore,
}

func TestMajorityElement(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{"Simple majority", []int{3, 2, 3}, 3},
		{"All same", []int{1, 1, 1, 1}, 1},
		{"Majority at end", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"Single element", []int{1}, 1},
		{"Two elements same", []int{5, 5}, 5},
		{"Majority more than half", []int{6, 5, 5}, 5},
		{"Negative numbers", []int{-1, -1, 2}, -1},
		{"Zero as majority", []int{0, 0, 1}, 0},
		{"Large array", []int{1, 2, 1, 2, 1, 2, 1}, 1},
		{"Mixed values", []int{8, 8, 7, 7, 7, 8, 8}, 8},
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// Make a copy for sorting solution to avoid modifying original
					nums := make([]int, len(tc.nums))
					copy(nums, tc.nums)
					got := fn(nums)
					assert.Equal(t, tc.want, got, "%s: input=%v", fnName, tc.nums)
				})
			}
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := []int{1, 2, 1, 2, 1, 2, 1, 3, 1, 4, 1, 5, 1, 6, 1}

	for fnName, fn := range funcsToTest {
		b.Run(fnName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Make a copy for each iteration
				numsCopy := make([]int, len(nums))
				copy(numsCopy, nums)
				fn(numsCopy)
			}
		})
	}
}
