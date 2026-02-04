package rotate_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func copyMatrix(m [][]int) [][]int {
	if m == nil {
		return nil
	}
	out := make([][]int, len(m))
	for i := range m {
		out[i] = append([]int(nil), m[i]...)
	}

	return out
}

func TestRotateFlip(t *testing.T) {
	cases := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			"0x0 (empty)",
			[][]int{},
			[][]int{},
		},
		{
			"1x1",
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			"2x2",
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			"3x3",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			"4x4",
			[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	funcsToTest := map[string]func(matrix [][]int){
		"Naive": rotateWithExtraSpace,
		"flip":  rotateFlip,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					in := copyMatrix(tc.input)
					fn(in)
					assert.Equal(t, tc.want, in)
				})
			}
		})
	}
}
