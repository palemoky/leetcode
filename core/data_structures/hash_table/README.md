# 哈希表

哈希表（Hash Table）是一种基于**哈希函数**实现的数据结构，通过键（Key）直接访问值（Value），实现 \(O(1)\) 时间复杂度的查找、插入和删除操作。它是"**空间换时间**"思想的经典体现，常用于解决查找、计数、去重、分组、记忆优化、数组索引映射等问题。

!!! Note "空间换时间的智慧"

    "空间换时间"是计算机科学中的核心权衡思想，在软件和硬件设计中无处不在：

    - **CPU 缓存**：用 L1/L2/L3 缓存换取更快的数据访问
    - **数据库索引**：用额外的索引空间换取查询速度
    - **CDN**：在全球部署服务器换取更快的内容分发
    - **Redis/Memcached**：用内存存储换取毫秒级响应
    - **预计算**：提前计算结果（如前缀和）换取查询效率

## 核心概念

### 哈希函数

哈希函数将键映射到数组索引：

```
index = hash(key) % capacity
```

**理想哈希函数的特性**：

- **确定性**：相同的键总是产生相同的哈希值
- **均匀分布**：键均匀分布在哈希表中，减少冲突
- **高效计算**：哈希函数本身的计算速度要快

### 哈希冲突

当两个不同的键映射到同一个索引时，就发生了**哈希冲突**。

**解决方法**：

| 方法                                   | 原理                                               | 优点                      | 缺点                                    |
| -------------------------------------- | -------------------------------------------------- | ------------------------- | --------------------------------------- |
| **链地址法**<br/>（Separate Chaining） | 每个索引位置存储一个链表，冲突的键值对追加到链表中 | 实现简单<br/>不限容量     | 链表过长时性能退化<br/>需要额外指针空间 |
| **开放寻址法**<br/>（Open Addressing） | 冲突时按某种规则寻找下一个空位                     | 内存利用率高<br/>缓存友好 | 删除操作复杂<br/>负载因子高时性能下降   |

**开放寻址法的探测方式**：

- **线性探测**：`index = (hash + i) % capacity`
- **二次探测**：`index = (hash + i²) % capacity`
- **双重哈希**：`index = (hash1 + i × hash2) % capacity`

### 负载因子

```
负载因子 = 元素数量 / 哈希表容量
```

- **链地址法**：通常在负载因子 > 0.75 时扩容
- **开放寻址法**：通常在负载因子 > 0.5 时扩容

## 常见实现

### 不同语言的哈希表

| 语言           | 数据结构         | 底层实现               | 特点                  |
| -------------- | ---------------- | ---------------------- | --------------------- |
| **Go**         | `map[K]V`        | 开放寻址 + 链表        | 无序，并发不安全      |
| **Python**     | `dict`           | 开放寻址（二次探测）   | 3.7+ 保持插入顺序     |
| **Java**       | `HashMap`        | 链地址法（红黑树优化） | 链表长度 > 8 转红黑树 |
| **C++**        | `unordered_map`  | 链地址法               | 基于桶的实现          |
| **JavaScript** | `Map` / `Object` | 引擎优化               | `Map` 保持插入顺序    |

### Go 语言示例

```go
// 创建哈希表
m := make(map[string]int)

// 插入/更新
m["apple"] = 5
m["banana"] = 3

// 查找
value, exists := m["apple"]
if exists {
    fmt.Println(value) // 5
}

// 删除
delete(m, "banana")

// 遍历（无序）
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// 长度
fmt.Println(len(m))
```

## 时间复杂度

| 操作     | 平均时间复杂度 | 最坏时间复杂度 |
| -------- | -------------- | -------------- |
| **查找** | \(O(1)\)       | \(O(n)\)       |
| **插入** | \(O(1)\)       | \(O(n)\)       |
| **删除** | \(O(1)\)       | \(O(n)\)       |

!!! Warning "最坏情况"

    当所有键都发生哈希冲突时，退化为链表，时间复杂度为 \(O(n)\)。

## 经典应用场景

### 1. 快速查找

**问题**：判断元素是否存在

```go
// 判断数组中是否存在重复元素
func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
```

**经典题目**：

- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)
- [1. Two Sum](https://leetcode.com/problems/two-sum/)

### 2. 计数统计

**问题**：统计元素出现次数

```go
// 统计字符串中每个字符的出现次数
func countChars(s string) map[rune]int {
    count := make(map[rune]int)
    for _, ch := range s {
        count[ch]++
    }
    return count
}
```

**经典题目**：

- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)
- [383. Ransom Note](https://leetcode.com/problems/ransom-note/)
- [387. First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/)

### 3. 去重

**问题**：移除重复元素

```go
// 数组去重
func removeDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, num := range nums {
        if !seen[num] {
            seen[num] = true
            result = append(result, num)
        }
    }
    return result
}
```

### 4. 分组映射

**问题**：将元素按某种规则分组

```go
// 字母异位词分组
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, str := range strs {
        // 排序后的字符串作为键
        key := sortString(str)
        groups[key] = append(groups[key], str)
    }

    result := [][]string{}
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}
```

**经典题目**：

- [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

### 5. 缓存/记忆化

**问题**：存储计算结果避免重复计算

```go
// 斐波那契数列（记忆化递归）
func fib(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    if val, exists := memo[n]; exists {
        return val
    }
    memo[n] = fib(n-1, memo) + fib(n-2, memo)
    return memo[n]
}
```

### 6. 索引映射

**问题**：建立值到索引的映射

```go
// 两数之和
func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, exists := indexMap[complement]; exists {
            return []int{j, i}
        }
        indexMap[num] = i
    }
    return nil
}
```

**经典题目**：

- [1. Two Sum](https://leetcode.com/problems/two-sum/)
- [167. Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## 哈希表 vs 其他数据结构

| 特性         | 哈希表   | 数组     | 链表     | 二叉搜索树    |
| ------------ | -------- | -------- | -------- | ------------- |
| **查找**     | \(O(1)\) | \(O(n)\) | \(O(n)\) | \(O(\log n)\) |
| **插入**     | \(O(1)\) | \(O(n)\) | \(O(1)\) | \(O(\log n)\) |
| **删除**     | \(O(1)\) | \(O(n)\) | \(O(1)\) | \(O(\log n)\) |
| **有序性**   | ❌       | ✓        | ❌       | ✓             |
| **空间效率** | 低       | 高       | 低       | 中            |

## 使用技巧

### 1. 选择合适的键类型

**可哈希的类型**：

- ✅ 整数、浮点数、字符串
- ✅ 元组（不可变）
- ❌ 列表、字典（可变类型）

### 2. 自定义哈希键

当需要复杂的键时，可以将多个值组合成字符串或元组：

```go
// 使用字符串作为复合键
key := fmt.Sprintf("%d,%d", x, y)

// 或使用结构体（需实现 comparable）
type Point struct {
    X, Y int
}
m := make(map[Point]int)
```

### 3. 默认值处理

```go
// 方法 1：检查键是否存在
if val, exists := m[key]; exists {
    // 键存在
} else {
    // 键不存在，使用默认值
}

// 方法 2：直接使用零值
count := m[key] // 如果不存在，返回 0
m[key] = count + 1
```

### 4. 避免并发问题

Go 的 `map` 不是并发安全的，需要使用 `sync.Map` 或加锁：

```go
import "sync"

var (
    m  = make(map[string]int)
    mu sync.RWMutex
)

// 读操作
mu.RLock()
value := m[key]
mu.RUnlock()

// 写操作
mu.Lock()
m[key] = value
mu.Unlock()
```

## 常见陷阱

### 1. 遍历顺序不确定

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k, v := range m {
    // 每次遍历顺序可能不同！
    fmt.Println(k, v)
}
```

!!! Tip "解决方案"

    如果需要有序遍历，先将键排序：
    ```go
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Println(k, m[k])
    }
    ```

### 2. 遍历时修改

```go
// ❌ 错误：遍历时删除可能导致未定义行为
for k := range m {
    if someCondition(k) {
        delete(m, k) // 危险！
    }
}

// ✅ 正确：先收集要删除的键
toDelete := []string{}
for k := range m {
    if someCondition(k) {
        toDelete = append(toDelete, k)
    }
}
for _, k := range toDelete {
    delete(m, k)
}
```

### 3. 值是指针 vs 值类型

```go
// 值类型：无法直接修改
type Person struct {
    Name string
    Age  int
}
m := make(map[string]Person)
m["alice"] = Person{"Alice", 30}
m["alice"].Age = 31 // ❌ 编译错误！

// 解决方案 1：使用指针
m2 := make(map[string]*Person)
m2["alice"] = &Person{"Alice", 30}
m2["alice"].Age = 31 // ✅

// 解决方案 2：重新赋值
p := m["alice"]
p.Age = 31
m["alice"] = p // ✅
```

## 经典题目清单

### 基础题

- [1. Two Sum](https://leetcode.com/problems/two-sum/) — 两数之和（入门必做）
- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) — 存在重复元素
- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/) — 有效的字母异位词
- [383. Ransom Note](https://leetcode.com/problems/ransom-note/) — 赎金信

### 进阶题

- [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/) — 字母异位词分组
- [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) — 最长连续序列
- [146. LRU Cache](https://leetcode.com/problems/lru-cache/) — LRU 缓存（哈希表 + 双向链表）
- [387. First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/) — 字符串中的第一个唯一字符

### 高级题

- [149. Max Points on a Line](https://leetcode.com/problems/max-points-on-a-line/) — 直线上最多的点数
- [380. Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/) — 常数时间插入、删除和获取随机元素
- [460. LFU Cache](https://leetcode.com/problems/lfu-cache/) — LFU 缓存

## 总结

哈希表是算法面试中最常用的数据结构之一，掌握以下要点：

1. ✅ **理解原理**：哈希函数、冲突解决、负载因子
2. ✅ **熟练应用**：查找、计数、去重、分组、缓存
3. ✅ **注意陷阱**：遍历顺序、并发安全、值类型修改
4. ✅ **时间复杂度**：平均 \(O(1)\)，最坏 \(O(n)\)
5. ✅ **空间换时间**：用额外空间换取查找效率
