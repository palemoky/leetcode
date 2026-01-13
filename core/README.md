# 核心算法

## 算法思想

1. **空间换时间**：
   - **数据结构层面**：哈希表、前缀和、差分数组、记忆化递归（DP）
   - **系统层面**：Buffer（缓冲区，平滑速度差异）、Cache（缓存，避免重复计算/访问）
   - **核心思想**：用额外的存储空间来避免重复计算或加速访问
2. **边界条件**：差 1 问题、指针移动、数组越界、递归终止、整数溢出、循环条件（`<` vs `<=`）等，检查清单如下：

   | **通用边界**                                                                                                                                         | **数据结构特定边界**                                                                                                                                                           |
   | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
   | □ 空输入（空数组、空字符串、nil）<br>□ 单元素<br>□ 两个元素（最小有意义输入）<br>□ 全部相同元素<br>□ 全部不同元素<br>□ 最大/最小值<br>□ 负数/零/正数 | □ **树**：空树、单节点、只有左子树、只有右子树<br>□ **链表**：空链表、单节点、环形链表<br>□ **字符串**：空串、单字符、回文、全部相同字符<br>□ **图**：无边、单节点、环、不连通 |

3. 小技巧：数组原地修改->逆序修改
4. 许多题目的关键在于推导出公式，公式推导步骤如下：
   1. 求什么？
   2. 哪些参数是变量？
   3. 变量变化的时机是什么时候？
   4. 终止条件是什么？
5. 题目想不出来的时候，手动模拟小数据推导，画出执行过程，比如递归树、DP 表格等

技巧

- 指针往复时，可以定义一个值为 -1 的变量，在开始和结束的地方反转反向，比如 [6. Z 字形变换](https://leetcode.cn/problems/zigzag-conversion/)

  ```go
  direction := -1
  if i == 0 || i == len(s)-1 {
      direction = -direction
  }
  ```

## 链表

## 二叉树

### 二分查找

### 树的遍历

## 数组

### 前缀和

**解决的问题：** 通过预构建前缀和数组，将 **频繁查询区间和** 的时间复杂度从 $O(n)$ 降低到 $O(1)$，典型的空间换时间思想。

**核心技巧：**

- 前缀和数组长度比原数组多 1（`prefixSum[0] = 0`），类似链表的 dummy head，避免边界判断
- 构造公式：`prefixSum[i+1] = prefixSum[i] + nums[i]`
- 查询公式：`sum[left, right] = prefixSum[right+1] - prefixSum[left]`（左右都是闭区间）

**应用场景：**

- 销售数据分析：快速查询任意时间段的销售总额
- 流量统计：网站流量监控，快速计算任意时间区间的访问量
- 股票分析：计算股票在任意时间段内的累计涨跌幅
- 账户余额查询：快速计算任意时间段的收支总和
- 传感器数据：快速查询温度、湿度等传感器在某时间段的平均值

详见：[前缀和算法总结](patterns/prefix_sum/README.md)

### 差分数组

**解决的问题：** 将 **频繁对数组区间进行修改** 的时间复杂度从 $O(n)$ 降低到 $O(1)$，是前缀和的逆运算。

**核心技巧：**

- 差分数组长度与原数组相同
- 构造公式：`diff[0] = nums[0]`，`diff[i] = nums[i] - nums[i-1]` (i > 0)
- 区间修改：`diff[left] += val`，`diff[right+1] -= val`（需要注意 right+1 的边界）
- 还原数组：对差分数组求前缀和

**应用场景：**

- 区间修改问题：航班预订、拼车等
- 判断是否超载：火车、拼车等场景

详见：[差分数组算法总结](patterns/difference_array/README.md)

---

## 算法技巧与注意事项

### 循环不变量（Loop Invariant）

**循环不变量** 是理解和编写正确算法的关键。它是一个在循环执行过程中始终保持为真的条件。

**示例：快慢指针**

```go
// 循环不变量：[0, slow) 区间内的元素都不等于 val
slow := 0
for fast := range nums {
    if nums[fast] != val {
        nums[slow] = nums[fast]
        slow++
    }
}
// 循环结束时，不变量仍然成立
```

**示例：二分查找**

```go
// 循环不变量：target 如果存在，必在 [left, right] 区间内
left, right := 0, len(nums)-1
for left <= right {
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        left = mid + 1  // 维护不变量：target > nums[mid]
    } else {
        right = mid - 1  // 维护不变量：target < nums[mid]
    }
}
```

**好处**：

- 帮助理解算法的正确性
- 指导边界条件的处理
- 简化调试过程

---

### 常见边界条件陷阱

#### 1. 循环条件：`<` vs `<=`

**对撞指针**：

| 循环条件        | 使用场景                       | 原因                                     |
| --------------- | ------------------------------ | ---------------------------------------- |
| `left < right`  | 成对处理元素（回文、两数之和） | `left == right` 时指向同一元素，无需处理 |
| `left <= right` | 需要处理所有元素（二分查找）   | `left == right` 时还有一个元素未检查     |

**判断技巧**：问自己"当 `left == right` 时，我还需要处理这个元素吗？"

#### 2. 指针/索引移动时机

**双指针问题**：

```go
// ❌ 错误：右边换来的元素可能也需要处理
if nums[left] == val {
    nums[left] = nums[right]
    right--
    left++  // 错误！
}

// ✅ 正确：重新检查新换来的元素
if nums[left] == val {
    nums[left] = nums[right]
    right--
} else {
    left++
}
```

**递归问题**：

```go
// ❌ 错误：可能导致无限递归
func dfs(i int) {
    if i == n {  // 边界条件不完整
        return
    }
    dfs(i + 1)
}

// ✅ 正确：完整的边界检查
func dfs(i int) {
    if i < 0 || i >= n {  // 检查上下界
        return
    }
    dfs(i + 1)
}
```

#### 3. 数组/切片越界

常见场景：

- **二分查找**：`mid = left + (right - left) / 2` 避免溢出
- **区间操作**：注意 `right + 1` 是否越界
- **链表**：检查 `node != nil` 再访问 `node.Next`
- **矩阵遍历**：检查 `i < rows && j < cols`
- **滑动窗口**：确保 `right < len(arr)` 再访问

#### 4. 整数溢出

```go
// ❌ 错误：可能溢出
mid := (left + right) / 2

// ✅ 正确
mid := left + (right - left) / 2

// ❌ 错误：乘法可能溢出
result := a * b

// ✅ 正确：先检查
if a > math.MaxInt64 / b {
    // 处理溢出
}
```

#### 5. 递归终止条件

```go
// ❌ 错误：缺少边界检查
func fibonacci(n int) int {
    return fibonacci(n-1) + fibonacci(n-2)
}

// ✅ 正确：完整的基准情况
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

#### 6. 动态规划初始化

```go
// ❌ 错误：初始值不正确
dp := make([]int, n)  // 默认为 0，可能不适合求最小值

// ✅ 正确：根据问题初始化
dp := make([]int, n)
for i := range dp {
    dp[i] = math.MaxInt  // 求最小值时初始化为最大值
}
dp[0] = 0  // 设置基准情况
```

---

### 调试策略

#### 1. 手动模拟小数据

**数组/双指针**：用 2-3 个元素手动走一遍

```
nums = [2, 3, 2], val = 2
初始：left=0, right=2
     [2, 3, 2]
      ↑     ↑
...
```

**递归/树**：画出递归树

```
        fib(4)
       /      \
    fib(3)   fib(2)
    /   \     /   \
fib(2) fib(1) ...
```

**动态规划**：列出 DP 表格

```
i:    0  1  2  3  4
dp:   0  1  1  2  3
```

#### 2. 边界条件检查清单

**通用**：

```markdown
□ 空输入（空数组、空字符串、nil）
□ 单元素
□ 两个元素（最小有意义输入）
□ 全部相同元素
□ 全部不同元素
□ 最大/最小值
□ 负数/零/正数
```

**特定类型**：

```markdown
□ 树：空树、单节点、只有左子树、只有右子树
□ 链表：空链表、单节点、环形链表
□ 字符串：空串、单字符、回文、全部相同字符
□ 图：无边、单节点、环、不连通
```

#### 3. 使用测试驱动开发

```go
tests := []struct{
    name     string
    input    []int
    expected int
}{
    {"空数组", []int{}, 0},
    {"单元素", []int{1}, 1},
    {"全部移除", []int{2,2,2}, 0},
    {"边界元素", []int{2,1,3}, 2},
    {"连续重复", []int{1,2,2,2,3}, 3},
}
```

#### 4. 添加断言和日志

```go
// 开发时添加断言
if slow < 0 || slow > len(nums) {
    panic("slow out of bounds")
}

// 调试时添加日志
fmt.Printf("left=%d, right=%d, nums=%v\n", left, right, nums)
```

---

### 常用算法模板

#### 双指针

**快慢指针**（数组去重/移除）：

```go
slow := 0
for fast := range nums {
    if 满足条件 {
        nums[slow] = nums[fast]
        slow++
    }
}
```

**对撞指针**（两数之和）：

```go
left, right := 0, len(nums)-1
for left < right {
    sum := nums[left] + nums[right]
    if sum == target {
        return []int{left, right}
    } else if sum < target {
        left++
    } else {
        right--
    }
}
```

**滑动窗口**：

```go
left := 0
for right := range s {
    // 扩大窗口
    窗口状态更新

    // 收缩窗口
    for 窗口不满足条件 {
        窗口状态更新
        left++
    }

    // 更新结果
}
```

#### 二分查找

**标准二分**：

```go
left, right := 0, len(nums)-1
for left <= right {
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
return -1
```

**寻找左边界**：

```go
left, right := 0, len(nums)
for left < right {
    mid := left + (right-left)/2
    if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid
    }
}
return left
```

#### 递归

**标准递归模板**：

```go
func dfs(参数) 返回值 {
    // 1. 终止条件
    if 到达边界 {
        return 边界值
    }

    // 2. 递归调用
    result := dfs(更小的问题)

    // 3. 处理当前层
    当前层处理逻辑

    return 结果
}
```

**回溯模板**：

```go
func backtrack(路径, 选择列表) {
    if 满足结束条件 {
        result = append(result, 路径)
        return
    }

    for 选择 in 选择列表 {
        做选择
        backtrack(路径, 新的选择列表)
        撤销选择
    }
}
```

#### 动态规划

**一维 DP**：

```go
dp := make([]int, n+1)
dp[0] = 初始值

for i := 1; i <= n; i++ {
    dp[i] = 状态转移方程
}
return dp[n]
```

**二维 DP**：

```go
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}

// 初始化
for i := 0; i <= m; i++ {
    dp[i][0] = 初始值
}

// 状态转移
for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
        dp[i][j] = 状态转移方程
    }
}
return dp[m][n]
```

#### 树的遍历

**前序遍历**（递归）：

```go
func preorder(root *TreeNode) {
    if root == nil {
        return
    }
    处理当前节点
    preorder(root.Left)
    preorder(root.Right)
}
```

**层序遍历**（BFS）：

```go
queue := []*TreeNode{root}
for len(queue) > 0 {
    size := len(queue)
    for i := 0; i < size; i++ {
        node := queue[0]
        queue = queue[1:]

        处理当前节点

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

---

### 学习建议

1. **专注一个模式**：先掌握一种模式（如双指针），做 10-15 道类似题
2. **总结不变量**：每道题都明确循环不变量或递归不变量
3. **手动模拟**：遇到问题时，用小数据手动走一遍
4. **建立模板库**：总结常见模式的代码模板
5. **定期复习**：间隔重复，巩固记忆
6. **对比相似题**：找出同一模式下不同题目的共性和差异
7. **总结易错点**：记录自己容易出错的地方

**记住**：不要追求"一次写对"，通过测试发现问题、理解原因、总结规律，才是最有效的学习方式。
