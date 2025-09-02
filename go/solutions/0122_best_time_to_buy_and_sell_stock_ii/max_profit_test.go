package best_time_to_buy_and_sell_stock_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitImplementations(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		prices []int
		want   int
	}{
		{"multiple_buy_sell_opportunities", []int{7, 1, 5, 3, 6, 4}, 7},
		{"continuous_increasing", []int{1, 2, 3, 4, 5}, 4},
		{"continuous_decreasing", []int{7, 6, 4, 3, 1}, 0},
		{"single", []int{1}, 0},
		{"two", []int{2, 1}, 0},
		{"increase", []int{1, 2}, 1},
	}

	funcs := map[string]func(prices []int) int{
		"Greedy":    maxProfitNaive,
		"Optimized": maxProfitOptimized,
	}

	for fnName, fn := range funcs {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					in := append([]int{}, tc.prices...)
					got := fn(in)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
