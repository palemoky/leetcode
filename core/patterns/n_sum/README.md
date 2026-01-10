# n 数之和（k-sum）

目标：在数组中找出若干个数之和等于目标值的问题。常见变体：Two Sum / Three Sum / Four Sum / k-Sum。核心思想是把问题逐级化简到 Two Sum（双指针或哈希），并结合排序与去重进行剪枝与去重。

---

## 通用模式（适用场景）

- 要求返回值组合（值的集合）而非索引：可以对数组排序并使用双指针。
- 需要去重：排序后在枚举时跳过重复元素，或使用集合记录结果（代价较高）。
- k 较小（常见 k ≤ 4）：递归 + 双指针效率优秀。
- k 较大或 n 较小时可考虑回溯或 meet-in-the-middle。

---

## 基本策略（递归 + 双指针）

1. 先对数组排序（升序）。
2. 编写递归函数 kSum(nums, target, k, start)：
   - 若 k == 2：用左右双指针在有序数组范围 [start, n-1] 查找两数之和等于 target。
   - 否则：枚举 i 从 start 到 n-k，若 nums[i] 与 nums[i-1] 相同则跳过；否则固定 nums[i]，递归调用 (k-1)Sum 在 i+1 之后寻找 target-nums[i]，将 nums[i] 拼接到子解前加入结果。
3. 在递归中使用剪枝：
   - 若最小可能和（前 k 个数之和）大于 target 或最大可能和（后 k 个数之和）小于 target，则可以提前返回。
4. 每次找到解后要跳过重复值以避免重复组合。

---

## 复杂度

- 排序：$O(n log n)$
- k-Sum 最坏时间：$O(n^{k-1})$（递归降到 2-sum 后，外层有 k-2 层循环，每层最多 $O(n)$），空间：递归栈 $O(k)$ + 结果输出空间。
- 常见实例：ThreeSum -> $O(n^2)$，FourSum -> $O(n^3)$。

---

## 常见优化与变体

- 剪枝（bound checks）能显著降低常数项。
- Two-sum 可用哈希表实现 $O(n)$（但在 k-sum 模式下常用双指针以便去重与排序协同）。
- Meet-in-the-middle：当 k 很大时可把集合分为两半分别求和，再两两配对（时间-空间折中）。
- 若只需判断存在性，可用更强剪枝或位运算/数论技巧。
- 对于整数范围有限的情况，可用计数数组替代排序/哈希以加速。

---

## 去重要点

- 排序后：枚举时若 `nums[i] == nums[i-1]`（且 `i>start`）则跳过同值起点。
- Two-sum 找到一对后，移动 left/right 时跳过重复值：`for left<right && nums[left]==nums[left+1] { left++ }` 同理 right。
- 这样保证每组值仅被加入一次。

---

## Go 模板（递归 k-sum，返回值组合）

```go
package ksum

import "sort"

// kSum 主入口
func KSum(nums []int, target int, k int) [][]int {
    sort.Ints(nums)
    return kSum(nums, target, k, 0)
}

func kSum(nums []int, target, k, start int) [][]int {
    n := len(nums)
    res := [][]int{}
    if start >= n || k < 2 {
        return res
    }
    // 剪枝：最小和与最大和检查
    if k*nums[start] > target || k*nums[n-1] < target {
        return res
    }
    if k == 2 {
        // two-sum 双指针
        l, r := start, n-1
        for l < r {
            s := nums[l] + nums[r]
            if s == target {
                res = append(res, []int{nums[l], nums[r]})
                // 跳过重复
                for l < r && nums[l] == nums[l+1] { l++ }
                for l < r && nums[r] == nums[r-1] { r-- }
                l++; r--
            } else if s < target {
                l++
            } else {
                r--
            }
        }
        return res
    }

    for i := start; i <= n-k; i++ {
        if i > start && nums[i] == nums[i-1] { // 跳过重复起点
            continue
        }
        // 更强剪枝（可选）
        if nums[i] + (k-1)*nums[n-1] < target {
            continue
        }
        if nums[i] + (k-1)*nums[i+1] > target {
            break
        }
        sub := kSum(nums, target-nums[i], k-1, i+1)
        for _, comb := range sub {
            cur := make([]int, 0, len(comb)+1)
            cur = append(cur, nums[i])
            cur = append(cur, comb...)
            res = append(res, cur)
        }
    }
    return res
}
```

---

## 常见测试用例

- 空数组、元素少于 k
- 所有元素相同（如全 0）
- 有多个重复元素（检查去重）
- 大正/负数混合、极端边界（int 溢出注意）

---

## 总结

k-sum 的通用解法是排序 + 递归降维到 two-sum（双指针），并结合上下界剪枝与去重技巧。对于 k 小且 n 较大，此方法最实用；对于 k 很大或对时间/空间有特殊需求，可考虑 meet-in-the-middle 或特定数据结构优化。
