package add_binary

import "math/big"

// 由于输入的字符串长度最多可达10^4，因此不能直接将输入转换为int求解，这会导致计算溢出
// 因此需采用模拟加法计算

// 解法一：模拟手动二进制加法
// Time: O(max(N, M)), Space: O(max(N, M))
func addBinary(a string, b string) string {
	// 使用 byte 切片来构建结果，比字符串拼接更高效
	ans := []byte{}
	carry := 0

	// 使用两个指针从字符串末尾遍历
	i, j := len(a)-1, len(b)-1

	// 遍历完两字符串且没有进位
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		// 对有值的位加上值
		if i >= 0 {
			sum += int(a[i] - '0') // 通过 ASCII 码将字符转换为数字
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// 当前位的值是 sum % 2
		ans = append(ans, byte(sum%2+'0'))
		// 新的进位是 sum / 2
		carry = sum / 2 // 结果为浮点数时会自动丢弃小数部分
	}

	// 因为我们是从低位到高位追加的，所以需要反转结果
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}

	// 如果结果为空（例如 a="0", b="0" 的情况，在循环结束后 result 应该为 ['0']）
	// 但如果输入是空字符串（题目约束不会发生），可能需要处理。
	// 按题意，至少有一个字符，所以 result 不会为空。
	// 如果输入都是 "0"，我们的循环会产生一个 "0"，是正确的。
	// 如果没有结果（比如输入是 ""），返回 "0" 也是合理的。
	if len(ans) == 0 {
		return "0"
	}

	return string(ans)
}

// 拓展解法：日常这样的大数相加使用 math/big 包即可
func addBinaryWithBigInt(a string, b string) string {
	i := new(big.Int)
	j := new(big.Int)

	i.SetString(a, 2)
	j.SetString(b, 2)

	i.Add(i, j)

	return i.Text(2)
}
