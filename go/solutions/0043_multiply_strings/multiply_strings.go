package multiply_strings

import "strings"

// 跳过中间求和的巧妙标准解法
// Time:(O^2), Space:O(n)
func multiply(num1 string, num2 string) string {
	// 1. 处理特殊情况
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	// 2. 创建结果数组，长度为 m+n，初始化为0
	// 使用 int 切片防止中间计算溢出
	ans := make([]int, m+n)

	// 3. 双层循环，从后往前遍历
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 将字符转为数字后求乘积
			product := int(num1[i]-'0') * int(num2[j]-'0')

			// 确定结果存放的位置
			high := i + j    // 高位
			low := i + j + 1 // 低位

			// 将乘积加到已有的结果上
			sum := product + ans[low]

			// 更新低位和高位（进位）
			ans[low] = sum % 10
			ans[high] += sum / 10
		}
	}

	// 4. 将结果数组转换为字符串
	// 跳过前导零
	start := 0
	for start < len(ans)-1 && ans[start] == 0 {
		start++
	}

	// 使用 strings.Builder 高效构建
	var builder strings.Builder
	for i := start; i < len(ans); i++ {
		builder.WriteByte(byte(ans[i] + '0'))
	}

	return builder.String()
}
