package minimum_size_subarray_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "示例1：基本情况",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2, // [4,3]
		},
		{
			name:     "示例2：单个元素满足",
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1, // [4]
		},
		{
			name:     "示例3：无解",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "单个元素数组-满足",
			target:   5,
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "单个元素数组-不满足",
			target:   10,
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "整个数组才满足",
			target:   15,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "所有元素都很大",
			target:   10,
			nums:     []int{10, 20, 30},
			expected: 1,
		},
		{
			name:     "连续的小元素",
			target:   100,
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: 10, // [10,11,12,13,14,15] 或其他组合
		},
	}

	funcsToTest := map[string]func(int, []int) int{
		"minSubArrayLen": minSubArrayLen,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.target, tc.nums)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
