package lowest_common_ancestor_of_a_binary_search_tree

import "leetcode/go/solutions/utils"

// 迭代解法(推荐)
// Time: O(h), Space: O(1)
// BST LCA 问题，可以利用二分查找来剪枝处理
func lowestCommonAncestorIterative(root, p, q *utils.TreeNode) *utils.TreeNode {
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

// 递归解法
// Time: O(h), Space: O(h)
func lowestCommonAncestorRecursive(root, p, q *utils.TreeNode) *utils.TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}
