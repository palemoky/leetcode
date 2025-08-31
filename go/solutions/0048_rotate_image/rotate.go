package rotate_image

func rotateWithExtraSpace(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	newMatrix := make([][]int, n)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}

	// 原坐标 (row, col) -> 新坐标 (col, n - 1 - row)
	for row := range n {
		for col := range n {
			newMatrix[col][n-1-row] = matrix[row][col]
		}
	}

	for row := range n {
		for col := range n {
			matrix[row][col] = newMatrix[row][col]
		}
	}
}

func rotateFlip(matrix [][]int) {
	n := len(matrix)

	// Step 1: 沿主对角线翻转
	for i := range n {
		for j := i + 1; j < n; j++ {
			// 只遍历上三角或下三角区域，避免重复交换
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: 沿垂直中轴线翻转（翻转每一行）
	for i := range n {
		// 使用双指针翻转每一行
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
	}
}
