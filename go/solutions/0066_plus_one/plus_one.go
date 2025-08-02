package plus_one

const BASE = 10

func plusOne(digits []int) []int {
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
