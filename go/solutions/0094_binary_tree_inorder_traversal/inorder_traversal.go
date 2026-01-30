package binary_tree_inorder_traversal

import "leetcode/go/solutions/utils"

func inorderRecursive(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := inorderRecursive(root.Left)
	nums = append(nums, root.Val)
	nums = append(nums, inorderRecursive(root.Right)...)

	return nums
}

// 中序遍历：左 → 根 → 右
// Time: O(n), Space: O(h) - h为树的高度
func inorderIterative(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []*utils.TreeNode{}

	curr := root
	for curr != nil || len(stack) > 0 {
		// 把节点一路向左压入栈，一直到树底
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 开始倒序处理栈中的节点（即从下往上遍历树）
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问节点
		nums = append(nums, curr.Val)

		// 转向右子树，开始新的“一路向左”
		curr = curr.Right
	}

	return nums
}

func inorderIterativeWithColor(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []utils.ColorNode{{Color: utils.WHITE, Node: root}}
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cn.Node == nil {
			continue
		}

		if cn.Color == utils.WHITE {
			stack = append(stack, utils.ColorNode{Color: utils.WHITE, Node: cn.Node.Right})
			stack = append(stack, utils.ColorNode{Color: utils.BLACK, Node: cn.Node})
			stack = append(stack, utils.ColorNode{Color: utils.WHITE, Node: cn.Node.Left})
		} else {
			nums = append(nums, cn.Node.Val)
		}
	}

	return nums
}
