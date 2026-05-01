package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1 reachable",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "example 2 unreachable",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "stuck at start",
			nums: []int{0, 1},
			want: false,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := canJump(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}
}
