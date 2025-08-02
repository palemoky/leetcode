package plus_one

const BASE = 10

// 解法一：通用解法
// Time: O(n), Space: O(n)
func plusOneNaive(digits []int) []int {
	// 由于是加1操作，可以直接令进位为1
	carry := 1
	ans := []int{}
	// 1. 逆序扫描数组
	for i := len(digits) - 1; i >= 0 || carry > 0; i-- {
		sum := carry
		if i >= 0 {
			sum += digits[i]
		}
		// 2. 利用数学规律求和
		carry = sum / BASE
		ans = append(ans, sum%BASE)
	}

	// 3. 对结果进行反转
	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return ans
}

// 解法二：加一操作专有解法
// 针对解法一优化掉一次for循环和额外的数组空间
// Time: O(n), Space: O(1)
func plusOneOptimized(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 情况一：未发生进位
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 情况二：发生进位
		digits[i] = 0
	}

	// 情况三：全部发生进位
	return append([]int{1}, digits...)
}
