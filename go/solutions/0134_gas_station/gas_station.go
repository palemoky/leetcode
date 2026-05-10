package gas_station

// Solution 1: 贪心
// Time: O(n), Space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	total := 0   // 全局油箱盈余，判断最终是否有解
	current := 0 // 当前油箱盈余，用来判断从当前起点出发能否走到下一站
	start := 0   // 记录起点位置

	for i := 0; i < len(gas); i++ {
		netGas := gas[i] - cost[i]
		total += netGas
		current += netGas

		// 如果当前油箱欠费了，说明从 start 到 i 这一段都不能作为起点
		if current < 0 {
			// 将起点重置为下一站
			start = i + 1
			// 清空当前油箱，从新起点重新计算
			current = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
