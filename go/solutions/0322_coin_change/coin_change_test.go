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
		{"Example 1", []int{1, 2, 5}, 11, 3},      // 5 + 5 + 1
		{"Example 2", []int{2}, 3, -1},            // 无法凑成
		{"Example 3", []int{1}, 0, 0},             // 金额为0
		{"Greedy fails", []int{1, 3, 4}, 6, 2},    // 3 + 3 (贪心会得到 4 + 1 + 1 = 3)
		{"Single coin exact", []int{5}, 5, 1},     // 恰好一枚
		{"Multiple coins", []int{1, 2, 5}, 5, 1},  // 直接用 5
		{"Large amount", []int{1, 2, 5}, 100, 20}, // 20 个 5
	}

	// 暴力递归：仅测试小规模用例，避免超时
	smallTestCases := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"Example 1", []int{1, 2, 5}, 11, 3},
		{"Example 2", []int{2}, 3, -1},
		{"Example 3", []int{1}, 0, 0},
		{"Greedy fails", []int{1, 3, 4}, 6, 2},
		{"Single coin exact", []int{5}, 5, 1},
		{"Multiple coins", []int{1, 2, 5}, 5, 1},
	}

	funcsToTest := map[string]struct {
		fn        func([]int, int) int
		testCases []struct {
			name   string
			coins  []int
			amount int
			want   int
		}
	}{
		"BruteForce": {coinChangeBruteForce, smallTestCases}, // 仅测试小规模
		"Memoized":   {coinChangeMemo, testCases},
		"DP":         {coinChange, testCases},
	}

	for fnName, config := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range config.testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := config.fn(tc.coins, tc.amount)
					assert.Equal(t, tc.want, got, "coins=%v, amount=%d", tc.coins, tc.amount)
				})
			}
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
