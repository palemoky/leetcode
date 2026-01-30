# 队列

队列（Queue）是一种 **先进先出**（FIFO, First In First Out）的线性数据结构。它就像生活中排队买票，先来的人先买，后来的人排在队尾。

## 核心特性

队列的操作分别在两端进行：

1.  **Enqueue（入队）**：将元素添加到队尾。
2.  **Dequeue（出队）**：移除并返回队头元素。
3.  **Peek/Front（查看队头）**：返回队头元素但不移除。
4.  **IsEmpty（判空）**：检查队列是否为空。

## 复杂度分析

| 操作        | 时间复杂度 | 说明                                                    |
| :---------- | :--------- | :------------------------------------------------------ |
| **Enqueue** | $O(1)$     | 均摊复杂度（动态数组扩容时为 $O(n)$）                   |
| **Dequeue** | $O(1)$     | 链表实现为 $O(1)$；切片实现通常为 $O(1)$ 但可能触发缩容 |
| **Peek**    | $O(1)$     | 直接访问队头                                            |
| **Search**  | $O(n)$     | 需要遍历队列                                            |

## 实现方式 (Go)

在 Go 语言算法题中，通常使用切片（Slice）模拟队列。

=== "基础实现"

    ```go
    // 使用切片模拟队列
    queue := []int{}

    // 1. Enqueue - 入队
    queue = append(queue, 10)
    queue = append(queue, 20)

    // 2. Peek - 查看队头
    front := queue[0]
    // front = 10

    // 3. Dequeue - 出队
    val := queue[0]
    queue = queue[1:]
    // val = 10, queue = [20]

    // 4. IsEmpty - 判空
    if len(queue) == 0 {
        fmt.Println("Queue is empty")
    }
    ```

    !!! Warning "内存泄露注意"

        使用 `queue = queue[1:]` 出队后，底层数组的前部空间不会立即被回收。在长期运行的服务中，建议使用链表或循环数组，或者定期重新分配切片。但在算法题中，这种方式通常是最简便的。

=== "双端队列 (Deque)"

    双端队列（Double-Ended Queue）允许在两端进行入队和出队操作。Go 中没有内置 Deque，通常也用切片模拟，或者自行封装链表。

    ```go
    // 队头入队
    queue = append([]int{val}, queue...) // O(n) 开销较大慎用

    // 队头出队
    val = queue[0]
    queue = queue[1:]

    // 队尾入队
    queue = append(queue, val)

    // 队尾出队
    val = queue[len(queue)-1]
    queue = queue[:len(queue)-1]
    ```

## 变种

### 1. 循环队列 (Circular Queue)

为了解决普通数组队列出队后空间浪费的问题，使用固定大小的数组和头尾指针，逻辑上形成环状。

- **经典题目**：[622. Design Circular Queue](https://leetcode.cn/problems/design-circular-queue/)

### 2. 优先队列 (Priority Queue)

元素按照优先级出队，而不是进入顺序。通常使用 **堆 (Heap)** 实现。

- **详见**：[堆 (Heap)](../tree/README.md#堆heap)

### 3. 单调队列 (Monotonic Queue)

队列内元素保持单调递增或递减，常用于解决滑动窗口最大值问题。

- **经典题目**：[239. Sliding Window Maximum](https://leetcode.cn/problems/sliding-window-maximum/)

## 应用场景

1.  **广度优先搜索 (BFS)**：层序遍历树或图，寻找最短路径。
2.  **任务调度**：操作系统中的进程调度、打印机任务队列。
3.  **消息队列**：系统解耦、削峰填谷（如 Kafka, RabbitMQ）。
4.  **缓存**：LRU 缓存（最近最少使用）通常结合哈希表和双向链表（类似队列）实现。

## 经典题目

=== "基础题"

    - [225. Implement Stack using Queues](https://leetcode.cn/problems/implement-stack-using-queues/) — 用队列实现栈
    - [387. First Unique Character in a String](https://leetcode.cn/problems/first-unique-character-in-a-string/) — 字符串中的第一个唯一字符
    - [933. Number of Recent Calls](https://leetcode.cn/problems/number-of-recent-calls/) — 最近的请求次数

=== "进阶题 (BFS)"

    - [102. Binary Tree Level Order Traversal](https://leetcode.cn/problems/binary-tree-level-order-traversal/) — 二叉树的层序遍历
    - [200. Number of Islands](https://leetcode.cn/problems/number-of-islands/) — 岛屿数量 (BFS解法)
    - [279. Perfect Squares](https://leetcode.cn/problems/perfect-squares/) — 完全平方数 (BFS求最短路径)
    - [752. Open the Lock](https://leetcode.cn/problems/open-the-lock/) — 打开转盘锁

=== "高级题"

    - [239. Sliding Window Maximum](https://leetcode.cn/problems/sliding-window-maximum/) — 滑动窗口最大值 (单调队列)
    - [622. Design Circular Queue](https://leetcode.cn/problems/design-circular-queue/) — 设计循环队列
    - [641. Design Circular Deque](https://leetcode.cn/problems/design-circular-deque/) — 设计循环双端队列
    - [862. Shortest Subarray with Sum at Least K](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/) — 和至少为 K 的最短子数组

## 总结

- **FIFO 原则**：先进先出，BFS 的基石。
- **切片注意事项**：在 Go 中用切片模拟队列简单高效，但要注意 `append([]int{val}, queue...)` 是 $O(n)$ 操作，尽量避免在切片头部插入，或者使用双指针/循环数组/链表优化。
