package invert_binary_tree

import "leetcode/go/solutions/utils"

// 推荐解法：前序遍历
// 优点：逻辑直观（先交换当前节点，再递归子树），代码简洁
// Time: O(n), Space: O(h)
func invertTreePreorder(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// 先交换当前节点的左右子树
	root.Left, root.Right = root.Right, root.Left

	// 再递归翻转子树
	invertTreePreorder(root.Left)
	invertTreePreorder(root.Right)

	return root
}

// 备选解法：后序遍历
// 特点：先递归翻转子树，再交换当前节点
// Time: O(n), Space: O(h)
func invertTreePostorder(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// 先递归翻转左右子树
	invertTreePostorder(root.Left)
	invertTreePostorder(root.Right)

	// 再交换当前节点的左右子树
	root.Left, root.Right = root.Right, root.Left

	return root
}
