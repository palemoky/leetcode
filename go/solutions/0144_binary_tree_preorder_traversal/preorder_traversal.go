package binary_tree_preorder_traversal

import "leetcode/go/solutions/utils"

// 解法一：递归
// 优点：代码简洁
// 缺点：重复计算，爆栈
// Time: O(n), Space: O(h)
func preorderRecursive(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{root.Val}
	nums = append(nums, preorderRecursive(root.Left)...)
	nums = append(nums, preorderRecursive(root.Right)...)

	return nums
}

// 解法二（⭐推荐）：迭代 + 栈
// Time: O(n), Space: O(h) - h为树的高度
func preorderIterativeStack(root *utils.TreeNode) []int {
	nums := []int{}
	if root == nil {
		return nums
	}

	stack := []*utils.TreeNode{root}
	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问节点
		nums = append(nums, node.Val)

		// 先压右子节点，再压左子节点（栈是LIFO，所以先压右）
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return nums
}

// 解法三：代码简洁，与中序遍历迭代写法类似，但理解门槛略高
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

// 解法四：颜色标记法，可通杀前、中、后序遍历
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
