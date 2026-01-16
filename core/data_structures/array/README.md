# 数组

数组（Array）是最基础、最常用的数据结构，它在 **连续的内存空间** 中存储相同类型的元素，通过 **索引** 实现 \(O(1)\) 时间复杂度的随机访问。

![Complete Binary Tree Array Index](../tree/complete_binary_tree_array_index.png){ align=right width=35% }

数组是典型的线性结构，最贴近真实的内存和磁盘连续地址（逻辑上），因此是所有数据结构的基础。其他高级数据结构由于受限于线性的内存结构，往往需要通过数组来实现底层存储。例如：栈和队列可以用数组实现，哈希表的桶用数组存储，堆是基于数组的完全二叉树，图的邻接矩阵也是二维数组。可以说，**掌握数组就是掌握了数据结构的基石**。

## 核心特性

### 1. 连续内存存储

数组元素在内存中连续排列，元素地址可以通过公式计算：

```
address(arr[i]) = base_address + i × element_size
```

**优势**：

- ✅ 缓存友好（Cache-friendly）
- ✅ 支持随机访问
- ✅ 内存局部性好

**劣势**：

- ❌ 插入/删除需要移动元素
- ❌ 固定大小（静态数组）

!!! Tip

    要求原地修改数组元素时可逆向遍历修改

### 2. 索引访问

索引从 0 开始（大多数语言）：

```go
arr := []int{10, 20, 30, 40, 50}
//     index: 0   1   2   3   4

fmt.Println(arr[0])  // 10
fmt.Println(arr[2])  // 30
fmt.Println(arr[4])  // 50
```

### 3. 固定类型

数组中的所有元素必须是相同类型。

## 时间复杂度

| 操作             | 时间复杂度 | 说明               |
| ---------------- | ---------- | ------------------ |
| **访问**         | \(O(1)\)   | 通过索引直接访问   |
| **查找**         | \(O(n)\)   | 需要遍历数组       |
| **插入（末尾）** | \(O(1)\)   | 动态数组摊销复杂度 |
| **插入（中间）** | \(O(n)\)   | 需要移动后续元素   |
| **删除（末尾）** | \(O(1)\)   | 直接删除           |
| **删除（中间）** | \(O(n)\)   | 需要移动后续元素   |

## 静态数组 vs 动态数组

| 特性            | 静态数组      | 动态数组      |
| --------------- | ------------- | ------------- |
| **大小**        | 固定          | 可变          |
| **内存分配**    | 编译时        | 运行时        |
| **扩容**        | 不支持        | 自动扩容      |
| **Go 实现**     | `[n]T`        | `[]T` (slice) |
| **C++ 实现**    | `int arr[10]` | `vector<int>` |
| **Python 实现** | -             | `list`        |

### Go 语言示例

```go
// 静态数组（固定大小）
var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

// 动态数组（切片）
arr2 := []int{1, 2, 3, 4, 5}
arr2 = append(arr2, 6)  // 可以追加元素

// 创建指定容量的切片
arr3 := make([]int, 0, 10)  // 长度 0，容量 10
```

## 常见操作

### 1. 遍历

```go
arr := []int{1, 2, 3, 4, 5}

// 方法 1：使用索引
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// 方法 2：使用 range
for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}

// 方法 3：只要值
for _, value := range arr {
    fmt.Println(value)
}
```

### 2. 插入元素

```go
// 在末尾插入
arr = append(arr, 6)

// 在开头插入
arr = append([]int{0}, arr...)

// 在中间插入（在索引 2 处插入 99）
index := 2
arr = append(arr[:index], append([]int{99}, arr[index:]...)...)
```

### 3. 删除元素

```go
// 删除末尾元素
arr = arr[:len(arr)-1]

// 删除开头元素
arr = arr[1:]

// 删除中间元素（删除索引 2）
index := 2
arr = append(arr[:index], arr[index+1:]...)
```

### 4. 切片操作

```go
arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 获取子数组 [start:end)，左闭右开
sub1 := arr[2:5]     // [2, 3, 4]
sub2 := arr[:3]      // [0, 1, 2]
sub3 := arr[7:]      // [7, 8, 9]
sub4 := arr[:]       // 完整复制

// 切片共享底层数组！
sub1[0] = 99
fmt.Println(arr)     // [0, 1, 99, 3, 4, 5, 6, 7, 8, 9]
```

!!! Warning "切片陷阱"

    切片操作返回的是**原数组的视图**，修改切片会影响原数组！如需独立副本，使用 `copy()`：
    ```go
    original := []int{1, 2, 3, 4, 5}
    copied := make([]int, len(original))
    copy(copied, original)
    ```

### 5. 反转数组

```go
func reverse(arr []int) {
    left, right := 0, len(arr)-1
    for left < right {
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
}
```

### 6. 数组排序

```go
import "sort"

arr := []int{5, 2, 8, 1, 9}

// 升序排序
sort.Ints(arr)  // [1, 2, 5, 8, 9]

// 降序排序
sort.Sort(sort.Reverse(sort.IntSlice(arr)))

// 自定义排序
sort.Slice(arr, func(i, j int) bool {
    return arr[i] > arr[j]  // 降序
})
```

## 经典算法模式

### 1. 双指针

**快慢指针**：

```go
// 移除数组中的重复元素（原地修改）
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}
```

**对撞指针**：

```go
// 两数之和 II（有序数组）
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return nil
}
```

**经典题目**：

- [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)
- [27. Remove Element](https://leetcode.com/problems/remove-element/)
- [167. Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

### 2. 滑动窗口

```go
// 长度为 k 的子数组的最大和
func maxSumSubarray(arr []int, k int) int {
    if len(arr) < k {
        return 0
    }

    // 初始化第一个窗口
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    maxSum := windowSum

    // 滑动窗口
    for i := k; i < len(arr); i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        maxSum = max(maxSum, windowSum)
    }

    return maxSum
}
```

**经典题目**：

- [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)
- [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

### 3. 前缀和

```go
// 构建前缀和数组
func buildPrefixSum(arr []int) []int {
    n := len(arr)
    prefixSum := make([]int, n+1)  // prefixSum[0] = 0

    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + arr[i]
    }

    return prefixSum
}

// 查询区间和 [left, right]
func rangeSum(prefixSum []int, left, right int) int {
    return prefixSum[right+1] - prefixSum[left]
}
```

**经典题目**：

- [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)
- [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/)

### 4. 原地修改

```go
// 将数组中的 0 移到末尾
func moveZeroes(nums []int) {
    slow := 0  // 指向下一个非零元素应该放置的位置

    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
    }
}
```

**经典题目**：

- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)
- [75. Sort Colors](https://leetcode.com/problems/sort-colors/)

### 5. 二分查找

```go
// 在有序数组中查找目标值
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1  // 未找到
}
```

**经典题目**：

- [704. Binary Search](https://leetcode.com/problems/binary-search/)
- [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)

## 多维数组

### 二维数组

```go
// 创建 3x4 的二维数组
matrix := make([][]int, 3)
for i := range matrix {
    matrix[i] = make([]int, 4)
}

// 初始化
matrix := [][]int{
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12},
}

// 访问元素
fmt.Println(matrix[1][2])  // 7

// 遍历
for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Print(matrix[i][j], " ")
    }
    fmt.Println()
}
```

**经典题目**：

- [48. Rotate Image](https://leetcode.com/problems/rotate-image/)
- [54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)
- [73. Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

## 常见陷阱

### 1. 数组越界

```go
arr := []int{1, 2, 3}

// ❌ 越界访问会 panic
// fmt.Println(arr[3])  // panic: index out of range

// ✅ 检查边界
if index >= 0 && index < len(arr) {
    fmt.Println(arr[index])
}
```

!!! Tip

    只要递归状态从 0 开始，数组就要多开 1 个以避免越界访问，如动态规划和前缀和。

### 2. 切片容量陷阱

```go
arr := make([]int, 0, 5)  // 长度 0，容量 5
fmt.Println(len(arr))     // 0
fmt.Println(cap(arr))     // 5

// ❌ 不能直接访问
// arr[0] = 1  // panic: index out of range

// ✅ 需要先 append
arr = append(arr, 1)
```

### 3. 切片共享底层数组

```go
original := []int{1, 2, 3, 4, 5}
slice := original[1:4]  // [2, 3, 4]

slice[0] = 99
fmt.Println(original)   // [1, 99, 3, 4, 5] ← 原数组被修改！

// ✅ 使用 copy 创建独立副本
slice2 := make([]int, 3)
copy(slice2, original[1:4])
slice2[0] = 88
fmt.Println(original)   // [1, 99, 3, 4, 5] ← 不受影响
```

### 4. append 可能重新分配

```go
arr1 := []int{1, 2, 3}
arr2 := arr1

arr1 = append(arr1, 4)  // 可能触发扩容，分配新内存
arr1[0] = 99

fmt.Println(arr1)  // [99, 2, 3, 4]
fmt.Println(arr2)  // [1, 2, 3] ← 不受影响（如果发生了扩容）
```

### 5. range 遍历的值是副本

```go
arr := []int{1, 2, 3}

// ❌ 修改 value 不会影响原数组
for _, value := range arr {
    value = value * 2  // 无效！
}
fmt.Println(arr)  // [1, 2, 3]

// ✅ 使用索引修改
for i := range arr {
    arr[i] = arr[i] * 2
}
fmt.Println(arr)  // [2, 4, 6]
```

## 性能优化技巧

### 1. 预分配容量

```go
// ❌ 频繁扩容，性能差
arr := []int{}
for i := 0; i < 10000; i++ {
    arr = append(arr, i)
}

// ✅ 预分配容量，避免扩容
arr := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    arr = append(arr, i)
}
```

### 2. 复用切片

```go
// ❌ 每次都创建新切片
for i := 0; i < 1000; i++ {
    temp := []int{}
    // ... 使用 temp
}

// ✅ 复用切片，重置长度
temp := make([]int, 0, 100)
for i := 0; i < 1000; i++ {
    temp = temp[:0]  // 重置长度，保留容量
    // ... 使用 temp
}
```

### 3. 批量操作

```go
// ❌ 逐个 append
for i := 0; i < 1000; i++ {
    arr = append(arr, i)
}

// ✅ 批量 append
batch := make([]int, 1000)
for i := 0; i < 1000; i++ {
    batch[i] = i
}
arr = append(arr, batch...)
```

## 经典题目清单

### 基础题

- [1. Two Sum](https://leetcode.com/problems/two-sum/) — 两数之和
- [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) — 删除有序数组中的重复项
- [27. Remove Element](https://leetcode.com/problems/remove-element/) — 移除元素
- [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) — 合并两个有序数组
- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) — 移动零

### 进阶题

- [15. 3Sum](https://leetcode.com/problems/3sum/) — 三数之和
- [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) — 最大子数组和
- [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) — 买卖股票的最佳时机
- [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) — 除自身以外数组的乘积
- [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/) — 和为 K 的子数组

### 高级题

- [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) — 接雨水
- [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) — 柱状图中最大的矩形
- [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) — 滑动窗口最大值

## 总结

数组是最基础但也是最重要的数据结构，掌握以下要点：

1. ✅ **理解特性**：连续内存、随机访问、固定类型
2. ✅ **熟练操作**：遍历、插入、删除、切片、排序
3. ✅ **掌握模式**：双指针、滑动窗口、前缀和、二分查找
4. ✅ **注意陷阱**：越界、切片共享、扩容、值副本
5. ✅ **性能优化**：预分配容量、复用切片、批量操作
