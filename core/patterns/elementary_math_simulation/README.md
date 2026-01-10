# 初等数学模拟

在算法问题中，当处理的数字长度超过了标准整数类型（如 int64）的表示范围时，我们必须回归最基本的方法：模拟人类手动进行竖式计算的过程。这一模式的核心是将数字（通常以字符串或链表形式给出）拆解成逐个的数位，然后模拟加、减、乘、除等运算中的 **进位**、**借位** 等逻辑。

**通用核心要素：**

- **指针**: 从低位到高位遍历数字（如字符串的末尾，链表的头部）。
- **状态变量**: 如加法的 carry（进位）或减法的 borrow（借位）。
- **结果构造器**: 使用 **可变缓冲区** 高效构建结果，避免在循环中重复创建和销毁不可变的字符对象。因为每次拼接都可能需要分配一块新内存并拷贝旧值，导致 `O(N²)` 的时间复杂度。而可变缓冲区则是将所有小片段追加进去（这个过程通常是摊销 `O(1)` 的），最后再从缓冲区一次性生成最终的字符串。不同语言的处理方式如下：
  - Go： `[]byte` 或 `strings.Builder`
  - Rust：`String::push_str()` 或 `String::push()`
  - Python：`list.append()` + `str.join()`

---

## 竖式加、减、乘、除

```
      99     |     ¹¹      |      123     |        246
    +  1     |     100     |    × 456     |       ----
     ¹¹¹     |    -  1     |      ¹¹      |    4 ) 987
    -----    |    -----    |    ------    |       -8
     100     |      99     |      738     |       ----
             |             |     ¹¹       |        18
             |             |     615      |       -16
             |             |    ¹¹        |       ----
             |             |    492       |         27
             |             |    -----     |        -24
             |             |    56088     |        ---
             |             |              |          3
```

## 不同运算规则对比

我们以 base 表示进制，如二进制为 2，十进制为 10。根据竖式可得，不同运算方式数学规律如下：

| 运算           | 当前位                               | 进位<br />（舍弃小数部分）         | 借位                                  | 去掉当前位            | 核心公式                             | 算法特点                                                                                                                                                                                                                                                                                                                                   | 相关题目                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| -------------- | ------------------------------------ | ---------------------------------- | ------------------------------------- | --------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 加法           | current = `sum % base`               | carry = `sum/ base`                |                                       |                       | sum = `x+y+carry`                    | 最基础、最常见的类型。<br />循环条件 `i >= 0 \|\| j >= 0 \|\| carry > 0` 应包含 `carry > 0`，以处理最高位的进位。对于字符串结果，通常需要最后反转。如果是链表，则使用 dummy 虚拟头节点和一个 current 指针来构建结果链表                                                                                                                    | LeetCode #415. [字符串相加](https://leetcode.cn/problems/add-strings) [Add Strings](https://leetcode.com/problems/add-strings) (十进制加法) <br />LeetCode #67. [二进制求和](https://leetcode.cn/problems/add-binary/) [Add Binary](https://leetcode.com/problems/add-binary/) (二进制加法, BASE=2) <br />LeetCode #2. [两数相加](https://leetcode.cn/problems/add-two-numbers) [Add Two Numbers](https://leetcode.com/problems/add-two-numbers) (链表形式，无需反转) <br />LeetCode #66. [加一](https://leetcode.cn/problems/plus-one/) [Plus One](https://leetcode.com/problems/plus-one/) (简化版加法，加一个常数 1) <br />LeetCode #989. [数组形式的整数加法](https://leetcode.cn/problems/add-to-array-form-of-integer/) [Add to Array-Form of Integer](https://leetcode.com/problems/add-to-array-form-of-integer/) (简化版字符串相加) |
| 减法           | current = `difference % base`        |                                    | borrow = `x - y - borrow < 0 ? 1 : 0` |                       | difference = `x - y - borrow + base` | 1. 减法是加法的逆运算，核心在于处理 **借位**。<br />2. 比较 a 和 b 的大小以确定符号，然后用大数减小数。需要先判断结果的正负，处理前导零。                                                                                                                                                                                                  | LeetCode 上没有直接的减法题，但这是大厂面试的常见变体，必须掌握                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| 乘法           | current = `(product + carry) % base` | carry = `(product + carry) / base` |                                       |                       | product = `num1[i] * num2[j]`        | 模拟的是一个数乘以另一个数的每一位，然后将所有中间结果相加的过程。<br />核心逻辑:<br />1. 初始化一个足够长的结果数组 res，长度为 `len(a) + len(b)`。<br />2. 用双层循环模拟乘法。`num1[i] * num2[j]` 的结果会影响 `res[i+j]` (高位) 和 `res[i+j+1]` (低位)。 <br />3. 遍历 res 数组，统一处理进位（将 `res[k] / 10` 加到 `res[k-1]` 上）。 | LeetCode #43. [字符串相乘](https://leetcode.cn/problems/multiply-strings/) [Multiply Strings](https://leetcode.com/problems/multiply-strings/)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| 除法           |                                      |                                    |                                       |                       |                                      | 最复杂的一种，模拟长除法。涉及大量的边界处理和循环逻辑，是模拟思想的终极考验。<br />从被除数的高位开始，截取一段比除数长的数字，然后试探这一段数字包含多少个除数（0-9 次）。将商写入结果，余数留给下一位继续构成新的被除数段。                                                                                                             | LeetCode #29. [两数相除](https://leetcode.cn/problems/divide-two-integers/) [Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)<br />LeetCode #166. [分数到小数](https://leetcode.cn/problems/fraction-to-recurring-decimal/) [Fraction to Recurring Decimal](https://leetcode.com/problems/fraction-to-recurring-decimal/)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| 整数拆到数组   | current = `num % base`               |                                    |                                       | newNum = `num / base` |                                      |                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| 数组组合为数字 |                                      |                                    |                                       |                       | result = `result * base + num`       | 逆序计算每位权重并累加当前位的值                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| 进制转换       |                                      |                                    |                                       |                       |                                      | 对大数反复执行“除以目标基数”和“取余数”的操作，直到商为 0，然后对结果逆序排列即可。                                                                                                                                                                                                                                                         | LeetCode #504. [七进制数](https://leetcode.cn/problems/base-7/) [Base 7](https://leetcode.com/problems/base-7/)<br />LeetCode #405. [数字转换为十六进制数](https://leetcode.cn/problems/convert-a-number-to-hexadecimal/) [Convert a Number to Hexadecimal](https://leetcode.com/problems/convert-a-number-to-hexadecimal/)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |

```go
// 将数字拆分到数组中
digits := []int{}
for n > 0 {
  digits = append(digits, n%base)
  n /= base
}
```

```go
// 将数组组合成数字
result := 0
for i := len(digits) - 1; i >= 0; i-- {
    result = result*base + digits[i]
}
```

在实现减法或除法时，比较大小是必不可少的前置步骤。核心逻辑：

1. 先比较长度，长者为大。
2. 若长度相等，则从高位到低位逐位比较，先遇到大数位者为大。

```go
// 大数比较
func compare(a, b string) int {
    if len(a) != len(b) {
        if len(a) > len(b) { return 1 }
        return -1
    }
    if a > b { return 1 }
    if a < b { return -1 }
    return 0
}
```

```go
// 在减法和除法中需要特别注意移除结果中的前导零
func removeLeadingZeros(s string) string {
    i := 0
    for i < len(s) && s[i] == '0' {
        i++
    }
    if i == len(s) {
        return "0"
    }
    return s[i:]
}
```

## 字符与数字的互相转换

以 Go 语言为例，转换规律如下：

<p align="center">
    <img src="type_convert.png" alt="convert-string-integer" width="25%" />
</p>

- 字符转数字：`ch - '0'`，利用 ASCII 码差值
- 数字转字符：`digit + '0'`，加上字符 '0' 的 ASCII 值

## 高精度小数

LeetCode #166. [分数到小数](https://leetcode.cn/problems/fraction-to-recurring-decimal/) [Fraction to Recurring Decimal](https://leetcode.com/problems/fraction-to-recurring-decimal/)

## 性能优优化

1. **预分配结果长度**：`make([]byte, 0, len(a)+len(b)+1)` 避免频繁扩容
2. **复用切片**：在循环中复用 `result = result[:0]` 重置长度但保留容量
3. **逆序构建**：从低位开始构建，最后一次性反转，避免频繁插入
