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
		currentDiameter := leftDepth + rightDepth
		if currentDiameter > maxDiameter {
			maxDiameter = currentDiameter // 闭包可以修改外层变量
		}

		// 返回当前节点的深度给父节点
		// 深度 = 1 (当前节点) + 左右子树中较深的那个
		if leftDepth > rightDepth {
			return 1 + leftDepth
		}
		return 1 + rightDepth
	}

	depth(root)
	return maxDiameter
}
