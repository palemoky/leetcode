package lowest_common_ancestor_of_a_binary_search_tree

import "leetcode/go/solutions/utils"

// BST LCA 问题，可以利用二分查找来剪枝处理
func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}
