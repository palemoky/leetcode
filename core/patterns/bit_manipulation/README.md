# 位运算

位运算（Bit Manipulation）是直接对整数在内存中的二进制位进行操作的技巧。它是计算机底层运算的基础，具有 **执行速度快、代码简洁** 的特点，在算法优化、状态压缩、权限管理等场景中有广泛应用。

!!! Quote "核心优势"

    位运算直接操作二进制位，**无需进位和借位**，是 CPU 最原始、最快速的运算方式。一个位运算往往能替代多行条件判断或循环代码。

## 基础位运算符

| 运算符   | 名称            | 说明                        | 示例                            | 结果          |
| -------- | --------------- | --------------------------- | ------------------------------- | ------------- |
| `&`      | 按位与（AND）   | 两位都为 1 时结果为 1       | `5 & 3` (`101 & 011`)           | `1` (`001`)   |
| `&#124;` | 按位或（OR）    | 任一位为 1 时结果为 1       | `5 &#124; 3` (`101 &#124; 011`) | `7` (`111`)   |
| `^`      | 按位异或（XOR） | 两位不同时结果为 1          | `5 ^ 3` (`101 ^ 011`)           | `6` (`110`)   |
| `~`      | 按位取反（NOT） | 0 变 1，1 变 0              | `~5` (`~0101`)                  | `-6` (`1010`) |
| `<<`     | 左移            | 向左移动 n 位，右边补 0     | `5 << 1` (`101 << 1`)           | `10` (`1010`) |
| `>>`     | 右移            | 向右移动 n 位，左边补符号位 | `5 >> 1` (`101 >> 1`)           | `2` (`010`)   |

!!! Note "符号位说明"

    - **有符号右移 `>>`**：左边补符号位（正数补 0，负数补 1）
    - **无符号右移 `>>>`**（部分语言如 Java）：左边始终补 0
    - Go 语言中，右移行为取决于操作数类型（有符号/无符号）

## 核心技巧

### 1. 判断奇偶

```go
// 判断 n 是否为奇数
func isOdd(n int) bool {
    return n & 1 == 1  // 最低位为 1 则为奇数
}
```

**原理**：奇数的二进制最低位必为 1，偶数为 0

### 2. 交换两数（无需临时变量）

```go
func swap(a, b int) (int, int) {
    a ^= b
    b ^= a
    a ^= b
    return a, b
}
```

**原理**：异或满足 `a ^ a = 0` 和 `a ^ 0 = a`

### 3. 获取/设置/清除第 k 位

```go
// 获取第 k 位（从右往左，从 0 开始）
func getBit(n, k int) int {
    return (n >> k) & 1
}

// 设置第 k 位为 1
func setBit(n, k int) int {
    return n | (1 << k)
}

// 清除第 k 位（设为 0）
func clearBit(n, k int) int {
    return n & ^(1 << k)
}

// 切换第 k 位（0 变 1，1 变 0）
func toggleBit(n, k int) int {
    return n ^ (1 << k)
}
```

### 4. 清除最低位的 1

```go
func clearLowestBit(n int) int {
    return n & (n - 1)
}
```

**应用**：

- 统计二进制中 1 的个数（Brian Kernighan 算法）
- 判断是否为 2 的幂

**示例**：

```
n     = 12 = 1100
n - 1 = 11 = 1011
n & (n-1) = 1000 = 8
```

### 5. 获取最低位的 1

```go
func getLowestBit(n int) int {
    return n & -n
}
```

**原理**：`-n` 是 `n` 的补码（取反加 1），与 `n` 按位与只保留最低位的 1

**应用**：树状数组（Fenwick Tree）

### 6. 判断是否为 2 的幂

```go
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n - 1)) == 0
}
```

**原理**：2 的幂的二进制只有一个 1，减 1 后所有位取反

### 7. 统计二进制中 1 的个数

```go
// 方法 1：Brian Kernighan 算法
func countOnes(n int) int {
    count := 0
    for n > 0 {
        n = n & (n - 1)  // 每次清除最低位的 1
        count++
    }
    return count
}

// 方法 2：逐位检查
func countOnes2(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
```

### 8. 找出唯一的数（其他数都出现两次）

```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num  // 相同的数异或为 0
    }
    return result
}
```

**原理**：`a ^ a = 0`，`a ^ 0 = a`，异或满足交换律和结合律

## 进阶技巧

### 1. 位掩码（Bitmask）

用一个整数的每一位表示一个状态，常用于状态压缩 DP。

```go
// 示例：集合的子集枚举
func subsets(nums []int) [][]int {
    n := len(nums)
    result := [][]int{}

    // 遍历所有可能的子集（2^n 个）
    for mask := 0; mask < (1 << n); mask++ {
        subset := []int{}
        for i := 0; i < n; i++ {
            if mask & (1 << i) != 0 {  // 检查第 i 位是否为 1
                subset = append(subset, nums[i])
            }
        }
        result = append(result, subset)
    }

    return result
}
```

### 2. 找出两个只出现一次的数（其他数都出现两次）

```go
func singleNumber(nums []int) []int {
    // 第一步：所有数异或，得到两个唯一数的异或结果
    xor := 0
    for _, num := range nums {
        xor ^= num
    }

    // 第二步：找到 xor 中任意一个为 1 的位（两数在此位不同）
    diff := xor & -xor  // 获取最低位的 1

    // 第三步：根据这一位将数组分成两组
    a, b := 0, 0
    for _, num := range nums {
        if num & diff == 0 {
            a ^= num
        } else {
            b ^= num
        }
    }

    return []int{a, b}
}
```

### 3. 位运算实现加法（无进位）

```go
func add(a, b int) int {
    for b != 0 {
        carry := (a & b) << 1  // 计算进位
        a = a ^ b              // 无进位加法
        b = carry              // 将进位赋给 b
    }
    return a
}
```

## 常见应用场景

### 1. 权限管理

```go
const (
    READ   = 1 << 0  // 0001 = 1
    WRITE  = 1 << 1  // 0010 = 2
    EXECUTE = 1 << 2  // 0100 = 4
    DELETE = 1 << 3  // 1000 = 8
)

// 添加权限
func addPermission(perm, newPerm int) int {
    return perm | newPerm
}

// 移除权限
func removePermission(perm, removePerm int) int {
    return perm & ^removePerm
}

// 检查权限
func hasPermission(perm, checkPerm int) bool {
    return perm & checkPerm == checkPerm
}

// 示例
perm := READ | WRITE  // 0011 = 3
perm = addPermission(perm, EXECUTE)  // 0111 = 7
hasRead := hasPermission(perm, READ)  // true
```

### 2. 状态压缩 DP

在动态规划中，用一个整数表示多个状态，节省空间。

**经典题目**：

- [LeetCode 698. Partition to K Equal Sum Subsets](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/)
- [LeetCode 847. Shortest Path Visiting All Nodes](https://leetcode.com/problems/shortest-path-visiting-all-nodes/)

### 3. 快速幂运算

```go
func pow(x, n int) int {
    result := 1
    for n > 0 {
        if n & 1 == 1 {  // 如果当前位为 1
            result *= x
        }
        x *= x
        n >>= 1
    }
    return result
}
```

**时间复杂度**：\(O(\log n)\)

## 位运算技巧总结表

| 技巧           | 公式                        | 说明           |
| -------------- | --------------------------- | -------------- |
| 去除最后一位   | `n >> 1`                    | 相当于 `n / 2` |
| 最后一位变 0   | `n & (n - 1)`               | 清除最低位的 1 |
| 最后一位变 1   | `n &#124; 1`                | 设置最低位为 1 |
| 最后一位取反   | `n ^ 1`                     | 0 变 1，1 变 0 |
| 获取最后一位   | `n & 1`                     | 判断奇偶       |
| 获取最低位的 1 | `n & -n`                    | 树状数组核心   |
| 判断 2 的幂    | `n > 0 && (n & (n-1)) == 0` | 只有一个 1     |
| 乘以 2^k       | `n << k`                    | 左移 k 位      |
| 除以 2^k       | `n >> k`                    | 右移 k 位      |

## 经典题目清单

### 基础题

- [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) — 统计二进制中 1 的个数
- [136. Single Number](https://leetcode.com/problems/single-number/) — 找出唯一的数
- [231. Power of Two](https://leetcode.com/problems/power-of-two/) — 判断是否为 2 的幂
- [338. Counting Bits](https://leetcode.com/problems/counting-bits/) — 计算 0 到 n 每个数的 1 的个数

### 进阶题

- [137. Single Number II](https://leetcode.com/problems/single-number-ii/) — 其他数出现 3 次
- [260. Single Number III](https://leetcode.com/problems/single-number-iii/) — 找出两个唯一的数
- [201. Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/) — 区间按位与
- [371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) — 不用加减法实现加法

### 高级题

- [78. Subsets](https://leetcode.com/problems/subsets/) — 子集枚举（位掩码）
- [698. Partition to K Equal Sum Subsets](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/) — 状态压缩 DP
- [847. Shortest Path Visiting All Nodes](https://leetcode.com/problems/shortest-path-visiting-all-nodes/) — 状态压缩 BFS

## 性能对比

| 操作     | 普通方法       | 位运算方法           | 性能提升  |
| -------- | -------------- | -------------------- | --------- |
| 判断奇偶 | `n % 2 == 1`   | `n & 1 == 1`         | 约 2-3 倍 |
| 乘以 2   | `n * 2`        | `n << 1`             | 约 2 倍   |
| 除以 2   | `n / 2`        | `n >> 1`             | 约 2 倍   |
| 取模 2^k | `n % (1 << k)` | `n & ((1 << k) - 1)` | 约 3-5 倍 |

!!! Warning "注意事项"

    - 位运算优先级较低，建议加括号：`(a & b) == 0` 而非 `a & b == 0`
    - 负数的位运算涉及补码，需要特别注意
    - 过度使用位运算会降低代码可读性，应在性能关键处使用

## 总结

位运算是算法优化的利器，掌握以下要点：

1. ✅ **基础运算符**：`&`、`|`、`^`、`~`、`<<`、`>>`
2. ✅ **核心技巧**：清除最低位 `n & (n-1)`、获取最低位 `n & -n`
3. ✅ **经典应用**：权限管理、状态压缩、快速幂
4. ✅ **刷题策略**：从 Single Number 系列开始，逐步掌握位掩码和状态压缩

!!! Tip "学习建议"

    位运算的精髓在于**用二进制的视角思考问题**。建议：

    1. 手动推导几个例子的二进制过程
    2. 熟记常用技巧（如 `n & (n-1)`）
    3. 从简单题开始，逐步理解位运算的威力
