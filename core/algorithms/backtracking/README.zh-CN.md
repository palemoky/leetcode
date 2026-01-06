
回溯模板
```go
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

死磕三道经典题，它们代表了所有回溯问题：
组合 (Combination): [LeetCode 77. Combinations](https://leetcode.com/prombles/combinations/)
排列 (Permutation): [LeetCode 46. Permutations](https://leetcode.com/prombles/permutations)
子集 (Subset): [LeetCode 78. Subsets](https://leetcode.com/prombles/subset)
理解这三者在“做选择”和“撤销选择”上的细微差别，你就掌握了回溯80%的精髓

汉诺塔、N皇后、图着色、旅行商问题
