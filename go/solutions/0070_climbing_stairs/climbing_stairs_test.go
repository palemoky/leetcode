package climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Base case: n = 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "Base case: n = 2",
			input:    2,
			expected: 2,
		},
		{
			name:     "Standard case: n = 3",
			input:    3,
			expected: 3,
		},
		{
			name:     "Standard case: n = 5",
			input:    5,
			expected: 8,
		},
		{
			name:     "Larger number: n = 10",
			input:    10,
			expected: 89,
		},
		{
			name:     "LeetCode max constraint: n = 45",
			input:    45,
			expected: 1836311903,
		},
		{
			name:     "Edge case: n = 0",
			input:    0,
			expected: 0, // 根据我们的函数定义，0个台阶有0种方法
		},
		{
			name:     "Edge case: negative input",
			input:    -5,
			expected: 0, // 无效输入应该返回0
		},
	}

	funcsToTest := map[string]func(int) int{
		"Iterative": climbStairsIterative,
		"Memoized":  climbStairsRecursiveMemo,
		"Recursive": climbStairsRecursive, // 朴素递归对于 n=45 会超时，通常不测试它
	}

	for funcName, climbFunc := range funcsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := climbFunc(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
