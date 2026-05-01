package jump_game_ii

// Solution 1: Greedy
// 核心思路：维护当前跳能到达的边界 curEnd 和全局最远覆盖范围 cover。
// 每当走到 curEnd 时，说明当前这跳已经用完，必须再跳一次，
// 将 curEnd 更新为目前能到达的最远位置 cover。
// Time: O(n), Space: O(1)
func jump(nums []int) int {
	jumps, curEnd, cover := 0, 0, 0
	for i, jump := range nums {
		// 更新从当前位置出发能到达的最远覆盖范围
		cover = max(cover, i+jump)
		// 走到当前跳的边界，且尚未到达终点，必须跳一次
		if i == curEnd && i < len(nums)-1 {
			jumps++
			curEnd = cover // 下一跳的边界扩展到目前最远处
		}
	}

	return jumps
}
