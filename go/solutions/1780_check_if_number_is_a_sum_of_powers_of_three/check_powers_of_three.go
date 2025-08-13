package check_powers_of_three

// 注意本题要求每个3的幂次（3^0, 3^1, 3^2, ...）最多只能出现一次
// 由此条件可得，余数为2时一定不符合要求（2由 3^0+3^0 组成）
// 因此，本题的判断条件就转换为了判断余数是否为2
func checkPowersOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n > 0 {
		if n%3 == 2 {
			return false
		}

		n /= 3
	}

	return true
}
