package maximum_depth_of_binary_tree

import "leetcode/go/solutions/utils"

// 推荐解法：DFS 后序遍历（递归）
// 优点：
//   - 代码简洁（3行核心逻辑）
//   - 空间复杂度更优：O(h) vs BFS 的 O(w)
//     对于平衡树：h = log(n)，w ≈ n/2，差距明显
//
// 适用场景：求深度/高度
// Time: O(n), Space: O(h) h 为树高，最坏情况（退化链表）O(n)
func maxDepthDFS(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepthDFS(root.Left)
	right := maxDepthDFS(root.Right)

	return max(left, right) + 1
}

// 备选解法：BFS 层序遍历（迭代）
// 优点：
//   - 逻辑直观（直接数层数）
//   - 求最小深度时可提前终止（遇到第一个叶子节点）
//
// 缺点：
//   - 空间复杂度较高：O(w)，完全二叉树最后一层有 n/2 个节点
//
// 适用场景：求宽度、最小深度、层序输出
// Time: O(n), Space: O(w) w 为树的最大宽度
func maxDepthBFS(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
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
