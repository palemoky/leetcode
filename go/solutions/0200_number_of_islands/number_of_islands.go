package number_of_islands

import "leetcode/go/solutions/utils"

// 推荐解法：DFS（原地修改）
// Time: O(m×n), Space: O(m×n) 递归栈
func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	// 边界检查
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	// 标记为已访问
	grid[i][j] = '0'

	// 递归访问四个方向
	dfs(grid, i-1, j) // 上
	dfs(grid, i+1, j) // 下
	dfs(grid, i, j-1) // 左
	dfs(grid, i, j+1) // 右
}

// 解法二：BFS（原地修改）
// Time: O(m×n), Space: O(m×n) 队列
func numIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				count++
				bfs(grid, i, j)
			}
		}
	}
	return count
}

func bfs(grid [][]byte, i, j int) {
	queue := [][2]int{{i, j}}
	grid[i][j] = '0'

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			ni, nj := pos[0]+dir[0], pos[1]+dir[1]
			if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == '1' {
				grid[ni][nj] = '0'
				queue = append(queue, [2]int{ni, nj})
			}
		}
	}
}

// 解法三：并查集（不修改原数组）
// Time: O(m×n), Space: O(m×n) 并查集
func numIslandsUnionFind(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	uf := utils.NewUnionFind(m * n)

	// 统计水的数量
	water := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '0' {
				water++
			} else {
				// 向右和向下合并
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j, i*n+j+1)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union(i*n+j, (i+1)*n+j)
				}
			}
		}
	}

	return uf.Count - water
}
