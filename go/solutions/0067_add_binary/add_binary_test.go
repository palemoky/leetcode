package add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "LeetCode Example 1",
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			name: "LeetCode Example 2",
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
		{
			name: "Addition without carry",
			a:    "101",
			b:    "010",
			want: "111",
		},
		{
			name: "Unequal lengths (a is longer)",
			a:    "1110",
			b:    "1",
			want: "1111",
		},
		{
			name: "Unequal lengths (b is longer)",
			a:    "10",
			b:    "1101",
			want: "1111",
		},
		{
			name: "Result has a final carry",
			a:    "1",
			b:    "111",
			want: "1000",
		},
		{
			name: "Both inputs are zero",
			a:    "0",
			b:    "0",
			want: "0",
		},
		{
			name: "One input is zero",
			a:    "11011",
			b:    "0",
			want: "11011",
		},
	}

	funcsToTest := map[string]func(a, b string) string{
		"Simulation": addBinary,
		"BigInt":     addBinaryWithBigInt,
	}

	for funcName, addFunc := range funcsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					got := addFunc(tc.a, tc.b)
					assert.Equal(t, tc.want, got, "addBinary(%q, %q) ailed", tc.a, tc.b)
				})
			}
		})
	}
}
