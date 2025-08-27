package invert_binary_tree

import "leetcode/go/solutions/utils"

func invertTreePostorder(root *utils.TreeNode) *utils.TreeNode {
	// 触底反弹条件
	if root == nil {
		return nil
	}

	// 递归地翻转左子树和右子树
	invertTreePostorder(root.Left)
	invertTreePostorder(root.Right)

	// 交换当前节点的左右子节点
	root.Left, root.Right = root.Right, root.Left

	return root
}

func invertTreePreorder(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// 交换当前节点的左右子节点
	root.Left, root.Right = root.Right, root.Left

	// 递归地翻转新的左子树（原来是右子树）和新的右子树（原来是左子树）
	invertTreePreorder(root.Left)
	invertTreePreorder(root.Right)

	return root
}
