package binary_tree_level_order_traversal

import "leetcode/go/solutions/utils"

// 层序遍历顺序就像Z字形，通过两个嵌套循环，外层控制深度，内层控制宽度
// Time: O(n), Space: O(n)
func levelOrder(root *utils.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*utils.TreeNode{root} // 初始化队列，放入根节点
	for len(queue) > 0 {             // 遍历树的深度
		level := make([]int, 0, len(queue))
		for range len(queue) { // 遍历当前层的宽度
			// 从队列头部弹出节点
			node := queue[0]
			queue = queue[1:]

			// 收集当前层的值
			level = append(level, node.Val)

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
