package path_sum

import "leetcode/go/solutions/utils"

// 解法一：前序遍历
func hasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 到达叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	remainingSum := targetSum - root.Val
	// 只需要左、右子树其中一个满足条件即可
	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}
