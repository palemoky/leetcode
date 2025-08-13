package check_powers_of_three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPowersOfThree(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want bool
	}{
		// True Cases
		{
			name: "Input is 1 (3^0)",
			num:  1,
			want: true,
		},
		{
			name: "Input is 3 (3^1)",
			num:  3,
			want: true,
		},
		{
			name: "Input is 4 (3^1 + 3^0)",
			num:  4,
			want: true,
		},
		{
			name: "Input is 10 (3^2 + 3^0)",
			num:  10,
			want: true,
		},
		{
			name: "Input is 12 (3^2 + 3^1)",
			num:  12,
			want: true,
		},
		{
			name: "Input is 13 (3^2 + 3^1 + 3^0)",
			num:  13,
			want: true,
		},
		{
			name: "A larger valid number",
			num:  91, // 3^4 + 3^2 + 3^0, trinary: 10101
			want: true,
		},

		// False Cases
		{
			name: "Input is 2 (smallest failing number)",
			num:  2, // trinary: 2
			want: false,
		},
		{
			name: "Input is 5",
			num:  5, // trinary: 12
			want: false,
		},
		{
			name: "Input is 6",
			num:  6, // trinary: 20
			want: false,
		},
		{
			name: "Input is 7",
			num:  7, // trinary: 21
			want: false,
		},
		{
			name: "Input is 21",
			num:  21, // trinary: 210
			want: false,
		},

		// Edge Cases
		{
			name: "Input is 0 (out of problem scope, but should be false)",
			num:  0,
			want: false,
		},
		{
			name: "Input is a negative number",
			num:  -10,
			want: false,
		},
		{
			name: "Max constraint from LeetCode (10^7)",
			num:  10000000,
			want: false, // Its trinary representation contains '2'
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := checkPowersOfThree(tc.num)
			assert.Equal(t, tc.want, got, "checkPowersOfThree(%d) failed", tc.num)
		})
	}
}
