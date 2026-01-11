package remove_duplicates_from_sorted_array_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		input         []int
		expected      int
		expectedArray []int
	}{
		{
			name:          "example 1: [1,1,1,2,2,3]",
			input:         []int{1, 1, 1, 2, 2, 3},
			expected:      5,
			expectedArray: []int{1, 1, 2, 2, 3},
		},
		{
			name:          "example 2: [0,0,1,1,1,1,2,3,3]",
			input:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected:      7,
			expectedArray: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:          "empty array",
			input:         []int{},
			expected:      0,
			expectedArray: []int{},
		},
		{
			name:          "single element",
			input:         []int{1},
			expected:      1,
			expectedArray: []int{1},
		},
		{
			name:          "two elements same",
			input:         []int{1, 1},
			expected:      2,
			expectedArray: []int{1, 1},
		},
		{
			name:          "two elements different",
			input:         []int{1, 2},
			expected:      2,
			expectedArray: []int{1, 2},
		},
		{
			name:          "all same elements",
			input:         []int{1, 1, 1, 1, 1},
			expected:      2,
			expectedArray: []int{1, 1},
		},
		{
			name:          "no duplicates",
			input:         []int{1, 2, 3, 4, 5},
			expected:      5,
			expectedArray: []int{1, 2, 3, 4, 5},
		},
		{
			name:          "three occurrences at end",
			input:         []int{1, 2, 2, 2},
			expected:      3,
			expectedArray: []int{1, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// 复制输入数组，因为函数会修改原数组
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			result := removeDuplicates(inputCopy)
			assert.Equal(t, tc.expected, result, "返回的长度不正确")

			// 验证前 result 个元素是否正确
			if result > 0 {
				assert.Equal(t, tc.expectedArray, inputCopy[:result], "修改后的数组内容不正确")
			}
		})
	}
}
