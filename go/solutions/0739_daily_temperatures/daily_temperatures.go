package daily_temperatures

// 解法一：暴力解法
// Time: O(n^2), Space: O(n)
func dailyTemperaturesBruteForce(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	for i := range n {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				ans[i] = j - i
				break
			}
		}
	}

	return ans
}

// 解法二（推荐）：单调递减栈，通过不断维护一个栈顶小于当前元素的栈来求解。求解时最好画个栈的过程
// 出栈更新
// Time: O(n), Space: O(n)
func dailyTemperaturesStackLeftToRight(temperatures []int) []int {
	tLen := len(temperatures)
	ans := make([]int, tLen)

	stack := []int{} // 栈内存放 temperatures 的索引
	for i := range tLen {
		// 如果当前元素 > 栈顶元素，则不断弹出栈中的元素，直至当前元素 < 栈顶元素
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]   // 取栈顶值，即 temperatures 数组中 i 之前的索引
			ans[top] = i - top           // 将差值更新到对应的位置
			stack = stack[:len(stack)-1] // pop
		}
		stack = append(stack, i) // push
	}

	return ans
}

// 入栈更新
func dailyTemperaturesStackRightToLeft(temperatures []int) []int {
	tLen := len(temperatures)
	ans := make([]int, tLen)

	stack := []int{} // 栈内存放 temperatures 的索引
	for i := tLen - 1; i >= 0; i-- {
		// 弹出所有小于等于当前温度的元素
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1] // pop
		}

		// 如果栈不为空，栈顶就是下一个更高温度的索引
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i) // push
	}

	return ans
}
