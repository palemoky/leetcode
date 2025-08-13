package power_of_three

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		num  int
		want bool
	}{
		// Positive Cases
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
			name: "Input is 9 (3^2)",
			num:  9,
			want: true,
		},
		{
			name: "Input is 27 (3^3)",
			num:  27,
			want: true,
		},
		{
			name: "Largest power of 3 within int32",
			num:  1162261467, // 3^19
			want: true,
		},

		// Negative Cases
		{
			name: "Input is 0",
			num:  0,
			want: false,
		},
		{
			name: "Input is 2",
			num:  2,
			want: false,
		},
		{
			name: "Input is 6 (divisible by 3 but not a power)",
			num:  6,
			want: false,
		},
		{
			name: "Input is 12",
			num:  12,
			want: false,
		},
		{
			name: "A large number that is not a power of 3",
			num:  1162261466,
			want: false,
		},

		// --- 3. 负数和边界 ---
		{
			name: "Input is a negative number",
			num:  -1,
			want: false,
		},
		{
			name: "Input is -3",
			num:  -3,
			want: false,
		},
		{
			name: "Input is MaxInt32",
			num:  math.MaxInt32,
			want: false,
		},
	}

	funcToTest := map[string]func(n int) bool{
		"Iterative": isPowerOfThreeIterative,
		"MathLog":   isPowerOfThreeLog,
		"MathMax":   isPowerOfThreeMax,
	}

	for funcName, isPowerOfThreeFunc := range funcToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					t.Run(tc.name, func(t *testing.T) {
						got := isPowerOfThreeFunc(tc.num)
						assert.Equal(t, tc.want, got, "Magic number solution failed for %d", tc.num)
					})
				})
			}
		})
	}
}
