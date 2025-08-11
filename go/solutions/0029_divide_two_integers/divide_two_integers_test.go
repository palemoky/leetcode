package divide_two_integers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{
			name:     "OverflowCase", // 必须处理的溢出
			dividend: math.MinInt32,
			divisor:  -1,
			want:     math.MaxInt32,
		},
		{
			name:     "MinInt32Case", // 边界值
			dividend: math.MinInt32,
			divisor:  1,
			want:     math.MinInt32,
		},
		{
			name:     "NormalPositive", // 基本情况
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "NormalNegative", // 符号处理
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},
		{
			name:     "ZeroDividend", // 被除数为0
			dividend: 0,
			divisor:  1,
			want:     0,
		},
		{
			name:     "DivisorIsOne", // 测试朴素算法的超时点
			dividend: 100,
			divisor:  1,
			want:     100,
		},
	}

	funcsToTest := map[string]func(dividend, divisor int) int{
		"divide": divide,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.dividend, tc.divisor)
					assert.Equal(t, tc.want, got, "input: %s * %s", tc.dividend, tc.divisor)
				})
			}
		})
	}
}
