package power_of_two

import "math"

// 解法一：最容易想到的不断整除取余解法
// Time: O(logn), Space: O(1)
func isPowerOfTwoIterative(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}

// 解法二：将指数问题转换为对数问题
// n = 2^k ==> lg(n) = lg(2^k) = klg(2) ==> k = lg(n)/lg(2)
// Time: O(1), Space: O(1)
func isPowerOfTwoLog(n int) bool {
	if n <= 0 {
		return false
	}

	k := math.Log10(float64(n)) / math.Log10(2)

	return math.Abs(k-math.Round(k)) < 1e-10
}

// 解法三：转变思路，n如果是2的幂，那么必然能被有效范围内的最大数整除
// Time: O(1), Space: O(1)
func isPowerOfTwoMax(n int) bool {
	// 注意：在代码中计算指数用 pow()，^是异或！
	return n > 0 && 1073741824%n == 0
}

// 解法四：二进制独有规律
func isPowerOfTwoBinary(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
