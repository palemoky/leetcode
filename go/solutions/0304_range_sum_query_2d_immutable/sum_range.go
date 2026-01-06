package range_sum_query_2d_immutable

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	if row == 0 || col == 0 {
		return NumMatrix{}
	}

	preSum := make([][]int, row+1)
	for i := range preSum {
		preSum[i] = make([]int, col+1)
	}

	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			preSum[r][c] = preSum[r-1][c] + preSum[r][c-1] - preSum[r-1][c-1] + matrix[r-1][c-1]
		}
	}

	return NumMatrix{preSum}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.preSum[row2+1][col2+1] - nm.preSum[row1][col2+1] - nm.preSum[row2+1][col1] + nm.preSum[row1][col1]
}
