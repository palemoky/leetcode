package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingAlgorithms(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Typical unsorted",
			nums: []int{5, 2, 9, 1, 5, 6},
			want: []int{1, 2, 5, 5, 6, 9},
		},
		{
			name: "Already sorted",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse sorted",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "All equal",
			nums: []int{7, 7, 7, 7},
			want: []int{7, 7, 7, 7},
		},
		{
			name: "Empty",
			nums: []int{},
			want: []int{},
		},
		{
			name: "Single element",
			nums: []int{42},
			want: []int{42},
		},
	}

	funcsToTest := map[string]func([]int) []int{
		"Bubble":    bubble,
		"Selection": selection,
		"Insertion": insertion,
		"Shell":     shell,
		"Quick":     quick,
		"Merge":     merge,
		"Heap":      heapSorting,
		"Counting":  counting,
		"Bucket":    bucket,
		"Radix":     radix,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// 传入副本，避免原数组被修改影响其他测试
					input := append([]int(nil), tc.nums...)
					got := fn(input)
					assert.Equal(t, tc.want, got, "algorithm=%s, input=%v", fnName, tc.nums)
				})
			}
		})
	}
}
