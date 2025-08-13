package power_of_two

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
			name: "Input is 1 (2^0)",
			num:  1,
			want: true,
		},
		{
			name: "Input is 2 (2^1)",
			num:  2,
			want: true,
		},
		{
			name: "Input is 4 (2^2)",
			num:  4,
			want: true,
		},
		{
			name: "Input is 8 (2^10)",
			num:  1024,
			want: true,
		},
		{
			name: "Largest power of 2 within int32",
			num:  1073741824, // 2^30
			want: true,
		},

		// Negative Cases
		{
			name: "Input is 0",
			num:  0,
			want: false,
		},
		{
			name: "Input is 3",
			num:  3,
			want: false,
		},
		{
			name: "Input is 6 (divisible by 2 but not a power)",
			num:  6,
			want: false,
		},
		{
			name: "Input is 12",
			num:  12,
			want: false,
		},
		{
			name: "A large number that is not a power of 2",
			num:  2147483646,
			want: false,
		},

		// --- 3. 负数和边界 ---
		{
			name: "Input is a negative number",
			num:  -1,
			want: false,
		},
		{
			name: "Input is -2",
			num:  -2,
			want: false,
		},
		{
			name: "Input is MaxInt32",
			num:  math.MaxInt32,
			want: false,
		},
	}

	funcToTest := map[string]func(n int) bool{
		"Iterative": isPowerOfTwoIterative,
		"MathLog":   isPowerOfTwoLog,
		"MathMax":   isPowerOfTwoMax,
		"Binary":    isPowerOfTwoBinary,
	}

	for fnName, fn := range funcToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					t.Run(tc.name, func(t *testing.T) {
						got := fn(tc.num)
						assert.Equal(t, tc.want, got, "Magic number solution failed for %d", tc.num)
					})
				})
			}
		})
	}
}
