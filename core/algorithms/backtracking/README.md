# 回溯

## 简介

回溯是一种通用的算法技术，通过递归尝试逐步构建解决方案来解决问题，在任何时候移除不满足约束条件的解决方案（即回溯）。它特别适用于组合问题，如排列、组合和子集。

## 典型应用

- 排列和组合
- 子集和幂集
- N 皇后问题
- 数独求解器
- 单词搜索
- 图着色
- 迷宫寻路
- 分割问题

## 经典问题

- [39. 组合总和](https://leetcode.cn/problems/combination-sum/)
- [46. 全排列](https://leetcode.cn/problems/permutations/)
- [78. 子集](https://leetcode.cn/problems/subsets/)
- [77. 组合](https://leetcode.cn/problems/combinations/)
- [90. 子集 II](https://leetcode.cn/problems/subsets-ii/)
- [40. 组合总和 II](https://leetcode.cn/problems/combination-sum-ii/)
- [47. 全排列 II](https://leetcode.cn/problems/permutations-ii/)
- [51. N 皇后](https://leetcode.cn/problems/n-queens/)
- [79. 单词搜索](https://leetcode.cn/problems/word-search/)
- [37. 解数独](https://leetcode.cn/problems/sudoku-solver/)

## 回溯解题模板

```go
// 示例：生成所有子集
func backtrackSubsets(nums []int) [][]int {
    res := [][]int{}
    var dfs func(path []int, start int)
    dfs = func(path []int, start int) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)
        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            dfs(path, i+1)
            path = path[:len(path)-1] // 回溯
        }
    }
    dfs([]int{}, 0)
    return res
}
```

## 关键要点与陷阱

- 递归调用后始终要回溯（撤销最后的选择）
- 如果违反约束条件，尽早剪枝（提高效率）
- 排列问题使用 visited/used 数组
- 组合/子集问题使用起始索引避免重复
- 注意 Go 中的引用类型（slice、map）；必要时进行复制

## 参考资料

- [LeetCode 回溯标签](https://leetcode.cn/tag/backtracking/)
- [OI Wiki: 回溯](https://oi-wiki.org/basic/backtracking/)
- [CP Algorithms: 回溯](https://cp-algorithms.com/backtracking.html)

## 回溯模板

```
func backtrack(路径, 选择列表) {
    if 满足结束条件 {
        记录结果
        return
    }
    for 选择 in 选择列表 {
        做选择
        backtrack(路径, 新的选择列表) // 递归
        撤销选择 // 状态重置，这是回溯的精髓！
    }
}
```

多叉树遍历框架：

```
func traverse(root *TreeNode) {
    if root == nil {
        return
    }
    traverse(root.Left)
    traverse(root.Right)
}
```

## 核心三题

**掌握三道核心题，理解所有回溯问题**：

- **组合 (Combination)**: [LeetCode 77. 组合](https://leetcode.cn/problems/combinations/)
- **排列 (Permutation)**: [LeetCode 46. 全排列](https://leetcode.cn/problems/permutations/)
- **子集 (Subset)**: [LeetCode 78. 子集](https://leetcode.cn/problems/subsets/)

这三类问题的本质都是 **遍历多叉树**：

1. **构建决策树**：每个节点代表一个选择，从根节点到叶子节点的路径就是一个解
2. **遍历求解**：用回溯框架遍历这棵多叉树，收集所有路径
3. **关键差异**：三者在"做选择"和"撤销选择"的细节上略有不同

理解这些差异，你就掌握了回溯算法 80% 的精髓。

**示例：`[1,2,3]` 的全排列决策树**

<figure>
    <img src="permutation_decision_tree.webp" alt="全排列决策树" width="50%" />
</figure>

## 其他经典问题

汉诺塔、N 皇后、图着色、旅行商问题
