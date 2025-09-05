package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateImplementations(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"basic", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"k_eq_len", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k_gt_len", []int{1, 2, 3}, 4, []int{3, 1, 2}},
		{"single", []int{1}, 0, []int{1}},
		// {"empty", []int{}, 0, []int{}}, // 对 0 取余等价于除 0
	}

	funcs := map[string]func([]int, int){
		"BruteForce": rotateBruteForce,
		"Slice":      rotateSlice,
	}

	for fnName, fn := range funcs {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					in := append([]int{}, tc.nums...)
					fn(in, tc.k)
					assert.Equal(t, tc.want, in)
				})
			}
		})
	}
}
