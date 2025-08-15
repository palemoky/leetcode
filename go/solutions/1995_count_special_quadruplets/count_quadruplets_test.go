package count_special_quadruplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountQuadruplets(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "Too short",
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "Simple example",
			nums: []int{1, 2, 3, 6},
			want: 1, // 1+2+3 = 6
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0, 0},
			want: 1, // one quadruplet (0,0,0,0)
		},
		{
			name: "Duplicates produce multiple",
			nums: []int{1, 1, 1, 3, 5},
			want: 4, // explained in discussion: total 4 valid quadruplets
		},
		{
			name: "No solution",
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
	}

	funcsToTest := map[string]func([]int) int{
		"BruteForce": countQuadruplets,
		"HashMap":    countQuadrupletsHashMap,
	}

	for fnName, fn := range funcsToTest {
		fnName, fn := fnName, fn
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// pass a copy to avoid in-place modification issues
					input := append([]int(nil), tc.nums...)
					got := fn(input)
					assert.Equal(t, tc.want, got, "nums=%v", tc.nums)
				})
			}
		})
	}
}
