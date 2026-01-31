package symmetric_tree

import (
	"math"

	"leetcode/go/solutions/utils"
)

// 推荐解法：镜像递归
// 核心思路：对称树 = 左子树的左孩子 vs 右子树的右孩子 && 左子树的右孩子 vs 右子树的左孩子
// 优点：代码简洁，逻辑清晰，面试首选
// Time: O(n), Space: O(h) h 为树高（递归栈深度）
func isSymmetricMirrorRecursive(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *utils.TreeNode) bool {
	// 递归终止条件
	// 检查节点存在的对称性
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	// 递归处理逻辑
	// 检查节点值的对称性
	if left.Val != right.Val {
		return false
	}

	// 递归处理：交叉比较子树（镜像对称）
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 解法一：层序遍历，收集每层的节点（包括nil节点）后判断是否为回文
// 缺点：需要额外空间存储每一层的节点值
// Time: O(n), Space(n)
func isSymmetricLevelOrder(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		row := []int{}
		for i := range levelSize {
			if queue[i] == nil {
				row = append(row, math.MinInt32) // 用特殊值表示空节点
				continue
			}
			row = append(row, queue[i].Val)
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}

		left, right := 0, len(row)-1
		for left < right {
			if row[left] != row[right] {
				return false
			}
			left++
			right--
		}

		queue = queue[levelSize:]
	}

	return true
}

// 解法二：直接将左右子树成对放入队列比较（迭代版镜像比较）
// 优点：无需存储整层所有节点的值，只存储待比较的节点对，因此效率更高
// Time: O(n), Space(n)
func isSymmetricTwoQueues(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*utils.TreeNode{root.Left, root.Right} // 一开始就把左右子节点成对放入
	for len(queue) > 0 {
		// 每次从队列中取出两个节点进行比较
		leftNode := queue[0]
		rightNode := queue[1]
		queue = queue[2:]

		// 1. 两个都为空，对称，继续
		if leftNode == nil && rightNode == nil {
			continue
		}
		// 2. 一个为空，一个不为空，或者值不同，不对称
		if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			return false
		}

		// 3. 关键：按镜像顺序将子节点成对入队
		// leftNode 的左子树应该和 rightNode 的右子树 对称
		queue = append(queue, leftNode.Left)
		queue = append(queue, rightNode.Right)

		// leftNode 的右子树应该和 rightNode 的左子树 对称
		queue = append(queue, leftNode.Right)
		queue = append(queue, rightNode.Left)
	}

	return true
}
