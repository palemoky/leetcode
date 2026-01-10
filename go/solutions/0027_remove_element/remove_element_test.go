package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		nums     []int
		val      int
		expected []int
		length   int
	}{
		{
			name:     "example 1: remove 3",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
			length:   2,
		},
		{
			name:     "example 2: remove 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 3, 0, 4},
			length:   5,
		},
		{
			name:     "no elements to remove",
			nums:     []int{1, 2, 3, 4, 5},
			val:      6,
			expected: []int{1, 2, 3, 4, 5},
			length:   5,
		},
		{
			name:     "all elements to remove",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			expected: []int{},
			length:   0,
		},
		{
			name:     "single element - remove",
			nums:     []int{1},
			val:      1,
			expected: []int{},
			length:   0,
		},
		{
			name:     "single element - keep",
			nums:     []int{1},
			val:      2,
			expected: []int{1},
			length:   1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			val:      1,
			expected: []int{},
			length:   0,
		},
		{
			name:     "remove first element",
			nums:     []int{1, 2, 3, 4},
			val:      1,
			expected: []int{2, 3, 4},
			length:   3,
		},
		{
			name:     "remove last element",
			nums:     []int{1, 2, 3, 4},
			val:      4,
			expected: []int{1, 2, 3},
			length:   3,
		},
		{
			name:     "remove multiple occurrences",
			nums:     []int{2, 2, 2, 2, 2},
			val:      2,
			expected: []int{},
			length:   0,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -2, -4},
			val:      -2,
			expected: []int{-1, -3, -4},
			length:   3,
		},
	}

	funcsToTest := map[string]func([]int, int) int{
		"removeElementTwoPointers": removeElementTwoPointers,
		"removeElementSwap":        removeElementSwap,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					// 复制数组，因为函数会修改原数组
					numsCopy := append([]int{}, tc.nums...)

					length := fn(numsCopy, tc.val)

					// 验证返回的长度
					assert.Equal(t, tc.length, length, "returned length should match expected")

					// 验证前 length 个元素（不关心顺序，因为 Swap 方法不保证顺序）
					if fnName == "removeElementTwoPointers" {
						// TwoPointers 方法保持顺序
						assert.Equal(t, tc.expected, numsCopy[:length], "elements should match expected")
					} else {
						// Swap 方法不保证顺序，只验证元素集合
						assert.ElementsMatch(t, tc.expected, numsCopy[:length], "elements should match expected (order may differ)")
					}

					// 验证结果中不包含 val
					for i := range length {
						assert.NotEqual(t, tc.val, numsCopy[i], "result should not contain val")
					}
				})
			}
		})
	}
}
