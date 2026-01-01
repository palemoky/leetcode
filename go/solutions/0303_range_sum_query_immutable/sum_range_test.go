package range_sum_query_immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		nums       []int
		operations []struct {
			left  int
			right int
			want  int
		}
	}{
		{
			name: "Example 1",
			nums: []int{-2, 0, 3, -5, 2, -1},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 2, 1},  // sum([-2, 0, 3]) = 1
				{2, 5, -1}, // sum([3, -5, 2, -1]) = -1
				{0, 5, -3}, // sum([-2, 0, 3, -5, 2, -1]) = -3
			},
		},
		{
			name: "Single element",
			nums: []int{5},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 0, 5}, // sum([5]) = 5
			},
		},
		{
			name: "All positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 4, 15}, // sum([1, 2, 3, 4, 5]) = 15
				{1, 3, 9},  // sum([2, 3, 4]) = 9
				{0, 0, 1},  // sum([1]) = 1
				{4, 4, 5},  // sum([5]) = 5
			},
		},
		{
			name: "All negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 4, -15}, // sum([-1, -2, -3, -4, -5]) = -15
				{1, 3, -9},  // sum([-2, -3, -4]) = -9
			},
		},
		{
			name: "With zeros",
			nums: []int{0, 0, 0, 0, 0},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 4, 0}, // sum([0, 0, 0, 0, 0]) = 0
				{2, 3, 0}, // sum([0, 0]) = 0
			},
		},
		{
			name: "Mixed values",
			nums: []int{-2, 5, -1, 7, -3, 4},
			operations: []struct {
				left  int
				right int
				want  int
			}{
				{0, 1, 3},  // sum([-2, 5]) = 3
				{1, 4, 8},  // sum([5, -1, 7, -3]) = 8
				{2, 5, 7},  // sum([-1, 7, -3, 4]) = 7
				{0, 5, 10}, // sum([-2, 5, -1, 7, -3, 4]) = 10
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			numArray := Constructor(tc.nums)
			for i, op := range tc.operations {
				got := numArray.SumRange(op.left, op.right)
				assert.Equal(t, op.want, got,
					"Operation %d: SumRange(%d, %d) failed, nums=%v",
					i+1, op.left, op.right, tc.nums)
			}
		})
	}
}

func BenchmarkNumArray(b *testing.B) {
	nums := []int{-2, 0, 3, -5, 2, -1, 4, 7, -3, 8, 1, -6, 5, 9, -4, 2}
	numArray := Constructor(nums)

	b.Run("SumRange", func(b *testing.B) {
		for b.Loop() {
			numArray.SumRange(0, 5)
			numArray.SumRange(3, 10)
			numArray.SumRange(7, 15)
		}
	})

	b.Run("Constructor", func(b *testing.B) {
		for b.Loop() {
			Constructor(nums)
		}
	})
}
