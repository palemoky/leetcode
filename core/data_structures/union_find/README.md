# 并查集

并查集（Union-Find / Disjoint Set Union, DSU）是一种用于处理 **不相交集合** 的合并与查询问题的数据结构。它能够高效地判断两个元素是否属于同一集合，以及将两个集合合并。

**核心操作：**

- **Find（查找）**：查找元素所属集合的代表元素（根节点）
- **Union（合并）**：将两个集合合并为一个集合

**关键特性：**

- 近乎 $O(1)$ 的查找和合并操作（使用优化后）
- 动态维护连通性关系
- 空间复杂度 $O(n)$

并查集在图论中的连通性问题、最小生成树算法（Kruskal）、网络连接判断等场景中有广泛应用。

| 操作      | 朴素实现 | 路径压缩 | 按秩合并   | 路径压缩 + 按秩合并 |
| --------- | -------- | -------- | ---------- | ------------------- |
| **Find**  | $O(n)$   | $O(n)$   | $O(log n)$ | $O(\alpha(n))$      |
| **Union** | $O(n)$   | $O(n)$   | $O(log n)$ | $O(\alpha(n))$      |

其中 $\alpha(n)$ 是阿克曼函数的反函数，增长极其缓慢，实际应用中可视为常数。

---

## 实现

### 基本结构

```python
class UnionFind:
    def __init__(self, n: int):
        # parent[i] 表示元素 i 的父节点
        self.parent = list(range(n))
        # rank[i] 表示以 i 为根的树的高度（秩）
        self.rank = [0] * n
        # 集合数量
        self.count = n
```

### 核心操作

=== "查找（路径压缩）"

    ```python
    def find(self, x: int) -> int:
        """查找 x 的根节点，并进行路径压缩"""
        if self.parent[x] != x:
            # 递归查找根节点，并将路径上的所有节点直接连接到根节点
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]
    ```

    **路径压缩原理：** 在查找过程中，将路径上的所有节点直接连接到根节点，使树变得扁平化。

    **时间复杂度**: 均摊 $O(\alpha(n))$
    **空间复杂度**: $O(h)$，其中 $h$ 是递归深度

    !!! tip "迭代版本"

        ```python
        def find(self, x: int) -> int:
            """迭代版本的路径压缩"""
            root = x
            # 找到根节点
            while self.parent[root] != root:
                root = self.parent[root]

            # 路径压缩：将路径上所有节点直接连接到根节点
            while x != root:
                next_node = self.parent[x]
                self.parent[x] = root
                x = next_node

            return root
        ```

=== "合并（按秩合并）"

    ```python
    def union(self, x: int, y: int) -> bool:
        """合并 x 和 y 所在的集合"""
        root_x = self.find(x)
        root_y = self.find(y)

        # 已经在同一集合中
        if root_x == root_y:
            return False

        # 按秩合并：将秩小的树连接到秩大的树下
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            # 秩相同时，任意选择一个作为根，并增加其秩
            self.parent[root_y] = root_x
            self.rank[root_x] += 1

        self.count -= 1
        return True
    ```

    **按秩合并原理：** 总是将高度较小的树连接到高度较大的树下，避免树退化成链表。

    **时间复杂度**: 均摊 $O(\alpha(n))$
    **空间复杂度**: $O(1)$

=== "判断连通性"

    ```python
    def is_connected(self, x: int, y: int) -> bool:
        """判断 x 和 y 是否在同一集合中"""
        return self.find(x) == self.find(y)
    ```

    **时间复杂度**: 均摊 $O(\alpha(n))$
    **空间复杂度**: $O(1)$

=== "获取集合数量"

    ```python
    def get_count(self) -> int:
        """获取当前集合的数量"""
        return self.count
    ```

    **时间复杂度**: $O(1)$
    **空间复杂度**: $O(1)$

## 优化技巧

### 1. 路径压缩（Path Compression）

在 `find` 操作中，将查找路径上的所有节点直接连接到根节点。

**效果：** 使树的高度趋近于 1，后续查找操作接近 $O(1)$。

```python
# 递归版本
def find(self, x: int) -> int:
    if self.parent[x] != x:
        self.parent[x] = self.find(self.parent[x])  # 路径压缩
    return self.parent[x]
```

### 2. 按秩合并（Union by Rank）

**秩（Rank）通常指树的高度**。在 `union` 操作中，总是将高度小的树连接到高度大的树下。

**效果：** 控制树的高度，避免退化成链表。

```python
def union(self, x: int, y: int) -> bool:
    root_x, root_y = self.find(x), self.find(y)
    if root_x == root_y:
        return False

    # 按秩合并
    if self.rank[root_x] < self.rank[root_y]:
        self.parent[root_x] = root_y
    elif self.rank[root_x] > self.rank[root_y]:
        self.parent[root_y] = root_x
    else:
        self.parent[root_y] = root_x
        self.rank[root_x] += 1

    return True
```

### 3. 按大小合并（Union by Size）

这是按秩合并的一个变体，**使用集合大小作为"秩"的度量**，将小集合合并到大集合。与按高度合并相比，按大小合并的优势是可以直接获取集合大小。

```python
class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.size = [1] * n  # size[i] 表示以 i 为根的集合大小

    def union(self, x: int, y: int) -> bool:
        root_x, root_y = self.find(x), self.find(y)
        if root_x == root_y:
            return False

        # 按大小合并
        if self.size[root_x] < self.size[root_y]:
            self.parent[root_x] = root_y
            self.size[root_y] += self.size[root_x]
        else:
            self.parent[root_y] = root_x
            self.size[root_x] += self.size[root_y]

        return True

    def get_size(self, x: int) -> int:
        """获取 x 所在集合的大小"""
        return self.size[self.find(x)]
```

## 进阶操作

=== "撤销操作（可持久化并查集）"

    使用栈记录每次操作，支持撤销。

    ```python
    class UndoableUnionFind:
        def __init__(self, n: int):
            self.parent = list(range(n))
            self.rank = [0] * n
            self.history = []  # 记录操作历史

        def union(self, x: int, y: int) -> bool:
            root_x, root_y = self.find(x), self.find(y)
            if root_x == root_y:
                return False

            # 记录操作前的状态
            if self.rank[root_x] < self.rank[root_y]:
                self.history.append((root_x, self.parent[root_x], self.rank[root_x]))
                self.parent[root_x] = root_y
            elif self.rank[root_x] > self.rank[root_y]:
                self.history.append((root_y, self.parent[root_y], self.rank[root_y]))
                self.parent[root_y] = root_x
            else:
                self.history.append((root_y, self.parent[root_y], self.rank[root_y]))
                self.history.append((root_x, self.parent[root_x], self.rank[root_x]))
                self.parent[root_y] = root_x
                self.rank[root_x] += 1

            return True

        def undo(self) -> None:
            """撤销上一次 union 操作"""
            if not self.history:
                return

            node, parent, rank = self.history.pop()
            self.parent[node] = parent
            self.rank[node] = rank
    ```

=== "带权并查集"

    在边上维护权值，用于处理带权关系的问题。

    ```python
    class WeightedUnionFind:
        def __init__(self, n: int):
            self.parent = list(range(n))
            self.weight = [0] * n  # weight[i] 表示 i 到 parent[i] 的权值

        def find(self, x: int) -> int:
            if self.parent[x] != x:
                root = self.find(self.parent[x])
                # 更新权值：累加路径上的权值
                self.weight[x] += self.weight[self.parent[x]]
                self.parent[x] = root
            return self.parent[x]

        def union(self, x: int, y: int, w: int) -> bool:
            """合并 x 和 y，x 到 y 的权值为 w"""
            root_x, root_y = self.find(x), self.find(y)
            if root_x == root_y:
                return False

            # 计算 root_x 到 root_y 的权值
            self.parent[root_x] = root_y
            self.weight[root_x] = self.weight[y] - self.weight[x] + w
            return True

        def diff(self, x: int, y: int) -> int:
            """返回 x 到 y 的权值差"""
            if self.find(x) != self.find(y):
                return float('inf')  # 不在同一集合
            return self.weight[x] - self.weight[y]
    ```

=== "动态添加节点"

    支持动态添加新节点到并查集。

    ```python
    class DynamicUnionFind:
        def __init__(self):
            self.parent = {}
            self.rank = {}

        def add(self, x: int) -> None:
            """添加新节点"""
            if x not in self.parent:
                self.parent[x] = x
                self.rank[x] = 0

        def find(self, x: int) -> int:
            self.add(x)  # 确保节点存在
            if self.parent[x] != x:
                self.parent[x] = self.find(self.parent[x])
            return self.parent[x]

        def union(self, x: int, y: int) -> bool:
            self.add(x)
            self.add(y)
            root_x, root_y = self.find(x), self.find(y)
            if root_x == root_y:
                return False

            if self.rank[root_x] < self.rank[root_y]:
                self.parent[root_x] = root_y
            elif self.rank[root_x] > self.rank[root_y]:
                self.parent[root_y] = root_x
            else:
                self.parent[root_y] = root_x
                self.rank[root_x] += 1

            return True
    ```

## 典型应用场景

=== "判断图的连通性"

    ```python
    def count_components(n: int, edges: list[list[int]]) -> int:
        """统计无向图中的连通分量数量"""
        uf = UnionFind(n)

        for u, v in edges:
            uf.union(u, v)

        return uf.get_count()
    ```

=== "检测环"

    ```python
    def has_cycle(n: int, edges: list[list[int]]) -> bool:
        """判断无向图是否有环"""
        uf = UnionFind(n)

        for u, v in edges:
            # 如果两个节点已经连通，再添加边会形成环
            if uf.is_connected(u, v):
                return True
            uf.union(u, v)

        return False
    ```

=== "最小生成树（Kruskal 算法）"

    ```python
    def minimum_spanning_tree(n: int, edges: list[tuple[int, int, int]]) -> int:
        """
        计算最小生成树的权值和
        edges: [(u, v, weight), ...]
        """
        # 按权值排序
        edges.sort(key=lambda x: x[2])

        uf = UnionFind(n)
        total_weight = 0
        edge_count = 0

        for u, v, weight in edges:
            # 如果两个节点不连通，添加这条边
            if uf.union(u, v):
                total_weight += weight
                edge_count += 1

                # 最小生成树有 n-1 条边
                if edge_count == n - 1:
                    break

        return total_weight
    ```

=== "账户合并"

    ```python
    def accounts_merge(accounts: list[list[str]]) -> list[list[str]]:
        """
        合并具有相同邮箱的账户
        LeetCode 721. 账户合并
        """
        from collections import defaultdict

        uf = UnionFind(len(accounts))
        email_to_id = {}  # 邮箱 -> 账户ID

        # 建立邮箱到账户的映射
        for i, account in enumerate(accounts):
            for email in account[1:]:
                if email in email_to_id:
                    uf.union(i, email_to_id[email])
                else:
                    email_to_id[email] = i

        # 合并账户
        merged = defaultdict(set)
        for email, account_id in email_to_id.items():
            root = uf.find(account_id)
            merged[root].add(email)

        # 构建结果
        result = []
        for account_id, emails in merged.items():
            name = accounts[account_id][0]
            result.append([name] + sorted(emails))

        return result
    ```

## 经典题目

=== "基础题"

    - [547. 省份数量](https://leetcode.cn/problems/number-of-provinces/) — 基本连通性
    - [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/) — 二维并查集
    - [684. 冗余连接](https://leetcode.cn/problems/redundant-connection/) — 检测环
    - [990. 等式方程的可满足性](https://leetcode.cn/problems/satisfiability-of-equality-equations/) — 关系判断

=== "进阶题"

    - [721. 账户合并](https://leetcode.cn/problems/accounts-merge/) — 字符串映射
    - [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/) — 动态合并
    - [1319. 连通网络的操作次数](https://leetcode.cn/problems/number-of-operations-to-make-network-connected/) — 连通分量
    - [1584. 连接所有点的最小费用](https://leetcode.cn/problems/min-cost-to-connect-all-points/) — 最小生成树

=== "高级题"

    - [1579. 保证图可完全遍历](https://leetcode.cn/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/) — 双并查集
    - [952. 按公因数计算最大组件大小](https://leetcode.cn/problems/largest-component-size-by-common-factor/) — 数学 + 并查集
    - [399. 除法求值](https://leetcode.cn/problems/evaluate-division/) — 带权并查集
    - [1697. 检查边长度限制的路径是否存在](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/) — 离线查询

## 并查集 VS 其他数据结构

| 场景         | 推荐数据结构    | 原因                             |
| ------------ | --------------- | -------------------------------- |
| 动态连通性   | **并查集**      | 近乎 $O(1)$ 的查询和合并         |
| 静态连通性   | DFS/BFS         | 一次性计算，无需动态维护         |
| 最短路径     | Dijkstra/BFS    | 并查集只能判断连通性，不能求距离 |
| 最小生成树   | **并查集**      | Kruskal 算法的核心               |
| 强连通分量   | Tarjan/Kosaraju | 有向图的强连通性                 |
| 等价关系判断 | **并查集**      | 天然支持传递性                   |

## 实现要点

!!! warning "常见陷阱"

    1. **忘记路径压缩**：导致树退化成链表，时间复杂度退化到 $O(n)$
    2. **union 时未判断是否已连通**：可能导致集合数量统计错误
    3. **索引越界**：确保节点编号在 `[0, n)` 范围内
    4. **递归深度过大**：对于极端情况，迭代版本的 `find` 更安全

!!! tip "优化建议"

    - **路径压缩 + 按秩合并** 是标准组合，时间复杂度接近 $O(1)$
    - 如果需要获取集合大小，使用 **按大小合并** 代替按秩合并
    - 对于带权问题，使用 **带权并查集**
    - 对于需要撤销的场景，使用 **可持久化并查集**

## 总结

- **核心思想**：用树结构维护集合关系，每个集合用根节点代表
- **关键优化**：路径压缩 + 按秩合并，使操作接近 $O(1)$
- **典型应用**：连通性判断、环检测、最小生成树、等价关系
- **变体扩展**：带权并查集、可持久化并查集、动态并查集
- **实现要点**：注意路径压缩、集合数量维护、索引范围检查
