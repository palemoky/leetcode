package lowest_common_ancestor_of_a_binary_tree

import "leetcode/go/solutions/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	// Base Case: 搜到底了，或者直接撞上目标
	if root == nil || root == p || root == q {
		return root
	}

	// 1. 后序遍历：先去左子树找
	leftResult := lowestCommonAncestor(root.Left, p, q)

	// 2. 后序遍历：再去右子树找
	rightResult := lowestCommonAncestor(root.Right, p, q)

	// 3. 处理根节点：根据左右子树的报告做决策
	if leftResult != nil && rightResult != nil {
		// p 和 q 分散在两侧，当前 root 就是 LCA
		return root
	}
	if leftResult != nil {
		// 两个目标都在左子树，LCA也在左子树，把左子树的报告传上去
		return leftResult
	}

	// 两个目标都在右子树，或只有一个目标在右子树
	return rightResult
}
