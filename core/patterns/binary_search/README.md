# 二分查找（Binary Search）

二分查找是一种高效的查找算法，适用于 **有序数组或区间** ，通过每次将搜索范围缩小一半，快速定位目标元素或插入位置。时间复杂度为 $O(log n)$，空间复杂度为 $O(1)$。

---

## 适用场景

- 有序数组查找元素
- 查找插入位置（如 lower_bound/upper_bound）
- 判定区间最值（如最小化最大值、最大化最小值等）
- 单调性判定（如答案具有单调性时的“答案二分”）

---

## 基本实现

### 1. 查找目标元素

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        // mid 位置为偶数长度的中间偏左
        mid := left + (right-left)/2
        if target == nums[mid] {
            return mid
        } else if target > nums[mid] { // 位于右侧区间
            left = mid + 1
        } else { // 位于左侧区间
            right = mid - 1
        }
    }
    return -1 // 未找到
}
```

### 2. 查找插入位置（lower_bound）

```go
func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if target > nums[mid] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left // 返回第一个 >= target 的位置
}
```

## 变体与扩展

- 查找区间边界：如第一个大于等于/小于等于目标的位置
- 答案二分：在单调性判定问题中，二分答案空间（如 LeetCode 410、875）
- 旋转数组查找：如 LeetCode 33、81
- 递归实现：可用递归方式实现，但需注意递归深度和空间

## 常见错误与注意事项

- 循环条件：`left <= right`（闭区间）或 `left < right`（半开区间），根据需求选择
- 防止溢出：`mid := left + (right-left)/2` 和 `mid := (left+right)/2` 都可以求中值，但前者可以避免溢出
- 边界处理：空数组、目标不存在、插入点在头尾
- 不要无限循环：每次循环必须收缩区间

## 经典题目

- [LeetCode 704. Binary Search](https://leetcode.cn/problems/binary-search/)
- [LeetCode 35. Search Insert Position](https://leetcode.cn/problems/search-insert-position/)
- [LeetCode 278. First Bad Version](https://leetcode.cn/problems/first-bad-version/)
- [LeetCode 33. Search in Rotated Sorted Array](https://leetcode.cn/problems/search-in-rotated-sorted-array/)
- [LeetCode 69. Sqrt(x)](https://leetcode.cn/problems/sqrtx/)
- [LeetCode 875. Koko Eating Bananas](https://leetcode.cn/problems/koko-eating-bananas/)

## 总结

二分查找是算法面试和实际开发中的基础技巧，适用于所有有序和单调性问题。掌握不同区间写法和边界处理，能高效解决大量查找与判定类题目。
