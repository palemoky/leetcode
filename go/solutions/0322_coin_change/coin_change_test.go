package coin_change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "Example 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3, // 5 + 5 + 1
		},
		{
			name:   "Example 2",
			coins:  []int{2},
			amount: 3,
			want:   -1, // 无法凑成
		},
		{
			name:   "Example 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "Greedy fails",
			coins:  []int{1, 3, 4},
			amount: 6,
			want:   2, // 3 + 3 (贪心会得到 4 + 1 + 1 = 3)
		},
		{
			name:   "Single coin exact",
			coins:  []int{5},
			amount: 5,
			want:   1,
		},
		{
			name:   "Multiple coins same result",
			coins:  []int{1, 2, 5},
			amount: 5,
			want:   1, // 直接用 5
		},
		{
			name:   "Large amount",
			coins:  []int{1, 2, 5},
			amount: 100,
			want:   20, // 20 个 5
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := coinChange(tc.coins, tc.amount)
			assert.Equal(t, tc.want, got, "coins=%v, amount=%d", tc.coins, tc.amount)
		})
	}
}

func TestCoinChangeErr(t *testing.T) {
	t.Parallel()

	t.Run("Greedy algorithm fails", func(t *testing.T) {
		coins := []int{1, 3, 4}
		amount := 6

		got := coinChangeErr(coins, amount)
		// 贪心会得到 3 (4 + 1 + 1)，但正确答案是 2 (3 + 3)
		assert.NotEqual(t, 2, got, "Greedy algorithm gives wrong answer")
		assert.Equal(t, 3, got, "Greedy gives 3 coins instead of optimal 2")
	})
}
