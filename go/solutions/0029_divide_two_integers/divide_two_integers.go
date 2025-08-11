package divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	// 1. 处理唯一会导致商溢出的情况
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 2. 确定符号，并将数字转为正数处理
	// 为了避免 abs(MinInt32) 溢出，我们使用 int64 来存储绝对值
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	dvd := int64(dividend)
	if dvd < 0 {
		dvd = -dvd
	}
	dvs := int64(divisor)
	if dvs < 0 {
		dvs = -dvs
	}

	// 3. 主循环，模拟二进制长除法
	var quotient int64 = 0
	for dvd >= dvs {
		// 内层循环：找到最大的 2^k * dvs 块
		tempDivisor := dvs
		powerOf2 := int64(1)

		// 关键：tempDivisor << 1 <= dvd
		// 防止 tempDivisor 本身在左移时溢出 int64
		// (dvd-tempDivisor >= tempDivisor) 是 (dvd >= tempDivisor * 2) 的安全写法
		for dvd-tempDivisor >= tempDivisor {
			tempDivisor <<= 1
			powerOf2 <<= 1
		}

		// 更新 dividend 和商
		dvd -= tempDivisor
		quotient += powerOf2
	}

	// 4. 应用符号并返回
	result := int(sign * int(quotient))
	return result
}
