# 单调栈（Monotonic Stack）

单调栈是一种特殊的栈结构，栈内元素始终保持单调递增或递减，常用于解决“下一个更大/更小元素”、“区间最值”、“直方图最大矩形”等问题。

## 适用场景

- 需要在 O(n) 时间内找到每个元素左/右侧第一个比它大（小）的元素
- 解决区间最值、滑动窗口最值、直方图最大矩形等问题
- 常见于数组、序列、温度、股票等题型

## 模板代码（递减栈，找下一个更大元素）

```go
stack := []int{} // 存下标或值
for i := 0; i < len(nums); i++ {
    for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
        top := stack[len(stack)-1]
        // 处理 nums[top]，如记录下一个更大元素的位置/值
        stack = stack[:len(stack)-1]
    }
    stack = append(stack, i)
}
```

- 递增栈（找下一个更小元素）只需把 `>` 换成 `<`

## 常见题型

- [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)（下一个更高温度）
- [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)
- [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)
- [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)
- [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)

## 技巧与注意事项

- 栈内通常存下标，便于回溯和计算距离
- 处理“循环数组”时可遍历两遍或用取模技巧
- 递增/递减栈的选择取决于题目需求（更大/更小）
- 画图模拟栈的变化有助于理解

## 相关模式

- 单调队列（Monotonic Queue）：用于滑动窗口最值
- 双指针、前缀和等可与单调栈结合
