package power_of_three

import "math"

// 最容易想到的就是不断对3整除取余
// Time: O(log₃n), Space: O(1)
func isPowerOfThreeIterative(n int) bool {
	// 负数一定不满足条件
	if n <= 0 {
		return false
	}

	// 不断对3整除
	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

// 求3的幂可以根据数学规律转换为对数
// lgn = lg(3^k) = klg3 ==> k = lgn/lg3
// 因此，如果n是3的幂，k一定是个整数
// Time: O(1), Space: O(1)
func isPowerOfThreeLog(n int) bool {
	if n <= 0 {
		return false
	}
	// 计算 k = log3(n)
	k := math.Log10(float64(n)) / math.Log10(3.0)

	// 判断 k 是否为整数。注意：由于浮点数精度问题，不能直接用 k % 1 == 0
	// 最好是检查 k 与其最近的整数的差值是否在一个极小的范围内
	return math.Abs(k-math.Round(k)) < 1e-10
}

// 由于LeetCode环境是32位，因此 1162261467 是 32位有符号整数范围内最大的3的幂 (3^19)
// 因此，更巧妙的解法则是，转变思路，如果 n 是 3 的幂，那么一定可以被最大值整除
// Time: O(1), Space: O(1)
func isPowerOfThreeMax(n int) bool {
	// 1162261467 是 32位有符号整数范围内最大的3的幂 (3^19)
	// 条件：n必须是正数，并且能被这个最大的3的幂整除
	return n > 0 && 1162261467%n == 0
}
