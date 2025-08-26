package maximum_depth_of_binary_tree

import "leetcode/go/solutions/utils"

// 解法一：层序遍历
// Time: O(n), Space: O(n)
func maxDepthBFS(root *utils.TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := range levelSize {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelSize:]
		depth++
	}

	return depth
}

// 解法二：递归求解
// Time: O(n), Space: O(h) h 为树高，最坏 O(n)
func maxDepthDFS(root *utils.TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	left, right := maxDepthDFS(root.Left), maxDepthDFS(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
