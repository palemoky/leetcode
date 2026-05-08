package maximum_matching_of_players_with_trainers

import "sort"

// Solution 1: 双指针
// 本题和分饼干一样，教练的水平>运动员水平才能匹配成功
// Time: O(mlogm+nlogn), Space: O(logm+logn)
func matchPlayersAndTrainers(players []int, trainers []int) int {
	ans := 0

	// 排序后用双指针
	sort.Ints(players)
	sort.Ints(trainers)

	m, n := len(players), len(trainers)
	for i, j := 0, 0; i < m && j < n; i++ {
		// 查找高于当前运动员水平的教练
		for j < n && trainers[j] < players[i] {
			j++
		}

		// 匹配成功
		if j < n { // 注意 j 在 i 下的越界
			ans++
			j++
		}
	}

	return ans
}
