package base7

import "strconv"

const BASE = 7

func convertToBase7Naive(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := num < 0
	if isNegative {
		num = -num
	}

	ans := []byte{}
	for num > 0 {
		ans = append(ans, byte(num%BASE+'0'))
		num /= BASE
	}

	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	if isNegative {
		return "-" + string(ans)
	}

	return string(ans)
}

func convertToBase7Recursive(num int) string {
	if num < 0 {
		return "-" + convertToBase7Recursive(-num)
	}

	if num < 7 {
		return strconv.Itoa(num)
	}

	// 递归调用得到高位部分，然后拼接上当前位的余数
	return convertToBase7Recursive(num/7) + strconv.Itoa(num%7)
}
