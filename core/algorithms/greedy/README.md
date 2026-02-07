# 贪心算法

## 核心思想

贪心算法在每一步都做出 **局部最优选择**，期望通过局部最优达到全局最优。它广泛应用于优化问题，在这些问题中，每一步选择最佳选项能够导致整体最优解。

## 典型应用场景

- **区间调度** (如活动选择问题)
- **找零问题** (当面额是规范的)
- **霍夫曼编码** (数据压缩)
- **最小生成树** (Kruskal、Prim 算法)
- **最短路径** (Dijkstra 算法,非负权重)
- **任务分配与资源调度**
- **字符串和数组问题** (如跳跃游戏、分割问题)

## 经典题目

- [122. 买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)
- [435. 无重叠区间](https://leetcode.cn/problems/non-overlapping-intervals/)
- [452. 用最少数量的箭引爆气球](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/)
- [860. 柠檬水找零](https://leetcode.cn/problems/lemonade-change/)
- [455. 分发饼干](https://leetcode.cn/problems/assign-cookies/)
- [406. 根据身高重建队列](https://leetcode.cn/problems/queue-reconstruction-by-height/)
- [135. 分发糖果](https://leetcode.cn/problems/candy/)
- [605. 种花问题](https://leetcode.cn/problems/can-place-flowers/)

## 贪心算法模板

```python
def greedy_solve(intervals: list[list[int]]) -> int:
    """
    示例: 区间调度问题
    按结束时间排序,选择不重叠的最多区间
    """
    # 按结束时间排序
    intervals.sort(key=lambda x: x[1])

    count = 0
    end = float('-inf')

    for interval in intervals:
        if interval[0] >= end:
            count += 1
            end = interval[1]

    return count
```

## 关键要点与常见陷阱

### 使用贪心的条件

1. **贪心选择性质**: 局部最优选择能导致全局最优
2. **最优子结构**: 问题的最优解包含子问题的最优解

### 注意事项

- **必须证明正确性**: 贪心选择必须被证明能导致全局最优(通常通过交换论证或反证法)
- **不是所有问题都适用**: 并非所有涉及局部选择的问题都适合贪心,务必检查反例
- **排序是关键**: 排序往往是贪心解法的关键步骤
- **可能有多种策略**: 有时存在多种贪心策略,需要测试并证明正确性

### 如何判断是否用贪心?

**使用信号**:

- 问题要求 **最优解**（最大、最小、最多、最少）
- 可以通过 **排序** 简化问题
- 局部最优选择 **不影响** 后续选择
- 问题有明显的 **贪心策略**（如总是选最小/最大）

**典型关键词**: 最多、最少、最大、最小、区间、调度、分配

## 贪心 vs 动态规划

| 特性           | 贪心算法           | 动态规划                |
| -------------- | ------------------ | ----------------------- |
| **决策方式**   | 每步做局部最优选择 | 考虑所有可能,选全局最优 |
| **是否回溯**   | 不回溯,一次决策    | 可能需要回溯            |
| **适用条件**   | 贪心选择性质       | 重叠子问题 + 最优子结构 |
| **时间复杂度** | 通常 $O(n \log n)$ | 通常 $O(n^2)$ 或更高    |
| **正确性**     | 需要严格证明       | 状态转移方程保证        |

**何时用贪心而非 DP?**

- 贪心能证明正确性时，优先用贪心(更快)
- 找零问题: 规范面额用贪心，非规范用 DP
- 区间问题: 通常用贪心
- 背包问题: 0-1 背包用 DP，分数背包用贪心

## 常见贪心策略

1. **按某个属性排序**: 结束时间、开始时间、长度等
2. **选择极值**: 总是选最大/最小的
3. **优先队列**: 动态维护最优选择
4. **双指针**: 从两端向中间贪心选择

## 参考资料

- [LeetCode 贪心标签](https://leetcode.cn/tag/greedy/)
- [OI Wiki: 贪心算法](https://oi-wiki.org/basic/greedy/)
- [CP Algorithms: Greedy](https://cp-algorithms.com/greedy.html)
