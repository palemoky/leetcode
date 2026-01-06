package binary_tree_level_order_traversal

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(n)
func levelOrder(root *utils.TreeNode) [][]int {
	nums := [][]int{}
	if root == nil {
		return nums
	}

	queue := []*utils.TreeNode{root} // 将整棵树放入队列
	// 遍历树的深度
	for len(queue) > 0 { // 由于最后不断地弹出已遍历元素，此处必须是 len(queue)
		levelSize := len(queue)
		row := []int{} // 收集每层的结果
		// 当前行元素的最大长度就等于树的深度，依次读取当前深度的节点
		for i := range levelSize {
			// 左右节点可能为空，因此非空时放入下一层
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left) // enqueue
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right) // enqueue
			}

			// 读取当前层的值
			row = append(row, queue[i].Val)
		}

		// 裁切掉已扫描过的节点
		queue = queue[levelSize:] // dequeue
		// 将当前深度的节点放入结果集中
		nums = append(nums, row)
	}

	return nums
}
