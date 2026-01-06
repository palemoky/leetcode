package balanced_binary_tree

import "leetcode/go/solutions/utils"

// 解法一：自顶向下递归（Top-Down）
// 对每个节点都计算左右子树高度，效率较低
// Time: O(n^2) 最坏情况，Space: O(n) 递归栈
func isBalancedTopDown(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	// 检查当前节点的左右子树高度差
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	// 递归检查左右子树
	return isBalancedTopDown(root.Left) && isBalancedTopDown(root.Right)
}

// 计算树的高度
func height(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 解法二：自底向上递归（Bottom-Up，推荐）
// 一次遍历，同时计算高度和检查平衡性
// Time: O(n), Space: O(n) 递归栈
func isBalanced(root *utils.TreeNode) bool {
	return checkHeight(root) != -1
}

// 返回树的高度，如果不平衡则返回 -1
func checkHeight(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归检查左子树
	leftHeight := checkHeight(root.Left)
	if leftHeight == -1 {
		return -1 // 左子树不平衡
	}

	// 递归检查右子树
	rightHeight := checkHeight(root.Right)
	if rightHeight == -1 {
		return -1 // 右子树不平衡
	}

	// 检查当前节点是否平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1 // 当前节点不平衡
	}

	// 返回当前节点的高度
	return max(leftHeight, rightHeight) + 1
}
