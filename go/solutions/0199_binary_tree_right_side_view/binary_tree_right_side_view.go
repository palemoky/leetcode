package binary_tree_right_side_view

import "leetcode/go/solutions/utils"

// Solution 1: 层序遍历
// Time: O(n), Space: O(n)
func rightSideView(root *utils.TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		width := len(queue)
		for i := range width {
			node := queue[0]
			queue = queue[1:]

			// 只把每层的最后一个节点加入结果集
			if i == width-1 {
				ans = append(ans, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return ans
}
