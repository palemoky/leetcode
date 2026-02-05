package utils

// UnionFind 结构体
type UnionFind struct {
	parent []int
	Count  int // 连通分量数量
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, Count: n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.Count-- // 合并后连通分量减1
	}
}
