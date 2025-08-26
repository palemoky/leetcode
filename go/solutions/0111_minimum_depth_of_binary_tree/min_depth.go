package minimum_depth_of_binary_tree

import "leetcode/go/solutions/utils"

func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1 // 非空节点的最小深度为 1
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := range levelSize {
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}

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
