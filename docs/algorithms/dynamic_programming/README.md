# 动态规划

动态规划（Dynamic Programming, DP）是一种将复杂问题分解为子问题，通过保存子问题的最优解来避免重复计算，从而高效求解全局最优解的算法思想。

动态规划有两种形式：

- 自顶向下（Top-Down）：记忆化递归
- 自底向上（Bottom-Up）：迭代填表

## 核心思想

动态规划的本质是 **用空间换时间**，通过记录子问题的解来避免重复计算，将指数级时间复杂度优化为多项式级。

**DP vs 递归**：

| 特征           | 朴素递归         | 动态规划                      |
| -------------- | ---------------- | ----------------------------- |
| **子问题重叠** | 重复计算         | 记忆化存储                    |
| **时间复杂度** | 通常 $O(2^n)$    | 通常 $O(n)$ 或 $O(n^2)$       |
| **空间复杂度** | $O(n)$（递归栈） | $O(n)$ 或 $O(n^2)$（dp 数组） |
| **实现方式**   | 自顶向下         | 自底向上或记忆化递归          |

!!! Tip

    只要递归状态从 0 开始，备忘录的大小就要 +1，比如 Fibonacci、爬楼梯、前缀和，因为 `dp[0]` 也需要存储。

    动态规划先找到状态转移方程，用暴力递归实现，再用备忘录优化掉重复计算，最后改为自底向上的 dp 数组（消除递归栈）。

!!! Note "动态规划学习路径"

    1. **掌握递归思维**：理解如何将问题分解为子问题
    2. **从前缀和入门**：最简单的 DP，状态转移方程清晰
    3. **学习经典问题**：斐波那契、爬楼梯、硬币找零
    4. **进阶复杂场景**：背包问题、股票买卖、最长子序列

## DP 解题五步曲

掌握以下框架，可以系统性地解决大部分 DP 问题：

1.  **定义 dp 数组的含义**
    - `dp[i]` 或 `dp[i][j]` 表示什么？
    - 明确状态的物理意义

2.  **找出状态转移方程**
    - 当前状态如何由之前的状态推导而来？
    - 这是 DP 的核心，也是最难的部分

3.  **初始化 dp 数组**
    - 确定边界条件（base case）
    - 通常是 `dp[0]` 或 `dp[0][0]`

4.  **确定遍历顺序**
    - 从前往后还是从后往前？
    - 一维还是二维遍历？

5.  **举例推导 dp 数组**
    - 用具体例子验证状态转移方程
    - 调试时打印 dp 数组观察规律

## 经典例子：斐波那契数列

### 问题

计算斐波那契数列的第 n 项：$F(n) = F(n-1) + F(n-2)$，其中 $F(0) = 0, F(1) = 1$

### 朴素递归（$O(2^n)$）

```go
func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)  // 大量重复计算
}
```

### DP 优化（$O(n)$）

=== "自底向上（迭代）"

    ```go
    func fib(n int) int {
        if n <= 1 {
            return n
        }

        // 1. 定义 dp 数组：dp[i] 表示第 i 个斐波那契数
        dp := make([]int, n+1)

        // 3. 初始化
        dp[0], dp[1] = 0, 1

        // 4. 遍历顺序：从前往后
        for i := 2; i <= n; i++ {
            // 2. 状态转移方程
            dp[i] = dp[i-1] + dp[i-2]
        }

        return dp[n]
    }
    ```

=== "记忆化递归（自顶向下）"

    ```go
    func fib(n int) int {
        memo := make(map[int]int)
        return fibHelper(n, memo)
    }

    func fibHelper(n int, memo map[int]int) int {
        if n <= 1 {
            return n
        }

        // 检查是否已计算过
        if v, ok := memo[n]; ok {
            return v
        }

        // 计算并存储
        memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
        return memo[n]
    }
    ```

=== "空间优化（$O(1)$）"

    ```go
    func fib(n int) int {
        if n <= 1 {
            return n
        }

        // 只需要保存前两个状态
        prev2, prev1 := 0, 1

        for i := 2; i <= n; i++ {
            curr := prev1 + prev2
            prev2, prev1 = prev1, curr
        }

        return prev1
    }
    ```

## DP 核心题型

### 1. 线性 DP

**特征**：状态转移只依赖于前面的若干个状态

**经典题目**：

- [70. Climbing Stairs](https://leetcode.cn/problems/climbing-stairs/) — 爬楼梯（入门题）
- [198. House Robber](https://leetcode.cn/problems/house-robber/) — 打家劫舍
- [300. Longest Increasing Subsequence](https://leetcode.cn/problems/longest-increasing-subsequence/) — 最长递增子序列

### 2. 背包问题

**特征**：在限制条件下选择物品，使得价值最大

| 类型         | 特点                 | 经典题目                                                                                    |
| ------------ | -------------------- | ------------------------------------------------------------------------------------------- |
| **0-1 背包** | 每个物品只能选一次   | [416. Partition Equal Subset Sum](https://leetcode.cn/problems/partition-equal-subset-sum/) |
| **完全背包** | 每个物品可以选无限次 | [322. Coin Change](https://leetcode.cn/problems/coin-change/)                               |
| **多重背包** | 每个物品有数量限制   | -                                                                                           |

### 3. 区间 DP

**特征**：在一个区间上进行决策，通常需要二维 dp 数组

**经典题目**：

- [5. Longest Palindromic Substring](https://leetcode.cn/problems/longest-palindromic-substring/) — 最长回文子串
- [516. Longest Palindromic Subsequence](https://leetcode.cn/problems/longest-palindromic-subsequence/) — 最长回文子序列

### 4. 路径问题

**特征**：在网格或图上寻找路径，通常是二维 dp

**经典题目**：

- [62. Unique Paths](https://leetcode.cn/problems/unique-paths/) — 不同路径
- [64. Minimum Path Sum](https://leetcode.cn/problems/minimum-path-sum/) — 最小路径和
- [120. Triangle](https://leetcode.cn/problems/triangle/) — 三角形最小路径和

### 5. 字符串 DP

**特征**：涉及字符串匹配、编辑距离等

**经典题目**：

- [72. Edit Distance](https://leetcode.cn/problems/edit-distance/) — 编辑距离
- [1143. Longest Common Subsequence](https://leetcode.cn/problems/longest-common-subsequence/) — 最长公共子序列

## 0-1 背包详解

0-1 背包是 DP 中最经典的问题，掌握它可以解决很多变体。

### 问题描述

有 n 个物品，每个物品有重量 `w[i]` 和价值 `v[i]`。背包容量为 `capacity`。每个物品只能选一次，求背包能装下的最大价值。

### 状态定义

`dp[i][j]` 表示前 i 个物品，背包容量为 j 时的最大价值。

### 状态转移方程

```
dp[i][j] = max(
    dp[i-1][j],              // 不选第 i 个物品
    dp[i-1][j-w[i]] + v[i]   // 选第 i 个物品（前提是 j >= w[i]）
)
```

### 代码实现

=== "二维 DP"

    ```go
    func knapsack(weights []int, values []int, capacity int) int {
        n := len(weights)
        // dp[i][j]: 前 i 个物品，容量为 j 的最大价值
        dp := make([][]int, n+1)
        for i := range dp {
            dp[i] = make([]int, capacity+1)
        }

        for i := 1; i <= n; i++ {
            for j := 0; j <= capacity; j++ {
                // 不选第 i 个物品
                dp[i][j] = dp[i-1][j]

                // 选第 i 个物品（如果放得下）
                if j >= weights[i-1] {
                    dp[i][j] = max(dp[i][j], dp[i-1][j-weights[i-1]] + values[i-1])
                }
            }
        }

        return dp[n][capacity]
    }
    ```

=== "一维 DP（空间优化）"

    ```go
    func knapsack(weights []int, values []int, capacity int) int {
        n := len(weights)
        dp := make([]int, capacity+1)

        for i := 0; i < n; i++ {
            // 必须从后往前遍历，避免重复使用同一物品
            for j := capacity; j >= weights[i]; j-- {
                dp[j] = max(dp[j], dp[j-weights[i]] + values[i])
            }
        }

        return dp[capacity]
    }
    ```

## DP 优化技巧

### 1. 空间优化

如果 `dp[i]` 只依赖于 `dp[i-1]`，可以用滚动数组或两个变量优化空间：

```go
// 从 O(n) 优化到 O(1)
prev, curr := 0, 1
for i := 2; i <= n; i++ {
    prev, curr = curr, prev + curr
}
```

### 2. 状态压缩

使用位运算压缩状态，适用于状态数量较少的情况。

### 3. 单调队列优化

在某些 DP 问题中，可以用单调队列优化转移过程，将 $O(n^2)$ 降至 $O(n)$。

## 学习建议

1. **从简单题开始**：先做爬楼梯、斐波那契等入门题
2. **掌握经典模型**：0-1 背包、完全背包、最长公共子序列
3. **多画图推导**：手动推导 dp 数组，理解状态转移
4. **总结模板**：归纳不同类型 DP 的解题模板
5. **刷题顺序**：线性 DP → 背包 DP → 区间 DP → 树形 DP

## 经典题目清单

- [70. Climbing Stairs](https://leetcode.cn/problems/climbing-stairs/) — 爬楼梯（入门）
- [198. House Robber](https://leetcode.cn/problems/house-robber/) — 打家劫舍
- [322. Coin Change](https://leetcode.cn/problems/coin-change/) — 零钱兑换（完全背包）
- [416. Partition Equal Subset Sum](https://leetcode.cn/problems/partition-equal-subset-sum/) — 分割等和子集（0-1 背包）
- [62. Unique Paths](https://leetcode.cn/problems/unique-paths/) — 不同路径
- [64. Minimum Path Sum](https://leetcode.cn/problems/minimum-path-sum/) — 最小路径和
- [300. Longest Increasing Subsequence](https://leetcode.cn/problems/longest-increasing-subsequence/) — 最长递增子序列
- [1143. Longest Common Subsequence](https://leetcode.cn/problems/longest-common-subsequence/) — 最长公共子序列
- [72. Edit Distance](https://leetcode.cn/problems/edit-distance/) — 编辑距离
- [5. Longest Palindromic Substring](https://leetcode.cn/problems/longest-palindromic-substring/) — 最长回文子串
