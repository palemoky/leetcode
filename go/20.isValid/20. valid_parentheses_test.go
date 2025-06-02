package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name   string
		braces string
		want   bool
	}{
		{"Example 1", "", true},
		{"Example 2", "{}", true},
		{"Example 3", "[]", true},
		{"Example 4", "()", true},
		{"Example 5", "{[()]}", true},
		{"Example 6", "(", false},
		{"Example 7", "{([)]}", false},
		{"Example 8", "{([)]", false},
	}

	algorithms := []struct {
		name string
		fn   func(string) bool
	}{
		{"IfElse", isValidIfElse},
		{"SwitchCase", isValidSwitchCase},
	}

	for _, algo := range algorithms {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := algo.fn(tc.braces)
					assert.Equal(t, tc.want, got, "%s: input=%v", algo.name, tc.braces)
				})
			}
		})
	}
}

func BenchmarkIsValidIfElse(b *testing.B) {
	for b.Loop() {
		isValidIfElse("{[()]}[]{}({[]})")
	}
}

func BenchmarkIsValidSwitchCase(b *testing.B) {
	for b.Loop() {
		isValidSwitchCase("{[()]}[]{}({[]})")
	}
}
