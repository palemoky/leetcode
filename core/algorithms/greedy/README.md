# 贪心

## Introduction

Greedy algorithms make the locally optimal choice at each step, hoping to find the global optimum. They are widely used for optimization problems where choosing the best option at each stage leads to an overall optimal solution.

## Typical Applications

- Interval scheduling (e.g., activity selection)
- Coin change (when denominations are canonical)
- Huffman coding
- Minimum spanning tree (Kruskal, Prim)
- Dijkstra's shortest path (with non-negative weights)
- Task assignment and resource allocation
- String and array problems (e.g., jump game, partitioning)

## Classic Problems

- [122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)
- [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)
- [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)
- [860. Lemonade Change](https://leetcode.com/problems/lemonade-change/)
- [455. Assign Cookies](https://leetcode.com/problems/assign-cookies/)
- [406. Queue Reconstruction by Height](https://leetcode.com/problems/queue-reconstruction-by-height/)
- [135. Candy](https://leetcode.com/problems/candy/)
- [605. Can Place Flowers](https://leetcode.com/problems/can-place-flowers/)

## Greedy Solution Template

```go
// Example: Interval scheduling
func greedySolve(intervals [][]int) int {
    // Sort intervals by end time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    count, end := 0, math.MinInt32
    for _, interval := range intervals {
        if interval[0] >= end {
            count++
            end = interval[1]
        }
    }
    return count
}
```

## Key Points & Pitfalls

- Greedy choice must be proven to lead to global optimum (often via exchange argument or proof by contradiction)
- Not all problems with local choices are suitable for greedy; always check for counterexamples
- Sorting is often a key step in greedy solutions
- Sometimes multiple greedy strategies exist; test and prove correctness

## References

- [LeetCode Greedy Tag](https://leetcode.com/tag/greedy/)
- [OI Wiki: Greedy Algorithm](https://oi-wiki.org/basic/greedy/)
- [CP Algorithms: Greedy](https://cp-algorithms.com/greedy.html)
