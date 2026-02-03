package binary_tree_level_order_traversal_ii

import "leetcode/go/solutions/utils"

// Solution 1: 要求反向输出，在原层序遍历基础上用栈或者直接把结果插入头部即可实现
// Time: O(n), Space: O(n)
func levelOrderBottom(root *utils.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0, len(queue))
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append([][]int{level}, result...)
	}

	return result
}
