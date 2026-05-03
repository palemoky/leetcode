package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "zero", input: 0, expected: 0},
		{name: "one", input: 1, expected: 1},
		{name: "perfect square", input: 4, expected: 2},
		{name: "perfect square", input: 9, expected: 3},
		{name: "truncated", input: 8, expected: 2},
		{name: "large perfect square", input: 2147395600, expected: 46340},
		{name: "max int", input: 2147483647, expected: 46340},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, mySqrt(tc.input))
		})
	}
}
