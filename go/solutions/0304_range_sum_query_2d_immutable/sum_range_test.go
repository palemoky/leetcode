package range_sum_query_2d_immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMatrix(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		matrix     [][]int
		operations []struct {
			row1 int
			col1 int
			row2 int
			col2 int
			want int
		}
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{2, 1, 4, 3, 8},  // sum of [[2, 0, 1], [1, 0, 1], [0, 3, 0]] = 8
				{1, 1, 2, 2, 11}, // sum of [[6, 3], [2, 0]] = 11
				{1, 2, 2, 4, 12}, // sum of [[3, 2, 1], [0, 1, 5]] = 12
			},
		},
		{
			name: "Single element",
			matrix: [][]int{
				{5},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 0, 0, 5}, // sum of [[5]] = 5
			},
		},
		{
			name: "Single row",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 0, 4, 15}, // sum of [[1, 2, 3, 4, 5]] = 15
				{0, 1, 0, 3, 9},  // sum of [[2, 3, 4]] = 9
				{0, 0, 0, 0, 1},  // sum of [[1]] = 1
			},
		},
		{
			name: "Single column",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 4, 0, 15}, // sum of [[1], [2], [3], [4], [5]] = 15
				{1, 0, 3, 0, 9},  // sum of [[2], [3], [4]] = 9
				{0, 0, 0, 0, 1},  // sum of [[1]] = 1
			},
		},
		{
			name: "All zeros",
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 2, 2, 0}, // sum of all zeros = 0
				{1, 1, 2, 2, 0}, // sum of sub-matrix zeros = 0
			},
		},
		{
			name: "Negative numbers",
			matrix: [][]int{
				{-1, -2, -3},
				{-4, -5, -6},
				{-7, -8, -9},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 2, 2, -45}, // sum of all = -45
				{1, 1, 2, 2, -28}, // sum of [[-5, -6], [-8, -9]] = -28
				{0, 0, 0, 0, -1},  // sum of [[-1]] = -1
			},
		},
		{
			name: "Mixed values",
			matrix: [][]int{
				{-2, 5, -1},
				{7, -3, 4},
				{1, 6, -2},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 2, 2, 15}, // sum of entire matrix = 15
				{1, 1, 2, 2, 5},  // sum of [[-3, 4], [6, -2]] = 5
				{0, 1, 1, 2, 5},  // sum of [[5, -1], [-3, 4]] = 5
			},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			operations: []struct {
				row1 int
				col1 int
				row2 int
				col2 int
				want int
			}{
				{0, 0, 1, 1, 10}, // sum of [[1, 2], [3, 4]] = 10
				{0, 0, 0, 1, 3},  // sum of [[1, 2]] = 3
				{0, 0, 1, 0, 4},  // sum of [[1], [3]] = 4
				{1, 1, 1, 1, 4},  // sum of [[4]] = 4
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			numMatrix := Constructor(tc.matrix)
			for i, op := range tc.operations {
				got := numMatrix.SumRegion(op.row1, op.col1, op.row2, op.col2)
				assert.Equal(t, op.want, got,
					"Operation %d: SumRegion(%d, %d, %d, %d) failed",
					i+1, op.row1, op.col1, op.row2, op.col2)
			}
		})
	}
}

func BenchmarkNumMatrix(b *testing.B) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)

	b.Run("SumRegion", func(b *testing.B) {
		for b.Loop() {
			numMatrix.SumRegion(2, 1, 4, 3)
			numMatrix.SumRegion(1, 1, 2, 2)
			numMatrix.SumRegion(1, 2, 2, 4)
		}
	})

	b.Run("Constructor", func(b *testing.B) {
		for b.Loop() {
			Constructor(matrix)
		}
	})
}
