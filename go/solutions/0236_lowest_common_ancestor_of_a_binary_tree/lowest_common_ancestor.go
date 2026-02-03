package lowest_common_ancestor_of_a_binary_tree

import "leetcode/go/solutions/utils"

// 推荐解法：递归 + 自底向上（后序遍历）
// Time: O(n), Space: O(h)
func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	// 递归终止条件：
	// 1. 搜到底了（nil）
	// 2. 找到目标节点（p 或 q）
	if root == nil || root == p || root == q {
		return root
	}

	// 后序遍历：先递归左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 根据左右子树的返回值判断 LCA 位置
	// 情况1：p 和 q 分散在左右两侧 → 当前节点就是 LCA
	if left != nil && right != nil {
		return root
	}

	// 情况2：p 和 q 都在左子树 → 返回左子树的结果
	if left != nil {
		return left
	}

	// 情况3：p 和 q 都在右子树（或右子树找到一个）→ 返回右子树的结果
	return right
}
