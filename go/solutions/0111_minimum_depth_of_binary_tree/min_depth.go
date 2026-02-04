package minimum_depth_of_binary_tree

import "leetcode/go/solutions/utils"

// Solution 1: BFS 找到第一个叶子节点立即返回
// Time: O(n), Space: O(n)
func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1 // 非空节点的最小深度为 1
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]

			// 找到第一个叶子节点，立即返回
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}
