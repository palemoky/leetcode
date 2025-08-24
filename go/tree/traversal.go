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
				queue = append(queue, queue[i].Left) // enqueue
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right) //enqueue
			}

			// 读取当前层的值
			row = append(row, queue[i].Val)
		}

		// 裁切掉已扫描过的节点
		queue = queue[levelSize:] // dequeue
		// 将当前深度的节点放入结果集中
		nums = append(nums, row)
	}

	return nums
}

// DFS
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

// 迭代解法就是在手动维护栈结构，因此比递归更复杂
// Preorder Traversal: the simplest
// 面试刷题推荐写法，易于理解
func preorderIterativeStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
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
func preorderIterative(root *TreeNode) []int {
	nums := []int{}
	stack := []*TreeNode{}
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

// Inorder Traversal: Most challenging, most frequently tested
// Time: O(n), Space: O(n)
func inorderIterative(root *TreeNode) []int {
	nums := []int{}
	stack := []*TreeNode{}
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

// Postorder Traversal: tricky
// 面试刷题推荐，复用前序遍历逻辑，易于实现，额外的反转并非严格的后序遍历
// Time: O(n), Space: O(n)
func postorderTraversalWithReverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
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

// 颜色标记法，只需调整入栈顺序即可可通杀前、中、后序的遍历
const (
	WHITE = 0
	BLACK = 1
)

type ColorNode struct {
	Color int
	Node  *TreeNode
}

func inorderIterativeWithColor(root *TreeNode) []int {
	nums := []int{}
	stack := []ColorNode{{WHITE, root}}
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cn.Node == nil {
			continue
		}

		if cn.Color == WHITE {
			stack = append(stack, ColorNode{WHITE, cn.Node.Right})
			stack = append(stack, ColorNode{BLACK, cn.Node})
			stack = append(stack, ColorNode{WHITE, cn.Node.Left})
		} else {
			nums = append(nums, cn.Node.Val)
		}
	}

	return nums
}
