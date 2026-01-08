package car_pooling

// 题目解读：
// trips[i] = [numPassengers, from, to] 表示第 i 个行程：numPassengers 个乘客从 from 站上车，到 to 站下车
// 解题思路：使用差分数组记录每个站点的乘客数量变化
//   - 在 from 站，乘客数量 +numPassengers（上车）
//   - 在 to 站，乘客数量 -numPassengers（下车）
//
// 求解关键：遍历所有站点，累加乘客数量变化，检查是否超过 capacity
func carPooling(trips [][]int, capacity int) bool {
	// 找到最远的站点位置，题目约束 to < 1000
	maxLocation := 0
	for _, trip := range trips {
		if trip[2] > maxLocation {
			maxLocation = trip[2]
		}
	}

	// 构建差分数组，记录每个站点的乘客数量变化
	diff := make([]int, maxLocation+1)

	// 对每个行程，在上车点增加乘客，在下车点减少乘客
	for _, trip := range trips {
		numPassengers, from, to := trip[0], trip[1], trip[2]
		diff[from] += numPassengers // 上车
		diff[to] -= numPassengers   // 下车
	}

	// 遍历所有站点，累加乘客数量变化，检查是否超载
	currentPassengers := 0
	for _, change := range diff {
		currentPassengers += change
		if currentPassengers > capacity {
			return false
		}
	}

	return true
}
