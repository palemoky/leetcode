package leetcode

import (
	"testing"
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

	t.Run("isValid", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := isValidIfElse(tc.braces)
				if got != tc.want {
					t.Errorf("isValid(%q) = %v; want %v", tc.braces, got, tc.want)
				}
			})
		}
	})

		t.Run("isValid", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := isValidSwitchCase(tc.braces)
				if got != tc.want {
					t.Errorf("isValid(%q) = %v; want %v", tc.braces, got, tc.want)
				}
			})
		}
	})
}

func BenchmarkIsValidIfElse(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValidIfElse("{[()]}[]{}({[]})")
    }
}

func BenchmarkIsValidSwitchCase(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isValidSwitchCase("{[()]}[]{}({[]})")
    }
}