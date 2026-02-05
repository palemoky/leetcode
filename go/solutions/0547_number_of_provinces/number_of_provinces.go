package number_of_provinces

import "leetcode/go/solutions/utils"

// Solution 1:
// Time: O(n²·α(n)), Space: O(n)
// α(n) 是阿克曼函数的反函数，可以认为是常数
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := utils.NewUnionFind(n)

	// 遍历上三角矩阵（因为是对称矩阵）
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count
}
