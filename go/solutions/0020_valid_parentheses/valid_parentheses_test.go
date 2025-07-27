package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithms = []struct {
	name string
	fn   func(string) bool
}{
	{"IfElse", isValidIfElse},
	{"SwitchCase", isValidSwitchCase},
}

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name   string
		braces string
		want   bool
	}{
		{"Empty string", "", true},
		{"Single left", "(", false},
		{"Single right", ")", false},
		{"Single pair", "()", true},
		{"All types", "()[]{}", true},
		{"Nested", "{[()]}", true},
		{"Wrong order", "(]", false},
		{"Wrong nested", "([)]", false},
		{"Odd length", "(()", false},
		{"Only lefts", "(((", false},
		{"Only rights", ")))", false},
		{"Starts with right", ")()", false},
		{"Ends with left", "(()", false},
		{"Multiple valid", "()(){}", true},
		{"Multiple invalid", "(()))(", false},
		{"Deep nested valid", "{[({[()]})]}", true},
		{"Deep nested invalid", "{[({[()]}]}", false},
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

func BenchmarkIsValid(b *testing.B) {
	braces := "{[()]}[]{}({[]})"
	for _, algo := range algorithms {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				algo.fn(braces)
			}
		})
	}
}
