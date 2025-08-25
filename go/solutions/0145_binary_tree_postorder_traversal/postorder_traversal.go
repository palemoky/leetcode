package binary_tree_postorder_traversal

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(n)
func postorderRecursive(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := postorderRecursive(root.Left)
	nums = append(nums, postorderRecursive(root.Right)...)
	nums = append(nums, root.Val)

	return nums
}

// 面试刷题推荐，复用前序遍历逻辑，易于实现，额外的反转并非严格的后序遍历
// Time: O(n), Space: O(n)
func postorderTraversalWithReverse(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*utils.TreeNode{root}
	nums := []int{}
	for len(stack) > 0 {
		root := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		nums = append(nums, root.Val) // 收集结果

		// 和前序相反，先压左，再压右
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}

	// 反转结果
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

// 更标准，严格按照后序遍历顺序处理节点，实现复杂、易出错、理解门槛高
// Time: O(n), Space: O(n)
func postorderIterative(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []*utils.TreeNode{}
	var prev *utils.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			nums = append(nums, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}

	return nums
}

func postorderIterativeWithColor(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []utils.ColorNode{{Color: utils.WHITE, Node: root}}
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cn.Node == nil {
			continue
		}

		if cn.Color == utils.WHITE {
			stack = append(stack, utils.ColorNode{Color: utils.BLACK, Node: cn.Node})
			stack = append(stack, utils.ColorNode{Color: utils.WHITE, Node: cn.Node.Right})
			stack = append(stack, utils.ColorNode{Color: utils.WHITE, Node: cn.Node.Left})
		} else {
			nums = append(nums, cn.Node.Val)
		}
	}

	return nums
}
