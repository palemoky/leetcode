package number_of_provinces

// Solution 1:
// Time: O(), Space: O()
type UnionFind struct {
	parent []int
	count  int // 连通分量数量
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, count: n}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count-- // 合并后连通分量减1
	}
}

// Time: O(n²·α(n)), Space: O(n)
// α(n) 是阿克曼函数的反函数，可以认为是常数
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)

	// 遍历上三角矩阵（因为是对称矩阵）
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	return uf.count
}
