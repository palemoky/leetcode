package binary_tree_zigzag_level_order_traversal

import "leetcode/go/solutions/utils"

// Solution 1: 通过奇偶层来判断插入方向
// Time: O(n), Space: O(n)
func zigzagLevelOrder(root *utils.TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	leftToRight := true
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := range levelSize {
			node := queue[0]
			queue = queue[1:]

			// 根据方向决定插入位置
			if !leftToRight {
				i = levelSize - 1 - i
			}
			level[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, level)
		leftToRight = !leftToRight
	}

	return ans
}
