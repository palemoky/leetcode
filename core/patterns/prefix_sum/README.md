# 前缀和算法总结

前缀和（Prefix Sum）是一种预处理技巧，通过**空间换时间**的方式，将区间和查询的时间复杂度从 O(n) 优化到 O(1)。适用于**频繁查询数组区间和**的场景。

---

## 核心思想

| index                 |  0  |  1  |  2  |  3  |  4  | Formula                                                                                                                           |
| --------------------- | :-: | :-: | :-: | :-: | :-: | --------------------------------------------------------------------------------------------------------------------------------- |
| `nums`                |  3  |  5  |  2  |  7  |     |                                                                                                                                   |
| **`preSum[0] = 0`**   |  0  |  3  |  8  | 10  | 17  | **`sum[left, right] = preSum[right+1] - preSum[left]`**                                                                           |
| `preSum[0] = nums[0]` |  3  |  8  | 10  | 17  |     | `sum[left, right] = preSum[right] - preSum[left-1]   # left>0`<br/>`sum[0, right] = preSum[right]                       # left=0` |

从上表可以看出，**`preSum[0] = 0` 就像链表中的 dummy head**，使操作统一，避免容易出错的的边界判断。

---

## 典型应用场景

- 频繁查询数组区间和（如 LeetCode 303、304）
- 子数组和问题（如和为 K 的子数组个数）
- 二维矩阵区域和查询
- 结合哈希表优化查找（如两数之和的变体）

---

## 实现模板

### 1. 一维前缀和

```go
// 构建前缀和数组
func buildPrefixSum(nums []int) []int {
    prefixSum := make([]int, len(nums)+1) // prefixSum[0] = 0

    for i := range nums {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }

    return prefixSum
}

// 查询区间 [left, right] 的和
func rangeSum(prefixSum []int, left, right int) int {
    return prefixSum[right+1] - prefixSum[left]
}
```

**时间复杂度：**

- 构建：O(n)
- 查询：O(1)

**空间复杂度：** O(n)

### 2. 二维前缀和

用于二维矩阵的区域和查询。

```go
// 构建二维前缀和
func buildPrefixSum2D(matrix [][]int) [][]int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }

    m, n := len(matrix), len(matrix[0])
    prefixSum := make([][]int, m+1)
    for i := range prefixSum {
        prefixSum[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            prefixSum[i][j] = prefixSum[i-1][j] +
                             prefixSum[i][j-1] -
                             prefixSum[i-1][j-1] +
                             matrix[i-1][j-1]
        }
    }

    return prefixSum
}

// 查询矩形区域 (row1, col1) 到 (row2, col2) 的和
func regionSum(prefixSum [][]int, row1, col1, row2, col2 int) int {
    return prefixSum[row2+1][col2+1] -
           prefixSum[row1][col2+1] -
           prefixSum[row2+1][col1] +
           prefixSum[row1][col1]
}
```

**时间复杂度：**

- 构建：O(m × n)
- 查询：O(1)

**空间复杂度：** O(m × n)

---

## 进阶技巧

### 结合哈希表

用于查找和为 K 的子数组个数（LeetCode 560）。

```go
func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0
    sumCount := map[int]int{0: 1} // 前缀和 -> 出现次数

    for _, num := range nums {
        prefixSum += num

        // 查找是否存在 prefixSum - k
        if cnt, exists := sumCount[prefixSum-k]; exists {
            count += cnt
        }

        sumCount[prefixSum]++
    }

    return count
}
```

**核心思想：**

- `prefixSum[j] - prefixSum[i] = k`
- 转化为：`prefixSum[i] = prefixSum[j] - k`
- 用哈希表记录前缀和出现次数，边遍历边查找

---

## 优势

- **查询高效**：O(1) 时间复杂度查询区间和
- **实现简单**：只需一次预处理即可
- **适用广泛**：一维、二维、甚至高维数组

---

## 注意事项

- **数组越界**：前缀和数组长度为 `n+1`，注意索引范围
- **初始值**：`prefixSum[0] = 0`，表示空区间的和
- **区间端点**：查询 `[left, right]` 时，公式为 `prefixSum[right+1] - prefixSum[left]`
- **整数溢出**：如果数组元素很大，考虑使用 `int64` 或取模

---

## 经典题目

### 基础题

- [LeetCode 303. 区域和检索 - 数组不可变](https://leetcode.com/problems/range-sum-query-immutable/)
- [LeetCode 304. 二维区域和检索 - 矩阵不可变](https://leetcode.com/problems/range-sum-query-2d-immutable/)

### 进阶题

- [LeetCode 560. 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k/)
- [LeetCode 974. 和可被 K 整除的子数组](https://leetcode.com/problems/subarray-sums-divisible-by-k/)
- [LeetCode 525. 连续数组](https://leetcode.com/problems/contiguous-array/)
- [LeetCode 1314. 矩阵区域和](https://leetcode.com/problems/matrix-block-sum/)

---

## 相关模式

- **差分数组**：前缀和的逆运算，用于区间修改
- **滑动窗口**：动态维护区间和，适用于可变窗口
- **线段树**：支持区间查询和修改，但实现复杂
