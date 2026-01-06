package validate_binary_search_tree

import (
	"math"

	"leetcode/go/solutions/utils"
)

// 解法一思路：中序遍历读取二叉树的值，然后在数组中判断是否有序
// Time: O(n), Space: O(n)
func isValidBSTSortedArray(root *utils.TreeNode) bool {
	nums := []int{}
	inorderTraversal(root, &nums)

	for i := 1; i < len(nums); i++ {
		// 必须是严格递增，相等也不行
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true
}

func inorderTraversal(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
}

// 解法二：递归约束
// Time: O(n), Space: O(h)
func isValidBSTDFS(root *utils.TreeNode) bool {
	return isValid(root, int64(math.MinInt64), int64(math.MaxInt64))
}

func isValid(node *utils.TreeNode, lower, upper int64) bool {
	if node == nil {
		return true
	}

	val := int64(node.Val)
	// 检查当前节点的值是否在 (lower, upper) 范围内
	if val <= lower || val >= upper {
		return false
	}

	// 递归检查左右子树，并更新边界（类似于滑动窗口）
	// 对于左子树，上界是当前节点的值
	if !isValid(node.Left, lower, val) {
		return false
	}
	// 对于右子树，下界是当前节点的值
	if !isValid(node.Right, val, upper) {
		return false
	}

	return true
}
