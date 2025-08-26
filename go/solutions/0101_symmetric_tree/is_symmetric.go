package symmetric_tree

import (
	"leetcode/go/solutions/utils"
	"math"
)

// 思路：层序遍历，收集每层的节点（包括nil节点）后判断是否为回文
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

// 思路：直接将左右子树成对放入队列比较
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
