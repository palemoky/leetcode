package binary_tree_preorder_traversal

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(n)
func preorderRecursive(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{root.Val}
	nums = append(nums, preorderRecursive(root.Left)...)
	nums = append(nums, preorderRecursive(root.Right)...)

	return nums
}

// 面试刷题推荐写法，易于理解
func preorderIterativeStack(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*utils.TreeNode{root}
	nums := []int{}
	for len(stack) > 0 {
		root := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop

		nums = append(nums, root.Val) // 收集结果

		if root.Right != nil {
			stack = append(stack, root.Right) // 先压右
		}
		if root.Left != nil {
			stack = append(stack, root.Left) // 再压左
		}
	}

	return nums
}

// 代码简洁，与中序遍历迭代写法类似，但理解门槛略高
// Time: O(n), Space: O(n)
func preorderIterative(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []*utils.TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root) // push to stack
			root = root.Left
		}

		root = stack[len(stack)-1].Right // 从栈中取出右子树逐一处理
		stack = stack[:len(stack)-1]     // pop from stack
	}

	return nums
}

func preorderIterativeWithColor(root *utils.TreeNode) []int {
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
			stack = append(stack, utils.ColorNode{Color: utils.WHITE, Node: cn.Node.Left})
			stack = append(stack, utils.ColorNode{Color: utils.BLACK, Node: cn.Node})
		} else {
			nums = append(nums, cn.Node.Val)
		}
	}

	return nums
}
