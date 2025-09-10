package path_sum_ii

import (
	"leetcode/go/solutions/utils"
	"slices"
)

func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	path := []int{}

	var dfs func(*utils.TreeNode, int)
	dfs = func(node *utils.TreeNode, left int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		left -= node.Val
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, slices.Clone(path))
		} else {
			dfs(node.Left, left)
			dfs(node.Right, left)
		}
		path = path[:len(path)-1] // 恢复现场
	}

	dfs(root, targetSum)
	return ans
}
