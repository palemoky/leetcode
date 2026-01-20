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

## 二分图

二分图是一种特殊的图,其顶点可以分成两个互不相交的集合,使得 **同一集合内的顶点之间没有边**。

### 二分图判定

使用 **染色法**(DFS/BFS):尝试用两种颜色给图着色,如果能成功着色且相邻节点颜色不同,则是二分图。

```go
func isBipartite(graph [][]int) bool {
    n := len(graph)
    color := make([]int, n) // 0: 未染色, 1: 颜色1, -1: 颜色2

    // 处理非连通图
    for i := 0; i < n; i++ {
        if color[i] == 0 {
            if !dfs(graph, i, 1, color) {
                return false
            }
        }
    }
    return true
}

func dfs(graph [][]int, node, c int, color []int) bool {
    color[node] = c

    for _, neighbor := range graph[node] {
        if color[neighbor] == c {
            // 相邻节点颜色相同,不是二分图
            return false
        }
        if color[neighbor] == 0 {
            // 染成相反的颜色
            if !dfs(graph, neighbor, -c, color) {
                return false
            }
        }
    }
    return true
}
```

**应用场景:**

- 任务分配问题
- 配对问题
- LeetCode: [785. Is Graph Bipartite?](https://leetcode.com/problems/is-graph-bipartite/), [886. Possible Bipartition](https://leetcode.com/problems/possible-bipartition/)

---

## 环检测

### 无向图环检测

使用 DFS,记录父节点。如果访问到已访问的节点且不是父节点,则存在环。

```go
func hasCycle(graph [][]int) bool {
    n := len(graph)
    visited := make([]bool, n)

    for i := 0; i < n; i++ {
        if !visited[i] {
            if dfsCycle(graph, i, -1, visited) {
                return true
            }
        }
    }
    return false
}

func dfsCycle(graph [][]int, node, parent int, visited []bool) bool {
    visited[node] = true

    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            if dfsCycle(graph, neighbor, node, visited) {
                return true
            }
        } else if neighbor != parent {
            // 访问到已访问的节点且不是父节点,存在环
            return true
        }
    }
    return false
}
```

### 有向图环检测

使用 DFS + **路径标记**。需要三种状态:未访问、访问中(在当前路径上)、已完成。

```go
func hasCycleDirected(graph [][]int) bool {
    n := len(graph)
    // 0: 未访问, 1: 访问中(在路径上), 2: 已完成
    state := make([]int, n)

    for i := 0; i < n; i++ {
        if state[i] == 0 {
            if dfsCycleDirected(graph, i, state) {
                return true
            }
        }
    }
    return false
}

func dfsCycleDirected(graph [][]int, node int, state []int) bool {
    state[node] = 1 // 标记为访问中

    for _, neighbor := range graph[node] {
        if state[neighbor] == 1 {
            // 遇到访问中的节点,存在环
            return true
        }
        if state[neighbor] == 0 {
            if dfsCycleDirected(graph, neighbor, state) {
                return true
            }
        }
    }

    state[node] = 2 // 标记为已完成
    return false
}
```

---

## 拓扑排序

拓扑排序是对 **有向无环图(DAG)** 的所有顶点进行线性排序,使得对于任何有向边 `u → v`,`u` 在排序中都出现在 `v` 之前。

### 方法 1: DFS + 后序遍历反转

**核心思想**: DFS 后序遍历的结果反转就是拓扑排序。

```go
func topologicalSort(graph [][]int) []int {
    n := len(graph)
    visited := make([]bool, n)
    result := []int{}

    var dfs func(int)
    dfs = func(node int) {
        visited[node] = true

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }

        // 后序位置:所有子节点都已访问完
        result = append(result, node)
    }

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
        }
    }

    // 反转结果
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }

    return result
}
```

### 方法 2: Kahn 算法(BFS + 入度)

```go
func topologicalSortKahn(graph [][]int) []int {
    n := len(graph)
    inDegree := make([]int, n)

    // 计算入度
    for i := 0; i < n; i++ {
        for _, neighbor := range graph[i] {
            inDegree[neighbor]++
        }
    }

    // 将入度为 0 的节点加入队列
    queue := []int{}
    for i := 0; i < n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    result := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)

        // 删除该节点的所有出边
        for _, neighbor := range graph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    // 如果结果长度不等于节点数,说明存在环
    if len(result) != n {
        return []int{} // 存在环,无法拓扑排序
    }

    return result
}
```

**应用场景:**

- 课程安排(有先修课程要求)
- 任务调度
- 编译依赖
- LeetCode: [207. Course Schedule](https://leetcode.com/problems/course-schedule/), [210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii/)

---

## 无权图

无权图的边没有权重,只表示节点之间的连接关系。这是最基础的图结构,适合表示二元关系。

### 无权图的表示

```go
// 邻接表表示无权图
type UnweightedGraph struct {
    graph [][]int  // graph[i] 存储节点 i 的所有邻居
}

// 邻接矩阵表示无权图
type UnweightedGraphMatrix struct {
    matrix [][]bool  // matrix[i][j] = true 表示 i 和 j 之间有边
}
```

### BFS 求最短路径

在**无权图**中,BFS 可以找到从起点到任意节点的**最短路径**(边数最少)。

```go
func shortestPath(graph [][]int, start, target int) int {
    n := len(graph)
    visited := make([]bool, n)
    queue := []int{start}
    visited[start] = true
    step := 0

    for len(queue) > 0 {
        size := len(queue)

        // 遍历当前层的所有节点
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            if node == target {
                return step
            }

            // 将邻居加入队列
            for _, neighbor := range graph[node] {
                if !visited[neighbor] {
                    visited[neighbor] = true
                    queue = append(queue, neighbor)
                }
            }
        }

        step++
    }

    return -1  // 无法到达
}
```

### 双向 BFS 优化

当起点和终点都已知时,可以使用**双向 BFS** 来优化搜索效率。

```go
func bidirectionalBFS(graph [][]int, start, target int) int {
    if start == target {
        return 0
    }

    // 从起点和终点同时开始搜索
    visitedStart := map[int]bool{start: true}
    visitedTarget := map[int]bool{target: true}
    queueStart := []int{start}
    queueTarget := []int{target}
    step := 0

    for len(queueStart) > 0 && len(queueTarget) > 0 {
        step++

        // 优化:总是扩展较小的队列
        if len(queueStart) > len(queueTarget) {
            queueStart, queueTarget = queueTarget, queueStart
            visitedStart, visitedTarget = visitedTarget, visitedStart
        }

        size := len(queueStart)
        for i := 0; i < size; i++ {
            node := queueStart[0]
            queueStart = queueStart[1:]

            for _, neighbor := range graph[node] {
                if visitedTarget[neighbor] {
                    // 两个方向相遇
                    return step
                }
                if !visitedStart[neighbor] {
                    visitedStart[neighbor] = true
                    queueStart = append(queueStart, neighbor)
                }
            }
        }
    }

    return -1
}
```

**双向 BFS 的优势:**

- 时间复杂度从 $O(b^d)$ 降低到 $O(b^{d/2})$ (b 为分支因子,d 为深度)
- 适合搜索空间很大的场景

### 应用场景

**无权图的典型应用:**

- **社交网络** - 好友关系、关注关系
  - 好友推荐(共同好友)
  - 六度分离理论
  - 影响力传播
- **网络拓扑** - 计算机网络、互联网连接
- **迷宫/棋盘** - 最短路径问题
- **单词接龙** - 词汇转换问题
- **基因序列** - DNA 序列相似度

**LeetCode 经典题目:**

- [127. Word Ladder](https://leetcode.com/problems/word-ladder/) - 单词接龙(BFS 最短路径)
- [433. Minimum Genetic Mutation](https://leetcode.com/problems/minimum-genetic-mutation/) - 基因序列(BFS)
- [1091. Shortest Path in Binary Matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/) - 二进制矩阵最短路径
- [752. Open the Lock](https://leetcode.com/problems/open-the-lock/) - 开锁(双向 BFS)

---

## 加权图

加权图的边带有权重，常用于最短路径、最小生成树等问题。

### 加权图的表示

```go
// 方法 1: 邻接表 + 边结构
type Edge struct {
    to     int
    weight int
}

type WeightedGraph struct {
    graph [][]Edge
}

// 方法 2: 邻接表 + 二维数组
// graph[i] = [[neighbor1, weight1], [neighbor2, weight2], ...]
type WeightedGraph2 struct {
    graph [][][2]int
}
```

### Dijkstra 算法(单源最短路径)

适用于 **非负权重** 的图,使用优先队列(最小堆)。

```go
import "container/heap"

type Item struct {
    node int
    dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func dijkstra(graph [][]Edge, start int) []int {
    n := len(graph)
    dist := make([]int, n)
    for i := range dist {
        dist[i] = 1<<31 - 1 // 初始化为无穷大
    }
    dist[start] = 0

    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, Item{start, 0})

    for pq.Len() > 0 {
        item := heap.Pop(pq).(Item)
        node, d := item.node, item.dist

        if d > dist[node] {
            continue
        }

        for _, edge := range graph[node] {
            newDist := dist[node] + edge.weight
            if newDist < dist[edge.to] {
                dist[edge.to] = newDist
                heap.Push(pq, Item{edge.to, newDist})
            }
        }
    }

    return dist
}
```

### Bellman-Ford 算法

适用于 **有负权重** 的图,可以检测负权环。

```go
func bellmanFord(edges [][3]int, n, start int) []int {
    // edges[i] = [from, to, weight]
    dist := make([]int, n)
    for i := range dist {
        dist[i] = 1<<31 - 1
    }
    dist[start] = 0

    // 松弛 n-1 次
    for i := 0; i < n-1; i++ {
        for _, edge := range edges {
            from, to, weight := edge[0], edge[1], edge[2]
            if dist[from] != 1<<31-1 && dist[from]+weight < dist[to] {
                dist[to] = dist[from] + weight
            }
        }
    }

    // 检测负权环
    for _, edge := range edges {
        from, to, weight := edge[0], edge[1], edge[2]
        if dist[from] != 1<<31-1 && dist[from]+weight < dist[to] {
            // 存在负权环
            return []int{}
        }
    }

    return dist
}
```

### 最短路径算法对比

| 算法               | 适用场景                         | 时间复杂度        | 空间复杂度 |
| ------------------ | -------------------------------- | ----------------- | ---------- |
| **Dijkstra**       | 非负权重,单源最短路径            | $O((V+E) \log V)$ | $O(V)$     |
| **Bellman-Ford**   | 有负权重,单源最短路径,检测负权环 | $O(VE)$           | $O(V)$     |
| **Floyd-Warshall** | 所有节点对最短路径               | $O(V^3)$          | $O(V^2)$   |

**应用场景:**

- 地图导航(最短路径) - 权重是距离/时间
- 网络路由 - 权重是延迟/带宽
- 航班网络 - 权重是票价/飞行时间
- 物流配送 - 权重是运输成本
- 社交网络影响力 - 权重是亲密度/互动频率
- LeetCode: [743. Network Delay Time](https://leetcode.com/problems/network-delay-time/), [787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)
