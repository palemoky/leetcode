package number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "示例1：三个岛屿",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "示例2：一个岛屿",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name:     "空网格",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name: "全是水",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "全是陆地",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 1,
		},
		{
			name: "单个岛屿",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "多个独立岛屿",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			expected: 5,
		},
	}

	funcsToTest := map[string]func([][]byte) int{
		"numIslandsDFS":       numIslandsDFS,
		"numIslandsBFS":       numIslandsBFS,
		"numIslandsUnionFind": numIslandsUnionFind,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					// 深拷贝 grid，因为 DFS 和 BFS 会修改原数组
					gridCopy := make([][]byte, len(tc.grid))
					for i := range tc.grid {
						gridCopy[i] = make([]byte, len(tc.grid[i]))
						copy(gridCopy[i], tc.grid[i])
					}
					result := fn(gridCopy)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
