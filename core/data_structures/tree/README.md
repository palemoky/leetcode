# 树

## 基本概念

树是一种层次化的数据结构，在现实生活和计算机科学中都有广泛应用:

- **现实生活**: 组织架构、家族谱系、生物分类等
- **计算机领域**: 文件系统、DOM、抽象语法树（AST）、数据库索引（B/B+树）、路由表（前缀树/Trie）、决策树（机器学习）等

树的遍历是许多算法的基础，如动态规划（DFS 关注节点状态）、回溯（关注决策路径）等。掌握树的操作，对理解算法设计至关重要。

### 树的性质

- 第 i 层最多有 $2^{i-1}$ 个结点
- 如果二叉树的深度为 k，那么此二叉树 **最多有 $2^k-1$ 个结点**

### 二叉树分类

<div align="center">
  <table>
    <tr>
      <td align="center" style="vertical-align: bottom;" style="padding-left: 20px;">
        <img src="perfect_binary_tree.webp" alt="Perfect Binary Tree" width="200" /><br />
        <sub class="img-caption">Perfect Binary Tree</sub>
      </td>
      <td align="center" style="vertical-align: bottom;" style="padding-right: 20px;">
        <img src="complete_binary_tree.webp" alt="Complete Binary Tree" width="200" /><br />
        <sub class="img-caption">Complete Binary Tree</sub>
      </td>
      <td align="center" style="vertical-align: bottom;" style="padding-left: 20px;">
        <img src="full_binary_tree.webp" alt="Full Binary Tree" width="200" /><br />
        <sub class="img-caption">Full Binary Tree</sub>
      </td>
    </tr>
  </table>
</div>

#### 完美二叉树

所有层的节点都被完全填满。在完美二叉树中，若树的高度为 $h$，则节点总数为 $2^{(h+1)}-1$，呈现标准的指数级关系，反映了自然界中常见的细胞分裂现象。

#### 完全二叉树

仅允许最底层的节点不完全填满，且最底层的节点必须从左至右依次连续填充。二叉堆是完全二叉树的重要应用场景。

由于完全二叉树的性质，使其在数组中可以被连续存储（数组索引可按 1-based 或 0-based 约定）。

![Complete Binary Tree Array Index](complete_binary_tree_array_index.webp){ align=right width=35% }

| 节点关系   | 1-based 索引                             | 0-based 索引                               |
| ---------- | ---------------------------------------- | ------------------------------------------ |
| **父节点** | $\left\lfloor \frac{i}{2} \right\rfloor$ | $\left\lfloor \frac{i-1}{2} \right\rfloor$ |
| **左孩子** | $2 \times i$                             | $2 \times i + 1$                           |
| **右孩子** | $2 \times i + 1$                         | $2 \times i + 2$                           |

#### 满二叉树

非叶节点都有 0 个或两个子节点。一个有趣的规律是：叶节点的数量=内部节点数量+1

## 遍历方式

二叉树的遍历分为广度优先（BFS）和深度优先（DFS）两种，其中DFS又分为前、中、后序遍历。这4种遍历方式极其重要，之后的动态规划、回溯等都依赖于此。

<div align="center">
  <table>
    <tr>
      <td align="center" style="vertical-align: bottom;" >
        <img src="binary_tree_bfs.webp" alt="Binary Tree BFS" /><br />
        <sub class="img-caption">Binary Tree BFS</sub>
      </td>
      <td align="center" style="vertical-align: bottom;" >
        <img src="binary_tree_dfs.webp" alt="Binary Tree DFS" /><br />
        <sub class="img-caption">Binary Tree DFS</sub>
      </td>
    </tr>
  </table>
</div>

| 遍历方式                | 顺序     | 特点                                                 | 实现    | 应用场景                       |
| ----------------------- | -------- | ---------------------------------------------------- | ------- | ------------------------------ |
| **广度优先遍历（BFS）** | 层序遍历 | 按层从上到下、从左到右依次访问每一层的节点           | 队列    | 树的深度、宽度、层、最短路径   |
| **前序遍历（DFS）**     | 根-左-右 | **自顶向下** 从根节点到叶子节点传递信息              | 递归/栈 | 复制树、序列化、路径查找       |
| **中序遍历（DFS）**     | 左-根-右 | 纵向一条线从左向右扫描。最难，最常考，**专用于 BST** | 递归/栈 | BST 有序输出、验证 BST         |
| **后序遍历（DFS）**     | 左-右-根 | **自底向上** 从叶节点到根节点返回信息                | 递归/栈 | 删除树、计算高度、最近公共祖先 |

!!! Note

    在 DFS 的遍历中，**前、中、后指的是根节点的位置**。递归解法更容易理解和编写，因为直接调用了系统栈，而迭代解法则是在手动维护栈结构，控制节点入栈/出栈顺序。

    快速排序就是前序遍历，归并排序就是后序遍历。

<div align="center">
  <table>
    <tr>
      <td align="center">
        <img src="bfs_traversal.webp" alt="BFS 遍历" /><br />
        <sub>BFS: 横向扩散，逐层遍历</sub>
      </td>
      <td align="center">
        <img src="dfs_traversal.webp" alt="DFS 前序遍历" /><br />
        <sub>DFS (前序遍历): 纵向深入，一条路走到底</sub>
      </td>
    </tr>
  </table>
</div>

- **BFS**: 从根节点向四周扩散，逐层遍历，类似水中的波纹，这使得 BFS 适合求最短路径
- **DFS**: 纵向深入，沿着一条路径走到底，再回溯探索其他路径

BFS 与 DFS 的区别：

- **最短路径**: 在 **无权图** 中，BFS 找到的路径一定是最短的(边数最少)
- **空间复杂度**: BFS 的空间复杂度为 $O(w)$(w 为最大宽度)，DFS 为 $O(h)$(h 为树高)。对于 **完全二叉树**，BFS 的空间开销远大于 DFS

=== "递归解题模板"

    递归解法可以分为两类思路：

    - **遍历思路**：遍历整棵树，在遍历过程中更新外部变量或执行操作（类似回溯算法的思维）
    - **分解思路**：将问题分解为子问题，通过子问题的解推导出原问题的解（类似动态规划的思维）

    !!! Abstract "递归 -> 下地下室取东西"

        如果我们把递归的过程想象成去地下室每层房间取东西，那么就有三种情况：

        1. 下去时打包好带上 → 前序遍历
        2. 下去时打包好，上来时带上 → 中序遍历
        3. 触底返回时打包好带上 → 后序遍历

    ```go
    func traversal(root *TreeNode) []int {
        if root == nil {
            return []int{}
        }

        // Preorder: 根 -> 左 -> 右
        // nums := []int{root.Val}
        // nums = append(nums, traversal(root.Left)...)
        // nums = append(nums, traversal(root.Right)...)

        // Inorder: 左 -> 根 -> 右
        nums := traversal(root.Left)
        nums = append(nums, root.Val)
        nums = append(nums, traversal(root.Right)...)

        // Postorder: 左 -> 右 -> 根
        // nums := traversal(root.Left)
        // nums = append(nums, traversal(root.Right)...)
        // nums = append(nums, root.Val)

        return nums
    }
    ```

    !!! Note

        如果要倒序打印单链表上的所有节点的值，可以采用后序递归操作。

        ```go
        func reversePrint(head *ListNode) []int {
            if head == nil {
                return []int{}
            }

            return append(reversePrint(head.Next), head.Val)
        }
        ```

=== "迭代解题模板"

    颜色标记法通过调整入栈顺序即可通杀前、中、后序遍历

    ```go
    const (
        WHITE = 0
        BLACK = 1
    )

    type ColorNode struct {
        Color int
        Node  *TreeNode
    }

    func iterative(root *TreeNode) []int {
        nums := []int{}
        stack := []ColorNode{{WHITE, root}}
        for len(stack) > 0 {
            cn := stack[len(stack)-1] // cn is colorNode
            stack = stack[:len(stack)-1]

            if cn.Node == nil {
                continue
            }

            if cn.Color == WHITE {
                // 前序的压入顺序：右-左-根（BLACK）
                // stack = append(stack, ColorNode{WHITE, cn.Node.Right})
                // stack = append(stack, ColorNode{WHITE, cn.Node.Left})
                // stack = append(stack, ColorNode{BLACK, cn.Node})

                // 中序的压入顺序：右-根（BLACK）-左
                stack = append(stack, ColorNode{WHITE, cn.Node.Right})
                stack = append(stack, ColorNode{BLACK, cn.Node})
                stack = append(stack, ColorNode{WHITE, cn.Node.Left})

                // 后序的压入顺序：根（BLACK）-右-左
                // stack = append(stack, ColorNode{BLACK, cn.Node})
                // stack = append(stack, ColorNode{WHITE, cn.Node.Right})
                // stack = append(stack, ColorNode{WHITE, cn.Node.Left})
            } else {
                nums = append(nums, cn.Node.Val)
            }
        }

        return nums
    }
    ```

## 特殊树结构

### 堆（Heap）

堆是一种满足堆性质的完全二叉树，常用来实现优先队列。堆分为小顶堆和大顶堆两种类型。

- 小顶堆（min-heap）：任意父节点的值小于等于其子节点的值。
- 大顶堆（max-heap）：任意父节点的值大于等于其子节点的值。

| 堆类型 | 1-based 索引                         | 0-based 索引                           |
| ------ | ------------------------------------ | -------------------------------------- |
| 小顶堆 | `a[i] <= a[2*i] && a[i] <= a[2*i+1]` | `a[i] <= a[2*i+1] && a[i] <= a[2*i+2]` |
| 大顶堆 | `a[i] >= a[2*i] && a[i] >= a[2*i+1]` | `a[i] >= a[2*i+1] && a[i] >= a[2*i+2]` |

<figure>
    <img src="min_heap_and_max_heap.webp" alt="Min Heap and Max Heap" width="65%" />
</figure>

常见操作的时间复杂度：

[![Heap Sort](heap_sort_video_thumbnail.webp){ align=right width=20% }](https://www.bilibili.com/video/BV1HYtseiEQ8)

- push / pop（上浮或下沉）: $O(log n)$。建堆时，需要不断比较插入元素与插入位置元素的大小，插入后还要对子节点沉浮以保证堆的有序性。
- peek / top: $O(1)$
- size / isEmpty: $O(1)$

应用场景：

- 优先队列（任务调度、事件驱动、作业排队等）：按优先级快速取出最小/最大元素。
- 路径与图算法：Dijkstra、A\* 等需要按最小代价取节点的算法（通常使用小顶堆）。
- Top‑k 问题：遍历元素并维护一个大小为 k 的小顶堆，堆顶即为第 k 大元素；时间复杂度 $O(n log k)$，空间 $O(k)$。
- 流式数据 / 实时排行榜：数据持续到达时，用小顶堆维护前 k 个最大值（当堆大小超过 k 时弹出堆顶）。
- 合并 k 个有序流 / 滑动窗口：合并 k 个有序数组/链表常用小顶堆；滑动窗口可以用堆维护窗口内的最大/最小值，但要注意堆不擅长做任意位置删除，常配合延迟删除或索引映射。
- 堆排序（Heap Sort）：先建堆（$O(n)$）再持续 pop 得到有序序列，总体 $O(n log n)$，可原地实现。

小顶堆 vs 大顶堆：

- **小顶堆**：快速获取最小值，或维护前 k 个最大值（堆顶为第 k 大元素）
- **大顶堆**：快速获取最大值，或维护前 k 个最小值（堆顶为第 k 小元素）

实现注意事项：

- 数组索引基（0-based / 1-based）会影响父/子索引计算，文中已列出两种约定的公式。
- 堆原地操作会修改输入数组；若需要保留原数组，请先拷贝。
- 堆不擅长高效删除任意元素；在需要频繁随机删除的场景，可考虑配合哈希索引（记录元素位置）或使用延迟删除（标记失效元素，遇到堆顶时再清理）。

经典题目

- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) — 数组中第 K 个最大元素（Top-k / 堆）
- [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) — 前 K 个高频元素（哈希 + 堆）

### 二叉搜索树（BST）

每个节点的数值比左子树上的节点大，比右子树上的节点小。排序二叉树相比于其他数据结构的优势在于查找、插入的时间复杂度较低，为 $O(logn)$。二叉搜索树相比于有序数组的二分搜索，在保证搜索效率不变的情况下，插入和删除不需要移动大量的元素，从而提升整体效率。

!!! Tip

    为了避免 `left+right` 导致 int 溢出，通常使用 `left+(right-left)/2` 的方式计算 mid。

    由于主流编程语言 **整数除法向下取整**，因此偶数长度的数组，mid 会 **居中偏左**，奇数长度的数组则在正中间。

BST 有两种不同的写法，处理细节也不同：

|           写法           | 初始化 right  |    循环条件     |    更新 right     |
| :----------------------: | :-----------: | :-------------: | :---------------: |
|  闭区间 `[left, right]`  | `len(nums)-1` | `left <= right` | `right = mid - 1` |
| 左闭右开 `[left, right)` |  `len(nums)`  | `left < right`  |   `right = mid`   |

=== "闭区间 [left, right]"

    ```go
    left, right := 0, len(nums)-1
    for left <= right {  // ← 闭区间必须用 <=，以确保所有元素都被搜索
        mid := left + (right-left)/2 // ← 避免 int 溢出
        if nums[mid] < target {
            left = mid + 1   // 排除 mid 并查找右侧区间
        } else if nums[mid] > target {
            right = mid - 1  // 排除 mid 并查找左侧区间
        } else {
            return mid
        }
    }
    ```

=== "左闭右开 [left, right)"

    ```go
    left, right := 0, len(nums)  // ← right 初始化为 len(nums)
    for left < right {  // ← 用 <
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid  // ← 不减1，因为 right 本身不包含
        } else {
            return mid
        }
    }
    ```

以上是基础的二分查找，当找到目标值后立即返回。但在实际应用中，**数组可能包含重复元素**，此时我们需要找到目标值的 **左边界**（第一次出现的位置）或 **右边界**（最后一次出现的位置）。这两种变体只需在找到目标后，继续向左或向右收缩搜索区间即可实现：

!!! Note

    同样的，BST 查找左、右边界也有闭区间和开区间两种写法，以下为闭区间的写法。

=== "查找左侧边界"

    ```go
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            right = mid - 1  // ← 关键：找到后继续向左收缩
        }
    }

    // 检查越界和是否找到
    if left >= len(nums) || nums[left] != target {
        return -1
    }

    return left
    ```

=== "查找右侧边界"

    ```go
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1  // ← 关键：找到后继续向右收缩
        }
    }

    // 检查越界和是否找到
    if right < 0 || nums[right] != target {
        return -1
    }

    return right
    ```

**左右边界的应用场景：**

- 在有序数组中查找元素的第一次和最后一次出现位置
- 统计某个元素在有序数组中出现的次数：`count = rightBound - leftBound + 1`
- 经典题目：[34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

### 平衡二叉树

<figure>
    <img src="binary_tree_special_cases.webp" alt="Binary Tree Special Cases" width="60%" />
</figure>

#### AVL树

> AVL 得名于两位苏联犹太人发明者：Adelson-Velsky and Landis

由于任何二叉树都可能退化为链表，所有操作的时间复杂度将从 $O(log n)$ 劣化为 $O(n)$ 。而 AVL 树则解决了二叉树退化为链表的问题，可以确保在持续添加和删除节点后，AVL 树不会退化，从而使得各种操作的时间复杂度保持在 $O(log n)$ 级别。

每个节点的左右两子树高度差都不超过 1 的二叉树。它能在 $O(log n)$ 内完成插入、查找和删除操作。

- 如果二叉树不做维护，会退化为链表，从而导致查找效率的急剧下降，因此需要不断维护二叉树尽可能的左右平衡，从而保证查找效率。
- 在 AVL 树中，任一节点左右子树高差只会有-1、0、1 三个值
- 四种旋转情况的选择条件：

<div style="margin-left: 2em;" markdown="1">

| 失衡节点的平衡因子 | 子节点的平衡因子 | 应采用的旋转方法 |
| :----------------- | :--------------- | :--------------- |
| $\gt 1$（左偏树）  | $\geq 0$         | 右旋             |
| $\gt 1$（左偏树）  | $\lt 0$          | 先左旋后右旋     |
| $\lt -1$（右偏树） | $\leq 0$         | 左旋             |
| $\lt -1$（右偏树） | $\gt 0$          | 先右旋后左旋     |

</div>

<figure>
    <img src="avl_tree_rotation_cases.webp" alt="AVL Tree Rotation" width="60%" />
</figure>

<div align="center">
  <table>
    <tr>
      <td align="center" style="vertical-align: bottom;" style="padding-right: 20px;">
        <img src="balanced_binary_tree.webp" alt="Balanced Binary Tree" width="300" /><br />
        <sub class="img-caption">Balanced Binary Tree</sub>
      </td>
      <td align="center" style="vertical-align: bottom;" style="padding-right: 20px;">
        <img src="imbalanced_binary_tree.webp" alt="Imbalanced Binary Tree" width="300" /><br />
        <sub class="img-caption">Imbalanced Binary Tree</sub>
      </td>
    </tr>
  </table>
</div>

AVL 树的应用场景：

组织和存储大型数据，适用于 **高频查找、低频增删** 的场景，如用于构建数据库中的索引系统。

红黑树相比于 AVL 树来说，平衡条件更宽松，插入与删除节点所需的旋转操作更少，节点增删操作的平均效率更高。

#### 红黑树

**红黑树在保留 AVL 树高效查找特性的同时，对 AVL 树的再平衡性能做了优化。**

一种自平衡二叉搜索树，常用于关联数组、字典。红黑树相对于 AVL 树来说，牺牲了部分平衡性以换取插入/删除操作时少量的旋转操作，整体来说性能要优于 AVL 树，因此标准库大都采用红黑树。

红黑树的规则如下：

1. 节点是红色或黑色。
2. 根是黑色。
3. 所有叶子都是黑色（叶子是 NIL 节点）。
4. 每个红色节点必须有两个黑色的子节点。（或者说不存在两个相邻的红色节点）（或者说红色节点的父节点和子节点均是黑色的）
5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点（保证平衡的关键）

<figure>
    <img src="red_black_tree_example.webp" alt="Red Black Tree Example" width="50%" />
</figure>

推论：

- 最长路径 $<=$ 最短路径\*2
- 红色的目的是触发平衡机制
- RBT 中，空节点是有效节点，因此叶节点是空节点
- AVL 对比 RBT 来看，只在搜索方面有微弱优势：由于 AVL 的平均高度比 RBT 更矮，因此理论查找速度更快，但在实际工程里，这点差距往往可以忽略（比如缓存命中、磁盘 I/O 延迟远远大于树的查找开销）

### B/B+树

#### B 树

> B 树和红黑树都是由德国计算机科学家 Rudolf Bayer 在 1972 年发明

由于内存对随机访问友好，磁盘对顺序访问友好，因此红黑树可以在内存环境中高效运行，但随着数据规模的增大，磁盘管理数据的效率就大幅下降，因此出现了针对磁盘优化的 B 树。

B 树是一个一般化的二叉搜索树，每个节点可以拥有两个以上的子节点。该树是自平衡树，能保持数据有序，让数据的查找、顺序访问、插入、删除都在 $O(log n)$ 时间内完成。

一个Ｂ树通过约束所有叶子节点在相同深度来保持平衡。深度在元素添加至树的过程中缓慢增长，而整体深度极少地增长，并导致所有叶子节点与根节点距离加 1。

![B Tree](B_Tree.webp)

B 树的生长方式：**向上生长**。B 树的所有叶子节点必须在同一层，因此 B 树永远不会在一个叶子节点下挂载新的节点来增加数据，这会破坏平衡。
B 树插入数据时，会发生分裂和提升，真正的修改发生在父节点，它接受了被提升上来的键，并更新了其子节点指针。

#### B+ 树

B+ 树通过将数据存储在叶子节点、叶子节点间通过链表相连解决了 B 树的以下问题：

- 在内部节点也存储数据，这导致 B 树的查询性能不稳定。查询性能的稳定性和可预测对于数据库的性能估算和优化至关重要
- 顺序访问和范围访问效率极低
- B+树仅在叶子节点存储数据，内部节点仅存储指针，相比 B 树能存储更多索引键，这使得 **树的阶数更高，树更矮胖，最终减少了磁盘的 I/O 次数** 。一个拥有千万级数据的 B+树，其高度通常只有 3 到 4 层。这意味着任何一次查询，最多只需要 3 到 4 次磁盘 I/O，性能极高。

基于 B+树的以上优点，是数据库索引（MySQL 的 InnoDB）、文件系统的常用结构。

![B+ Tree](B+_Tree.webp)

### 前缀树（Trie）

Trie 是一种专门用于高效处理字符串集合的多叉树结构，在搜索引擎的自动补全、拼写检查、IP 路由等场景中有广泛应用。
**核心特性：**

- 前缀搜索时间复杂度：$O(m)$（m 为字符串长度）
- 适用于字符串集合的快速查找和前缀匹配

详细内容请参考：[Trie 数据结构详解](../trie/README.md)

### 平衡树进化之路

![The Evolution of Balanced Trees](the_evolution_of_balanced_trees.webp)

## 经典题目

=== "基础题"

    - [144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) — 前序遍历
    - [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) — 中序遍历
    - [145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) — 后序遍历
    - [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) — 层序遍历
    - [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) — 最大深度
    - [111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/) — 最小深度
    - [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) — 翻转二叉树
    - [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) — 对称二叉树
    - [112. Path Sum](https://leetcode.com/problems/path-sum/) — 路径总和
    - [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/) — 合并二叉树

=== "进阶题"

    - [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) — 平衡二叉树
    - [98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) — 验证二叉搜索树
    - [235. Lowest Common Ancestor of a BST](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) — 二叉搜索树的最近公共祖先
    - [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) — 二叉树的最近公共祖先
    - [113. Path Sum II](https://leetcode.com/problems/path-sum-ii/) — 路径总和 II
    - [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) — 从前序与中序遍历序列构造二叉树
    - [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) — 从中序与后序遍历序列构造二叉树
    - [543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) — 二叉树的直径

=== "高级题"

    - [124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) — 二叉树中的最大路径和
    - [297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) — 二叉树的序列化与反序列化
