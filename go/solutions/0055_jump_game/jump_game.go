package jump_game

// Solution 1: 贪心算法
// 思路：维护当前可达的最远下标 cover。
// 若遍历到位置 i 时 i > cover，说明出现断层，无法到达终点；
// 若 cover >= len(nums)-1，说明已能覆盖终点。
// Time: O(n), Space: O(1)
func canJump(nums []int) bool {
	// 当前可达的最远下标
	cover := 0
	for i, jump := range nums {
		// 当前位置不可达，后续位置也不可能可达
		if i > cover {
			return false
		}

		// 用当前位置更新可达范围上限
		cover = max(cover, i+jump) // 从 i 最远可到 i+jump

		// 已覆盖终点，答案确定为 true，可提前结束
		if cover >= len(nums)-1 {
			break
		}
	}

	return true
}
