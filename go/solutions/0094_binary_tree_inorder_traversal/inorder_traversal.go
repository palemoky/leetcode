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

// Time: O(n), Space: O(n)
func inorderIterative(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []*utils.TreeNode{}
	for root != nil || len(stack) > 0 {
		// 一路向左，不停把左子树压入栈中
		for root != nil {
			stack = append(stack, root) // push to stack
			root = root.Left
		}

		// 左边到头了，从栈中取出压入的左子树，逆序取值
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop from stack

		// 收集结果
		nums = append(nums, root.Val)

		// 转向右子树，开始新的“一路向左”
		root = root.Right
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
