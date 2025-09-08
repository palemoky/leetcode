# 单调栈（Monotonic Stack）

单调栈是一种特殊的栈结构，栈内元素始终保持单调递增或递减，常用于解决“下一个更大/更小元素”、“区间最值”、“直方图最大矩形”等问题。

其本质是通过局部有序性，求得相邻的更大/更小数。

<div align="center">
  <table>
    <tr>
      <td align="center" valign="bottom" >
        <img src="monotonous-stack-before.svg" alt="Before pushing 14" /><br />
        <sub style="font-size: 14px;">Before pushing 14</sub>
      </td>
      <td align="center" valign="bottom" >
        <img src="monotonous-stack-after.svg" alt="After pushing 14" /><br />
        <sub style="font-size: 14px;">After pushing 14</sub>
      </td>
    </tr>
  </table>
</div>

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

# LeetCode 经典题目

- [739. 每日温度 (Daily Temperatures)](https://leetcode.com/problems/daily-temperatures/)
- [84. 柱状图中最大的矩形 (Largest Rectangle in Histogram)](https://leetcode.com/problems/largest-rectangle-in-histogram/)
- [42. 接雨水 (Trapping Rain Water)](https://leetcode.com/problems/trapping-rain-water/)
- [496. 下一个更大元素 I (Next Greater Element I)](https://leetcode.com/problems/next-greater-element-i/)
- [503. 下一个更大元素 II (Next Greater Element II)](https://leetcode.com/problems/next-greater-element-ii/)
- [85. 最大矩形 (Maximal Rectangle)](https://leetcode.com/problems/maximal-rectangle/)
- [907. 子数组的最小值之和 (Sum of Subarray Minimums)](https://leetcode.com/problems/sum-of-subarray-minimums/)
- [901. 股票价格跨度 (Online Stock Span)](https://leetcode.com/problems/online-stock-span/)
- [1019. 链表中的下一个更大节点 (Next Greater Node In Linked List)](https://leetcode.com/problems/next-greater-node-in-linked-list/)
- [1944. 队列中可以看到的人数 (Number of Visible People in a Queue)](https://leetcode.com/problems/number-of-visible-people-in-a-queue/)
- [2289. 使数组按非递减顺序排列 (Steps to Make Array Non-decreasing)](https://leetcode.com/problems/steps-to-make-array-non-decreasing/)
- [20. 有效的括号 (Valid Parentheses)](https://leetcode.com/problems/valid-parentheses/)
