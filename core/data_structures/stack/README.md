# 栈

栈（Stack）是一种**后进先出**（LIFO, Last In First Out）的线性数据结构。它就像一摞盘子，你只能在顶端放盘子（Push），也只能从顶端取盘子（Pop）。

## 核心特性

栈的操作受到严格限制，只能在**栈顶**进行：

1.  **Push（入栈）**：将元素放入栈顶。
2.  **Pop（出栈）**：移除并返回栈顶元素。
3.  **Peek/Top（查看栈顶）**：返回栈顶元素但不移除。
4.  **IsEmpty（判空）**：检查栈是否为空。

## 复杂度分析

| 操作       | 时间复杂度 | 说明                                                          |
| :--------- | :--------- | :------------------------------------------------------------ |
| **Push**   | $O(1)$     | 仅操作栈顶，不涉及移动其他元素（动态数组扩容摊销后为 $O(1)$） |
| **Pop**    | $O(1)$     | 仅操作栈顶                                                    |
| **Peek**   | $O(1)$     | 直接通过索引访问                                              |
| **Search** | $O(n)$     | 需要遍历栈中元素                                              |

## 实现方式 (Go)

在 Go 语言中，通常直接使用切片（Slice）模拟栈，这是最常用且高效的方式。

=== "基础实现"

    ```go
    // 使用切片模拟栈
    stack := []int{}

    // 1. Push - 入栈
    stack = append(stack, 10)
    stack = append(stack, 20)

    // 2. Peek - 查看栈顶
    top := stack[len(stack)-1]
    // top = 20

    // 3. Pop - 出栈
    val := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    // val = 20, stack = [10]

    // 4. IsEmpty - 判空
    if len(stack) == 0 {
        fmt.Println("Stack is empty")
    }
    ```

=== "封装实现"

    ```go
    type Stack struct {
        elements []int
    }

    func NewStack() *Stack {
        return &Stack{elements: []int{}}
    }

    func (s *Stack) Push(val int) {
        s.elements = append(s.elements, val)
    }

    func (s *Stack) Pop() (int, bool) {
        if s.IsEmpty() {
            return 0, false
        }
        index := len(s.elements) - 1
        val := s.elements[index]
        s.elements = s.elements[:index]
        return val, true
    }

    func (s *Stack) Peek() (int, bool) {
        if s.IsEmpty() {
            return 0, false
        }
        return s.elements[len(s.elements)-1], true
    }

    func (s *Stack) IsEmpty() bool {
        return len(s.elements) == 0
    }
    ```

## 应用场景

1.  **函数调用栈**：操作系统维护函数调用关系，实现递归。
2.  **括号匹配**：IDE 检查代码中的括号是否闭合。
3.  **表达式求值**：计算逆波兰表达式（后缀表达式）。
4.  **浏览器历史**：后退按钮（两个栈实现前进后退）。
5.  **撤销操作（Undo）**：编辑器中的 Ctrl+Z。

## 经典算法模式

### [单调栈](../../patterns/monotonic_stack/README.md)

单调栈是一种特殊的栈应用，栈内元素保持单调递增或递减。主要用于解决 **寻找最近的更大/更小元素** 这类问题。

- **适用场景**：找左边/右边第一个比当前元素大/小的元素。
- **经典例题**：[739. 每日温度](https://leetcode.com/problems/daily-temperatures/)

## 经典题目

=== "基础题"

    - [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) — 有效的括号（栈的经典应用）
    - [155. Min Stack](https://leetcode.com/problems/min-stack/) — 最小栈（辅助栈思想）
    - [232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) — 用栈实现队列
    - [1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/) — 删除字符串中的所有相邻重复项

=== "进阶题"

    - [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) — 逆波兰表达式求值
    - [71. Simplify Path](https://leetcode.com/problems/simplify-path/) — 简化路径
    - [394. Decode String](https://leetcode.com/problems/decode-string/) — 字符串解码（辅助栈）
    - [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) — 每日温度（单调栈入门）

=== "高级题"

    - [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) — 柱状图中最大的矩形（单调栈经典）
    - [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) — 接雨水（单调栈解法）
    - [32. Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/) — 最长有效括号
    - [224. Basic Calculator](https://leetcode.com/problems/basic-calculator/) — 基本计算器（处理括号和优先级）

## 总结

- **操作受限**：牢记 LIFO 特性，所有操作都在栈顶。
- **切片即栈**：在 Go 算法题中，通常直接用切片操作，无需专门封装类。
- **递归的本质**：理解递归就是隐式地使用系统栈，有时可以用显式栈将递归转化为迭代。
