# 回溯

## Introduction

Backtracking is a general algorithmic technique for solving problems recursively by trying to build a solution incrementally, removing solutions that fail to satisfy the constraints at any point (i.e., backtracking). It is especially useful for combinatorial problems, such as permutations, combinations, and subsets.

## Typical Applications

- Permutations and combinations
- Subsets and powersets
- N-Queens problem
- Sudoku solver
- Word search
- Graph coloring
- Path finding in mazes
- Partitioning problems

## Classic Problems

- [39. Combination Sum](https://leetcode.com/problems/combination-sum/)
- [46. Permutations](https://leetcode.com/problems/permutations/)
- [78. Subsets](https://leetcode.com/problems/subsets/)
- [77. Combinations](https://leetcode.com/problems/combinations/)
- [90. Subsets II](https://leetcode.com/problems/subsets-ii/)
- [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)
- [47. Permutations II](https://leetcode.com/problems/permutations-ii/)
- [51. N-Queens](https://leetcode.com/problems/n-queens/)
- [79. Word Search](https://leetcode.com/problems/word-search/)
- [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

## Backtracking Solution Template

```go
// Example: Generate all subsets
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
            path = path[:len(path)-1] // backtrack
        }
    }
    dfs([]int{}, 0)
    return res
}
```

## Key Points & Pitfalls

- Always backtrack (undo the last choice) after recursive call
- Prune branches early if constraints are violated (improves efficiency)
- Use visited/used arrays for permutation problems
- For combinations/subsets, use start index to avoid duplicates
- Be careful with reference types (slice, map) in Go; copy when needed

## References

- [LeetCode Backtracking Tag](https://leetcode.com/tag/backtracking/)
- [OI Wiki: Backtracking](https://oi-wiki.org/basic/backtracking/)
- [CP Algorithms: Backtracking](https://cp-algorithms.com/backtracking.html)
