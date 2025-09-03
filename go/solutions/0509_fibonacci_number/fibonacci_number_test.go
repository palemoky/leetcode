package fibonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Base case: n = 0",
			input:    0,
			expected: 0,
		},
		{
			name:     "Base case: n = 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "Small number: n = 2",
			input:    2,
			expected: 1,
		},
		{
			name:     "Standard case: n = 5",
			input:    5,
			expected: 5,
		},
		{
			name:     "Larger number: n = 10",
			input:    10,
			expected: 55,
		},
		{
			name:     "Even larger number: n = 20",
			input:    20,
			expected: 6765,
		},
		// {
		// 	name:     "Edge case: negative input",
		// 	input:    -10,
		// 	expected: 0, // 假设我们定义负数输入的斐波那契数为0
		// },
	}

	funcsToTest := map[string]func(int) int{
		"Recursive": fibRecursive, // 朴素递归对于 n=20 会很慢，可以选择性地不测试它
		"Memoized":  fibRecursiveMemo,
		"DP":        fibDP,
		"Iterative": fibIterative,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					// 由于采用了全局变量记忆优化，导致无法使用并发测试
					result := fn(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
