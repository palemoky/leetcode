# 图

核心理念：由顶点(Vertex)和边(Edge)组成的网络结构。
为何重要：用于表示各种网络关系，如社交网络、地图路线、依赖关系等。

图有邻接表和邻接矩阵两种形式，算法中通常用邻接表。

必练操作：

- 图的表示方法：邻接矩阵和邻接表。
- 图的遍历：深度优先搜索(DFS)和广度优先搜索(BFS)。
- （进阶）最小生成树算法 (Prim, Kruskal)、最短路径算法 (Dijkstra)。

## 图的遍历框架

图的遍历与树的遍历最大的区别在于:**图可能存在环**,因此需要使用 `visited` 数组来避免重复访问节点。

### 图的表示(邻接表)

```go
// 邻接表表示图
type Graph struct {
    // graph[i] 存储节点 i 的所有邻居节点
    graph [][]int
}
```

### 深度优先搜索(DFS)

DFS 使用递归或栈来实现,适合解决路径、连通性、环检测等问题。

```go
// DFS 遍历框架
func dfs(graph [][]int, start int, visited []bool) {
    // 标记当前节点已访问
    visited[start] = true

    // 前序位置:进入节点时的操作
    // ...

    // 遍历所有邻居节点
    for _, neighbor := range graph[start] {
        if !visited[neighbor] {
            dfs(graph, neighbor, visited)
        }
    }

    // 后序位置:离开节点时的操作
    // ...
}

// 遍历整个图(处理非连通图)
func traverseGraph(graph [][]int) {
    n := len(graph)
    visited := make([]bool, n)

    // 遍历所有节点,确保访问到所有连通分量
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(graph, i, visited)
        }
    }
}
```

### 广度优先搜索(BFS)

BFS 使用队列来实现,适合解决最短路径、层级遍历等问题。

```go
// BFS 遍历框架
func bfs(graph [][]int, start int) {
    n := len(graph)
    visited := make([]bool, n)
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        // 取出队首节点
        node := queue[0]
        queue = queue[1:]

        // 访问当前节点
        // ...

        // 将所有未访问的邻居加入队列
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}
```

### DFS vs BFS

| 特性           | DFS                        | BFS                   |
| -------------- | -------------------------- | --------------------- |
| **数据结构**   | 栈(递归)                   | 队列                  |
| **空间复杂度** | $O(h)$ (h 为图的深度)      | $O(w)$ (w 为图的宽度) |
| **适用场景**   | 路径问题、拓扑排序、环检测 | 最短路径、层级遍历    |
| **遍历顺序**   | 深度优先,一条路走到底      | 广度优先,逐层扩散     |

### 常见应用场景

**DFS 应用:**

- 检测图中是否有环
- 拓扑排序
- 寻找所有路径
- 连通分量计数

**BFS 应用:**

- 无权图的最短路径
- 层级遍历
- 最小步数问题
- 二分图检测

### 关键注意事项

1. **visited 数组**: 防止重复访问,避免死循环
2. **非连通图**: 需要遍历所有节点作为起点,确保访问到所有连通分量
3. **有向图 vs 无向图**: 邻接表的构建方式不同
   - 无向图: `graph[u]` 包含 `v`,`graph[v]` 也包含 `u`
   - 有向图: 只在 `graph[u]` 中添加 `v`

## 图的 DFS vs 回溯 DFS

图的 DFS 和回溯算法都使用深度优先搜索,但它们的目的和实现细节有重要区别。

### 核心区别

| 维度             | 图的 DFS (遍历)       | 回溯 (找所有路径)               |
| ---------------- | --------------------- | ------------------------------- |
| **目的**         | 遍历每个节点一次      | 找所有可能的路径/组合           |
| **visited 含义** | "这个节点被访问过了"  | "这个节点在当前路径中"          |
| **做选择**       | ✅ for 循环外         | ✅ for 循环外                   |
| **撤销选择**     | ❌ 不撤销             | ✅ for 循环后撤销               |
| **节点重复访问** | ❌ 每个节点只访问一次 | ✅ 不同路径可以重复访问同一节点 |

### 代码对比

**图的 DFS (遍历):**

```go
func graphDFS(graph [][]int, start int, visited []bool) {
    // 做选择(在 for 循环外)
    visited[start] = true

    // 前序位置:访问节点
    fmt.Println(start)

    // 遍历所有邻居
    for _, neighbor := range graph[start] {
        if !visited[neighbor] {
            graphDFS(graph, neighbor, visited)
        }
    }

    // ❌ 不撤销选择
}
```

**回溯 (找所有路径):**

```go
func backtrackDFS(graph [][]int, start, target int, visited []bool, path []int, result *[][]int) {
    // 做选择(在 for 循环外)
    path = append(path, start)
    visited[start] = true

    // 到达目标,记录路径
    if start == target {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *result = append(*result, tmp)
    }

    // 遍历所有邻居
    for _, neighbor := range graph[start] {
        if !visited[neighbor] {
            backtrackDFS(graph, neighbor, target, visited, path, result)
        }
    }

    // ✅ 撤销选择(在 for 循环外) - 这是回溯的精髓!
    visited[start] = false
    path = path[:len(path)-1]
}
```

### 具体例子

假设有这样一个图:

```
    1
   / \
  2   3
   \ /
    4
```

**图的 DFS (遍历):**

- 访问顺序: `1 → 2 → 4 → 3`
- 每个节点只访问一次
- `visited[4] = true` 后,从 3 到 4 的边不会再走

**回溯 (找从 1 到 4 的所有路径):**

- 路径 1: `1 → 2 → 4`
- 路径 2: `1 → 3 → 4`
- 节点 4 在两条路径中都被访问
- 第一条路径结束后,`visited[4]` 被撤销,允许第二条路径再次访问

### 使用场景

**图的 DFS:**

- 遍历所有节点
- 检测环
- 拓扑排序
- 连通分量计数

**回溯:**

- 找所有路径
- 找所有组合/排列
- N 皇后问题
- 数独求解

### 记忆要点

> **回溯 = DFS + 撤销选择**

图的 DFS 关注"是否访问过"，回溯关注"当前路径的状态"。撤销选择是回溯的精髓，它允许算法探索所有可能的解空间。
