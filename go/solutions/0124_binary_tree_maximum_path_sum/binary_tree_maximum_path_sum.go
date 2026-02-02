package binary_tree_maximum_path_sum

import (
	"math"

	"leetcode/go/solutions/utils"
)

// 推荐解法：递归 + 自底向上（后序遍历）
// 核心思路：对于每个节点，路径和 = 左子树贡献 + 右子树贡献 + 当前节点值
// 但返回给父节点的只能是"单边路径"（左或右其中一条）
// Time: O(n), Space: O(h)
func maxPathSum(root *utils.TreeNode) int {
	maxSum := math.MinInt32 // 初始化为最小值，因为节点值可能为负

	var maxGain func(*utils.TreeNode) int
	maxGain = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}

		// 后序遍历：先递归计算左右子树的最大贡献
		// 如果子树贡献为负，则不选择该子树（取 0）
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 以当前节点为"拐点"的路径和
		// 路径和 = 左子树贡献 + 右子树贡献 + 当前节点值
		currentPathSum := leftGain + rightGain + node.Val
		maxSum = max(maxSum, currentPathSum)

		// 返回给父节点的最大贡献：只能选择左或右其中一条路径
		// 贡献 = 当前节点值 + max(左子树贡献, 右子树贡献)
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
