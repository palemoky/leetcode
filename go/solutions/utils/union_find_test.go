package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	t.Parallel()

	t.Run("初始化", func(t *testing.T) {
		uf := NewUnionFind(5)
		assert.Equal(t, 5, uf.Count, "初始连通分量应该等于元素数量")

		// 每个元素的根应该是自己
		for i := range 5 {
			assert.Equal(t, i, uf.Find(i), "初始时每个元素的根应该是自己")
		}
	})

	t.Run("合并操作", func(t *testing.T) {
		uf := NewUnionFind(5)

		// 合并 0 和 1
		uf.Union(0, 1)
		assert.Equal(t, 4, uf.Count, "合并后连通分量应该减1")
		assert.Equal(t, uf.Find(0), uf.Find(1), "0和1应该在同一个集合")

		// 合并 2 和 3
		uf.Union(2, 3)
		assert.Equal(t, 3, uf.Count, "再次合并后连通分量应该再减1")
		assert.Equal(t, uf.Find(2), uf.Find(3), "2和3应该在同一个集合")

		// 合并 0 和 2（连接两个集合）
		uf.Union(0, 2)
		assert.Equal(t, 2, uf.Count, "合并两个集合后连通分量应该再减1")
		assert.Equal(t, uf.Find(0), uf.Find(3), "0和3应该在同一个集合")
		assert.Equal(t, uf.Find(1), uf.Find(2), "1和2应该在同一个集合")
	})

	t.Run("重复合并", func(t *testing.T) {
		uf := NewUnionFind(3)

		uf.Union(0, 1)
		assert.Equal(t, 2, uf.Count)

		// 重复合并同一对元素
		uf.Union(0, 1)
		assert.Equal(t, 2, uf.Count, "重复合并不应该改变连通分量数")

		uf.Union(1, 0)
		assert.Equal(t, 2, uf.Count, "反向合并也不应该改变连通分量数")
	})

	t.Run("路径压缩", func(t *testing.T) {
		uf := NewUnionFind(5)

		// 创建一条链：0 -> 1 -> 2 -> 3 -> 4
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)
		uf.Union(3, 4)

		assert.Equal(t, 1, uf.Count, "所有元素应该在同一个集合")

		// 查找会触发路径压缩
		root := uf.Find(0)
		assert.Equal(t, root, uf.Find(4), "所有元素应该有相同的根")

		// 路径压缩后，再次查找应该更快（虽然无法直接测试性能）
		for i := 0; i < 5; i++ {
			assert.Equal(t, root, uf.Find(i), "路径压缩后所有元素应该直接指向根")
		}
	})

	t.Run("完全连通", func(t *testing.T) {
		uf := NewUnionFind(4)

		// 将所有元素连接成一个集合
		uf.Union(0, 1)
		uf.Union(1, 2)
		uf.Union(2, 3)

		assert.Equal(t, 1, uf.Count, "应该只有一个连通分量")

		root := uf.Find(0)
		for i := 1; i < 4; i++ {
			assert.Equal(t, root, uf.Find(i), "所有元素应该在同一个集合")
		}
	})

	t.Run("完全独立", func(t *testing.T) {
		uf := NewUnionFind(5)

		// 不进行任何合并
		assert.Equal(t, 5, uf.Count, "应该有5个独立的连通分量")

		// 每个元素应该是独立的
		for i := 0; i < 5; i++ {
			for j := i + 1; j < 5; j++ {
				assert.NotEqual(t, uf.Find(i), uf.Find(j), "不同元素应该在不同集合")
			}
		}
	})
}
