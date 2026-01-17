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
		{"Base case: n = 0", 0, 0},
		{"Base case: n = 1", 1, 1},
		{"Small number: n = 2", 2, 1},
		{"Standard case: n = 5", 5, 5},
		{"Larger number: n = 10", 10, 55},
		{"Even larger number: n = 20", 20, 6765},
	}

	funcsToTest := map[string]func(int) int{
		"Recursive": fibRecursive,
		"Memoized":  fibRecursiveMemo,
		"DP":        fibDP,
		"Iterative": fibIterative,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
