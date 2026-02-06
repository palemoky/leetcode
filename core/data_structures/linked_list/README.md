# 链表

链表（Linked List）是一种线性数据结构，由一系列节点组成，每个节点包含数据和指向下一个节点的指针。与数组不同，链表的元素在内存中不连续存储。

<figure align="center">
    <img src="linked_list.webp" alt="Linked List Overview" width="60%" />
</figure>

## 核心概念

### 虚拟头节点

虚拟头节点(Dummy Node)是链表操作中最重要的技巧之一，用于 **统一对头节点和非头节点的处理逻辑**。

```go
// 创建虚拟头节点
dummy := &ListNode{Next: head}

// 操作完成后返回新的头节点
return dummy.Next
```

**何时使用：**

- 可能需要删除头节点
- 需要在头部插入节点
- 需要修改头节点的指向

### 遍历条件

选择正确的遍历条件决定了循环能访问到哪些节点。

| 条件                  | 访问范围                   | 循环结束时                 | 适用场景                               |
| :-------------------- | :------------------------- | :------------------------- | :------------------------------------- |
| `current != nil`      | 所有节点（包括最后一个）   | `current` 为 `nil`         | 读取/处理每个节点的值                  |
| `current.Next != nil` | 除最后一个节点外的所有节点 | `current` 指向最后一个节点 | 修改节点间的链接（需要访问前一个节点） |

**安全检查：**

- `current != nil`：只需检查 `head == nil`
- `current.Next != nil`：需检查 `head == nil` 和 `head.Next == nil`

!!! Note "高效遍历"

    如果需要访问索引为 `i` 的节点，可以利用 `list.Len` 进行优化：如果 `i < len/2`，从头部开始遍历；否则，从尾部开始遍历。

### 指针操作原则

**黄金法则：先连接新指针，再断开旧指针**

```go
// ❌ 错误：先断开会丢失后续节点
node.Next = newNode
newNode.Next = oldNext  // oldNext 已经丢失！

// ✅ 正确：先保存，再连接
oldNext := node.Next
newNode.Next = oldNext
node.Next = newNode
```

**最佳实践：画图！** 在纸上画出节点和指针，模拟每一步操作。

!!! tip "插入位置的理解"

    一个长度为 `len` 的链表，有 `len + 1` 个可供插入的位置（从头部之前到尾部之后）。当题目描述"在索引 `i` 处插入"时，通常意味着将新节点插入到原索引 `i-1` 和 `i` 的节点之间，即新节点将成为新的索引 `i`。

### 边界条件清单

编写链表代码时，务必考虑以下边界情况：

- **空链表** (`head == nil`)
- **单节点链表** (`head.Next == nil`)
- **头节点操作**（使用 dummy 节点可简化）
- **尾节点操作**
- **无效输入**（如 n 超出链表长度）

## 常见算法模式

### 1. 快慢指针

快慢指针是链表中最常用的技巧，通过两个移动速度不同的指针解决问题。

**应用场景：**

=== "找中间节点"

    ```go
    func findMiddle(head *ListNode) *ListNode {
        slow, fast := head, head
        for fast != nil && fast.Next != nil {
            slow = slow.Next      // 慢指针走一步
            fast = fast.Next.Next // 快指针走两步
        }
        return slow // 慢指针指向中间节点
    }
    ```

=== "检测环"

    ```go
    func hasCycle(head *ListNode) bool {
        slow, fast := head, head
        for fast != nil && fast.Next != nil {
            slow = slow.Next
            fast = fast.Next.Next
            if slow == fast {
                return true // 快慢指针相遇，有环
            }
        }
        return false
    }
    ```

=== "删除倒数第N个节点"

    ```go
    func removeNthFromEnd(head *ListNode, n int) *ListNode {
        dummy := &ListNode{Next: head}
        slow, fast := dummy, dummy

        // 快指针先走 n+1 步
        for i := 0; i <= n; i++ {
            fast = fast.Next
        }

        // 快慢指针同步前进
        for fast != nil {
            slow = slow.Next
            fast = fast.Next
        }

        // 删除节点
        slow.Next = slow.Next.Next
        return dummy.Next
    }
    ```

### 2. 递归

递归是解决链表问题的优雅方式，核心思想是将问题分解为子问题。

**递归三要素：**

1. **定义函数功能**：明确递归函数要完成什么任务
2. **基本情况 (Base Case)**：`head == nil` 或 `head.Next == nil`
3. **递归步骤**：假设子问题已解决，处理当前节点

**示例：反转链表（递归法）**

```go
func reverseList(head *ListNode) *ListNode {
    // Base case
    if head == nil || head.Next == nil {
        return head
    }

    // 递归反转后续链表
    newHead := reverseList(head.Next)

    // 处理当前节点
    head.Next.Next = head
    head.Next = nil

    return newHead
}
```

#### 链表反转的两种方法

**迭代法** - 适合全部反转（如 [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)）

- 像翻书，逐个改变指针方向：`1→2→3` ⇒ `1←2 3` ⇒ `1←2←3`

**头插法** - 适合局部反转（如 [92. 反转链表 II](https://leetcode.cn/problems/reverse-linked-list-ii/)）

- 像抽扑克牌，依次插到头部：`prev→1→2→3→4` ⇒ `prev→2→1→3→4` ⇒ `prev→3→2→1→4`

<div align="center">
    <table>
    <tr>
        <td align="center">
        <img src="reverse_by_iteration.webp" alt="迭代法反转链表" /><br />
        <sub>迭代法反转链表</sub>
        </td>
        <td align="center">
        <img src="reverse_by_head_insert.webp" alt="头插法反转链表" /><br />
        <sub>头插法反转链表</sub>
        </td>
    </tr>
    </table>
</div>

### 3. 双指针

用于处理两个链表的问题，如合并、相交等。

**示例：合并两个有序链表**

```go
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }

    // 连接剩余节点
    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }

    return dummy.Next
}
```

### 4. 哈希表

用于需要快速查找或记录已访问节点的场景。

**应用：**

- 检测环的入口
- 删除未排序链表的重复元素
- 复制带随机指针的链表

---

## 跳表

### 实现原理

跳表是基于 **链表+随机索引** 的数据结构，通过多层索引链表，实现与平衡树相同的 **$O(log n)$** 级别的读写操作。由美国计算机科学家 William Pugh 于 1989 年发明。

<div align="center">
  <table>
    <tr>
      <td align="center"  style="vertical-align: bottom;">
        <img src="skip_list.webp" alt="Skip List"  /><br />
        <br /><sub class="img-caption">Skip List</sub><br />
      </td>
      <td align="center" style="vertical-align: bottom;">
        <img src="skip_list_add_element.webp" alt="Insert element to skip list"  /><br />
        <sub class="img-caption">Insert element to skip list</sub>
      </td>
    </tr>
  </table>
</div>

跳表查找时从顶部最稀疏的子序列向下进行, 直至需要查找的元素在该层两个相邻的元素中间。

写入时就像 **插扑克牌** 一样放在指定位置，然后通过 **抛硬币** 来决定是否在上层插入索引：不停地抛硬币，直到首次出现反面为止，连续抛出正面的次数，就是新节点索引的总高度。例如：
假设写入值在最底层为 Level 0，那么，

1. 第1次抛硬币结果为正面：在 Level 1 建立索引
2. 第2次抛硬币结果为正面：在 Level 2 建立索引
3. 第3次抛硬币结果为反面：停止建立索引

连续抛硬币的随机过程在数学上保证了跳表索引的随机性：

- 一个节点有 1 层索引的概率是 $\frac{1}{2}$。
- 一个节点有 2 层索引的概率是 $\frac{1}{2} \times \frac{1}{2} = \frac{1}{4}$。
- 一个节点有 k 层索引的概率是 $(\frac{1}{2})^k$。

这意味着，层越高，索引就越稀疏。每一层的节点数大约是下一层的一半。这种结构在宏观上就极其类似一棵完美的二叉搜索树，从而保证了其平均查找、插入、删除的时间复杂度都是 $O(log n)$。

### 跳表 VS 平衡二叉树

跳表用一个极其优雅和简单的随机化思想，达到了与极其复杂的确定性算法（如红黑树）相媲美的 $O(log n)$性能。相比于平衡二叉树，跳表通过 **链表+抛硬币** 的结构具有轻量、直观、实现简单的特点。跳表在保证同等性能的前提下，在并发场景下更具优势：

- **局部性操作**：当插入或删除一个跳表节点时，只需要局部修改少数几个前驱节点的指针。
- **低锁粒度**：在多线程环境下，这意味着只需要锁定这几个前驱节点即可完成操作，其他线程可以同时在树的其他部分进行读写，并发性能极高。
- **全局性操作**：平衡二叉树的一个“旋转”操作，可能会牵扯到根节点或者树的很大一部分。在并发环境下，执行旋转可能需要锁定整个树或者一个巨大的子树，这会成为严重的性能瓶颈。

基于以上因素，跳表正在许多场景中逐步取代红黑树，如 Redis 的 zset。另外，由于跳表同样是有序的数据结构，因此在涉及快速查找、顺序遍历、范围查询、有序任务（如定时任务、倒排索引、消息队列）时都非常适合。

跳表相比于平衡二叉树严格的、确定的时间复杂度来说，其性能是基于概率的，即抛硬币时产生极其不平衡的结构，但这个概率可以低到忽略不计。跳表正是采用工程思维牺牲部分严谨性来提升效率，比特币的设计哲学与此类似：

| 设计哲学                                        | 跳表 (Skip List)                                                                                                                                                                                                                                                             | 比特币 (Bitcoin)                                                                                                                                                                                                                                                                                          |
| ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **1. 接受概率，放弃确定性**                     | 放弃了平衡二叉树那种确定性的、严格的平衡。它通过抛硬币（随机性）来构建索引层。它不保证树在任何时刻都处于完美平衡，但在概率上，它能以极高的可能性维持 $O(log n)$的性能。                                                                                                      | 放弃了传统分布式系统（如银行）那种确定性的、中心化的共识。它通过工作量证明（PoW，一个概率性谜题）来决定记账权。它不保证交易的瞬间最终性，但随着区块不断叠加，交易被推翻的概率会以指数级下降，趋近于零。                                                                                                   |
| **2. 牺牲次要，保全核心**                       | 核心矛盾：如何在保持高性能的同时，让实现变得极其简单，并拥有超强的并发能力？<br><br>牺牲：<br>1. 内存空间（比红黑树占用更多指针）。<br>2. 理论上的最坏情况保证（有极小概率退化）。<br><br>保全：<br>1. 实现极简（相比红黑树的旋转变色）。<br>2. 并发性能极高（锁粒度极小）。 | 核心矛盾：如何在完全没有信任的去中心化网络中，实现一个不可篡改的、安全的公共账本？<br><br>牺牲：<br>1. 效率（PoW 是巨大的能源消耗）。<br>2. 交易速度/TPS（每秒只能处理个位数交易）。<br>3. 可扩展性（区块大小受限）。<br><br>保全：<br>1. 去中心化（没有任何单点故障或控制）。<br>2. 安全性与不可篡改性。 |
| **3. 优雅的简单暴力**                           | 平衡二叉树的“旋转”是一种非常精巧、复杂的操作。而跳表的“抛硬币、加一层索引”则是一种优雅的、基于概率的“简单暴力”，它用最简单的方式解决了平衡问题。                                                                                                                             | 传统分布式共识算法（如 Paxos, Raft）非常复杂，需要节点间多轮通信投票。而比特币的“谁先算出题谁记账，其他人抄作业”的 PoW 机制，是一种极其创新的、基于算力竞争的“简单暴力”，它用最直接的方式解决了“拜占庭将军问题”。                                                                                         |
| **4. 最终结果：解决了“完美方案”解决不了的问题** | 红黑树虽然理论完美，但其复杂的实现和糟糕的并发性能，使其在很多现代高并发场景下（如 Redis）并不适用。跳表以其“不完美”的设计，成为了这些场景下的最优解。                                                                                                                       | 传统的中心化系统（如 Visa）虽然高效，但无法解决信任和审查的问题。比特币以其“笨拙”的设计，创造了人类历史上第一个无需信任、抗审查的价值存储和转移网络，解决了传统金融无法解决的问题。                                                                                                                       |

### 跳表 VS 堆

堆适合最大/最小，Top-1 或 Top-k 很小的场景，而跳表能同时满足 Top-k、范围查询、rank 查询、动态更新，比堆更灵活高效。以取前 100 个元素为例，在堆中，需要弹出 100 次或维护大小为 100 的堆，但跳表只需要在起始位置向前顺序访问链表的 100 个元素即可。

---

## 经典题目

=== "基础题"

    - [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/) — 迭代/递归
    - [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/) — 双指针
    - [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/) — 遍历
    - [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/) — 快慢指针
    - [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/) — 快慢指针

=== "进阶题"

    - [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) — 快慢指针 + dummy
    - [92. 反转链表 II](https://leetcode.cn/problems/reverse-linked-list-ii/) — 部分反转
    - [143. 重排链表](https://leetcode.cn/problems/reorder-list/) — 找中点 + 反转 + 合并
    - [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/) — 双指针
    - [234. 回文链表](https://leetcode.cn/problems/palindrome-linked-list/) — 快慢指针 + 反转
    - [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/) — 模拟加法

=== "高级题"

    - [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/) — 分治/堆
    - [25. K 个一组翻转链表](https://leetcode.cn/problems/reverse-nodes-in-k-group/) — 分组反转
    - [138. 复制带随机指针的链表](https://leetcode.cn/problems/copy-list-with-random-pointer/) — 哈希表/原地修改
    - [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/) — 快慢指针数学性质

## 总结

- **虚拟头节点**：简化头节点操作，避免特殊处理
- **快慢指针**：找中点、检测环、删除倒数第N个节点
- **递归思维**：将问题分解为子问题，优雅解决反转、合并等问题
- **画图调试**：链表问题的最佳调试方法，将抽象的指针操作可视化
- **边界检查**：空链表、单节点、头尾节点操作都要考虑
