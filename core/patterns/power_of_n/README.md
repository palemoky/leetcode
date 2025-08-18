# n 的幂问题总结

“n 的幂”问题是指判断一个数是否为 n 的某次幂，或与幂运算相关的题型。常见的有 2 的幂、3 的幂、4 的幂等。此类问题在面试和算法竞赛中非常常见，考察对数论、位运算、循环、递归等基础知识的掌握。

---

## 常见题型

- 判断一个数是否为 n 的幂（如 LeetCode 231、326、342）
- 找出小于等于某个数的所有 n 的幂
- 计算 n 的 k 次幂
- 判断一个数是否可以表示为若干 n 的幂之和

---

## 常用解法

### 1. 循环/递归除法

不断用 n 去整除，最后判断是否能整除到 1。

```go
func isPowerOfN(n, base int) bool {
    if n < 1 {
        return false
    }
    for n%base == 0 {
        n /= base
    }
    return n == 1
}
```

### 2. 对数法

利用对数的性质：logₙ(x) 如果是整数，则 x 是 n 的幂。

```go
import "math"

func isPowerOfN_Log(x, base int) bool {
    if x < 1 {
        return false
    }

    logVal := math.Log(float64(x)) / math.Log(float64(base))

    return math.Abs(logVal-math.Round(logVal)) < 1e-10
}
```

### 3. 位运算（仅适用于 2 的幂）

由于 2 的幂的二进制表示只有一个 1，其余全是 0。例如：

- 1 = 0001
- 2 = 0010
- 4 = 0100
- 8 = 1000
- 16 = 10000

n-1 则会将唯一的那个 1 变成 0，右边的所有 0 变成 1。类似于对 n 取反码，例如：

- n = 8 (1000)，n-1 = 7 (0111)
- n = 4 (0100)，n-1 = 3 (0011)
- n = 16 (10000)，n-1 = 15 (01111)

而通过按位与运算，则可以确定是否为 2 的幂。

```go
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n-1)) == 0
}
```

### 4. 最大幂整除法（只适用于质数，如 2、3）

对于正整数 n，如果 n 是 base 的幂，则 base 的最大幂能被 n 整除且无余数。

```go
import "math"

func isPowerOfN_MaxDiv(n, base int) bool {
    maxPow := maxPowerLE(base, 2147483647) // 最大限制为2^31-1
    return n > 0 && maxPow%n == 0
}

// base^k = limit ==> lgbase(limit) = k
// 根据换底公式 log_a(b) = log_c(b)/log_c(a) 得 k = lg(limit)/lg(base)
func maxPowerLE(base, limit int) int {
    if base <= 1 {
        return base
    }

    k := int(math.Floor(math.Log(float64(limit)) / math.Log(float64(base))))
    maxPow := int(math.Pow(float64(base), float64(k)))

    return maxPow
}
```

---

## 相关题目

- [LeetCode 231. Power of Two](https://leetcode.com/problems/power-of-two/)
- [LeetCode 326. Power of Three](https://leetcode.com/problems/power-of-three/)
- [LeetCode 342. Power of Four](https://leetcode.com/problems/power-of-four/)

---

## 常见变体

- 判断是否为 4 的幂、8 的幂等
- 判断是否为 n 的 k 次幂
- 判断能否用 n 的幂之和表示一个数
