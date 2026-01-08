package range_addition

// getModifiedArray 使用差分数组算法处理区间更新操作
// 时间复杂度: O(n + k), 其中 n 是数组长度, k 是更新操作数
// 空间复杂度: O(n)
//
// 差分数组原理:
// 1. 对于区间 [start, end] 加上 inc, 在差分数组中:
//   - diff[start] += inc   (从 start 开始增加)
//   - diff[end+1] -= inc   (从 end+1 开始恢复)
//
// 2. 通过前缀和还原原数组: result[i] = sum(diff[0:i+1])
func getModifiedArray(length int, updates [][]int) []int {
	// 创建差分数组, 长度为 length+1 以避免边界检查
	diff := make([]int, length+1)

	// 在差分数组中标记每个更新操作的起止位置
	for _, update := range updates {
		start, end, inc := update[0], update[1], update[2]
		diff[start] += inc // 区间起点: 增加 inc
		diff[end+1] -= inc // 区间终点的下一位: 减少 inc
	}

	// 通过前缀和还原结果数组
	result := make([]int, length)
	result[0] = diff[0]
	for i := 1; i < length; i++ {
		result[i] = result[i-1] + diff[i]
	}
	return result
}
