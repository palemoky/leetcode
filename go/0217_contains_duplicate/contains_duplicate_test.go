package solution

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
		{"Empty array", []int{}, false},
		{"Single element", []int{0}, false},
		{"Two same", []int{1, 1}, true},
		{"Two different", []int{1, 2}, false},
		{"Duplicate at start", []int{2, 2, 3, 4}, true},
		{"Duplicate at end", []int{1, 2, 3, 3}, true},
		{"All unique", []int{1, 2, 3, 4, 5}, false},
		{"All same", []int{7, 7, 7, 7}, true},
		{"Negative numbers", []int{-1, -2, -3, -1}, true},
		{"Mixed positive and negative", []int{-1, 2, -3, 2}, true},
		{"Large numbers", []int{1, 2, 3, 1 << 30, 1 << 30}, true},
		{"No duplicate, large", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				tc := tc // Capture range variable to avoid data race
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
