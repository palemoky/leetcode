package maximum_matching_of_players_with_trainers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		players  []int
		trainers []int
		expected int
	}{
		{
			name:     "example 1",
			players:  []int{4, 7, 9},
			trainers: []int{8, 2, 5, 8},
			expected: 2,
		},
		{
			name:     "example 2",
			players:  []int{1, 1, 1},
			trainers: []int{10},
			expected: 1,
		},
		{
			name:     "no players",
			players:  []int{},
			trainers: []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "no trainers",
			players:  []int{1, 2, 3},
			trainers: []int{},
			expected: 0,
		},
		{
			name:     "all unmatched",
			players:  []int{5, 6, 7},
			trainers: []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "all matched with duplicates",
			players:  []int{2, 2, 2},
			trainers: []int{2, 2, 2},
			expected: 3,
		},
		{
			name:     "greedy pairing required",
			players:  []int{3, 4, 4, 9},
			trainers: []int{3, 3, 5, 10},
			expected: 3,
		},
	}

	funcsToTest := map[string]func([]int, []int) int{
		"matchPlayersAndTrainers": matchPlayersAndTrainers,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					players := append([]int(nil), tc.players...)
					trainers := append([]int(nil), tc.trainers...)
					result := fn(players, trainers)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
