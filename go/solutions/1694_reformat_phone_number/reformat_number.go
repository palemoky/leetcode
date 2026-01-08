package reformat_phonenumber

func reformatNumberBasic(number string) string {
	// 先删除'-'和' '
	digits := []byte{}
	for i := range number {
		if number[i] == '-' || number[i] == ' ' {
			continue
		}
		digits = append(digits, number[i])
	}

	n := len(digits)
	if n < 4 {
		return string(digits)
	}

	// 根据号码长度处理，余1为末尾两两一组，余2为两个一组
	ans := []byte{}
	mod := n % 3
	counter := 0
	// 根据题意，以余数的思想来解题。当余数为0和2时，处理逻辑是相同的，当余数为1时，需要末尾两两分组
	for i := range digits {
		if counter == 3 {
			ans = append(ans, '-')
			counter = 0
		}

		if mod == 1 && i == n-4 {
			break
		}

		ans = append(ans, digits[i])
		counter++
	}

	if mod == 1 {
		ans = append(ans, digits[n-4:n-2]...)
		ans = append(ans, '-')
		ans = append(ans, digits[n-2:]...)
	}

	return string(ans)
}

// 优化点：
// 1. 直接筛选出数字，过滤更健壮
// 2. 以剩余长度4为处理分界线，逻辑更清晰，代码更健壮
// 3. 循环中直接以3为步长，效率更高
// 4. 高效的 append 相比 basic 版更优雅
func reformatNumberOptimized(number string) string {
	digits := []byte{}
	for i := range number {
		if number[i] >= '0' && number[i] <= '9' {
			digits = append(digits, number[i])
		}
	}

	n := len(digits)
	ans := []byte{}
	i := 0
	for n-i > 4 {
		ans = append(ans, digits[i], digits[i+1], digits[i+2], '-')
		i += 3
	}

	if n-i == 4 {
		ans = append(ans, digits[i], digits[i+1], '-', digits[i+2], digits[i+3])
	} else {
		for ; i < n; i++ {
			ans = append(ans, digits[i])
		}
	}

	return string(ans)
}
