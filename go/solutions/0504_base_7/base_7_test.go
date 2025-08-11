package base7

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "Typical Positive Number",
			num:  100,
			want: "202", // 100 = 2*49 + 0*7 + 2*1
		},
		{
			name: "Typical Negative Number",
			num:  -7,
			want: "-10",
		},
		{
			name: "Input is Zero",
			num:  0,
			want: "0", // 核心边界：零
		},
		{
			name: "Input is exactly the base",
			num:  7,
			want: "10",
		},
		{
			name: "Input is negative base",
			num:  -7,
			want: "-10",
		},
		{
			name: "Input is less than the base (positive)",
			num:  6,
			want: "6",
		},
		{
			name: "Input is less than the base (negative)",
			num:  -6,
			want: "-6",
		},
		{
			name: "Input is a large number",
			num:  7 * 7 * 7 * 5, // 5 * 7^3 = 1715
			want: "5000",
		},
		{
			name: "Input is MaxInt32", // 测试大数，确保循环正常
			num:  math.MaxInt32,       // 2147483647
			want: "104134211161",      // (这个结果可以通过在线转换器验证)
		},
		{
			name: "Input is MinInt32", // 测试负向大数
			num:  math.MinInt32,       // -2147483648
			want: "-104134211162",
		},
	}

	funcsToTest := map[string]func(num int) string{
		"naive":     convertToBase7Naive,
		"recursive": convertToBase7Recursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.num)
					assert.Equal(t, tc.want, got, "input: %d", tc.num)
				})
			}
		})
	}
}
