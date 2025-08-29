## 目录

- [目录](#目录)
- [树的基本性质](#树的基本性质)
- [二叉树类型](#二叉树类型)
  - [二叉树](#二叉树)
    - [完美二叉树（Perfect Binary Tree）](#完美二叉树perfect-binary-tree)
    - [完全二叉树（Complete Binary Tree）](#完全二叉树complete-binary-tree)
    - [满二叉树（Full Binary Tree）](#满二叉树full-binary-tree)
    - [平衡搜索树](#平衡搜索树)
      - [二叉搜索树（Binary Search Tree）](#二叉搜索树binary-search-tree)
      - [平衡二叉搜索树（AVL Tree）](#平衡二叉搜索树avl-tree)
      - [红黑树（Red-Black Tree）](#红黑树red-black-tree)
  - [跳表](#跳表)
    - [简介与实现原理](#简介与实现原理)
    - [跳表 VS 平衡二叉树（红黑树）](#跳表-vs-平衡二叉树红黑树)
  - [多叉树](#多叉树)
    - [B 树（B-Tree）](#b-树b-tree)
    - [B+ 树](#b-树)
  - [平衡树进化之路](#平衡树进化之路)
- [二叉树的遍历](#二叉树的遍历)
  - [递归解题模板](#递归解题模板)
  - [迭代解题模板](#迭代解题模板)
- [LeetCode 经典题目](#leetcode-经典题目)

树用来模拟现实世界中的层级关系，是数据库索引、文件系统等的基础。算法面试的绝对主角

必练操作：

- 二叉树的遍历：前序、中序、后序的递归与非递归（用栈）实现，以及层序遍历（用队列）实现。
- 特殊二叉树：
  - 二叉搜索树 (BST)：掌握其定义（左<中<右）和增、删、查操作。
  - 平衡二叉搜索树 (AVL Tree, Red-Black Tree)：理解其平衡的概念即可，通常不需要手写。
- 解决与树相关的问题：最大深度、判断是否对称、路径总和等。

## 树的基本性质

- 第 i 层最多有 $2^{i-1}$  个结点
- 如果二叉树的深度为 k，那么此二叉树**最多有 $2^k-1$ 个结点**

## 二叉树类型

### 二叉树

#### 完美二叉树（Perfect Binary Tree）

所有层的节点都被完全填满。在完美二叉树中，若树的高度为 $h$，则节点总数为 $2^{(h+1)}-1$，呈现标准的指数级关系，反映了自然界中常见的细胞分裂现象。

#### 完全二叉树（Complete Binary Tree）

仅允许最底层的节点不完全填满，且最底层的节点必须从左至右依次连续填充。二叉堆是完全二叉树的重要应用场景。

#### 满二叉树（Full Binary Tree）

非叶节点都有 0 个或两个子节点。一个有趣的规律是：叶节点的数量=内部节点数量+1

<div align="center">
  <table>
    <tr>
      <td align="center" valign="bottom" style="padding-left: 20px;">
        <img src="perfect_binary_tree.png" alt="Perfect Binary Tree" width="200" /><br />
        <sub style="font-size: 14px;">Perfect Binary Tree</sub>
      </td>
      <td align="center" valign="bottom" style="padding-right: 20px;">
        <img src="complete_binary_tree.png" alt="Complete Binary Tree" width="200" /><br />
        <sub style="font-size: 14px;">Complete Binary Tree</sub>
      </td>
      <td align="center" valign="bottom" style="padding-left: 20px;">
        <img src="full_binary_tree.png" alt="Full Binary Tree" width="200" /><br />
        <sub style="font-size: 14px;">Full Binary Tree</sub>
      </td>
    </tr>
  </table>
</div>

#### 平衡搜索树

<p align="center">
    <img src="binary_tree_special_cases.png" alt="Binary Tree Special Cases" width="60%" />
</p>

##### 二叉搜索树（Binary Search Tree）

每个节点的数值比左子树上的节点大，比右子树上的节点小。排序二叉树相比于其他数据结构的优势在于查找、插入的时间复杂度较低，为 O(logn)。二叉搜索树相比于有序数组的二分搜索，在保证搜索效率不变的情况下，插入和删除不需要移动大量的元素，从而提升整体效率。

##### 平衡二叉搜索树（AVL Tree）

> AVL 得名于两位苏联犹太人发明者：Adelson-Velsky and Landis

由于任何二叉树都可能退化为链表，所有操作的时间复杂度将从 $O(log n)$ 劣化为 $O(n)$ 。而 AVL 树则解决了二叉树退化为链表的问题，可以确保在持续添加和删除节点后，AVL 树不会退化，从而使得各种操作的时间复杂度保持在 $O(log n)$ 级别。

每个节点的左右两子树高度差都不超过 1 的二叉树。它能在 $O(log n)$ 内完成插入、查找和删除操作。

- 如果二叉树不做维护，会退化为链表，从而导致查找效率的急剧下降，因此需要不断维护二叉树尽可能的左右平衡，从而保证查找效率。
- 在 AVL 树中，任一节点左右子树高差只会有-1、0、1 三个值
- 四种旋转情况的选择条件：
  | 失衡节点的平衡因子 | 子节点的平衡因子 | 应采用的旋转方法 |
  | :----------------- | :--------------- | :--------------- |
  | $>1$（左偏树） | $>=0$ | 右旋 |
  | $>1$（左偏树） | $<0$ | 先左旋后右旋 |
  | $<-1$（右偏树） | $<=0$ | 左旋 |
  | $<-1$（右偏树） | $>0$ | 先右旋后左旋 |

<p align="center">
    <img src="avl_tree_rotation_cases.png" alt="AVL Tree Rotation" width="60%" />
</p>

<div align="center">
  <table>
    <tr>
      <td align="center" valign="bottom" style="padding-right: 20px;">
        <img src="balanced_binary_tree.png" alt="Balanced Binary Tree" width="300" /><br />
        <sub style="font-size: 14px;">Balanced Binary Tree</sub>
      </td>
      <td align="center" valign="bottom" style="padding-right: 20px;">
        <img src="imbalanced_binary_tree.png" alt="Imbalanced Binary Tree" width="300" /><br />
        <sub style="font-size: 14px;">Imbalanced Binary Tree</sub>
      </td>
    </tr>
  </table>
</div>

AVL 树的应用场景：

组织和存储大型数据，适用于**高频查找、低频增删**的场景，如用于构建数据库中的索引系统。

红黑树相比于 AVL 树来说，平衡条件更宽松，插入与删除节点所需的旋转操作更少，节点增删操作的平均效率更高。

##### 红黑树（Red-Black Tree）

**红黑树在保留 AVL 树高效查找特性的同时，对 AVL 树的再平衡性能做了优化。**

一种自平衡二叉搜索树，常用于关联数组、字典。红黑树相对于 AVL 树来说，牺牲了部分平衡性以换取插入/删除操作时少量的旋转操作，整体来说性能要优于 AVL 树，因此标准库大都采用红黑树。

红黑树的规则如下：

1. 节点是红色或黑色。
2. 根是黑色。
3. 所有叶子都是黑色（叶子是 NIL 节点）。
4. 每个红色节点必须有两个黑色的子节点。（或者说不存在两个相邻的红色节点）（或者说红色节点的父节点和子节点均是黑色的）
5. 从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点（保证平衡的关键）

<p align="center">
    <img src="red_black_tree_example.png" alt="Red Black Tree Example" width="50%" />
</p>

推论：

- 最长路径 $<=$ 最短路径\*2
- 红色的目的是触发平衡机制
- RBT 中，空节点是有效节点，因此叶节点是空节点
- AVL 对比 RBT 来看，只在搜索方面有微弱优势：由于 AVL 的平均高度比 RBT 更矮，因此理论查找速度更快，但在实际工程里，这点差距往往可以忽略（比如缓存命中、磁盘 I/O 延迟远远大于树的查找开销）

### 跳表

#### 简介与实现原理

跳表是基于**链表+随机索引**的数据结构，通过多层索引链表，实现与平衡树相同的 **O(log n)** 级别的读写操作。由美国计算机科学家 William Pugh 于 1989 年发明。

<div align="center">
  <table>
    <tr>
      <td align="center"  style="padding-right: 20px;">
        <img src="skip_list.png" alt="Skip List"  /><br />
        <br /><sub style="font-size: 14px;">Skip List</sub><br />
      </td>
      <td align="center" style="padding-right: 20px;">
        <img src="skip_list_add_element.gif" alt="Insert element to skip list"  /><br />
        <sub style="font-size: 14px;">Insert element to skip list</sub>
      </td>
    </tr>
  </table>
</div>

跳表查找时从顶部最稀疏的子序列向下进行, 直至需要查找的元素在该层两个相邻的元素中间。

写入时就像**插扑克牌**一样放在指定位置，然后通过**抛硬币**来决定是否在上层插入索引：不停地抛硬币，直到首次出现反面为止，连续抛出正面的次数，就是新节点索引的总高度。例如：
假设写入值在最底层为 Level 0，那么，

1. 第一次抛硬币结果为正面：在 Level 1 建立索引
2. 第二次抛硬币结果为正面：在 Level 2 建立索引
3. 第三次抛硬币结果为反面：停止建立索引

连续抛硬币的随机过程在数学上保证了跳表索引的随机性：

- 一个节点有 1 层索引的概率是 $1/2$。
- 一个节点有 2 层索引的概率是 $1/2 * 1/2 = 1/4$。
- 一个节点有 k 层索引的概率是 $(1/2)^k$。

这意味着，层越高，索引就越稀疏。每一层的节点数大约是下一层的一半。这种结构在宏观上就极其类似一棵完美的二叉搜索树，从而保证了其平均查找、插入、删除的时间复杂度都是 $O(log n)$。

#### 跳表 VS 平衡二叉树（红黑树）

跳表用一个极其优雅和简单的随机化思想，达到了与极其复杂的确定性算法（如红黑树）相媲美的 O(log n)性能。相比于平衡二叉树，跳表通过**链表+抛硬币**的结构具有轻量、直观、实现简单的特点。跳表在保证同等性能的前提下，在并发场景下更具优势：

- **局部性操作**：当插入或删除一个跳表节点时，只需要局部修改少数几个前驱节点的指针。
- **低锁粒度**：在多线程环境下，这意味着只需要锁定这几个前驱节点即可完成操作，其他线程可以同时在树的其他部分进行读写，并发性能极高。
- **全局性操作**：平衡二叉树的一个“旋转”操作，可能会牵扯到根节点或者树的很大一部分。在并发环境下，执行旋转可能需要锁定整个树或者一个巨大的子树，这会成为严重的性能瓶颈。

基于以上因素，跳表正在许多场景中逐步取代红黑树，如 Redis 的 zset。另外，由于跳表同样是有序的数据结构，因此在涉及快速查找、顺序遍历、范围查询、有序任务（如定时任务、倒排索引、消息队列）时都非常适合。

跳表相比于平衡二叉树严格的、确定的时间复杂度来说，其性能是基于概率的，即抛硬币时产生极其不平衡的结构，但这个概率可以低到忽略不计。跳表正是采用工程思维牺牲部分严谨性来提升效率，比特币的设计哲学与此类似：

| 设计哲学                                        | 跳表 (Skip List)                                                                                                                                                                                                                                                             | 比特币 (Bitcoin)                                                                                                                                                                                                                                                                                          |
| ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **1. 接受概率，放弃确定性**                     | 放弃了平衡二叉树那种确定性的、严格的平衡。它通过抛硬币（随机性）来构建索引层。它不保证树在任何时刻都处于完美平衡，但在概率上，它能以极高的可能性维持 O(log n)的性能。                                                                                                        | 放弃了传统分布式系统（如银行）那种确定性的、中心化的共识。它通过工作量证明（PoW，一个概率性谜题）来决定记账权。它不保证交易的瞬间最终性，但随着区块不断叠加，交易被推翻的概率会以指数级下降，趋近于零。                                                                                                   |
| **2. 牺牲次要，保全核心**                       | 核心矛盾：如何在保持高性能的同时，让实现变得极其简单，并拥有超强的并发能力？<br><br>牺牲：<br>1. 内存空间（比红黑树占用更多指针）。<br>2. 理论上的最坏情况保证（有极小概率退化）。<br><br>保全：<br>1. 实现极简（相比红黑树的旋转变色）。<br>2. 并发性能极高（锁粒度极小）。 | 核心矛盾：如何在完全没有信任的去中心化网络中，实现一个不可篡改的、安全的公共账本？<br><br>牺牲：<br>1. 效率（PoW 是巨大的能源消耗）。<br>2. 交易速度/TPS（每秒只能处理个位数交易）。<br>3. 可扩展性（区块大小受限）。<br><br>保全：<br>1. 去中心化（没有任何单点故障或控制）。<br>2. 安全性与不可篡改性。 |
| **3. 优雅的简单暴力**                           | 平衡二叉树的“旋转”是一种非常精巧、复杂的操作。而跳表的“抛硬币、加一层索引”则是一种优雅的、基于概率的“简单暴力”，它用最简单的方式解决了平衡问题。                                                                                                                             | 传统分布式共识算法（如 Paxos, Raft）非常复杂，需要节点间多轮通信投票。而比特币的“谁先算出题谁记账，其他人抄作业”的 PoW 机制，是一种极其创新的、基于算力竞争的“简单暴力”，它用最直接的方式解决了“拜占庭将军问题”。                                                                                         |
| **4. 最终结果：解决了“完美方案”解决不了的问题** | 红黑树虽然理论完美，但其复杂的实现和糟糕的并发性能，使其在很多现代高并发场景下（如 Redis）并不适用。跳表以其“不完美”的设计，成为了这些场景下的最优解。                                                                                                                       | 传统的中心化系统（如 Visa）虽然高效，但无法解决信任和审查的问题。比特币以其“笨拙”的设计，创造了人类历史上第一个无需信任、抗审查的价值存储和转移网络，解决了传统金融无法解决的问题。                                                                                                                       |

### 多叉树

#### B 树（B-Tree）

> B 树和红黑树都是由德国计算机科学家 Rudolf Bayer 在 1972 年发明

由于内存对随机访问友好，磁盘对顺序访问友好，因此红黑树可以在内存环境中高效运行，但随着数据规模的增大，磁盘管理数据的效率就大幅下降，因此出现了针对磁盘优化的 B 树。

B 树是一个一般化的二叉搜索树，每个节点可以拥有两个以上的子节点。该树是自平衡树，能保持数据有序，让数据的查找、顺序访问、插入、删除都在 $O(log n)$ 时间内完成。

一个Ｂ树通过约束所有叶子节点在相同深度来保持平衡。深度在元素添加至树的过程中缓慢增长，而整体深度极少地增长，并导致所有叶子节点与根节点距离加 1。

B 树中每一个内部节点会包含一定数量的键，键将节点的子树分开。例如，如果一个内部节点有 3 个子节点（子树），那么它就必须有两个键： *a*1 和 *a*2（图中的 7 和 16） 。左边子树的所有值都必须小于 *a*1 ，中间子树的所有值都必须在 *a*1 和*a*2 之间，右边子树的所有值都必须大于 *a*2 。

![B Tree](B_Tree.png)
B 树的生长方式：**向上生长**。B 树的所有叶子节点必须在同一层，因此 B 树永远不会在一个叶子节点下挂载新的节点来增加数据，这会破坏平衡。
B 树插入数据时，会发生分裂和提升，真正的修改发生在父节点，它接受了被提升上来的键，并更新了其子节点指针。

#### B+ 树

B+ 树通过将数据存储在叶子节点、叶子节点间通过链表相连解决了 B 树的以下问题：

- 在内部节点也存储数据，这导致 B 树的查询性能不稳定。查询性能的稳定性和可预测对于数据库的性能估算和优化至关重要
- 顺序访问和范围访问效率极低
- B+树仅在叶子节点存储数据，内部节点仅存储指针，相比 B 树能存储更多索引键，这使得**树的阶数更高，树更矮胖，最终减少了磁盘的 I/O 次数**。一个拥有千万级数据的 B+树，其高度通常只有 3 到 4 层。这意味着任何一次查询，最多只需要 3 到 4 次磁盘 I/O，性能极高。

基于 B+树的以上优点，是数据库索引（MySQL 的 InnoDB）、文件系统的常用结构。

![B+ Tree](B+_Tree.png)

### 平衡树进化之路

![The Evolution of Balanced Trees](the_evolution_of_balanced_trees.png)

## 二叉树的遍历

**以下 4 种遍历是基本功，掌握它们，绝大多数二叉树问题都能迎刃而解。**

- 层序遍历（广度优先，BFS）：按层从上到下、从左到右依次访问每一层的节点，常用**队列**实现。应用于树的深度、宽度、层、最短路径等。
- 深度优先遍历（DFS）：
  - 前序遍历（根-左-右）：**自顶向下**从根节点到叶子节点传递信息
  - 中序遍历（左-根-右）：纵向一条线从左向右扫描。最难，**最常考**。**专用于 BST**
  - 后序遍历（左-右-根）：**自底向上**从叶节点到根节点返回信息

> 在 DFS 的遍历中，前、中、后指的是根节点的位置。递归解法更容易理解和编写，因为直接调用了系统栈，而迭代解法则是在手动维护栈结构，控制节点入栈/出栈顺序。

<div align="center">
  <table>
    <tr>
      <td align="center" valign="bottom" >
        <img src="binary_tree_bfs.png" alt="Binary Tree BFS" /><br />
        <sub style="font-size: 14px;">Binary Tree BFS</sub>
      </td>
      <td align="center" valign="bottom" >
        <img src="binary_tree_dfs.png" alt="Binary Tree DFS" /><br />
        <sub style="font-size: 14px;">Binary Tree DFS</sub>
      </td>
    </tr>
  </table>
</div>

我们可以用摘葡萄来类比二叉树的遍历过程：

- 层序遍历（BFS）：像一排排地摘葡萄，每次只摘同一层的所有葡萄，摘完一层再摘下一层，直到全部摘完。
- 迭代的 DFS（如前序、中序、后序）：像顺着葡萄藤一段一段地摘，先摘左边的，遇到分叉就把右边记下来（压栈），等左边摘完再回头摘右边。每次只处理当前能摘到的一串，剩下的留到后面。

### 递归解题模板

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

### 迭代解题模板

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

## LeetCode 经典题目

- [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) — 层序遍历（队列实现）
- [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) — 中序遍历（递归、迭代、颜色标记法）
- [144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) — 前序遍历（递归、迭代、颜色标记法）
- [145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) — 后序遍历（递归、迭代、颜色标记法）
- [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) — 翻转二叉树（前序/后序递归翻转）
- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) — 求最大深度（层序遍历）
- [111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/) — 最小深度（层序遍历）
- [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) — 判断是否对称（层序遍历后判断是否为回文，或者节点两两加入队列后比较）
- [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) — 判断是否平衡（左右子树的高度差不超过 1。后序遍历）
- [98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) — 验证二叉搜索树（中序遍历后数组比较，或者递归约束滑动窗口）
- [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) — 二叉树最近公共祖先（后序递归查找）
- [235. Lowest Common Ancestor of a BST](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) — BST 最近公共祖先（结果一定是处于 p 和 q 之间的值，迭代遍历）
- [112. Path Sum](https://leetcode.com/problems/path-sum/) — 路径总和（前序遍历）
- [543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) — 二叉树的直径（后序遍历）
- [113. Path Sum II](https://leetcode.com/problems/path-sum-ii/) — 所有路径总和（前序遍历+回溯）
- [124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) — 最大路径和（后序遍历+DP）
- [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) — 根据前序和中序构造二叉树
- [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) — 根据中序和后序构造二叉树
- [297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) — 序列化与反序列化
- [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/) — 合并二叉树
