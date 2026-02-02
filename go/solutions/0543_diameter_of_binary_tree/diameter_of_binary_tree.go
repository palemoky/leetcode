package diameter_of_binary_tree

import "leetcode/go/solutions/utils"

// 解法一：
// Time: O(n), Space: O(h)
func diameterOfBinaryTree(root *utils.TreeNode) int {
	maxDiameter := 0

	// 使用闭包捕获 maxDiameter
	var depth func(*utils.TreeNode) int
	depth = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}

		// 后序遍历：先递归计算左右子树的深度
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		// 以当前 node 为"拐点"，计算穿过它的直径
		// 直径 = 左子树深度 + 右子树深度
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		// 计算当前节点深度
		return max(leftDepth, rightDepth) + 1
	}

	depth(root)

	return maxDiameter
}
