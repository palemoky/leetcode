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
		{"k equals length", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k greater than length", []int{1, 2, 3}, 4, []int{3, 1, 2}},
		{"k is zero", []int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
		{"k is one", []int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
		{"single element", []int{1}, 0, []int{1}},
		{"single element with k", []int{5}, 1, []int{5}},
		{"two elements", []int{1, 2}, 1, []int{2, 1}},
		{"negative numbers", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{"large k", []int{1, 2, 3}, 10, []int{3, 1, 2}}, // 10 % 3 = 1
	}

	funcs := map[string]func([]int, int){
		"Shift":   rotateShift,
		"Slice":   rotateSlice,
		"Reverse": rotateReverse,
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
