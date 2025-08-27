package three_sum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// sort2DSlice 对二维切片进行排序，以便于比较。
// 首先按行内元素排序，然后按行排序。
func sort2DSlice(slice [][]int) {
	for _, s := range slice {
		sort.Ints(s)
	}
	sort.Slice(slice, func(i, j int) bool {
		for x := 0; x < len(slice[i]) && x < len(slice[j]); x++ {
			if slice[i][x] != slice[j][x] {
				return slice[i][x] < slice[j][x]
			}
		}
		return len(slice[i]) < len(slice[j])
	})
}

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want [][]int
	}{
		// --- 1. 基本情况 ---
		{
			name: "Empty Array",
			nums: []int{},
			want: nil, // 或者 [][]int{}
		},
		{
			name: "Array with less than 3 elements",
			nums: []int{1, 2},
			want: nil,
		},
		{
			name: "No solution found",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "All positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "All negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			want: [][]int{},
		},

		// --- 2. 典型情况 ---
		{
			name: "Typical case with negatives and duplicates",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},

		// --- 3. 去重逻辑的考验 ---
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Multiple zeros with one valid triplet",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Duplicates need to be skipped",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "Complex case with duplicates",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
	}

	funcsToTest := map[string]func(nums []int) [][]int{
		"BruteForce":  threeSumBruteForce,
		"TwoPointers": threeSumTwoPointers,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.nums)

					// 为了可靠地比较两个二维切片，我们需要先对它们进行排序
					sort2DSlice(got)
					sort2DSlice(tc.want)

					assert.Equal(t, tc.want, got, "threeSum(%v) failed", tc.nums)
				})
			}
		})
	}
}
