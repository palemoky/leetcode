# Go 语言核心技巧

本篇整理了 Go 语言在 LeetCode 刷题中的 **“生存指南”**，涵盖了从数据结构实现到语言陷阱的全方位技巧。

## 数据结构实现指南

Go 的标准库以“通用”著称，不像 C++ STL 或 Python `collections` 那样开箱即用，因此熟练掌握以下模板是刷题的基础。

### 堆 (Heap) / 优先队列

Go 需要实现 `heap.Interface` 接口。为了避免重复写模板，推荐使用 **函数闭包** 实现通用堆：

```go
import "container/heap"

// 通用堆模板：只需定义 Less 函数
type Hp struct {
    sort.IntSlice
    LessFunc func(i, j int) bool
}

func (h Hp) Less(i, j int) bool { return h.LessFunc(h.IntSlice[i], h.IntSlice[j]) }
func (h *Hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 使用示例：
func solve() {
    // 建立大顶堆
    pq := &Hp{IntSlice: []int{}, LessFunc: func(i, j int) bool { return i > j }}
    heap.Init(pq)
    heap.Push(pq, 10)
    top := heap.Pop(pq).(int)
}
```

### 栈 (Stack) & 队列 (Queue)

Go 使用 `slice` 模拟栈和队列。

| 结构     | 操作 | 代码                                      | 复杂度 | 备注               |
| :------- | :--- | :---------------------------------------- | :----- | :----------------- |
| **栈**   | 入栈 | `st = append(st, v)`                      | $O(1)$ |                    |
|          | 出栈 | `v := st[len(st)-1]; st = st[:len(st)-1]` | $O(1)$ | 记得判空           |
|          | 栈顶 | `st[len(st)-1]`                           | $O(1)$ |                    |
| **队列** | 入队 | `q = append(q, v)`                        | $O(1)$ |                    |
|          | 出队 | `v := q[0]; q = q[1:]`                    | $O(1)$ | **注意内存泄漏！** |

!!! warning "队列的内存陷阱"

    使用 `q = q[1:]` 会导致底层数组无法释放前半部分的空间。如果队列生命周期很长，建议在长度过大且有效元素较少时重建切片，或者使用环形数组。

### 集合 (Set)

Go 没有内置 `set`，通常用 `map[T]struct{}` 模拟。

```go
// 初始化
set := make(map[int]struct{})

// 添加
set[1] = struct{}{}

// 查找
if _, exists := set[1]; exists { ... }

// 删除
delete(set, 1)
```

**为什么用 `struct{}`？**
空结构体 `struct{}` 在 Go 中 **不占用内存空间**（size 为 0），比 `map[int]bool` 更节省内存。

---

## 避坑指南 (The Trap)

### Map 的随机性

**陷阱**：Go 的 `map` 遍历顺序是 **完全随机** 的！且每次运行都不一样。
**后果**：如果题目要求按顺序输出分组结果，直接遍历 `map` 必挂。
**解法**：先提取 Key 到切片，对 Key 排序，再遍历。

```go
keys := make([]int, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys) // 关键步骤
for _, k := range keys {
    val := m[k]
    // ...
}
```

### 循环变量引用陷阱 (Pre-Go 1.22)

**陷阱**：在 `for` 循环中获取元素指针或闭包引用。
**现状**：Go 1.22+ 已经修复了此问题，循环变量每次迭代都会创建新实例。但在老版本环境中（某些笔试平台），仍需注意：

```go
// 老版本解法：中间变量接引
for _, v := range nums {
    v := v // 影子变量
    result = append(result, &v)
}
```

### 二维切片初始化

**陷阱**：不能像 Python 那样 `[[0]*n]*m`（这是浅拷贝），Go 必须手动循环。

```go
// 初始化 dp[m][n]
dp := make([][]int, m)
for i := range dp {
    dp[i] = make([]int, n)
}
```

---

## LeetCode 实用技巧

### 字符串高效处理

Go 的 `string` 是不可变的，`[]byte` 是可变的。

- **大量拼接**：务必使用 `strings.Builder` (推荐) 或 `bytes.NewBuffer`。
  - `strings.Builder`：专门为构建字符串优化，零拷贝转换为 string。
  - `bytes.NewBuffer`：更通用，适合同时也需要字节流操作的场景。
- **修改字符**：先转 `[]byte`，改完再转回 `string`。
- **遍历字符**：
  - `for i := 0; i < len(s); i++`：按 **字节 (byte)** 遍历，适合 ASCII。
  - `for i, r := range s`：按 **字符 (rune)** 遍历，适合含中文等多字节字符。

### 常用简易辅助函数

LeetCode 环境通常没有 `generics` 版本的 `Max/Min/Abs`（Go 1.21 前），建议背诵以下三行：

```go
func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
func abs(a int) int    { if a < 0 { return -a }; return a }
```

### 排序黑科技

`sort.Slice` 极其灵活，配合匿名函数可以对任意结构排序：

```go
//按两个字段排序：先按 Age 升序，Age 相同按 Name 降序
sort.Slice(people, func(i, j int) bool {
    if people[i].Age != people[j].Age {
        return people[i].Age < people[j].Age
    }
    return people[i].Name > people[j].Name
})
```

---

### 二分查找神器

标准库 `sort.Search(n, f)` 是刷题神器，它返回 `[0, n)` 中 **第一个满足 `f(i) == true`** 的下标。

```go
// 二分查找第一个 >= target 的位置 (Lower Bound)
// 如果所有元素都 < target，返回 len(nums)
idx := sort.Search(len(nums), func(i int) bool {
    return nums[i] >= target
})

// 二分查找第一个 > target 的位置 (Upper Bound)
idx := sort.Search(len(nums), func(i int) bool {
    return nums[i] > target
})
```

#### 原理

Go 的 `sort.Search` 相当于 C++ 的 `std::lower_bound`。只要 `f(i)` 满足单调性（先 `false` 后 `true`），它就能稳定找到 `true/false` 的临界点。

### 位运算黑科技 (`math/bits`)

Go 提供了超级好用的位运算包 `math/bits`，无需手写位操作：

| 函数                    | 功能                             | 对应 C++                |
| :---------------------- | :------------------------------- | :---------------------- |
| `bits.OnesCount(x)`     | 统计二进制中 1 的个数            | `__builtin_popcount`    |
| `bits.Len(x)`           | 二进制长度 (最高位 1 的位置 + 1) | `32 - __builtin_clz(x)` |
| `bits.Reverse(x)`       | 二进制翻转                       | -                       |
| `bits.RotateLeft(x, k)` | 循环左移 k 位 (k<0 为右移)       | -                       |

```go
// 常用场景：状态压缩 DP 统计状态中 1 的个数
cnt := bits.OnesCount(state)
```

### 随机数 (`math/rand`)

对于需要随机化的算法（如 **快速选择**、**Treap**），Go 的 `rand` 需要注意随机种子：

```go
// Go 1.22+ (math/rand/v2) - 推荐
// 自动初始化 Seed，性能更好
idx := rand.IntN(len(nums)) // 生成 [0, n) 的随机数

// Go < 1.22 (math/rand)
// 注意函数名是 Intn (小写 n)
idx := rand.Intn(len(nums))
```

---

## 核心语法与底层陷阱

### 变量声明对比

| 声明方式            | 底层     | 优缺点                    | 场景            | 备注                        |
| ------------------- | -------- | ------------------------- | --------------- | --------------------------- |
| `var s []T`         | nil      | 不分配内存，可直接 append | 全局/延迟初始化 | 具名返回值 `[]T` 也是此类型 |
| `s := []T{}`        | 空数组   | 非 nil，序列化为 `[]`     | 明确需要空列表  | 扩容时会有性能开销          |
| `s := make([]T, n)` | 零值数组 | **性能最优**，一次分配    | 已知大小        |                             |

### 切片 (Slice) 的暗坑

切片不仅仅是一个动态数组，它的底层是一个 `struct { ptr, len, cap }`。

- **引用传递**：切片传参是“值传递”，但传递的是结构体副本（指针指向同一底层数组）。在函数内修改元素会影响原切片，但 **append 导致扩容** 后，会指向新数组，不再影响原切片。
- **内存泄漏**：`s2 := s1[100:101]`。若 `s1` 是 1GB 的大数组，只要 `s2` 还在，这 1GB 内存就不会被回收。
  - **解法**：使用 `copy` 复制所需数据到新切片。

### Map 的关键特性

- **无序性**：遍历顺序随机。
- **非并发安全**：多协程并发读写会直接 `fatal error`（panic 无法捕获）。需要加锁 `sync.RWMutex` 或使用 `sync.Map`。
- **内存不收缩**：删除 Key 不会释放内存，只有置 `nil` 才会由 GC 回收。

### Defer 的“三大天条”

1.  **LIFO 执行**：后 `defer` 的先执行（栈顺序）。
2.  **参数预计算**：`defer fmt.Println(i)` 会在 `defer`声明时**立即计算** `i` 的值并压栈。若要获取最终值，需使用闭包 `defer func() { fmt.Println(i) }()`。
3.  **修改返回值**：`defer` 可以读取并修改 **具名返回值**（Named Return Value）。

```go
func test() (result int) {
    defer func() { result++ }()
    return 1 // result = 1 -> defer -> result = 2
}
```

### Interface 的 Nil 判空

**记住：Interface 包含 `(Type, Value)` 两个部分。只有 Type 和 Value 都为 nil，Interface 才等于 nil。**

```go
var p *int = nil
var i interface{} = p
// i != nil，因为 i 的 Type 是 *int，Value 是 nil
if i == nil { exclude() } // ❌ 永远进不去
if v := reflect.ValueOf(i); v.IsNil() { } // ✅ 正确判断
```

### Channel 状态

| 操作             | nil channel | closed channel          | active channel        |
| :--------------- | :---------- | :---------------------- | :-------------------- |
| **close**        | panic       | panic                   | 成功关闭 (不能重复关) |
| **send (ch <-)** | 永久阻塞    | panic                   | 阻塞直到由接收方      |
| **recv (<- ch)** | 永久阻塞    | 立即返回零值 (v, false) | 阻塞直到有发送方      |

---

## 类型转换

### 基础转换

| 转换                 | 代码                      | 备注               |
| :------------------- | :------------------------ | :----------------- |
| `string` -> `int`    | `v, _ := strconv.Atoi(s)` | 忽略错误仅用于刷题 |
| `int` -> `string`    | `s := strconv.Itoa(v)`    | 最快               |
| `byte` -> `int`      | `int(c - '0')`            | 字符转数字         |
| `int` -> `byte`      | `byte(v + '0')`           | 数字转字符         |
| `string` -> `[]byte` | `[]byte(s)`               | 发生内存拷贝       |
| `[]byte` -> `string` | `string(b)`               | 发生内存拷贝       |

### 零拷贝转换 (Unsafe)

在追求极致性能（如千万级 QPS）时使用，普通刷题 **慎用**。

```go
import "unsafe"

// string -> []byte (只读，修改会导致 panic)
func StringToBytes(s string) []byte {
    return unsafe.Slice(unsafe.StringData(s), len(s))
}

// []byte -> string (零拷贝)
func BytesToString(b []byte) string {
    return unsafe.String(unsafe.SliceData(b), len(b))
}
```

## 指针

### 参数传递机制

**值传递（Pass by Value）**：

- C、Java、Go：所有参数都是值传递，包括指针（传递的是指针地址的副本）
- 特点：修改参数不影响原变量，除非传递指针并通过指针修改

**对象引用传递（Pass by Object-Reference）**：

- Python、JavaScript、Ruby：传递的是对象引用的副本
- 特点：可以修改对象内容，但重新赋值不影响原变量
- 注意：这 **不是真正的引用传递**，而是"共享传递"（Call by Sharing）

**引用传递（Pass by Reference）**：

- C++、C#、PHP、Swift：提供特殊语法，参数成为原变量的"别名"
- 语法：C++/PHP 用 `&`，Swift 用 `inout`，C# 用 `ref`/`out`
- 特点：修改参数直接影响原变量，包括重新赋值

---

### Go 的指针特性

虽然所有变量都会指向内存地址，但只有语言允许显式 **取地址**（`&`）、 **解引用**（`*`）、**操作地址本身** 时，才被认为支持指针操作。

Go 和 C 一样，支持 **多级指针**（如 `**T`、`***T`）。

![Pointer Memory Layout](pointer_memory_layout.webp)

---

### 多级指针示例

以验证二叉搜索树为例，需要在递归中修改外部变量 `prev`：

```go
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode  // 一级指针
	return inorder(root, &prev)  // 传递指针的地址（二级指针）
}

func inorder(node *TreeNode, prev **TreeNode) bool {  // 接收二级指针
	if node == nil {
		return true
	}

	if !inorder(node.Left, prev) {
		return false
	}

	// 通过二级指针修改外部变量
	if *prev != nil && node.Val <= (*prev).Val {
		return false
	}
	*prev = node  // 解引用赋值

	return inorder(node.Right, prev)
}
```

**问题**：

- 二级指针 `**TreeNode` 语法复杂
- 需要频繁解引用 `*prev`、`(*prev).Val`
- 容易出错，可读性差

---

### 用闭包避免多级指针

Go 社区不鼓励使用多级指针，推荐用 **闭包** 来捕获外部变量：

```go
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode  // 闭包自动捕获这个变量

    var inorder func(*TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }

        if !inorder(node.Left) {
            return false
        }

        // 直接访问外部变量，无需解引用
        if prev != nil && node.Val <= prev.Val {
            return false
        }
        prev = node  // 直接赋值

        return inorder(node.Right)
    }

    return inorder(root)
}
```

**优势**：

- 无需二级指针，代码更简洁
- 无需解引用，可读性更好
- 符合 Go 社区最佳实践

!!! Note

    Go、Python、Java、JavaScript 等语言会 **隐式自动捕获** 外部变量，而 PHP、C++ 则需要 **显式声明** 被捕获的变量（如 PHP 的 `use (&$var)`，C++ 的 `[&var]`）。
