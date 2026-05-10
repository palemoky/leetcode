package gas_station

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "example 1",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "example 2",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "single station enough gas",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "single station not enough gas",
			gas:      []int{1},
			cost:     []int{2},
			expected: -1,
		},
		{
			name:     "start shifts after deficit",
			gas:      []int{5, 1, 2, 3, 4},
			cost:     []int{4, 4, 1, 5, 1},
			expected: 4,
		},
		{
			name:     "all zeros",
			gas:      []int{0, 0, 0},
			cost:     []int{0, 0, 0},
			expected: 0,
		},
	}

	funcsToTest := map[string]func([]int, []int) int{
		"canCompleteCircuit": canCompleteCircuit,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.gas, tc.cost)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
