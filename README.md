# LeetCode Solutions

Welcome! This repository contains my solutions to LeetCode problems, implemented in **Go**, **Python**, and **Rust**.

Each problem is organized by language and directory. Solutions include clean code, unit tests, and concise explanations where appropriate.

## Directory Structure

- 📂 `go/` — Go solutions
- 📂 `python/` — Python solutions
- 📂 `rust/` — Rust solutions

Feel free to explore, learn, and discuss.
Contributions and suggestions are always welcome!

## Must-Know Classic LeetCode Problems by Category

### 1. Arrays & Strings

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
- [189. Rotate Array](https://leetcode.com/problems/rotate-array/)
- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)
- [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)
- [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
- [125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

### 2. Linked Lists

- [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)
- [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)
- [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)
- [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)
- [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

### 3. Stacks & Queues

- [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
- [155. Min Stack](https://leetcode.com/problems/min-stack/)
- [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)
- [232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)

### 4. Hash Tables

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)
- [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)
- [146. LRU Cache](https://leetcode.com/problems/lru-cache/)

### 5. Binary Search

- [704. Binary Search](https://leetcode.com/problems/binary-search/)
- [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
- [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)
- [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

### 6. Dynamic Programming

- [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)
- [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)
- [322. Coin Change](https://leetcode.com/problems/coin-change/)
- [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)
- [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

### 7. Backtracking

- [78. Subsets](https://leetcode.com/problems/subsets/)
- [46. Permutations](https://leetcode.com/problems/permutations/)
- [39. Combination Sum](https://leetcode.com/problems/combination-sum/)
- [79. Word Search](https://leetcode.com/problems/word-search/)

### 8. Trees & Recursion

- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
- [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)
- [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)
- [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)
- [235. Lowest Common Ancestor of a BST](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

### 9. Graphs (BFS & DFS)

- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/)
- [133. Clone Graph](https://leetcode.com/problems/clone-graph/)
- [207. Course Schedule](https://leetcode.com/problems/course-schedule/)
- [127. Word Ladder](https://leetcode.com/problems/word-ladder/)

### 10. Greedy Algorithms

- [55. Jump Game](https://leetcode.com/problems/jump-game/)
- [455. Assign Cookies](https://leetcode.com/problems/assign-cookies/)
- [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

### 11. Math & Bit Manipulation

- [231. Power of Two](https://leetcode.com/problems/power-of-two/)
- [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)
- [371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)

## 常用基础算法技巧总结

在算法与数据结构领域，有一些必须掌握的基础技巧，广泛应用于各类题目和实际开发。以下是最常见的几类：

### 链表

#### 反转链表（Reverse Linked List）

- 单链表反转，递归或迭代实现。
- 典型题目：LeetCode 206

#### 合并链表（Merge Linked List）

- 使用双指针逐步比较两个链表的当前节点，将较小的节点接到结果链表后面，直到某一链表为空，再将剩余部分接到结果链表
- 递归和迭代两种实现方式都很常见
- 典型题目：[LeetCode 21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

#### 判断链表是否有环（Linked List Has Cycle）

- 有环的链表一定会让移速差为 2 倍的快慢指针相遇（以相对运动来理解快慢指针的移动，这个速度不会跳过任何节点）
- 方法：快慢指针（Floyd 判圈算法）。慢指针每次走一步，快指针每次走两步。如果链表有环，快慢指针最终会相遇；如果无环，快指针会走到链表末尾。
- 典型题目：[LeetCode 141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)

### 数组与字符串

#### 双指针（Two Pointers）

- 用于有序数组、区间统计、去重等场景。
- 典型题目：LeetCode 15、167

##### 对撞指针

从数组两端向中间移动，常用于**有序数组**中二分查找值、反转字符串等，需要注意 `left <= right`（二分查找）和 `left < right`（反转字符串）的区别

##### 快慢指针/滑动窗口

两个指针同向移动，维护一个动态的“窗口”，用于解决子数组/子字符串、链表是否有环等相关的问题（如求最长、最短、满足特定条件的子串）

##### 数组中的快慢指针

用于原地修改数组，如移除重复项、移动零等

#### 二分查找（Binary Search）

用于**有序数组**或区间，快速定位目标或判定区间最值。典型题目：LeetCode 704、35

#### 前缀和/前缀积 (Prefix Sum/Product)

核心：预计算一个数组，`prefix[i]` 存储原数组 `nums[0...i]` 的和/积。

应用：可以在 O(1) 的时间内快速查询**任意子数组的和/积**。是解决子数组和相关问题的利器。

#### 差分数组 (Difference Array)

核心：构造一个新数组 diff，`diff[i] = nums[i] - nums[i-1]`。

应用：可以在 O(1) 的时间内对原数组的**一个区间**进行统一的增减操作。

#### 排序 + 贪心

常见于区间调度、会议室安排、最小箭头射爆气球等题。

### 数学与位运算

#### GCD/LCM

常用辗转相除法和公式法等[技巧](core/math/GCD&LCM/README.md)

#### 位运算技巧

- XOR (^)：查找只出现一次/两次的数字。
- AND (&):

  - `n & (n-1)` 去掉最低位 1，常用于判断一个数是否是 2 的幂
  - `n & -n` 得到最低位的 1

- << 和 >>：高效的乘除 2 运算
- 用异或找唯一不成对的数

#### 竖式模拟

竖式模拟中当前位、进位、借位的[技巧](core/patterns/elementary_math_simulation/README.md)

### 通用算法思想

#### 哈希表（HashMap）

核心：空间换时间

应用：

- 计数器：快速统计元素频率
- 快速查找：O(1) 时间判断元素是否存在
- 去重：利用 Set 特性

#### 分治（Divide and Conquer）

核心：将大问题分解成结构相同但规模更小的子问题，递归地解决子问题，然后合并结果。

应用：归并排序、快速排序、树的大部分递归问题

#### 回溯（Backtracking）

核心：一种通过探索所有可能路径来寻找解的算法。如果发现当前路径走不通，就“回溯”到上一步，尝试其他选择。它的模板是 **“做选择 -> 递归 -> 撤销选择”**。

应用：全排列、组合、子集、N 皇后问题、迷宫求解。


---


### 4. 滑动窗口（Sliding Window）

- 动态维护区间，解决子串、子数组统计问题。
- 典型题目：LeetCode 3、76、209

### 5. 哈希表（Hash Table）

- 快速查找、计数、去重，常用于 Two Sum、异位词等。
- 典型题目：LeetCode 1、49

### 6. 快慢指针（Fast & Slow Pointer）

- 链表环检测、找中点等。
- 典型题目：LeetCode 141、876

### 7. 分治（Divide and Conquer）

- 递归拆分问题，典型如归并排序、快速排序。
- 典型题目：LeetCode 23、53

### 8. 动态规划（Dynamic Programming）

- 记录子问题结果，避免重复计算。
- 典型题目：LeetCode 70、53、300

### 9. 贪心算法（Greedy Algorithm）

- 每步做出局部最优选择，常用于区间覆盖、跳跃游戏等。
- 典型题目：LeetCode 55、435

### 10. 堆与优先队列（Heap & Priority Queue）

- 用于动态维护最大/最小值，常见于 Top K 问题。
- 典型题目：LeetCode 215、295

### 11. 回溯（Backtracking）

- 枚举所有可能，常用于排列组合、子集、数独等。
- 典型题目：LeetCode 46、78、51

### 12. 并查集（Union Find）

- 处理连通性、集合合并问题。
- 典型题目：LeetCode 547、684

### 13. 树的遍历（Tree Traversal）

- 前序、中序、后序、层序遍历，递归或迭代实现。
- 典型题目：LeetCode 94、102

### 搜索与遍历

1. DFS（深度优先搜索）
   - 递归 / 栈，常用于回溯（组合、排列、子集、数独、N 皇后）。
2. BFS（广度优先搜索）
   - 队列实现，常用于最短路径（图、迷宫问题、最小步数）。
3. 二分答案
   - 在“答案具有单调性”的问题上用二分法（如最小最大值、最小船运载重）。
4. 拓扑排序（Kahn 算法 / DFS），用于有向无环图依赖问题。

### 动态规划（DP）

1. 线性 DP
   - 斐波那契数列、爬楼梯、打家劫舍。
2. 区间 DP
   - 石子合并、矩阵链乘。
3. 背包 DP
   - 0/1 背包、完全背包、多重背包。
4. 状态压缩 DP
   - N 皇后、旅行商问题。
5. 滚动数组优化空间。
