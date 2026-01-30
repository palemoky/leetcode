package binary_tree_postorder_traversal

import "leetcode/go/solutions/utils"

// 解法一：递归解法
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

// 解法二：反转法（巧妙但不标准）
// 思路：前序遍历的变体（根→右→左）+ 反转 = 后序遍历（左→右→根）
// Time: O(n), Space: O(n)
func postorderTraversalWithReverse(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*utils.TreeNode{root}
	nums := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nums = append(nums, node.Val)

		// 和前序相反：先压左，再压右
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// 反转结果
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

// 解法三（⭐推荐）：标准迭代法
// 后序遍历：左 → 右 → 根
// 核心：使用 prev 指针记录上一个访问的节点，判断右子树是否已访问
// Time: O(n), Space: O(h) - h为树的高度
func postorderIterative(root *utils.TreeNode) []int {
	nums := []int{}
	stack := []*utils.TreeNode{}
	var prev *utils.TreeNode
	curr := root // 使用 curr 避免与参数 root 混淆

	for curr != nil || len(stack) > 0 {
		// 一路向左，压栈所有左子节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 查看栈顶节点（不弹出）
		curr = stack[len(stack)-1]

		// 如果右子树为空或已访问，则访问当前节点
		if curr.Right == nil || curr.Right == prev {
			stack = stack[:len(stack)-1]  // 弹出
			nums = append(nums, curr.Val) // ← 第2次：真正访问
			prev = curr                   // 记录已访问
			curr = nil                    // 重置，避免重复访问左子树
		} else {
			// 右子树未访问，先访问右子树
			curr = curr.Right // ← 第1次：先去处理右子树
		}
	}

	return nums
}

// 解法四：颜色标记法
// Time: O(n), Space: O(h)
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
