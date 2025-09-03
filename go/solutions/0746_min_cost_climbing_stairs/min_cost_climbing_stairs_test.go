package min_cost_climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		cost []int
		want int
	}{
		{"classic_three_steps", []int{10, 15, 20}, 15},
		{"zigzag_min_path", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"single_step", []int{5}, 0},
		{"two_steps", []int{5, 6}, 5},
		{"all_zero", []int{0, 0, 0, 0}, 0},
		{"increasing_cost", []int{1, 2, 3, 4, 5}, 6},
		{"decreasing_cost", []int{5, 4, 3, 2, 1}, 6},
	}

	funcsToTest := map[string]func(cost []int) int{
		"DP":        minCostClimbingStairsDP,
		"Iterative": minCostClimbingStairsIterative,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(append([]int{}, tc.cost...))
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
