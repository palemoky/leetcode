package find_bottom_left_tree_value

import "leetcode/go/solutions/utils"

// Solution 1: 层序遍历
// Time: O(n), Space: O(n)
func findBottomLeftValue(root *utils.TreeNode) int {
	firstNode := 0
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		for i := range len(queue) {
			node := queue[0]
			queue = queue[1:]

			// 只记录每层的第一个节点
			if i == 0 {
				firstNode = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return firstNode
}
