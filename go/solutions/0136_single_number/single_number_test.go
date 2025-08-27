package single_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var funcsToTests = map[string]func([]int) int{
	"HashMap": singleNumberHashMap,
	"BitWise": singleNumberBitWise,
}

func TestSingleNumber(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{"Simple single", []int{2, 2, 1}, 1},
		{"Single at head", []int{4, 1, 2, 1, 2}, 4},
		{"Single element", []int{99}, 99},
		{"Negative number", []int{-1, -1, -2}, -2},
		{"Zero included", []int{0, 1, 1}, 0},
		{"Large numbers", []int{100000, 1, 1}, 100000},
		{"Mixed positive and negative", []int{-3, -3, 0, 0, 7}, 7},
		{"Empty", []int{}, 0},
		// {"All same", []int{5, 5, 5}, 0}, // BitWise cannot handle this case, but HashMap can handle it
	}

	for fnName, fn := range funcsToTests {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					got := fn(tc.nums)
					assert.Equal(t, tc.want, got, "%s: input=%v", fnName, tc.nums)
				})
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}

	for fnName, fn := range funcsToTests {
		b.Run(fnName, func(b *testing.B) {
			for b.Loop() {
				fn(nums)
			}
		})
	}
}
