# 初等数学模拟

在算法问题中，当处理的数字长度超过了标准整数类型（如 int64）的表示范围时，我们必须回归最基本的方法：模拟人类手动进行竖式计算的过程。这一模式的核心是将数字（通常以字符串或链表形式给出）拆解成逐个的数位，然后模拟加、减、乘、除等运算中的 **进位**、**借位** 等逻辑。

**通用核心要素：**

- **指针**: 从低位到高位遍历数字（如字符串的末尾，链表的头部）。
- **状态变量**: 如加法的 carry（进位）或减法的 borrow（借位）。
- **结果构造器**: 使用 **可变缓冲区** 高效构建结果，避免在循环中重复创建和销毁不可变的字符对象。因为每次拼接都可能需要分配一块新内存并拷贝旧值，导致 `$O(N²)$` 的时间复杂度。而可变缓冲区则是将所有小片段追加进去（这个过程通常是摊销 `$O(1)$` 的），最后再从缓冲区一次性生成最终的字符串。不同语言的处理方式如下：
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

我们以 base 表示进制，如二进制为 2，十进制为 10。根据竖式可得，不同运算方式数学规律如下:

=== "加法"

    **核心公式**: `sum = x + y + carry`

    | 项目 | 公式/说明 |
    |------|----------|
    | **当前位** | `current = sum % base` |
    | **进位** | `carry = sum / base` |
    | **算法特点** | 1. 最基础、最常见的类型<br>2. 循环条件应包含 `carry > 0`，以处理最高位的进位: `i >= 0 || j >= 0 || carry > 0`<br>3. 对于字符串结果，通常需要最后反转<br>4. 如果是链表，则使用 dummy 虚拟头节点和一个 current 指针来构建结果链表 |

    **相关题目**:

    - [415. 字符串相加](https://leetcode.com/problems/add-strings/) (十进制加法)
    - [67. 二进制求和](https://leetcode.com/problems/add-binary/) (二进制加法, BASE=2)
    - [2. 两数相加](https://leetcode.com/problems/add-two-numbers/) (链表形式,无需反转)
    - [66. 加一](https://leetcode.com/problems/plus-one/) (简化版加法,加一个常数 1)
    - [989. 数组形式的整数加法](https://leetcode.com/problems/add-to-array-form-of-integer/) (简化版字符串相加)

=== "减法"

    **核心公式**: `difference = x - y - borrow + base`

    | 项目 | 公式/说明 |
    |------|----------|
    | **当前位** | `current = difference % base` |
    | **借位** | `borrow = x - y - borrow < 0 ? 1 : 0` |
    | **算法特点** | 1. 减法是加法的逆运算，核心在于处理 **借位**<br>2. 比较 a 和 b 的大小以确定符号，然后用大数减小数<br>3. 需要先判断结果的正负，处理前导零 |

    **相关题目**:

    - LeetCode 上没有直接的减法题，但这是大厂面试的常见变体，必须掌握

=== "乘法"

    **核心公式**: `product = num1[i] * num2[j]`

    | 项目 | 公式/说明 |
    |------|----------|
    | **当前位** | `current = (product + carry) % base` |
    | **进位** | `carry = (product + carry) / base` |
    | **算法特点** | 模拟的是一个数乘以另一个数的每一位,然后将所有中间结果相加的过程<br><br>**核心逻辑**:<br>1. 初始化一个足够长的结果数组 res,长度为 `len(a) + len(b)`<br>2. 用双层循环模拟乘法。`num1[i] * num2[j]` 的结果会影响 `res[i+j]` (高位) 和 `res[i+j+1]` (低位)<br>3. 遍历 res 数组，统一处理进位(将 `res[k] / 10` 加到 `res[k-1]` 上) |

    **相关题目**:

    - [43. 字符串相乘](https://leetcode.com/problems/multiply-strings/)

=== "除法"

    **算法特点**:

    最复杂的一种，模拟长除法。涉及大量的边界处理和循环逻辑，是模拟思想的终极考验。

    从被除数的高位开始，截取一段比除数长的数字，然后试探这一段数字包含多少个除数(0-9 次)。将商写入结果，余数留给下一位继续构成新的被除数段。

    **相关题目**:

    - [29. 两数相除](https://leetcode.com/problems/divide-two-integers/)
    - [166. 分数到小数](https://leetcode.com/problems/fraction-to-recurring-decimal/)

=== "整数拆到数组"

    | 项目 | 公式/说明 |
    |------|----------|
    | **当前位** | `current = num % base` |
    | **去掉当前位** | `newNum = num / base` |

    ```go
    // 将数字拆分到数组中
    digits := []int{}
    for n > 0 {
      digits = append(digits, n%base)
      n /= base
    }
    ```

=== "数组组合为数字"

    **核心公式**: `result = result * base + num`

    **算法特点**: 逆序计算每位权重并累加当前位的值

    ```go
    // 将数组组合成数字
    result := 0
    for i := len(digits) - 1; i >= 0; i-- {
        result = result*base + digits[i]
    }
    ```

=== "进制转换"

    **算法特点**:

    对大数反复执行"除以目标基数"和"取余数"的操作，直到商为 0，然后对结果逆序排列即可。

    **相关题目**:

    - [504. 七进制数](https://leetcode.com/problems/base-7/)
    - [405. 数字转换为十六进制数](https://leetcode.com/problems/convert-a-number-to-hexadecimal/)

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

字符与数字的转换基于 ASCII 码差值:

$$
\text{字符} \xtofrom[\text{digit} + \text{'0'}]{\text{char} - \text{'0'}} \text{数字}
$$

- **字符 → 数字**: `char - '0'` (减去字符 `'0'` 的 ASCII 值)
- **数字 → 字符**: `digit + '0'` (加上字符 `'0'` 的 ASCII 值)

## 高精度小数

LeetCode [166.分数到小数](https://leetcode.cn/problems/fraction-to-recurring-decimal/) [Fraction to Recurring Decimal](https://leetcode.com/problems/fraction-to-recurring-decimal/)

## 性能优优化

1. **预分配结果长度**：`make([]byte, 0, len(a)+len(b)+1)` 避免频繁扩容
2. **复用切片**：在循环中复用 `result = result[:0]` 重置长度但保留容量
3. **逆序构建**：从低位开始构建，最后一次性反转，避免频繁插入
