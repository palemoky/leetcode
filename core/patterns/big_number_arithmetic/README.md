# 模拟加法类问题通用解法模板
## 概述
在 LeetCode 和算法面试中，“模拟加法”是一个非常经典的模式。这类问题要求我们对两个以非标准形式（如字符串、链表）表示的“大数”进行加法运算。由于数字可能非常大，无法直接装入 int64 等标准整数类型，我们必须模拟手动竖式加法的过程来解决。

适用问题类型:
- LeetCode #66. [加一](https://leetcode.cn/problems/plus-one/) ([Plus One](https://leetcode.com/problems/plus-one/))
- LeetCode #67. [二进制求和](https://leetcode.cn/problems/add-binary/) ([Add Binary](https://leetcode.com/problems/add-binary/))
- LeetCode #2. [两数相加](https://leetcode.cn/problems/add-two-numbers) ([Add Two Numbers](https://leetcode.com/problems/add-two-numbers))
- LeetCode #415. [字符串相加](https://leetcode.cn/problems/add-strings) ([Add Strings](https://leetcode.com/problems/add-strings))
- LeetCode #989. [数组形式的整数加法](https://leetcode.cn/problems/add-to-array-form-of-integer/) ([Add to Array-Form of Integer](https://leetcode.com/problems/add-to-array-form-of-integer/))
- 以及其他类似的大数加法变体。

---

## 核心思想：模拟竖式加法
### 数学规律
以十进制的`99+1`为例，我们得到以下竖式：
```
  99
+  1
 ¹¹¹
-----
 100
```
*注意最后的进位处理。*

我们用`base`来表示进制（如二进制、十进制、十六进制等），在进行每位计算时，
- `sum%base`：留在当前位的数字是
- `sum/base`：送往前一位的进位是（结果为浮点数时舍弃小数部分）

以上规律的逆运算即为将整数拆分到逆序数组的规律：
- `num%base`：取当前位
- `num/base`：去掉已处理的当前位
```go
digits := []int{}
for n > 0 {
  digits = append(digits, n%base)
  n /= base
}
```
以数字123推演上述代码执行过程：
1.	123 % 10 = 3 👉 得到 个位
2.	123 / 10 % 10 = 2 👉 得到 十位
3.	123 / 100 % 10 = 1 👉 得到 百位

### 操作步骤
算法的精髓可以分解为以下几个步骤：
1. 从最低位开始：无论是字符串的末尾还是链表的头部（取决于数字的表示顺序），我们总是从数字的最低有效位开始逐位相加。
2. 维护一个进位 (carry)：这是最关键的变量。它记录了前一位计算产生的进位（0 或 1），并参与到当前位的计算中。
3. 逐位计算：在每个位置上，执行 `sum = a的当前位 + b的当前位 + carry`。
确定当前位和新进位：
   - 写入结果的当前位数字 = `sum % base` （base 是进制基数，如 10 或 2）
   - 送往下一位的新进位 = `sum / base` （利用整数除法自动取整）
4. 循环直到最高位：持续这个过程，直到两个输入的所有位都被处理完毕，并且最后的进位也为 0。
5. 输出string和ListNode的处理：
   - 如果输出为string，则尽量使用数组结构类型高效构建字符串，避免在循环中使用低效的字符串拼接
     - 注意将结果字符串反转
   - 使用 dummy 虚拟头节点和一个 current 指针来构建结果链表
     - 因链表输入已经是逆序，因此无需反转

### 实现细节与注意事项
- 字符转数字：利用ASCII码的连续性，通过`int(char - '0')`巧妙实现将字符转换为对应的十进制数值
- 数字转字符：`byte(digit + '0')` 是上述操作的逆过程
- 循环条件：`for i >= 0 || j >= 0 || carry > 0` 是最健壮的循环条件。它优雅地处理了不等长输入和最终有进位的情况，无需在循环外单独处理最后的 carry
- 输出结果的边界处理：在输出结果中注意兼容两输入均为0或空
- 生产级代码: 在生产项目中，如果需要处理大数运算，应首选标准库 `math/big`，它提供了安全、高效且功能完备的大数支持
