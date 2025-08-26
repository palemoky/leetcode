package maximum_depth_of_binary_tree

import "leetcode/go/solutions/utils"

func maxDepth(root *utils.TreeNode) int {
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
