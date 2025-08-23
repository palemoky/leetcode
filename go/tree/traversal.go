package tree

// BFS 层序遍历：核心思路是将树放入队列遍历，不断将子节点append到队列中，并将已遍历节点裁切，直至队列为空
// Time: O(n), Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	nums := [][]int{}
	if root == nil {
		return nums
	}

	queue := []*TreeNode{root} // 将整棵树放入队列
	// 遍历树的深度
	for len(queue) > 0 { // 由于最后不断地弹出已遍历元素，此处必须是 len(queue)
		levelSize := len(queue)
		row := []int{} // 收集每层的结果
		// 当前行元素的最大长度就等于树的深度，依次读取当前深度的节点
		for i := range levelSize {
			// 左右节点可能为空，因此非空时放入下一层
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			// 读取当前层的值
			row = append(row, queue[i].Val)
		}

		// 裁切掉已扫描过的节点
		queue = queue[levelSize:]
		// 将当前深度的节点放入结果集中
		nums = append(nums, row)
	}

	return nums
}

// DFS
// Preorder Traversal
// Time: O(n), Space: O(n)
func preorderIterative(root *TreeNode) []int {
	nums := []int{}
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			nums = append(nums, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return nums
}

// Inorder Traversal
// Time: O(n), Space: O(n)
func inorderIterative(root *TreeNode) []int {
	nums := []int{}

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	return nums
}

// Postorder Traversal
// Time: O(n), Space: O(n)
func postorderIterative(root *TreeNode) []int {
	nums := []int{}
	stack := []*TreeNode{}
	var prev *TreeNode
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

type ColorNode struct {
	Color int
	Node  *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// 未访问节点为白色，已访问为黑色
	const WHITE, BLACK = 0, 1

	nums := []int{}
	stack := []ColorNode{{WHITE, root}}
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := cn.Node
		color := cn.Color
		if node == nil {
			continue
		}

		if color == WHITE {
			stack = append(stack, ColorNode{WHITE, node.Right})
			stack = append(stack, ColorNode{BLACK, node})
			stack = append(stack, ColorNode{WHITE, node.Left})
		} else {
			nums = append(nums, node.Val)
		}
	}

	return nums
}

// Preorder Traversal
// Time: O(n), Space: O(n)
func preorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{root.Val}
	nums = append(nums, preorderRecursive(root.Left)...)
	nums = append(nums, preorderRecursive(root.Right)...)

	return nums
}

// Inorder Traversal
// Time: O(n), Space: O(n)
func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := inorderRecursive(root.Left)
	nums = append(nums, root.Val)
	nums = append(nums, inorderRecursive(root.Right)...)

	return nums
}

// Postorder Traversal
// Time: O(n), Space: O(n)
func postorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := postorderRecursive(root.Left)
	nums = append(nums, postorderRecursive(root.Right)...)
	nums = append(nums, root.Val)

	return nums
}
