# 最大公约数与最小公倍数

最大公约数（Greatest Common Divisor，GCD）

最小公倍数（Least Common Multiple，LCM）

```
   120      <-- LCM(24, 60)
----------
  24  60
----------
    12      <-- GCD(24, 60)
```

## 求最大公约数

**通过 `GCD(a, b) == 1` 可以判断两数是否互质。**

### 短除法（分解质因数，适合手算）

用最小质因数不断除这两个数，直到商互质为止。

```
2| 24 60
  ------
 2 | 12 30
    ------
    3 | 6 15
       -----
        2  5
```

24 和 60 的最大公约数即为 $2\times2\times3=12$

### 辗转相除法（欧几里得算法，常用）

该算法的精髓在于 **不断取余**。

60 和 24 的最大公约数为 12

- 60 ÷ 24 = 2 ⋯⋯ 12
- 24 ÷ **12** = 2 ⋯⋯ 0 <-- 余数为 0 时，除数即为最大公约数

另一个例子，64 和 40 的最大公约数为 8

- 64 ÷ 40 = 1 ⋯⋯ 24
- 40 ÷ 24 = 1 ⋯⋯ 16
- 24 ÷ 16 = 1 ⋯⋯ 8
- 16 ÷ **8** = 2 ⋯⋯ 0

### 质因数分解

质因数分解可以通杀 GCD 和 LCM，但分解过程很慢，尤其对于大素数，因此适合小数或手算，大数时效率低，不推荐用于编程实现。。

## 求最小公倍数

最小公倍数的短除法和公式法求解都依赖最大公约数，短除法的本质是隐式提取 GCD 的质因数。

### 短除法（适合手算）

用最小质因数不断除这两个数，直到商互质为止。

```
2| 24 60
  ------
 2 | 12 30
    ------
    3 | 6 15
       -----
        2  5
```

24 和 60 的最小公倍数即为 $2\times2\times3\times2\times5=120$

### 公式法（常用）

在已经得出整数 a, b 的 GCD 基础上，可通过以下公式求得最小公倍数：

$$
LCM(a, b) = \frac{a × b}{GCD(a, b)}
$$

### 质因数分解

同上

## 代码实现

```go
// 欧几里得算法（递归）
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

// 多个数的 GCD
func gcdSlice(nums []int) int {
    res := nums[0]
    for _, n := range nums[1:] {
        res = gcd(res, n)
    }
    return res
}

// 最小公倍数
func lcm(a, b int) int {
    return a * b / gcd(a, b)
}
```

## 应用场景

| **场景**                                 | **用到的知识**             |
| ---------------------------------------- | -------------------------- |
| **数据对齐**（比如内存分块、批处理数据） | GCD/LCM 用来找公共分块大小 |
| **分数运算**（如 1/4 + 1/6）             | GCD 约分、LCM 找公共分母   |
| **定时任务同步**（不同周期任务一起触发） | LCM 用来找共同周期         |
| **图论/循环问题**（旋转、步长）          | GCD 判断循环次数           |
| **加密算法**（RSA）                      | GCD 检查互质，保证密钥合法 |

## 相关题目

- [1071. Greatest Common Divisor of Strings](https://leetcode.cn/problems/greatest-common-divisor-of-strings/) 直接与 GCD 相关
- [914. X of a Kind in a Deck of Cards](https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards/) 求多个数的 GCD
- [204. Count Primes](https://leetcode.cn/problems/count-primes/) 质数判断和筛法（与质因数分解相关）
